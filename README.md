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
	fmt.Println("Slack: ", githubmarkdownconvertergo.Slack(markdown, githubmarkdownconvertergo.SlackConvertOptions{
		// Headlines will define if GitHub headlines will be updated to be bold text in slack
		// there is no headlines as sucks in Slack
		Headlines: true,
		// Name of the git repo, used to link pull-requests/issues
		// repo name to be given in format "<owner>/<name>" (example: eritikass/githubmarkdownconvertergo)
		RepoName: "eritikass/githubmarkdownconvertergo",
	})
}

```
