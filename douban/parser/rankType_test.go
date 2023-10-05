package parser

import (
	"fmt"
	"os"
	"testing"
)

func TestParseRankType(t *testing.T) {
	content, err := os.ReadFile("typerank_test.json")
	if err != nil {
		t.Error(err)
	}

	result := ParseRankType(content)
	fmt.Printf("%v", result)
}
