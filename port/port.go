package port

import (
	"fmt"
	"net"
	"time"
	"sync"
)

var wg *sync.WaitGroup

type PortResult struct {
	Port    int
	State   bool
	Service string
}

type PortRange struct {
	Start int
	End   int
}

func ScanPorts(hostname string) []PortResult{
	var portResult []PortResult
	results := make(chan PortResult)

	for port, service := range common {
		
		wg.Add(1)
		go ScanPort(hostname, port, service, results, wg)
	}

	

	for result := range results {
		portResult = append(portResult, result)
	}
	
	wg.Wait()
		
	return portResult
}

func ScanPort(hostname string, port int, service string, result chan PortResult, wg *sync.WaitGroup) {
	defer wg.Done()	
	address := fmt.Sprintf("%s:%d", hostname, port)
	portResult := PortResult{
		Port: port,
		Service: service,
	}

	_, err := net.DialTimeout("tcp", address, 3 *time.Second)
	if err != nil {
		portResult.State = false
	}

	portResult.State = true

	result <- portResult
}