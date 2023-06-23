package service

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type PoolData struct {
	Token0Balance string `json:"token0Balance"`
	Token1Balance string `json:"token1Balance"`
	Tick          string `json:"tick"`
	BlockNumber   uint64 `json:"blockNumber"`
}

type DataService struct {
	client     *ethclient.Client
	poolConfig []common.Address
	dataStore  map[string][]PoolData
}

type UniswapV3Pool struct {
	address common.Address
	client  *ethclient.Client
}
