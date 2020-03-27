package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLines(t *testing.T) {
	input := [][]string{
		{"1+1", "2"},
		{"5+5", "10"},
	}
	problems := parseLines(input)

	for i, p := range problems {
		assert.Equal(t, problems[i].q, p.q, "The question should be the same")
		assert.Equal(t, problems[i].a, p.a, "The answer should be the same")
	}
}
