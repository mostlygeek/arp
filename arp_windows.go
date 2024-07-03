// +build windows

package arp

// Windows arp table reader added by Claudio Matsuoka.
// Tested only in Windows 8.1, hopefully the arp command output format
// is the same in other Windows versions.

import (
	"os/exec"
	"strings"
	"syscall"
)

func Table() ArpTable {
	cmd := exec.Command("arp", "-a")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	data, err := cmd.Output()
	if err != nil {
		return nil
	}

	var table = make(ArpTable)
	skipNext := false
	for _, line := range strings.Split(string(data), "\n") {
		// skip empty lines
		if len(line) <= 0 {
			continue
		}
		// skip Interface: lines
		if line[0] != ' ' {
			skipNext = true
			continue
		}
		// skip column headers
		if skipNext {
			skipNext = false
			continue
		}

		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}

		ip := fields[0]
		// Normalize MAC address to colon-separated format
		table[ip] = strings.Replace(fields[1], "-", ":", -1)
	}

	return table
}
