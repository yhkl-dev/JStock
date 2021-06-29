package common

import (
	"crypto/md5"
	"fmt"
)

func MD5Encrypt(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
