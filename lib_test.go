package hchat

import "testing"

func TestParse(t *testing.T) {
	testCases := []struct {
		msg  string
		json string
		err  error
	}{
		{
			"",
			"{}",
			nil,
		}, {
			"simple message",
			"{}",
			nil,
		},
	}

	for i, item := range testCases {
		json, err := Parse(item.msg)

		if err != nil && item.err == nil {
			t.Errorf("%d. Got error: %s", i, err)
		}
		if err == nil && item.err != nil {
			t.Errorf("%d. Not got error", i)
		}

		if json != item.json {
			t.Errorf("%d. Message: '%s', expected: %#v, real: %#v", i, item.msg, item.json, json)
		}
	}
}
