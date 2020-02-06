package githubmarkdownconvertergo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlack(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(Slack("a"), "a")
}
