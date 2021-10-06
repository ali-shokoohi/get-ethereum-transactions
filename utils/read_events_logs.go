package utils

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gitlab.com/shokoohi/get-ethereum-transactions/clients"
)

var client, err = clients.GetEthereumClient()

// ReadEventsLogs - Read contract's event logs
func ReadEventsLogs(ctx context.Context, hex string, fromBlock, toBlock int64) ([]types.Log, error) {
	if err != nil {
		return nil, err
	}
	address := common.HexToAddress(hex)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(fromBlock),
		ToBlock:   big.NewInt(toBlock),
		Addresses: []common.Address{address},
	}
	logs, err := client.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}
	return logs, nil
}
