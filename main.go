package main

import (

	"fmt"
	"github.com/gothello/go-scanner/port"
)



func main() {

	proxy := "169.159.131.146"

	scan := port.ScanPorts(proxy)
	
	fmt.Println(proxy)
	for _, value := range scan {
		fmt.Printf("%+v\n", value)
	}
}
