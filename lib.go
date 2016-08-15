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

	if mentions := parseMentions(msg); len(mentions) > 0 {
		parsedMessage.Mentions = mentions
	}

	if emoticons := parseEmoticons(msg); len(emoticons) > 0 {
		parsedMessage.Emoticons = emoticons
	}

	jsonBytes, err := json.Marshal(parsedMessage)
	return string(jsonBytes), err
}

// Mentions parse
var mentionsRe = regexp.MustCompile(`@\w+`)

func parseMentions(msg string) []string {
	mentions := mentionsRe.FindAllString(msg, -1)
	for i := range mentions {
		mentions[i] = strings.TrimPrefix(mentions[i], "@")
	}

	return mentions
}

// Emoticons parse
var emoticonsRe = regexp.MustCompile(`\(\w+\)`)

func parseEmoticons(msg string) []string {
	emoticons := emoticonsRe.FindAllString(msg, -1)
	for i := range emoticons {
		emoticons[i] = strings.TrimPrefix(strings.TrimSuffix(emoticons[i], ")"), "(")
	}

	return emoticons
}
