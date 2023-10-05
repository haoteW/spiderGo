package parser

import (
	"os"
	"testing"
)

func TestParseMovieInfo(t *testing.T) {
	content, err := os.ReadFile("movieInfo_test.html")
	if err != nil {
		t.Error(err)
	}

	ParseMovieInfo(content, "肖申克的救赎")
}
