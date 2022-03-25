package main

import (
	"fmt"
	"testing"

	"github.com/licat233/goutil/getpath"
)

func TestGetPath(t *testing.T) {
	res := getpath.CurrentPath()
	fmt.Println(res)
}
