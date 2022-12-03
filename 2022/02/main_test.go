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

func TestPart1(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		r := strings.NewReader(input)
		ans, err := part1(r)
		require.NoError(t, err)
		require.Equal(t, 15, ans)
	})
	t.Run("actual", func(t *testing.T) {
		f, err := os.Open("testdata/input.txt")
		require.NoError(t, err)
		defer f.Close()

		ans, err := part1(f)
		require.NoError(t, err)
		require.Equal(t, 14531, ans)
	})
	t.Run("actual1v2", func(t *testing.T) {
		f, err := os.Open("testdata/input.txt")
		require.NoError(t, err)
		defer f.Close()

		ans, err := part1v2(f)
		require.NoError(t, err)
		require.Equal(t, 14531, ans)
	})
}

func TestPart2(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		r := strings.NewReader(input)
		ans, err := part2(r)
		require.NoError(t, err)
		require.Equal(t, 12, ans)
	})

	t.Run("actual", func(t *testing.T) {
		f, err := os.Open("testdata/input.txt")
		require.NoError(t, err)
		defer f.Close()

		ans, err := part2(f)
		require.NoError(t, err)
		require.Equal(t, 11258, ans)
	})
}

func BenchmarkPart1(b *testing.B) {
	b.Run("part1", func(b *testing.B) {
		f, err := os.Open("testdata/input.txt")
		require.NoError(b, err)
		defer f.Close()

		for n := 0; n < b.N; n++ {
			_, _ = f.Seek(0, io.SeekStart)

			ans, err := part1(f)
			require.NoError(b, err)
			require.Equal(b, 14531, ans)
		}
	})
	b.Run("part1v2", func(b *testing.B) {
		f, err := os.Open("testdata/input.txt")
		require.NoError(b, err)
		defer f.Close()

		for n := 0; n < b.N; n++ {
			_, _ = f.Seek(0, io.SeekStart)

			ans, err := part1v2(f)
			require.NoError(b, err)
			require.Equal(b, 14531, ans)
		}
	})
}
