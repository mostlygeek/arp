package arp

import (
	"sync"
)

type ArpTable map[string]string

type cache struct {
	sync.RWMutex
	table ArpTable
}

func (c *cache) Refresh() {
	c.Lock()
	defer c.Unlock()
	c.table = Table()
}

func (c *cache) Search(ip string) string {
	c.RLock()
	defer c.RUnlock()

	mac, ok := c.table[ip]

	if !ok {
		c.RUnlock()
		c.Refresh()
		c.RLock()
		mac = c.table[ip]
	}

	return mac
}

var (
	arpCache = &cache{
		table: make(ArpTable),
	}
)

// Search looks up the MAC address for an IP address
// in the arp table
func Search(ip string) string {
	return arpCache.Search(ip)
}
