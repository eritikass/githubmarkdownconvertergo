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
	listGithub = `* test1
* test2`
	listSlack = `• test1
• test2`
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

func TestSlackCustomRefPatterns(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("https://xxx.atlassian.net/browse/JIRA-35", Slack("JIRA-35", SlackConvertOptions{
		CustomRefPatterns: map[string]string{
			`(?P<BOARD>JIRA|DEVOPS)-(?P<ID>\d+)`: "https://xxx.atlassian.net/browse/${BOARD}-${ID}",
		},
	}))

	assert.Equal(" https://xxx.atlassian.net/browse/JIRA-1   ", Slack(" JIRA-1   ", SlackConvertOptions{
		CustomRefPatterns: map[string]string{
			`JIRA-(?P<ID>\d+)`: "https://xxx.atlassian.net/browse/JIRA-${ID}",
		},
	}))

	assert.Equal(" [https://xxx.atlassian.net/browse/JIRA-1]   ", Slack(" [JIRA-1]   ", SlackConvertOptions{
		CustomRefPatterns: map[string]string{
			`JIRA-(?P<ID>\d+)`: "https://xxx.atlassian.net/browse/JIRA-${ID}",
		},
	}))

	assert.Equal(" (https://xxx.atlassian.net/browse/JIRA-1)   ", Slack(" (JIRA-1)   ", SlackConvertOptions{
		CustomRefPatterns: map[string]string{
			`JIRA-(?P<ID>\d+)`: "https://xxx.atlassian.net/browse/JIRA-${ID}",
		},
	}))
	assert.Equal(" (https://xxx.atlassian.net/browse/JIRA-1, https://xxx.atlassian.net/browse/JIRA-23)   ", Slack(" (JIRA-1, UPS-23)   ", SlackConvertOptions{
		CustomRefPatterns: map[string]string{
			`(?P<BOARD>JIRA|UPS)-(?P<ID>\d+)`: "https://xxx.atlassian.net/browse/JIRA-${ID}",
		},
	}))

	assert.Equal("https://xxx.atlassian.net/browse/JIRA-356", Slack("JIRA-356", SlackConvertOptions{
		CustomRefPatterns: map[string]string{
			`(?P<BOARD>JIRA|DEVOPS)-(?P<ID>\d{1,10})`: "https://xxx.atlassian.net/browse/${BOARD}-${ID}",
		},
	}))

	assert.Equal("XXX https://xxx.atlassian.net/browse/JIRA-12 UUU", Slack("XXX JIRA-12 UUU", SlackConvertOptions{
		CustomRefPatterns: map[string]string{
			`JIRA-(?P<ID>\d+)`: "https://xxx.atlassian.net/browse/JIRA-${ID}",
		},
	}))

	assert.Equal("XXXJIRA-2YYY", Slack("XXXJIRA-2YYY", SlackConvertOptions{
		CustomRefPatterns: map[string]string{
			`JIRA-(?P<ID>\d+)`: "https://xxx.atlassian.net/browse/JIRA-${ID}",
		},
	}))

	inputLong := `
	- JIRA-666: foo-booo (leg-123)
	- eventum-1335: cat was here (LEGAL-19)
	- ticket:555: lorem ipsum something-something`
	expectedLong := `
	- <https://xxx.atlassian.net/browse/JIRA-666|JIRA-666>: foo-booo (leg-123)
	- <https://eventum.example.com/issue.php?id=1335|eventum-1335>: cat was here (<https://xxx.atlassian.net/browse/LEGAL-19|LEGAL-19>)
	- <https://example.com/t/555|ticket-555>: lorem ipsum something-something`

	assert.Equal(expectedLong, Slack(inputLong, SlackConvertOptions{
		CustomRefPatterns: map[string]string{
			`(?P<BOARD>JIRA|DEVOPS|LEGAL|COPY|PASTA)-(?P<ID>\d+)`: "[${BOARD}-${ID}](https://xxx.atlassian.net/browse/${BOARD}-${ID})",
			`eventum-(?P<ID>\d+)`: "[eventum-${ID}](https://eventum.example.com/issue.php?id=${ID})",
			`ticket:(?P<ID>\d+)`:  "<https://example.com/t/${ID}|ticket-${ID}>",
		},
	}))

	assert.Equal(expectedLong, Slack(inputLong, SlackConvertOptions{
		CustomRefPatterns: map[string]string{
			`(?P<BOARD>[A-Z]{3,10})-(?P<ID>\d{2,5})`: "[${BOARD}-${ID}](https://xxx.atlassian.net/browse/${BOARD}-${ID})",
			`eventum-(?P<ID>\d+)`:                    "[eventum-${ID}](https://eventum.example.com/issue.php?id=${ID})",
			`ticket:(?P<ID>\d+)`:                     "<https://example.com/t/${ID}|ticket-${ID}>",
		},
	}))

	assert.Equal("https://example.com", Slack("example_com", SlackConvertOptions{
		CustomRefPatterns: map[string]string{
			`example_com`: "https://example.com",
		},
	}))
}
