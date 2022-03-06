package utils

import (
	"fmt"
	"crypto/md5"
)

func ToMd5(password string) string {
	data := []byte(password)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return  md5str
}