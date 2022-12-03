package main

import (
	"io"
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

func TestDay03Part1(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		ans := testPart1(t, strings.NewReader(input))
		require.Equal(t, 157, ans)
	})
	t.Run("actual", func(t *testing.T) {
		f, err := os.Open("testdata/input.txt")
		require.NoError(t, err)

		ans := testPart1(t, f)
		require.Equal(t, 7821, ans)
	})
}

func testPart1(t testing.TB, reader io.Reader) int {
	ans, err := part1(reader)
	require.NoError(t, err)
	return ans
}

func TestDay03Part2(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		ans := testPart2(t, strings.NewReader(input))
		require.Equal(t, 70, ans)
	})
	t.Run("actual", func(t *testing.T) {
		f, err := os.Open("testdata/input.txt")
		require.NoError(t, err)

		ans := testPart2(t, f)
		require.Equal(t, 2752, ans)
	})
}

func testPart2(t testing.TB, reader io.Reader) int {
	ans, err := part2(reader)
	require.NoError(t, err)
	return ans
}

func BenchmarkDay03Part1(b *testing.B) {
	f, err := os.Open("testdata/input.txt")
	require.NoError(b, err)

	for i := 0; i < b.N; i++ {
		_, _ = f.Seek(0, io.SeekStart)
		ans := testPart1(b, f)
		require.Equal(b, 7821, ans)
	}
}

func BenchmarkDay03Part2(b *testing.B) {
	f, err := os.Open("testdata/input.txt")
	require.NoError(b, err)

	for i := 0; i < b.N; i++ {
		_, _ = f.Seek(0, io.SeekStart)
		ans := testPart2(b, f)
		require.Equal(b, 2752, ans)
	}
}
