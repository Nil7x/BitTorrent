package main

import (
	"BitTorrentDemo/torrent"
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main()  {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("open file error")
		return
	}
	defer file.Close()
	tf, err := torrent.ParseFile(bufio.NewReader(file))
	if err != nil {
		fmt.Println("parse file error")
		return
	}

	var peerId [torrent.IDLEN]byte
	_, _ = rand.Read(peerId[:])
	peers := torrent.FindPeers(tf, peerId)
	if len(peers) == 0 {
		fmt.Println("can not find peers")
		return
	}

	task := &torrent.TorrentTask{
		PeerId:   peerId,
		PeerList: peers,
		InfoSHA:  tf.InfoSHA,
		FileName: tf.FileName,
		FileLen:  tf.FileLen,
		PieceLen: tf.PieceLen,
		PieceSHA: tf.PieceSHA,
	}
	torrent.Download(task)
}
