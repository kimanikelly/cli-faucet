package main

import (
	"cli-faucet/contract"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/miguelmota/go-ethutil"
	"github.com/urfave/cli/v2"
)

func main() {

	token := contract.ContractInstance()

	// Returns the Symbol of Token.sol
	tokenSymbol, symbolErr := token.Symbol(&bind.CallOpts{})

	if symbolErr != nil {
		log.Fatalf("Failed to return the Symbol %v", symbolErr)
	}

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "FundAmount",
				Usage: "View the amount of Test Tokens that will be transferred to a recipients wallet",
				Action: func(c *cli.Context) error {

					// Returns the amount of tokens a recipient can receive per call
					// Returns the amount as a wei value ne18
					fetchFundAmount, err := token.FundAmount(&bind.CallOpts{})

					// Converts the amount from wei to a decimal
					fundAmount := ethutil.ToDecimal(fetchFundAmount, 18)

					if err != nil {
						log.Fatalf("Failed to return the FundAmount %v", err)
					}

					fundAmtMessage := fmt.Sprintf("The fund amount is: %v %v ", fundAmount, tokenSymbol)

					fmt.Println(fundAmtMessage)

					return nil
				},
			},
			{
				Name:  "ContractBalance",
				Usage: "View the amount of Test Tokens Token.sol holds in the contract",
				Action: func(c *cli.Context) error {

					fetchContractBalance, err := token.BalanceOf(&bind.CallOpts{}, common.HexToAddress(contract.TokenAddress))

					contractBalance := ethutil.ToDecimal(fetchContractBalance, 18)

					if err != nil {
						log.Fatalf("Failed to return the contract balance %v", err)
					}

					contractBalStr := fmt.Sprintf("The contract balance is: %v %v", contractBalance, tokenSymbol)

					fmt.Println(contractBalStr)
					return nil
				},
			},
			{
				Name:  "BalanceOf",
				Usage: "View the amount of Test Tokens a public address holds in their wallet",
				Action: func(c *cli.Context) error {

					s := c.Args().First()

					balance, err := token.BalanceOf(&bind.CallOpts{}, common.HexToAddress(s))

					if err != nil {
						log.Fatalf("Failed to return the balance %v %v", balance, err)
					}

					g := fmt.Sprintf("The balance of %v is: %v %v", s, balance, tokenSymbol)

					fmt.Println(g)
					return nil
				},
			},
			{
				Name:  "FundAccount",
				Usage: "Transfers the FundAmount from the contract to the connected wallet",
				Action: func(c *cli.Context) error {

					fundAccount, err := token.FundAccount(&bind.TransactOpts{})

					if err != nil {
						log.Fatalf("Failed to fund the connected wallet %v", err)
					}
					fmt.Println(fundAccount)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}

}
