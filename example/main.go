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
		paystack.WithSecretKey("<your-secret-key-here>"),
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
