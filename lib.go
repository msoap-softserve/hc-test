// Package hchat - The hchat package implements parse chat messages
package hchat

import (
	"encoding/json"
	"regexp"
	"strings"
)

// Link - one link, url and title
type Link struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}

// ParsedMessage - parsed message
type ParsedMessage struct {
	Mentions  []string `json:"mentions,omitempty"`
	Emoticons []string `json:"emoticons,omitempty"`
	Links     []Link   `json:"links,omitempty"`
}

// Parse chat message, returns JSON string
func Parse(msg string) (JSON string, err error) {
	parsedMessage := ParsedMessage{}
	mentions := parseMentions(msg)
	if len(mentions) > 0 {
		parsedMessage.Mentions = mentions
	}

	jsonBytes, err := json.Marshal(parsedMessage)
	return string(jsonBytes), err
}

var mentionsRe = regexp.MustCompile(`@\w+`)

func parseMentions(msg string) []string {
	mentions := mentionsRe.FindAllString(msg, -1)
	for i := range mentions {
		mentions[i] = strings.TrimPrefix(mentions[i], "@")
	}

	return mentions
}
