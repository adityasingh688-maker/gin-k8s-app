package main

import (
	"crypto/md5"
	"fmt"
)

func weakHash(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
