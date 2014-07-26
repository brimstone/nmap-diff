package main

import (
	"fmt"
	//"github.com/davecgh/go-spew/spew"
	"github.com/brimstone/nmap-diff/nmaprun"
)

// http://golang.org/pkg/encoding/xml/#example_Unmarshal

func main() {
	n, err := nmaprun.NewFromFile("scan.xml")
	if err != nil {
		fmt.Println("got an error")
	}
	for _, host := range n.Host {
		fmt.Printf("Host")
		for _, addr := range host.Address {
			fmt.Printf(" %s", addr.Address)
		}
		fmt.Printf("\n")
		fmt.Printf("Open:")
		for _, port := range host.Ports.Port {
			fmt.Printf(" %d", port.PortID)
		}
		fmt.Printf("\n")
		fmt.Printf("\n")
	}
}
