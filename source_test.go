package grename

import (
	"bufio"
	"strings"
	"testing"
)

func TestMakeSourceFromStrings(t *testing.T) {
	var (
		inputSet = []string{"123", "µ", "X X    X", "", "\x00"}
		source   = MakeSourceFromStrings(inputSet...)
		result   []string
	)

	for elem := range source {
		result = append(result, elem)
	}

	// check length
	if len(inputSet) != len(result) {
		t.Errorf("len(inputSet) != len(result): %d!=  %d", len(inputSet), len(result))
	}
	// check contents
	for i, _ := range result {
		if inputSet[i] != result[i] {
			t.Error("Mismatch in position", i)
		}
	}
}

func TestMakeSourceFromScanner(t *testing.T) {
	var (
		inputSet = []string{"123", "µ", "X X    X", "", "\x00", "🙈🙉🙊🙌😸"}
		source   = MakeSourceFromScanner(
			strings.NewReader(
				strings.Join(inputSet, "\n")),
			bufio.ScanLines)
		result []string
	)

	for elem := range source {
		result = append(result, elem)
	}

	// check length
	if len(inputSet) != len(result) {
		t.Errorf("len(inputSet) != len(result): %d!=  %d", len(inputSet), len(result))
	}
	// check contents
	for i, _ := range result {
		if inputSet[i] != result[i] {
			t.Error("Mismatch in position", i)
		}
	}
}
