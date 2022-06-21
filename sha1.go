package enc

import (
	"crypto/sha1"
	"encoding/hex"
)

func SHA1(src string) string {
	bytes := sha1.Sum([]byte(src))
	return hex.EncodeToString(bytes[:])
}
