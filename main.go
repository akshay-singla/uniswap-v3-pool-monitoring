package main

import (
	"log"
	"net/http"
	"os"

	"github.com/akshay-singla/uniswap-v3-pool-monitoring/internal/service"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

func main() {
	poolConfig := []common.Address{
		common.HexToAddress("0x388C818CA8B9251b393131C08a736A67ccB19297"),
		common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"),
	}

	serviceURL := os.Getenv("INFURA_URL")
	if serviceURL == "" {
		log.Fatalf("Failed to fetch service url,please set the ENV for service url")
	}

	dataService, err := service.NewDataService(serviceURL, poolConfig)
	if err != nil {
		log.Fatalf("Failed to initialize DataService: %v", err)
	}

	dataService.StartMonitoring(12)

	router := gin.Default()

	router.GET("/v1/api/pool/:poolID", func(c *gin.Context) {
		poolID := c.Param("poolID")
		blockNumber := c.Query("block")

		if blockNumber == "" {
			poolData, err := dataService.GetPoolData(poolID, "latest")
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, poolData)
		} else if poolID == "" {
			poolDataList := dataService.GetPoolDataList()
			c.JSON(http.StatusOK, poolDataList)
		} else {
			poolData, err := dataService.GetPoolData(poolID, blockNumber)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, poolData)
		}
	})

	router.GET("/v1/api/pool/:poolID/historic", func(c *gin.Context) {
		poolID := c.Param("poolID")
		poolData, err := dataService.GetPoolDataHistoric(poolID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, poolData)
	})

	log.Fatal(router.Run(":8080"))
}
