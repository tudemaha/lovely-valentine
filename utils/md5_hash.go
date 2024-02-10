package utils

import (
	"crypto/md5"
	"fmt"
)

func CreateMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	hashString := fmt.Sprintf("%x", hash)
	return hashString
}
