package testpkg

import (
	"crypto/md5"
	"fmt"
)

func Md5New(text string) string {
	return fmt.Sprintf("X", md5.Sum([]byte(text)))
}
