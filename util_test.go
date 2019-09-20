package arp

import "testing"

func TestPadMacString(t *testing.T) {
	var validInput []string
	var expectedResult []string

	validInput = append(validInput,"0:0:0:0:0:0")
	expectedResult = append(expectedResult, "00:00:00:00:00:00")
	validInput = append(validInput,"00:00:00:00:00:00")
	expectedResult = append(expectedResult, "00:00:00:00:00:00")
	validInput = append(validInput,"a:b:c:D:E:F")
	expectedResult = append(expectedResult, "0a:0b:0c:0D:0E:0F")
	validInput = append(validInput,"0-0-0-0-0-0")
	expectedResult = append(expectedResult, "00-00-00-00-00-00")
	validInput = append(validInput,"00-00-00-00-00-00")
	expectedResult = append(expectedResult, "00-00-00-00-00-00")
	validInput = append(validInput,"a-b-c-D-E-F")
	expectedResult = append(expectedResult, "0a-0b-0c-0D-0E-0F")
	validInput = append(validInput,"0.00.000")
	expectedResult = append(expectedResult, "0000.0000.0000")
	validInput = append(validInput,"0000.0000.0000")
	expectedResult = append(expectedResult, "0000.0000.0000")
	validInput = append(validInput,"0.0a.abc")
	expectedResult = append(expectedResult, "0000.000a.0abc")

	for i, _ := range validInput {
		result := padMacString(validInput[i])
		if result != expectedResult[i] {
			t.Fatalf("expectedResult %s, got %s", expectedResult[i], result)
		}
	}

	var invalidInput []string

	invalidInput = append(invalidInput, "0000.0000.00:00")
	invalidInput = append(invalidInput, "0000.0000.00-00")
	invalidInput = append(invalidInput, "0000.0000:0000")
	invalidInput = append(invalidInput, "0000.0000-0000")
	invalidInput = append(invalidInput, "0000.0000.000g")
	invalidInput = append(invalidInput, "00:00:00:00:00.00")
	invalidInput = append(invalidInput, "00:00:00:00:00-00")
	invalidInput = append(invalidInput, "00:00:00:00:00:0g")
	invalidInput = append(invalidInput, "00-00-00-00-00:00")
	invalidInput = append(invalidInput, "00-00-00-00-00.00")
	invalidInput = append(invalidInput, "00-00-00-00-00-0g")

	for i, _ := range invalidInput {
		result := padMacString(invalidInput[i])
		if result != invalidInput[i] {
			t.Fatalf("invalid input %s should have been unmodified but got %s",
				invalidInput[i], result)
		}
	}
}