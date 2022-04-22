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

func ContractInstance() *Token {

	//
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
	// The "Failed to read the response body %v" will only log if there is an error reading the
	// response body returned from fetchContractData().Body
	if err != nil {
		log.Fatalf("Failed to read the response body %v", err)
	}

	json.Unmarshal([]byte(string(body)), &ContractData)

	// TokenAddress is assigned the value of the Rinkeby address stored in the ContractData struct
	TokenAddress = ContractData.Addresses.Rinkeby

	// rinkebyUrl is assigned the value of the RINKEBY_URL environment variable
	// rinkebyUrl := os.Getenv("RINKEBY_URL")

	// If err does not equal nil(zero value) throw an error and exit the process
	// The "Failed to connect to the Ethereum client: %v" will only log if there is an error
	// connecting to the Ethereum provider
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	token, err := NewToken(common.HexToAddress("0x0DCd1Bf9A1b36cE34237eEaFef220932846BCD82"), Connection())

	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}

	return token
}
