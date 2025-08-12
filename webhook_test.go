package paystack

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"testing"
)

func TestVerifyWebhookSignature(t *testing.T) {
	secret := "sk_test_xxxx"
	body := []byte(`{"event": "charge.success", "data": {"id": 12345}}`)

	mac := hmac.New(sha512.New, []byte(secret))
	mac.Write(body)
	signature := hex.EncodeToString(mac.Sum(nil))

	tests := []struct {
		name        string
		secret      string
		body        []byte
		signature   string
		expectError bool
	}{
		{
			name:        "can verify webhook with valid signature",
			secret:      secret,
			body:        body,
			signature:   signature,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := VerifyWebhookSignature(tt.body, tt.signature, tt.secret)
			if (err != nil) != tt.expectError {
				t.Errorf("expected error = %v, got %v", tt.expectError, err)
			}
		})
	}
}
