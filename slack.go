package githubmarkdownconvertergo

import (
	"fmt"
	"regexp"
	"strings"
)

// SlackConvertOptions contains options to fine-toon GitHub to Slack markdown convert
type SlackConvertOptions struct {
	// Headlines will define if GitHub headlines will be updated to be bold text in slack
	// there is no headlines as sucks in Slack
	Headlines bool
}

// Slack returns github markdown converted to Slack
func Slack(markdown string, options ...SlackConvertOptions) string {
	var re *regexp.Regexp

	opt := SlackConvertOptions{}
	if len(options) > 0 {
		opt = options[0]
	}

	// TODO: write proper regex
	linkRegex := ".*"

	// bold **TEXT**  -> *TEXT*
	re = regexp.MustCompile(`(?miU)((\*\*).+(\*\*))`)
	markdown = re.ReplaceAllStringFunc(markdown, func(s string) string {
		return strings.ReplaceAll(s, "**", "*")
	})

	// italic same for both _TEXT_

	// strikethrough ~~TEXT~~  -> ~TEXT~
	re = regexp.MustCompile(`(?miU)((~~).+(~~))`)
	markdown = re.ReplaceAllStringFunc(markdown, func(s string) string {
		return strings.ReplaceAll(s, "~~", "~")
	})

	// links [TEXT](link) -> <link|TEXT>
	re = regexp.MustCompile(`(?miU)(\[(?P<name>.+)\]\((?P<link>` + linkRegex + `)\))`)
	markdown = re.ReplaceAllStringFunc(markdown, func(s string) string {
		match := re.FindStringSubmatch(s)
		name := ""
		link := ""
		for i, n := range re.SubexpNames() {
			if i != 0 && n == "name" {
				name = match[i]
			}
			if i != 0 && n == "link" {
				link = match[i]
			}
		}
		if name != "" && link != "" {
			return fmt.Sprintf("<%s|%s>", link, name)
		}
		return s
	})

	// * -> •
	re = regexp.MustCompile(`(?s)([^\n][ ]{1,}\*)`)
	markdown = re.ReplaceAllStringFunc(markdown, func(s string) string {
		re2 := regexp.MustCompile(`^([ ]+)?(\*)`)
		return re2.ReplaceAllString(s, "$1•")
	})
	re = regexp.MustCompile(`(?m)((\n[ ]{0,})(\*))`)
	markdown = re.ReplaceAllString(markdown, "$2•")

	if opt.Headlines {
		re = regexp.MustCompile(`(?m)((^\t?[ ]{0,15}#{1,4}[ ]{1,})(.+))`)
		markdown = re.ReplaceAllString(markdown, "*$3*")
	}

	return markdown
}
