package utils

import "crypto/md5"

func MD5(str string) [16]byte {
	return md5.Sum([]byte(str))
}
