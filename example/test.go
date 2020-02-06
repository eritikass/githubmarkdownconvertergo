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
}
