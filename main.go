package main

import (

	"fmt"
	"github.com/gothello/go-scanner/port"
)



func main() {

	proxy := "138.121.113.164"

	scan := port.ScanPorts(proxy)
	
	fmt.Println(proxy)
	for _, value := range scan {
		fmt.Printf("%+v\n", value)
	}
}
