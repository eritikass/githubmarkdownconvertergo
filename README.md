# GitHub markdown converter

> `go get -u github.com/eritikass/githubmarkdownconvertergo`

```go
package main

import (
	"fmt"

	"github.com/eritikass/githubmarkdownconvertergo"
)

func main() {
	markdown := "random **bold** text"
	fmt.Println("GitHub: ", markdown)

	// convert to Slack markdown
	fmt.Println("Slack: ", githubmarkdownconvertergo.Slack(markdown))

	// optonally githubmarkdownconvertergo.Slack is also accepting 2'nd argument that that can be used to customize converted behavor
	markdown2 := "repo tickets in text, #6, #7"
	fmt.Println("Slack: ", githubmarkdownconvertergo.Slack(markdown2, githubmarkdownconvertergo.SlackConvertOptions{
		// Headlines will define if GitHub headlines will be updated to be bold text in slack
		// there is no headlines as sucks in Slack
		// optional: default false
		Headlines: true,
		// Name of the git repo, used to link pull-requests/issues
		// repo name to be given in format "<owner>/<name>" (example: eritikass/githubmarkdownconvertergo)
		// optional: default not used
		RepoName: "eritikass/githubmarkdownconvertergo",
		// CustomRefPatterns is optional list of replacement patterns to easily link tickets to alternative systems
		// optional: default not used
		CustomRefPatterns: map[string]string{
			// key is patterns that is searched using regex from markdown (try out using https://regex101.com/)
			// any capture groups you define there can be used in value. 
			// NB: only capture groups by name can be accessed, do not try to use via index.
			`JIRA-(?P<ID>\d+)`: "https://test.atlassian.net/browse/JIRA-${ID})",
			`(?P<BOARD>DEVOPS|LEGAL|COPY|PASTA)-(?P<ID>\d+)`: "[${BOARD}-${ID}](https://xxx.atlassian.net/browse/${BOARD}-${ID})",
			`eventum-(?P<ID>\d+)`: "https://eventum.example.com/issue.php?id=${ID}",
			`ticket-(?P<ID>\d+)`:  "<https://example.com/t/${ID}|ticket:${ID}>",
		},
	})
}

```
