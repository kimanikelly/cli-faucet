package contract

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func Connection() *ethclient.Client {
	client, err := ethclient.Dial("http://localhost:8545")

	// If err does not equal nil(zero value) throw an error and exit the process
	// The "Failed to connect to the Ethereum client: %v" will only log if there is an error
	// connecting to the Ethereum provider
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	return client

}
