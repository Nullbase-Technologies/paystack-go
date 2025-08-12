package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Nullbase-Technologies/paystack-go"
	"github.com/Nullbase-Technologies/paystack-go/customers"
)

func main() {
	client, err := paystack.New(
		paystack.WithSecretKey("sk_test_49b3451b3cb0887322e55ece94fb422c4a8c1fb7"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// transaction := transactions.New(client)

	// initResp, err := transaction.Initialize(context.TODO(), &transactions.InitializeTransactionOptions{
	// 	Amount: 1000,
	// 	Email:  "adedaramola@gmail.com",
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(initResp.AccessCode)
	// fmt.Println(initResp.Reference)
	// fmt.Println(initResp.AuthorizationURL)

	// t, err := transaction.List(context.TODO(), &transactions.ListTransactionsParams{
	// 	Status: "abandoned",
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	customer, err := customers.New(client).Create(context.TODO(), &customers.CreateCustomerOptions{
		Email:     "paranok2011@gmail.com",
		FirstName: "Adedaramola",
		LastName:  "Adetimehin",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(customer)
}
