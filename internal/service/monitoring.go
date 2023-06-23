package service

import "time"

func (ds *DataService) StartMonitoring(interval uint64) {
	go func() {
		for {
			ds.fetchData()
			time.Sleep(time.Duration(interval) * time.Second)
		}
	}()
}
