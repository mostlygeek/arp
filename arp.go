package arp

import (
	"time"
)

type ArpTable map[string]string

var (
	stop     = make(chan struct{})
	arpCache = &cache{
		table: make(ArpTable),
	}
)

func StartAutoRefresh(t time.Duration, onUpdate func()) {
	go func() {
		for {
			select {
			case <-time.After(t):
				arpCache.Refresh()
				if onUpdate != nil {
					onUpdate()
				}
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
