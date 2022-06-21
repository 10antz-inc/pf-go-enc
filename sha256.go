package enc

import (
	"crypto/sha256"
	"encoding/hex"
)

func SHA256(src string) string {
	bytes := sha256.Sum256([]byte(src))
	return hex.EncodeToString(bytes[:])
}
