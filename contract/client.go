package contract

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

// Connects to the provider and returns the instance
func ProviderConnection() *ethclient.Client {

	// Connects to the provider and creates the client instance
	client, err := ethclient.Dial(LoadEnvironment())

	// If err does not equal nil(zero value) throw an error and exit
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Returns the client instance
	return client

}
