package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"strconv"

	"github.com/SSSaaS/sssa-golang"
)

func createSecrets() {

	var shares, threshold int
	var arg string

	for {
		fmt.Print("How many shares? ")
		_, err := fmt.Scanln(&arg)
		res, err := strconv.ParseInt(arg, 10, 32)
		if err == nil {
			shares = int(res)
			break
		} else {
			fmt.Println("Invalid input - Numbers only")
			fmt.Println(err)
		}
	}

	for {
		fmt.Print("Threshold? ")
		_, err := fmt.Scanln(&arg)
		res, err := strconv.ParseInt(arg, 10, 32)
		if err == nil {
			threshold = int(res)
			break
		} else {
			fmt.Println("Invalid input - Numbers only")
			fmt.Println(err)
		}
	}

	// generate a random 256-bit secret
	max := new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil)
	x, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}
	src := x.Bytes()
	encodedStr := hex.EncodeToString(src)

	// fmt.Println("Secret: ", encodedStr)

	secrets, err := sssa.Create(threshold, shares, encodedStr)
	if err != nil {
		panic(err)
	}
	for i := range secrets {
		fmt.Printf("Secret %v: ", i+1)
		fmt.Print(secrets[i], "\n")
	}

}

func main() {

	args := os.Args[1:]

	if len(args) == 0 {
		createSecrets()
	} else {
		switch args[0] {
		case "combine":
			combine()
		}
	}

}
