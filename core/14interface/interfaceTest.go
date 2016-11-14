package main

import (
	"fmt"
	//"sync"
)

type USB interface {
	Name() string
	Connect()
}

type PhoneConnector struct {
}

func (p *PhoneConnector) Name() string {
	return ""
}

func (p *PhoneConnector) Connect() {

}

func main() {
	fmt.Println()
	//sync.Pool
}
