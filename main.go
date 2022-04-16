package main

import (
	"cli-faucet/contract"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func main() {

	token := contract.ContractInstance()

	tokenBalance, _ := token.BalanceOf(&bind.CallOpts{}, common.HexToAddress(contract.TokenAddress))

	fmt.Println("Token contract balance is:", tokenBalance)
}
