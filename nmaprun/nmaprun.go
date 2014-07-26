package nmaprun

import (
	"encoding/xml"
	//"github.com/davecgh/go-spew/spew"
	"os"
	"io"
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

func NewFromFile(path string) (*NmapRun, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil
	}
	return NewFromReader(file)
}

func NewFromReader(r io.Reader) (*NmapRun, error) {
	n := NmapRun{}
	xml.NewDecoder(r).Decode(&n)
	return &n, nil
}
