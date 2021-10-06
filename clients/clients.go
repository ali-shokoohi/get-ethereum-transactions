package clients

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"os"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var lock = &sync.Mutex{}

// GoClientBuilder - Ethereum client interface
type GoClientBuilder interface {
	BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error)
	FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error)
}


// GetEthereumClient - Return a ethereum client as your choice from system's environments
func GetEthereumClient() (GoClientBuilder, error) {
	// Get client type from  the system's environments
	client := os.Getenv("ETHEREUM_CLIENT")
	if client == "go-ethereum" || client == "" {
		return GetGoEthereumClient()
	} else {
		message := fmt.Sprintf("Can't found your ethereum '%s'! You can choice 'go-ethereum'.\n", client)
		err := errors.New(message)
		return nil, err
	}
}
