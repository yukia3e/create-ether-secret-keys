package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/crypto"
)

func generateEthereumKeys(count int) {
	for i := 0; i < count; i++ {
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			log.Fatal(err)
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		}

		address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

		fmt.Printf("Key Pair %d:\n", i+1)
		fmt.Printf("Private Key: 0x%x\n", crypto.FromECDSA(privateKey))
		fmt.Printf("Public Key: 0x%x\n", crypto.FromECDSAPub(publicKeyECDSA))
		fmt.Printf("Address: %s\n\n", address)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <number_of_keys>")
		os.Exit(1)
	}

	keyCount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error: Invalid number of keys. Please provide a valid integer.")
		os.Exit(1)
	}

	if keyCount <= 0 {
		fmt.Println("Error: Number of keys must be greater than 0.")
		os.Exit(1)
	}

	generateEthereumKeys(keyCount)
}
