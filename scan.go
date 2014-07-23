package main

import (
	"encoding/xml"
	"fmt"
	//"github.com/davecgh/go-spew/spew"
	"os"
)

// http://golang.org/pkg/encoding/xml/#example_Unmarshal

type NmapRun struct {
	Scanner  string   `xml:"scanner,attr"`
	Args     string   `xml:"args,attr"`
	Scaninfo Scaninfo `xml:"scaninfo"`
	Host     []Host   `xml:"host"`
}

type Scaninfo struct {
	Type     string `xml:"type,attr"`
	Protocol string `xml:"protocol,attr"`
}

type Host struct {
	Ports   Ports     `xml:"ports"`
	Address []Address `xml:"address"`
}

type Ports struct {
	Port []Port `xml:"port"`
}

type Port struct {
	PortID   int     `xml:"portid,attr"`
	Protocol string  `xml:"protocol,attr"`
	State    State   `xml:"state"`
	Service  Service `xml:"service"`
}

type State struct {
	State string `xml:"state,attr"`
}

type Service struct {
	Product string `xml:"product,attr"`
	Version string `xml:"version,attr"`
}

type Address struct {
	Address string `xml:"addr,attr"`
	Type    string `xml:"addrtype,attr"`
	Vendor  string `xml:"vendor,attr"`
}

func main() {
	file, err := os.Open("out.xml")
	if err != nil {
		return
	}
	n := NmapRun{}
	xml.NewDecoder(file).Decode(&n)
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
