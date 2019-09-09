package main

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
		NumberOfLinks int
	}
	testTable := []Test{
		Test{
			Name:          "finds correct links",
			Text:          anarchismText,
			ExpectedError: nil,
			NumberOfLinks: 497,
		},
	}
	for _, test := range testTable {
		e, links := ParseOutLinks(test.Text)
		assert.Equal(t, test.ExpectedError, e)
		assert.Equal(t, test.NumberOfLinks, len(links))
	}
}
