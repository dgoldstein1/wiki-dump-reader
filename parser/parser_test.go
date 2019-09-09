package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRun(t *testing.T) {}

func TestParseOutLinks(t *testing.T) {
	type Test struct {
		Name          string
		Text          string
		ExpectedError error
		ExpectedLinks []string
	}
	testTable := []Test{
		Test{
			Name:          "finds correct links",
			Text:          anarchismText,
			ExpectedError: nil,
			ExpectedLinks: []string{},
		},
	}
	for _, test := range testTable {
		assert.NotNil(t, test)
	}
}
