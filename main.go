package main

import (
	"cli-faucet/contract"
	"fmt"
)

func main() {
	token := contract.ContractInstance()
	fmt.Println(token)
}
