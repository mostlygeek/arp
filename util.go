package arp

import (
	"fmt"
	"strings"
)

const numberChars = "ABCDEFabcdef0123456789"

// onlyValidChars returns true if string "test" consists entirely of
// characters from string "expected". If any other characters it returns
// false.
func onlyValidChars(test string, expected string) bool {
	for _, char := range test {
		if !strings.Contains(expected, string(char)) {
			return false
		}
	}
	return true
}

// padMacString takes MAC addresses in string form, pads them to comply with
// the expectations of the standard library's net.ParseMAC(). For example,
// "0:0:c:7:ac:0" becomes "00:00:0c:07:ac:00". If input string cannot be
// understood/padded, then the string is returned without modification.
func padMacString(in string) string {
	var sep string
	var pad string

	switch {
	case strings.Contains(in, ":"):
		sep = ":"
		pad = "%2s"
		if !onlyValidChars(in, numberChars+sep) {
			return in
		}
	case strings.Contains(in, "-"):
		sep = "-"
		pad = "%2s"
		if !onlyValidChars(in, numberChars+sep) {
			return in
		}
	case strings.Contains(in, "."):
		sep = "."
		pad = "%4s"
		if !onlyValidChars(in, numberChars+sep) {
			return in
		}
	default:
		return in
	}

	s := strings.Split(in, sep)
	for i := range s {
		s[i] = strings.Replace(fmt.Sprintf(pad, s[i]), " ", "0", -1)
	}

	return strings.Join(s, sep)
}
