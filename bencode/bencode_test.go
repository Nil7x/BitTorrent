package bencode

import (
    "bytes"
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestString(t *testing.T) {
    val := "abc"
    buf := new(bytes.Buffer)
    wLen := EncodeString(buf, val)
    assert.Equal(t, 5, wLen)
    str, _ := DecodeString(buf)
    assert.Equal(t, val, str)

    val = ""
    for i := 0; i < 20; i++ {
        val += string(byte('a' + i))
    }
    fmt.Println(val)
    buf.Reset()
    wLen = EncodeString(buf, val)
    fmt.Println(buf)
    assert.Equal(t, 23, wLen)
    str, _ =DecodeString(buf)
    assert.Equal(t, str, val)
    fmt.Println(str)
}

func TestInt(t *testing.T)  {
    val := 999
    buf := new(bytes.Buffer)
    wLen := EncodeInt(buf, val)
    assert.Equal(t, 5, wLen)
    iv, _ := DecodeInt(buf)
    assert.Equal(t, val, iv)

    val = 0
    buf.Reset()
    wLen = EncodeInt(buf, val)
    assert.Equal(t, 3, wLen)
    iv, _ = DecodeInt(buf)
    assert.Equal(t, val, iv)

    val = -99
    buf.Reset()
    wLen = EncodeInt(buf, val)
    assert.Equal(t, 5, wLen)
    iv, _ = DecodeInt(buf)
    assert.Equal(t, val, iv)
}