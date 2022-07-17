package signer

import (
	"crypto/ecdsa"
	"log"

	"github.com/joho/godotenv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func Account() (common.Address, *ecdsa.PrivateKey) {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Failed to load the .env file %v", err)
	}

	privateKey, err := crypto.GenerateKey()

	if err != nil {
		log.Fatalf("Failed to parse the private key %v", err)
	}

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		log.Fatal("Failed assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	signerAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return signerAddress, privateKey
}
