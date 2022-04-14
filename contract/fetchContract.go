package contract

import (
	"log"
	"net/http"
)

func fetchContractData() *http.Response {

	resp, err := http.Get("https://kimanikelly-contractapi.herokuapp.com/tokenContract")

	if err != nil {
		log.Fatalf("Failed to connect with the /tokenContact endpoint %v", err)
	}

	return resp

}
