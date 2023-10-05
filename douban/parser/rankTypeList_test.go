package parser

import (
	"os"
	"testing"
)

func TestParseTypeList(t *testing.T) {
	contents, err := os.ReadFile("doubanrank_test.html")

	if err != nil {
		panic(err)
	}
	result := ParseTypeList(contents)

	expectedUrl := []string{
		"/typerank?type_name=剧情&type=11&interval_id=100:80&action=",
		"/typerank?type_name=喜剧&type=24&interval_id=100:90&action=",
		"/typerank?type_name=动作&type=5&interval_id=100:90&action=",
	}
	expectedType := []string{
		"剧情", "喜剧", "动作",
	}

	const resultSize = 28
	if len(result.Requests) != resultSize {
		t.Errorf("result Requests should has %d but only has %d", resultSize, len(result.Requests))
	}
	for i, url := range expectedUrl {
		if result.Requests[i].Url != url {
			t.Errorf("export city #%d: %s; but was %s",
				i, url, result.Requests[i].Url,
			)
		}
	}
	if len(result.Items) != resultSize {
		t.Errorf("result Items should has %d but only has %d", resultSize, len(result.Items))
	}
	for i, typeItem := range expectedType {
		if result.Items[i].(string) != typeItem {
			t.Errorf("export city #%d: %s; but was %s",
				i, typeItem, result.Items[i].(string),
			)
		}
	}

}
