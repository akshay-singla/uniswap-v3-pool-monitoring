package service

import (
	"errors"
	"log"
	"strconv"
)

func (ds *DataService) fetchData() {
	for _, poolAddress := range ds.poolConfig {
		pool, err := NewUniswapV3Pool(poolAddress, ds.client)
		if err != nil {
			log.Printf("Failed to create UniswapV3Pool contract instance: %v", err)
			continue
		}

		poolData, err := pool.FetchData()
		if err != nil {
			log.Printf("Failed to fetch pool data: %v", err)
			continue
		}
		ds.dataStore[poolAddress.Hex()] = append(ds.dataStore[poolAddress.Hex()], poolData)
	}
}

func (ds *DataService) GetPoolData(poolID string, blockNumber string) (*PoolData, error) {
	data, ok := ds.dataStore[poolID]
	if !ok {
		return nil, errors.New("pool not found")
	}

	if blockNumber == "latest" {
		n := len(data)
		return &data[n-1], nil
	}

	targetBlock, err := strconv.ParseUint(blockNumber, 10, 64)
	if err != nil {
		return nil, err
	}

	nearestData := getNearestData(data, targetBlock)

	return &nearestData, nil
}

func (ds *DataService) GetPoolDataList() []PoolData {

	poolData := []PoolData{}

	for _, value := range ds.dataStore {
		n := len(value)
		poolData = append(poolData, value[n-1])
	}

	return poolData
}

func (ds *DataService) GetPoolDataHistoric(poolID string) ([]PoolData, error) {
	data, ok := ds.dataStore[poolID]
	if !ok {
		return nil, errors.New("pool not found")
	}

	return data, nil
}

func getNearestData(data []PoolData, targetBlock uint64) PoolData {
	var nearestData PoolData

	for _, d := range data {
		if d.BlockNumber <= targetBlock {
			nearestData = d
		} else {
			break
		}
	}

	return nearestData
}
