package arp

import (
	"time"
	"errors"
)

type ArpTable map[string]string

var (
	stop     = make(chan struct{})
	arpCache = &cache{
		table: make(ArpTable),
	}
)

func AutoRefresh(t time.Duration) {
	go func() {
		for {
			select {
			case <-time.After(t):
				arpCache.Refresh()
			case <-stop:
				return
			}
		}
	}()
}

func StopAutoRefresh() {
	stop <- struct{}{}
}

func CacheUpdate() {
	arpCache.Refresh()
}

func CacheLastUpdate() time.Time {
	return arpCache.Updated
}

func CacheUpdateCount() int {
	return arpCache.UpdatedCount
}

// Search looks up the MAC address for an IP address
// in the arp table
func Search(ip string) string {
	return arpCache.Search(ip)
}

// Search looks up the IP address by MAC address
// in the arp table
func SearchByMac(target string) (string, error) {
	for ip, _ := range Table() {
		m := Search(ip)
		if(m == target) {
			return ip, nil
		}
	}
	return "", errors.New("Mac not found")
}
