package main

import (
	"cli-faucet/app"
	"cli-faucet/database"
)

func main() {

	app.StartApp()
	database.Database()

}
