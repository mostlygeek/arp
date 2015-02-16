package arp

type ArpTable map[string]string

var (
	cache = make(ArpTable)
)

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
