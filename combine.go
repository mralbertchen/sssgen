package main

import (
	"fmt"

	sssa "github.com/SSSaaS/sssa-golang"
)

func combine() {
	fmt.Println("Combining secrets!")

	var n int
	fmt.Print("How many secrets are you combining? ")
	fmt.Scanf("%d", &n)

	shares := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Enter secret %v: ", i+1)
		fmt.Scanf("%s", &shares[i])
	}

	secret, err := sssa.Combine(shares)

	if err != nil {
		panic(err)
	}

	fmt.Println(secret)

}
