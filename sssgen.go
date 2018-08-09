package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/SSSaaS/sssa-golang"
)

func main() {
	max := new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil)
	x, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}

	src := x.Bytes()
	encodedStr := hex.EncodeToString(src)

	fmt.Println("Secret: ", encodedStr)

	secrets, err := sssa.Create(2, 3, encodedStr)
	if err != nil {
		panic(err)
	}
	for i := range secrets {
		fmt.Printf("Secret %v: ", i)
		fmt.Print(secrets[i], "\n")
	}

	fmt.Println("Now we combine")

	var text1, text2 string

	fmt.Print("Enter secret 1: ")
	fmt.Scanf("%s", &text1)

	fmt.Print("Enter secret 2: ")
	fmt.Scanf("%s", &text2)

	a := make([]string, 2)

	a[0] = text1
	a[1] = text2

	secret, err := sssa.Combine(a)

	if err != nil {
		panic(err)
	}

	fmt.Println(secret)

}
