package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gitlab.com/shokoohi/get-ethereum-transactions/routers"
)

var ctx = context.Background()

// Start app
func main() {
	// Load system's environments
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Can't load .env file:", err)
	}
	// Load routers
	router := routers.LoadRouters()
	// Get listen address and port from system's environments
	var (
		address = os.Getenv("TOTEMFI_LISTEN_ADDRESS")
		port    = os.Getenv("TOTEMFI_LISTEN_PORT")
	)
	if port == "" {
		port = "7898"
	}
	listen := address + ":" + port
	log.Println("Listening at => ", listen)
	log.Fatal(router.Run(listen))
	// contractAbi, err := abi.JSON(strings.NewReader(string(store.StoreABI)))
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
