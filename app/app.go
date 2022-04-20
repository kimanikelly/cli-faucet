package app

import (
	"cli-faucet/contract"
	"cli-faucet/signer"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/miguelmota/go-ethutil"
	"github.com/urfave/cli/v2"
)

func StartApp() {

	// Returns the Token.sol contract instance
	token := contract.ContractInstance()

	// Returns the Symbol of Token.sol
	tokenSymbol, symbolErr := token.Symbol(&bind.CallOpts{})

	// If symbolErr is not nil(zero value) throw an error
	if symbolErr != nil {
		log.Fatalf("Failed to return the Symbol %v", symbolErr)
	}

	// The address of the connected signer account
	signerAddress := signer.Address()

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "FundAmount",
				Usage: "View the amount of Test Tokens that will be transferred to a recipients wallet",
				Action: func(c *cli.Context) error {

					// Returns the amount of Test Tokens a recipient can receive per call
					// Returns the amount as a wei value Ne18
					fetchFundAmount, err := token.FundAmount(&bind.CallOpts{})

					// If err is not nil(zero value) throw an error
					if err != nil {
						log.Fatalf("Failed to return the FundAmount %v", err)
					}

					// Converts the amount from wei to a decimal
					fundAmount := ethutil.ToDecimal(fetchFundAmount, 18)

					// Uses the fundAmount and tokenSymbol variable to format the message
					// The fund amount is: N TT as a string
					fundAmtMessage := fmt.Sprintf("The fund amount is: %v %v ", fundAmount, tokenSymbol)

					// Prints the message to the cli
					fmt.Println(fundAmtMessage)

					return nil
				},
			},
			{
				Name:  "ContractBalance",
				Usage: "View the amount of Test Tokens Token.sol holds in the contract",
				Action: func(c *cli.Context) error {

					// Returns the amount of Test Tokens Token.sol holds in the contract
					// Returns the amount as a wei value Ne18
					fetchContractBalance, err := token.BalanceOf(&bind.CallOpts{}, common.HexToAddress("0x0DCd1Bf9A1b36cE34237eEaFef220932846BCD82"))

					if err != nil {
						log.Fatalf("Failed to return the contract balance %v", err)
					}

					// Converts the amount from wei to a decimal
					contractBalance := ethutil.ToDecimal(fetchContractBalance, 18)

					contractBalMessage := fmt.Sprintf("The contract balance is: %v %v", contractBalance, tokenSymbol)

					fmt.Println(contractBalMessage)

					return nil
				},
			},
			{
				Name:  "BalanceOf",
				Usage: "View the amount of Test Tokens a public address holds in their wallet",
				Action: func(c *cli.Context) error {

					address := c.Args().First()

					fetchBalance, err := token.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address))

					if err != nil {
						log.Fatalf("Failed to return the balance %v %v", fetchBalance, err)
					}

					balance := ethutil.ToDecimal(fetchBalance, 18)

					balanceMessage := fmt.Sprintf("The balance of %v is: %v %v", address, balance, tokenSymbol)

					fmt.Println(balanceMessage)

					return nil
				},
			},
			{
				Name:  "FundAccount",
				Usage: "Transfers the FundAmount from the contract to the connected wallet",
				Action: func(c *cli.Context) error {

					fundAccount, err := token.FundAccount(&bind.TransactOpts{From: common.HexToAddress(signerAddress)})

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
