package contract

import (
	"log"
	"net/http"
)

// Performs a GET request to https://kimanikelly-contractapi.herokuapp.com/tokenContract
func fetchContractData() *http.Response {

	// GET request to access the ABI, Bytecode, and addresses of Token.sol
	resp, err := http.Get("https://kimanikelly-contractapi.herokuapp.com/tokenContract")

	// If err does not equal nil(zero value) throw an err
	if err != nil {
		log.Fatalf("Failed to connect with the /tokenContact endpoint %v", err)
	}

	// Returns the response made to the /tokenContract endpoint
	return resp

}
