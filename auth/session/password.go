package session

import (
	"crypto/md5"
	"encoding/hex"
)

type MD5HashFunc = func(text string) string

func MD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
