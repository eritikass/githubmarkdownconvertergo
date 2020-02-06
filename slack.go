package githubmarkdownconvertergo

import (
	"fmt"
	"regexp"
	"strings"
)

// Slack returns github markdown converted to slack
func Slack(markdown string) string {
	var re *regexp.Regexp

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

	return markdown
}
