package port

import (
	"fmt"
	"net"
	"time"
)

type PortResult struct {
	Port    int
	State   bool
	Service string
}

type PortRange struct {
	Start int
	End   int
}

type ScanResult struct {
	hostname string
	ip       []net.IP
	results  []PortResult
}

func GetOpenPorts(hostname string, ports PortRange) {
	scanned, err := ScanPorts(hostname, ports)
	if err != nil {
		fmt.Println(err)
	} else {
		Display(scanned)
	}
}

func Display(result ScanResult) {
	ip := result.ip[len(result.ip)-1]

	fmt.Printf("Open ports for %s (%s)\n", result.hostname, ip.String())

	for _, v := range result.results {
		if v.State {
			fmt.Printf("%d %s\n", v.Port, v.Service)
		}
	}
}

func ScanPorts(hostname string, ports PortRange) (ScanResult, error) {
	var results []PortResult
	var scanned ScanResult
	addr, err := net.LookupIP(hostname)
	if err != nil {
		return scanned, err
	}

	for p, s := range common {
		result := ScanPort(hostname, p)
		result.Service = s

		results = append(results, result)
	}

	scanned = ScanResult{
		hostname: hostname,
		ip:       addr,
		results:  results,
	}

	return scanned, nil
}

func ScanPort(hostname string, port int) PortResult {
	address := fmt.Sprintf("%s:%d", hostname, port)
	portResult := PortResult{
		Port: port,
	}

	conn, err := net.DialTimeout("tcp", address, 5*time.Second)
	if err != nil {
		portResult.State = false
	}

	defer conn.Close()
	portResult.State = true

	return portResult
}
