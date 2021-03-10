package port 

import (

	"fmt"
)

func Run(proxy string) {

	scanned := ScanPorts(proxy)

	for _, portResult := range scanned {
		fmt.Printf("%+v\n", portResult)
	}
}