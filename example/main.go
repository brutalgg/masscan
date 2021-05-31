package main

import (
	"fmt"

	"github.com/brutalgg/masscan"
)

func main() {
	f := "out.json"
	out, err := masscan.ParseJSON(f)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	fmt.Println(out[0].Port[0].PortId)
}
