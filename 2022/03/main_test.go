package main

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var input = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

func TestPart1(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		ans, err := part1(strings.NewReader(input))
		require.NoError(t, err)
		require.Equal(t, 157, ans)
	})
	t.Run("actual", func(t *testing.T) {
		f, err := os.Open("testdata/input.txt")
		require.NoError(t, err)

		ans, err := part1(f)
		require.NoError(t, err)
		require.Equal(t, 7821, ans)
	})
}

func TestPart2(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		ans, err := part2(strings.NewReader(input))
		require.NoError(t, err)
		require.Equal(t, 70, ans)
	})

	t.Run("actual", func(t *testing.T) {
		f, err := os.Open("testdata/input.txt")
		require.NoError(t, err)

		ans, err := part2(f)
		require.NoError(t, err)
		require.Equal(t, 2752, ans)
	})
}
