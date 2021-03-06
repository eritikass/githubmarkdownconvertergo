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


  markdown2 := "repo tickets in text, #6, #7"
	// optonally githubmarkdownconvertergo.Slack is also accepting 2'nd argument that that can be used to customize converted behavor
	fmt.Println("Slack: ", githubmarkdownconvertergo.Slack(markdown2, githubmarkdownconvertergo.SlackConvertOptions{
		// Headlines will define if GitHub headlines will be updated to be bold text in slack
		// there is no headlines as sucks in Slack
		// optional: default false
		Headlines: true,
		// Name of the git repo, used to link pull-requests/issues
		// repo name to be given in format "<owner>/<name>" (example: eritikass/githubmarkdownconvertergo)
		// optional: default not used
		RepoName: "eritikass/githubmarkdownconvertergo",
	})
}

```
