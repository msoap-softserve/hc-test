package hchat

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParse(t *testing.T) {
	// test http server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<html><title>2016 Rio Olympic Games | NBC Olympics</title><body><div>data</div></body></html>")
	}))

	defer ts.Close()
	URL := ts.URL

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
			"@chris you around? @john42",
			`{"mentions":["chris","john42"]}`,
			nil,
		}, {
			"Good morning! (megusta) (coffee)",
			`{"emoticons":["megusta","coffee"]}`,
			nil,
		}, {
			"Good morning! (megusta) (longlonglonglonglonglong) (coffee)",
			`{"emoticons":["megusta","coffee"]}`,
			nil,
		}, {
			"Olympics are starting soon; " + URL,
			`{"links":[{"url":"` + URL + `","title":"2016 Rio Olympic Games | NBC Olympics"}]}`,
			nil,
		}, { // test cache
			"Another message with same url: " + URL,
			`{"links":[{"url":"` + URL + `","title":"2016 Rio Olympic Games | NBC Olympics"}]}`,
			nil,
		}, {
			"@user! Olympics are starting soon; " + URL,
			`{"mentions":["user"],"links":[{"url":"` + URL + `","title":"2016 Rio Olympic Games | NBC Olympics"}]}`,
			nil,
		}, {
			"@user! Olympics are starting soon (smile); " + URL,
			`{"mentions":["user"],"emoticons":["smile"],"links":[{"url":"` + URL + `","title":"2016 Rio Olympic Games | NBC Olympics"}]}`,
			nil,
		}, {
			"Olympics are starting soon; http://host.fake.example",
			``,
			errors.New("error"),
		},
	}

	chatParser := New()
	for i, item := range testCases {
		json, err := chatParser.Parse(item.msg)

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
