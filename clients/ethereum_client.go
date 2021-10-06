package clients

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

var goEther *ethclient.Client

// GetGoEThereum - Return go-ethereum/ethclient client
func GetGoEthereumClient() (*ethclient.Client, error) {
	if goEther == nil {
		lock.Lock()
		defer lock.Unlock()
		if goEther == nil {
			projectID := os.Getenv("INFURA_PROJECT_ID")
			log.Println("Creating new go-ethereum client...")
			goEther, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/" + projectID)
			return goEther, err
		} else {
			log.Println("Returning the exists postgres client")
		}
	} else {
		log.Println("Returning the exists postgres client")
	}
	return goEther, nil
}
