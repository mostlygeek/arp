// +build linux

package arp

import (
	"bufio"
	"net"
	"os"
	"strings"
)

const (
	IPAddr int = iota
	HWType
	Flags
	HWAddr
	Mask
	Device
)

type ArpTable map[string]net.HardwareAddr

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

		ip := net.ParseIP(fields[IPAddr])
		if ip == nil {
			continue
		}

		mac, err := net.ParseMAC(fields[HWAddr])

		if err != nil {
			continue
		}

		table[fields[IPAddr]] = mac
	}

	return table
}
