# Paystack Golang SDK

Simple and un-obtrusive wrapper for Paystack APIs. Majorly exists cause of my distaste in existing packages.

## Installation

```bash
go get github.com/Nullbase-Technologies/paystack-go
```

## Usage

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Nullbase-Technologies/paystack-go"
	"github.com/Nullbase-Technologies/paystack-go/transactions"
)

func main() {
	client, err := paystack.New(
		paystack.WithSecretKey("sk_test_xxxxxxxxx"),
		// Optionally initialize client with your own http client
		paystack.WithHTTPClient(CustomHTTPClient{}),
	)
	if err != nil {
		log.Fatal(err)
	}

	transaction := transactions.New(client)

	initResp, err := transaction.Initialize(context.TODO(), &transactions.InitializeTransactionOptions{
		Amount: 1000,
		Email:  "adedaramola@gmail.com",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(initResp.AccessCode)
	fmt.Println(initResp.Reference)
	fmt.Println(initResp.AuthorizationURL)
}
```

## Webhook Validation

The package also provides a simple helper function to validate incoming paystack webhooks. A use case could be something like below

```go
func VerifyPaystackSignature(next http.Handler) http.Handler {
	return http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		// Load paystack secret key from environment
		// or however you choose to do so, no sha hardcode am!
		secret := os.Getenv("PAYSTACK_SECRET_KEY")

		signature := r.Header.Get("X-Paystack-Signature")

		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to read request body", http.StatusInternalServerError)
			return
		}

		if err := paystack.VerifyWebhookSignature(
			bodyBytes, signature, secret); err != nil {
			// You can type hint the actual invalid webhook error
			if errors.Is(err, paystack.ErrInvalidPaystackWebhook) {
				// do whatever you want
			}
			// maybe throw an internal server error
		}

		// everything is good
		// handle the webhook, however you will of course :)

		next.ServeHTTP(w, r)
	})
}
```

## SDK Coverage

- [ ] Transactions
	- [x] Initialize Transaction
	- [x] Verify Transaction
	- [x] List Transactions
	- [x] Fetch Transaction
	- [x] Charge Authorization
	- [ ] View Transaction Timeline
	- [ ] Transaction Totals
	- [ ] Export Transactions
	- [ ] Partial Debit 
- [x] Transaction Splits
	- [x] Create Split
	- [x] List/Search Split
	- [x] Fetch Split
	- [x] Update Split
	- [x] Add/Update Split Subaccount
	- [x] Remove Subaccount from Split
- [ ] Customers
	- [x] Create Customer
	- [x] List Customer
	- [x] Fetch Customer
	- [x] Update Customer
	- [x] Validate Customer
	- [x] Whitelist/Blacklist Customer
	- [x] Initialize Authorization
	- [x] Verify Authorization
	- [x] Deactivate Authorization
	- [ ] Initialize Direct Debit
	- [ ] Direct Debit Activation Charge
	- [ ] Fetch Mandate Authorization
- [x] Plans
	- [x] Create Plan
	- [x] List Plan
	- [x] Fetch Plan
	- [x] Update Plan
- [ ] Miscellaneous
	- [x] List Banks
	- [ ] List Countries
	- [ ] List States
