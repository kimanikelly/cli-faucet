package app

import (
	"cli-faucet/contract"
	"cli-faucet/database"
	"cli-faucet/signer"
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/miguelmota/go-ethutil"
	"github.com/urfave/cli/v2"
)

// Starts the CLI Faucet
func StartApp() {

	// Returns the Token.sol contract instance
	token := contract.ContractInstance()

	// Returns the Symbol of Token.sol
	tokenSymbol, symbolErr := token.Symbol(&bind.CallOpts{})

	// If symbolErr is not nil(zero value) throw an error
	if symbolErr != nil {
		log.Fatalf("Failed to return the Symbol %v", symbolErr)
	}

	// Returns the address and private key of the connected signer
	signerAddress, privateKey := signer.Account()

	// Returns the number of transactions the connected signer sent overtime
	// The nonce can only be used once per transaction and increments by 1 per transaction
	nonce, nonceErr := contract.ProviderConnection().PendingNonceAt(context.Background(), signerAddress)

	// If nonceErr does not equal nil(zero value) throw an error
	if nonceErr != nil {
		log.Fatal(nonceErr)
	}

	// Estimates the cost needed to perform a transaction
	gasPrice, gasErr := contract.ProviderConnection().SuggestGasPrice(context.Background())

	// If gasErr does not equal nil(zero value) throw an error
	if gasErr != nil {
		log.Fatal(gasErr)
	}

	// Returns the chainID of the connected provider
	chainID, chainIdErr := contract.ProviderConnection().ChainID(context.Background())

	// If chainIdErr does not equal nil(zero value) throw an error
	if chainIdErr != nil {
		log.Fatalf("Failed to return the chainID %v", chainIdErr)
	}

	// Binds the connected signer to the transaction options
	auth, authErr := bind.NewKeyedTransactorWithChainID(privateKey, chainID)

	// If authErr does not equal nil (zero value) throw an error
	if authErr != nil {
		log.Fatalf("Failed to build the NewKeyedTransactorWithChainID %v", authErr)
	}

	// Sets the nonce to send with a transaction
	auth.Nonce = big.NewInt(int64(nonce))

	// Sets the amount of test ETH to send with a transaction
	auth.Value = big.NewInt(0)

	// Sets the gasLimit for a transaction
	auth.GasLimit = uint64(300000)

	// Sets the gas price for a transaction
	auth.GasPrice = gasPrice

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
					// From Ne18 to N.0
					fundAmount := ethutil.ToDecimal(fetchFundAmount, 18)

					// Buils the fund amount message as a string by taking in fundAmount and tokenSymbol as args
					fundAmtMessage := fmt.Sprintf("The fund amount is: %v %v ", fundAmount, tokenSymbol)

					// Prints the fundAmount message
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

					// If err does not equal nil(zero value) throw an error
					if err != nil {
						log.Fatalf("Failed to return the contract balance %v", err)
					}

					// Converts the amount from wei to a decimal
					// From Ne18 to N.0
					contractBalance := ethutil.ToDecimal(fetchContractBalance, 18)

					// Builds the contract balance message as a string by taking in the contractBalance and tokenSymbol as args
					contractBalMessage := fmt.Sprintf("The TT contract balance is: %v %v", contractBalance, tokenSymbol)

					// Prints the contract balance message
					fmt.Println(contractBalMessage)

					return nil
				},
			},
			{
				Name:  "BalanceOf",
				Usage: "View the amount of Test Tokens a public address holds in their wallet",
				Action: func(c *cli.Context) error {

					// Stores the address to user inputs in the cli
					address := c.Args().First()

					// Checks if the address is an empty string
					if address == "" {

						// Throws an error if the address is an empty string
						log.Fatalf("BalanceOf Error: The address cannot be empty")

					} else {

						// Returns the amount of Test Tokens an address holds in their wallet
						// Returns the amount as a wei value Ne18
						fetchBalance, err := token.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address))

						// If err does not equal nil(zero value)
						if err != nil {
							log.Fatalf("Failed to return the balance %v %v", fetchBalance, err)
						}

						// Converts the amount from wei to a decimal
						// From Ne18 to N.0
						testTokenBalance := ethutil.ToDecimal(fetchBalance, 18)

						// Builds the balance message as a string by taking in the address, testTokenBalance, and tokenSymbol as args
						balanceMessage := fmt.Sprintf("The TT balance of %v is: %v %v", address, testTokenBalance, tokenSymbol)

						// Prints the balance message
						fmt.Println(balanceMessage)
					}

					return nil
				},
			},
			{
				Name:  "EthBalanceOf",
				Usage: "View the amount of Rinkeby ETH a public address holds in their wallet",
				Action: func(c *cli.Context) error {

					// Stores the address to user inputs in the cli
					address := c.Args().First()

					// Checks if the address is an empty string
					if address == "" {

						// Throws an error if the address is an empty string
						log.Fatalf("EthBalanceOf Error: The address cannot be empty")

					} else {

						// Returns the amount of testnet ETH an address holds in their wallet
						fetchBalance, err := contract.ProviderConnection().BalanceAt(context.Background(), common.HexToAddress(address), nil)

						// If err does not equal nil(zero value) throw an error
						if err != nil {
							log.Fatalf("Failed to return the balance %v %v", fetchBalance, err)
						}

						// Converts the amount from wei to a decimal
						// From Ne18 to N.0
						ethBalance := ethutil.ToDecimal(fetchBalance, 18)

						// Builds the balance message as a string by taking in address, ethBalance, and "ETH"
						balanceMessage := fmt.Sprintf("The ETH balance of %v is: %v %v", address, ethBalance, "ETH")

						// Prints the balance message
						fmt.Println(balanceMessage)
					}

					return nil
				},
			},
			{
				Name:  "FundAccount",
				Usage: "Transfers the FundAmount from the contract to the connected wallet",
				Action: func(c *cli.Context) error {

					// Transfers the Token.sol FundAmount to the connected signer
					fundAccount, err := token.FundAccount(auth)

					// If err does not equal nil(zero value) throw an error
					if err != nil {
						log.Fatalf("Failed to fund the connected wallet %v", err)
					}

					// Returns the Token.sol FundAmount as a wei value
					fetchAmountFunded, err := token.FundAmount(&bind.CallOpts{})

					// If err does not equal nil(zero value) throw an error
					if err != nil {
						log.Fatalf("Failed to get the fundAmount %v", err)
					}

					// Converts the returned FundAmount from a wei value to an ETH value
					amountFunded := ethutil.ToDecimal(fetchAmountFunded, 18)

					// Builds the string to insert the FundAccount transaction values into the recipients MySQL table
					insertFundAccountStr := fmt.Sprintf(
						"INSERT INTO recipients VALUES ('%v', %v, CURRENT_TIMESTAMP())", signerAddress, amountFunded)

					// Inserts the FundAccount transaction values into the recipients MySQL table
					insertFundAccount, err := database.Connection().Query(insertFundAccountStr)

					// If err does not equal nil(zero value) throw an error
					if err != nil {
						log.Fatalf("Failed to insert the FundAccount transaction values into the recipients MySQL table %v", err)
					}

					// Message logs the connected signers address and how many Test Tokens were transferred
					amountFundedMessage := fmt.Sprintf("%v was funded: %v %v", signerAddress, amountFunded, tokenSymbol)

					// Etherscan Rinkeby transaction hash
					txHash := fmt.Sprintf("https://rinkeby.etherscan.io/tx/%s", fundAccount.Hash())

					// Prints the message to the CLI when the transaction is successful
					fmt.Println(amountFundedMessage)

					// Prints the etherscan hash to the CLI when the transaction is successful
					fmt.Println("\nTransaction Hash:", txHash)

					// Closes the recipients table at the end of the FundAccount CLI command
					defer insertFundAccount.Close()

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
