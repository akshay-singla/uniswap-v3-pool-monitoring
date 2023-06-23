package service

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func NewDataService(service string, poolConfig []common.Address) (*DataService, error) {
	client, err := ethclient.Dial(service)
	if err != nil {
		return nil, err
	}

	dataStore := make(map[string][]PoolData)

	return &DataService{
		client:     client,
		poolConfig: poolConfig,
		dataStore:  dataStore,
	}, nil
}
