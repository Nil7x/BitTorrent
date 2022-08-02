package torrent

import (
	"fmt"
	"io"
)

const (
	Reserved	int = 8
	HsMsgLen	int = SHALEN + IDLEN + Reserved //20 20 8
)

type HandShakeMsg struct {
	PreStr	string
	InfoSHA	[SHALEN]byte
	PeerId	[IDLEN]byte
}

func NewHandShakeMsg(infoSHA, peerId [IDLEN]byte) *HandShakeMsg {
	return &HandShakeMsg{
		PreStr:  "BitTorrent protocol",
		InfoSHA: infoSHA,
		PeerId:  peerId,
	}
}

func WriteHandShake(w io.Writer, msg *HandShakeMsg) (int, error) {
	buf := make([]byte, len(msg.PreStr) + HsMsgLen + 1)
	buf[0] = byte(len(msg.PreStr))
	cur := 1
	cur += copy(buf[cur:], msg.PreStr) // check
	cur += copy(buf[cur:], make([]byte, Reserved))
	cur += copy(buf[cur:], msg.InfoSHA[:])
	cur += copy(buf[cur:], msg.PeerId[:])
	return w.Write(buf)
}

func ReadHandShake(r io.Reader) (*HandShakeMsg, error) {
	lenBuf := make([]byte, 1)
	_, err := io.ReadFull(r, lenBuf)
	if err != nil {
		return nil, err
	}
	prelen := int(lenBuf[0])
	if prelen == 0 {
		err := fmt.Errorf("prelen cannot be 0")
		return nil, err
	}

	msgBuf := make([]byte, HsMsgLen + prelen)
	_, err = io.ReadFull(r, msgBuf)
	if err != nil {
		return nil, err
	}

	var peerId [IDLEN]byte
	var infoSHA [SHALEN]byte

	copy(infoSHA[:], msgBuf[prelen+Reserved: prelen+Reserved+SHALEN])
	copy(peerId[:], msgBuf[prelen+Reserved+SHALEN:])

	return &HandShakeMsg{
		PreStr:  string(msgBuf[0:prelen]),
		InfoSHA: infoSHA,
		PeerId:  peerId,
	}, nil
}