package main

import (
	"fmt"
	"github.com/mostlygeek/arp"
)

func main() {

	table := arp.Table()

	for ip, mac := range table {
		fmt.Printf("%s = %s\n", ip, mac)
	}

}
