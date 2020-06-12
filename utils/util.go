package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GenMd5Key(id string) string {
	h := md5.New()
	md5Key := hex.EncodeToString(h.Sum(nil))
	return md5Key
}
