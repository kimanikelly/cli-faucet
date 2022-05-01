package contract

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Returns the provider url based on the "ENVIRONMENT" .env variable
func LoadEnvironment() string {

	// Provider url based on the ENVIRONMENT variable set in the .env file
	var providerUrl = ""

	// Loads the environment variables from the .env file
	err := godotenv.Load()

	// If err does not equal nil(zero value) throw an error
	if err != nil {
		log.Fatalf("Failed to load environment variables %v", err)
	}

	// Returns the value for the ENVIRONMENT
	environment := os.Getenv("ENVIRONMENT")

	// Checks if environment is set to Localhost
	if environment == "Localhost" {

		// Sets environment to the localhost connection
		providerUrl = "http://localhost:8545"

		// Checks if the environment is set to Rinkeby
	} else if environment == "Rinkeby" {

		// Sets environment to the Rinkeby connection
		providerUrl = os.Getenv("RINKEBY_URL")

	}

	// Returns the providerUrl variable
	return providerUrl
}
