package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"dbo-test/config"
	"encoding/hex"
)

func GenerateHmacSHA256(data string) string {
	h := hmac.New(sha256.New, []byte(config.Configure.Hash.Secret))
	h.Write([]byte(data))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}
