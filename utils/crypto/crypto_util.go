package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMd5(input string, slice string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum([]byte(slice)))
}
