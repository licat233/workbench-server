package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"testing"
)

func Test_MD5(t *testing.T) {
	str := []byte("abcd")
	b := byte('b')
	res := MD5V(str, b)
	fmt.Println(res)
}

func MD5V(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}
