package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func VerifySignature(signature string, body string) bool {
	signature = strings.TrimPrefix(signature, "sha256=")
	sig, err := hex.DecodeString(signature)
	if err != nil {
		fmt.Println("error decoding signature", err)
		return false
	}

	secret := os.Getenv("GITHUB_WEBHOOK_SECRET")
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(body))
	computedHash := h.Sum(nil)

	return hmac.Equal(sig, computedHash)
}
