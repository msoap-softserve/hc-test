// Package hchat - The hchat package implements parse chat messages
package hchat

import (
	"encoding/json"
	"regexp"
	"strings"

	"github.com/msoap/html2data"
)

// MaxEmoticonsLength - emoticons length
const MaxEmoticonsLength = 15

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

	links, err := parseLinks(msg)
	if err != nil {
		return "", err
	}
	if len(links) > 0 {
		parsedMessage.Links = links
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
	emoticons := []string{}
	for _, emoticon := range emoticonsRe.FindAllString(msg, -1) {
		emoticon = strings.TrimPrefix(strings.TrimSuffix(emoticon, ")"), "(")
		if len(emoticon) <= MaxEmoticonsLength {
			emoticons = append(emoticons, emoticon)
		}
	}

	return emoticons
}

// Links parse
var linksRe = regexp.MustCompile(`https?://[a-zA-Z0-9\-]{1,64}(\.[a-zA-Z0-9\-]{1,64})*(:\d+)?(\S+)?`)

func parseLinks(msg string) (result []Link, err error) {
	for _, link := range linksRe.FindAllString(msg, -1) {
		title, err := html2data.FromURL(link).GetDataSingle("title")
		if err != nil {
			return []Link{}, err
		}
		result = append(result, Link{URL: link, Title: title})
	}

	return result, nil
}
