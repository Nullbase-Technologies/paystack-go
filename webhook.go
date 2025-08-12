package paystack

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"errors"
)

var ErrInvalidPaystackWebhook = errors.New("invalid paystack webhook")

// VerifyWebhookSignature checks if a Paystack webhook is valid
func VerifyWebhookSignature(body []byte, signature, secret string) error {
	if body == nil {
		return errors.New("body cannot be empty")
	}

	if IsStringEmpty(signature) {
		return errors.New("signature cannot be empty")
	}

	if IsStringEmpty(secret) {
		return errors.New("provide a valid secret key")
	}

	mac := hmac.New(sha512.New, []byte(secret))
	mac.Write(body)

	expected := hex.EncodeToString(mac.Sum(nil))

	if !hmac.Equal([]byte(expected), []byte(signature)) {
		return ErrInvalidPaystackWebhook
	}

	return nil
}
