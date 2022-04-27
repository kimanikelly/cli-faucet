package contract

import (
	"encoding/json"
	"io"
	"log"

	// "os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
)

// Stores the UnMarshalled address of Token.sol
type AddressList struct {
	Rinkeby string
}

type ContractResponseData struct {
	Addresses AddressList
}

//
var ContractData ContractResponseData

// The deployed Token.sol address on the Rinkeby testnet
var TokenAddress string

// Creates and returns the Token.sol contract instance
func ContractInstance() *Token {

	err := godotenv.Load()

	// If err does not equal nil(zero value) throw an error and exit the process
	// The "Failed to load the .env file" message will only log if there is an error loading the .env file
	if err != nil {
		log.Fatalf("Failed to load the .env file %v", err)
	}

	// what is io?
	// fetchContractData() - A function that returns an HTTP response type *http.Response
	// fetchContractData().Body - Returns the HTTP response Body type *http.bodyEOFSignal

	// The ReadAll function
	body, err := io.ReadAll(fetchContractData().Body)

	// If err does not equal nil(zero value) throw an error and exit the process
	if err != nil {
		log.Fatalf("Failed to read the response body %v", err)
	}

	json.Unmarshal([]byte(string(body)), &ContractData)

	// TokenAddress is assigned the value of the Rinkeby address stored in the ContractData struct
	TokenAddress = ContractData.Addresses.Rinkeby

	// rinkebyUrl is assigned the value of the RINKEBY_URL environment variable
	// rinkebyUrl := os.Getenv("RINKEBY_URL")

	// If err does not equal nil(zero value) throw an error and exit the process
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Creates the Token.sol contract instance
	token, err := NewToken(common.HexToAddress("0x0DCd1Bf9A1b36cE34237eEaFef220932846BCD82"), ProviderConnection())

	// If err does not equal nil(zero value) throw an error and exit the process
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}

	// Returns the Token.sol contract instance
	return token
}
