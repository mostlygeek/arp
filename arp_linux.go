// +build linux

package arp

import (
	"bufio"
	"os"
	"strings"
)

const (
	f_IPAddr int = iota
	f_HWType
	f_Flags
	f_HWAddr
	f_Mask
	f_Device
)

var (
	cache = make(ArpTable)
)

type ArpTable map[string]string

func Table() ArpTable {
	f, err := os.Open("/proc/net/arp")

	if err != nil {
		return nil
	}

	defer f.Close()

	s := bufio.NewScanner(f)

	s.Scan() // skip the field descriptions

	var table = make(ArpTable)

	for s.Scan() {
		line := s.Text()
		fields := strings.Fields(line)
		table[fields[f_IPAddr]] = fields[f_HWAddr]
	}

	return table
}

// Search looks up the MAC address for an IP address
// in the arp table
func Search(ip string) string {

	mac, ok := cache[ip]
	if !ok {
		cache = Table()  // refresh the cache
		return cache[ip] // hope that it's there
	}

	return mac
}
