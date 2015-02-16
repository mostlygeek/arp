package arp

import (
	"testing"
)

func TestTable(t *testing.T) {

	table := Table()
	if table == nil {
		t.Errorf("Empty table")
	}
}

func TestSearch(t *testing.T) {

	table := Table()

	for ip, test := range table {

		result := Search(ip)
		if test != result {
			t.Errorf("expected %s got %s", test, result)
		}

	}
}

func BenchmarkSearch(b *testing.B) {
	table := Table()
	if len(table) == 0 {
		return
	}

	for ip, _ := range Table() {
		for i := 0; i < b.N; i++ {
			Search(ip)
		}

		// using the first key is enough
		break
	}
}
