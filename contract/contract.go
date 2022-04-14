package contract

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

type AddressList struct {
	Rinkeby string
}

type ContractResponseData struct {
	Addresses AddressList
}

var ContractData ContractResponseData

var TokenAddress string

func ContractInstance() *Token {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	body, err := io.ReadAll(fetchContractData().Body)

	if err != nil {
		log.Fatalf("Failed to read the response body %v", err)
	}

	json.Unmarshal([]byte(string(body)), &ContractData)

	TokenAddress = ContractData.Addresses.Rinkeby

	rinkebyUrl := os.Getenv("RINKEBY_URL")

	conn, err := ethclient.Dial(rinkebyUrl)

	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	token, err := NewToken(common.HexToAddress(TokenAddress), conn)

	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}

	return token
}
