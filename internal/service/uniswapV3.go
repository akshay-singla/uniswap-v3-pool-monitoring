package service

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func NewUniswapV3Pool(address common.Address, client *ethclient.Client) (*UniswapV3Pool, error) {
	pool := &UniswapV3Pool{
		address: address,
		client:  client,
	}
	return pool, nil
}

func (pool *UniswapV3Pool) FetchData() (PoolData, error) {

	var data PoolData

	// Fetch the balance of token0 in the pool
	token0Balance, err := pool.fetchToken0Balance()
	if err != nil {
		log.Println(err)
		return data, err
	}
	data.Token0Balance = token0Balance

	// Fetch the balance of token1 in the pool
	token1Balance, err := pool.fetchToken1Balance()
	if err != nil {
		log.Println(err)
		return data, err
	}
	data.Token1Balance = token1Balance

	// Fetch the price tick at block
	tick, err := pool.fetchTick()
	if err != nil {
		log.Println(err)
		return data, err
	}
	data.Tick = tick

	// Fetch the current block number
	blockNumber, err := pool.fetchBlockNumber()
	if err != nil {
		log.Println(err)
		return data, err
	}
	data.BlockNumber = blockNumber

	return data, nil
}

func (pool *UniswapV3Pool) fetchToken0Balance() (string, error) {
	balance, err := pool.client.BalanceAt(context.Background(), pool.address, nil)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return balance.String(), nil
}

func (pool *UniswapV3Pool) fetchToken1Balance() (string, error) {
	balance, err := pool.client.BalanceAt(context.Background(), pool.address, nil)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return balance.String(), nil
}

func (pool *UniswapV3Pool) fetchTick() (string, error) {
	// TODO:
	// Implement the logic to fetch the price tick at block using the pool address and Ethereum client.

	// Placeholder implementation, replace with actual code
	return "500", nil
}

func (pool *UniswapV3Pool) fetchBlockNumber() (uint64, error) {
	header, err := pool.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return header.Number.Uint64(), nil
}
