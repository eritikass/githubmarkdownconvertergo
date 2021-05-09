package githubmarkdownconvertergo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlack(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("a", Slack("a"))
	assert.Equal("text *bold* more-text", Slack("text **bold** more-text"))
	assert.Equal("text ~strike~ more-text", Slack("text ~~strike~~ more-text"))
	// mix of strike + bold
	assert.Equal("text ~strike~ more*bold*-te*x*t", Slack("text ~~strike~~ more**bold**-te**x**t"))
	// links
	assert.Equal("text <http://www.foo.com|This is link title> more-text", Slack("text [This is link title](http://www.foo.com) more-text"))
	assert.Equal("[text] <http://www.foo.com|This is link title> more-text", Slack("[text] [This is link title](http://www.foo.com) more-text"))
	// two links and bold
	assert.Equal("text <http://google.com/|Google> (<https://xxx.com/|x*BB*x>) more-text", Slack("text [Google](http://google.com/) ([x**BB**x](https://xxx.com/)) more-text"))
	// list
	listGithub := `   *
    aaa
				* *
    bbb  * cccc
*
    ddd
      *`
	listSlack := `   •
    aaa
				• *
    bbb  * cccc
•
    ddd
      •`

	assert.Equal(listSlack, Slack(listGithub))

}

func TestSlackHeadlinesOption(t *testing.T) {
	assert := assert.New(t)

	// release message
	msgGithub := `# [1.50.0](https://github.com/foo/boo/compare/v1.49.3...v1.50.0) (2015-02-12)
### Features
 * add GET /v1/events ([#134](https://github.com/foo/boo/issues/134)) ([1726806](https://github.com/foo/boo/commit/1726806))
 * remove DELETE /v1/message ([#121](https://github.com/foo/boo/issues/121)) ([3523r42](https://github.com/foo/boo/commit/3523r42))`
	msgSlack := `# <https://github.com/foo/boo/compare/v1.49.3...v1.50.0|1.50.0> (2015-02-12)
### Features
 • add GET /v1/events (<https://github.com/foo/boo/issues/134|#134>) (<https://github.com/foo/boo/commit/1726806|1726806>)
 • remove DELETE /v1/message (<https://github.com/foo/boo/issues/121|#121>) (<https://github.com/foo/boo/commit/3523r42|3523r42>)`

	assert.Equal(msgSlack, Slack(msgGithub))

	// test headlines parse
	optWithHeadlines := SlackConvertOptions{
		Headlines: true,
	}
	assert.Equal("*fooo*", Slack("### fooo", optWithHeadlines))
	assert.Equal("*Boo foo 123*", Slack(" # Boo foo 123", optWithHeadlines))
	assert.Equal(`
*Features*
`, Slack(`
	### Features
`, optWithHeadlines))
	assert.Equal("*Features*\n\nA feature", Slack("## Features\r\n\r\nA feature", optWithHeadlines))

	msgSlackHeadlinesBold := `*<https://github.com/foo/boo/compare/v1.49.3...v1.50.0|1.50.0> (2015-02-12)*
*Features*
 • add GET /v1/events (<https://github.com/foo/boo/issues/134|#134>) (<https://github.com/foo/boo/commit/1726806|1726806>)
 • remove DELETE /v1/message (<https://github.com/foo/boo/issues/121|#121>) (<https://github.com/foo/boo/commit/3523r42|3523r42>)`

	assert.Equal(msgSlackHeadlinesBold, Slack(msgGithub, optWithHeadlines))
}

func TestSlackRepoNameOption(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("Enhance link regexp <https://github.com/eritikass/githubmarkdownconvertergo/pull/134|#134>", Slack("Enhance link regexp #134", SlackConvertOptions{
		RepoName: "eritikass/githubmarkdownconvertergo",
	}))

	actualInput := `
	• add GET /v1/events (#134)
	• remove DELETE /v1/message, (#121)
	• remove DELETE /v1/message (#121)
	• fix UPDATE /v1/user/meta, #123`
	expected := `
	• add GET /v1/events (<https://github.com/foo-owner/boo-repo/pull/134|#134>)
	• remove DELETE /v1/message, (<https://github.com/foo-owner/boo-repo/pull/121|#121>)
	• remove DELETE /v1/message (<https://github.com/foo-owner/boo-repo/pull/121|#121>)
	• fix UPDATE /v1/user/meta, <https://github.com/foo-owner/boo-repo/pull/123|#123>`
	assert.Equal(expected, Slack(actualInput, SlackConvertOptions{
		RepoName: "foo-owner/boo-repo",
	}))

	assert.Equal("multiple refs, <https://github.com/eritikass/githubmarkdownconvertergo/pull/55|#55>, <https://github.com/eritikass/githubmarkdownconvertergo/pull/56|#56>", Slack("multiple refs, #55, #56", SlackConvertOptions{
		RepoName: "eritikass/githubmarkdownconvertergo",
	}))

	assert.Equal("multiple refs, <https://github.com/eritikass/githubmarkdownconvertergo/pull/55|#55>; <https://github.com/eritikass/githubmarkdownconvertergo/pull/56|#56>", Slack("multiple refs, #55; #56", SlackConvertOptions{
		RepoName: "eritikass/githubmarkdownconvertergo",
	}))

	assert.Equal("multiple refs, <https://github.com/eritikass/githubmarkdownconvertergo/pull/55|#55>, <https://github.com/eritikass/githubmarkdownconvertergo/pull/56|#56>, <https://github.com/eritikass/githubmarkdownconvertergo/pull/22225|#22225> ... and radom text", Slack("multiple refs, #55, #56, #22225 ... and radom text", SlackConvertOptions{
		RepoName: "eritikass/githubmarkdownconvertergo",
	}))

}
