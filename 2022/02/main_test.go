package main

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var input = `A Y
B X
C Z
`

func TestDay02Part1(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		ans := testPart1(t, strings.NewReader(input))
		require.Equal(t, 15, ans)
	})
	t.Run("actual", func(t *testing.T) {
		f, err := os.Open("testdata/input.txt")
		require.NoError(t, err)
		defer f.Close()

		ans := testPart1(t, f)
		require.Equal(t, 14531, ans)
	})
}

func TestDay02Part2(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		ans := testPart2(t, strings.NewReader(input))
		require.Equal(t, 12, ans)
	})

	t.Run("actual", func(t *testing.T) {
		f, err := os.Open("testdata/input.txt")
		require.NoError(t, err)
		defer f.Close()

		ans := testPart2(t, f)
		require.Equal(t, 11258, ans)
	})
}

func BenchmarkDay02Part1(b *testing.B) {
	f, err := os.Open("testdata/input.txt")
	require.NoError(b, err)

	for i := 0; i < b.N; i++ {
		_, _ = f.Seek(0, io.SeekStart)
		ans := testPart1(b, f)
		require.Equal(b, 14531, ans)
	}
}

func BenchmarkDay02Part2(b *testing.B) {
	f, err := os.Open("testdata/input.txt")
	require.NoError(b, err)

	for i := 0; i < b.N; i++ {
		_, _ = f.Seek(0, io.SeekStart)
		ans := testPart2(b, f)
		require.Equal(b, 11258, ans)
	}
}

func testPart1(t testing.TB, reader io.Reader) int {
	ans, err := part1(reader)
	require.NoError(t, err)
	return ans
}

func testPart2(t testing.TB, reader io.Reader) int {
	ans, err := part2(reader)
	require.NoError(t, err)
	return ans
}
