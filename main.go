package main

import (
	"github.com/gothello/go-scanner/port"
)

func main() {
	port.GetOpenPorts("189.52.154.213", port.PortRange{80, 8080})

}
