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
		}, {
			"@chris you around?",
			`{"mentions":["chris"]}`,
			nil,
		}, {
			"Good morning! (megusta) (coffee)",
			`{"emoticons":["megusta","coffee"]}`,
			nil,
		}, {
			"Olympics are starting soon; http://www.nbcolympics.com",
			`{"links":[{"url":"http://www.nbcolympics.com","title":"NBC Olympics | 2014 NBC Olympics in Sochi Russia"}]}`,
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
			t.Errorf("%d.\nmessage: %#v\nexpected: %#v\nreal:     %#v", i, item.msg, item.json, json)
		}
	}
}
