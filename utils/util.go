package utils

import (
	"crypto/md5"
	"fmt"
)

func GenMd5Key(id string) string {
	data := []byte(id)
	return fmt.Sprintf("%x", md5.Sum(data))
}
