package main

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var input = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func BenchmarkPart1(b *testing.B) {
	f, err := os.Open("testdata/input.txt")
	require.NoError(b, err)
	defer f.Close()

	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, io.SeekStart)

		ans, err := part1(f)
		require.NoError(b, err)
		require.Equal(b, 66306, ans)
	}
}

func BenchmarkPart2(b *testing.B) {
	f, err := os.Open("testdata/input.txt")
	require.NoError(b, err)
	defer f.Close()

	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, io.SeekStart)

		ans, err := part2(f)
		require.NoError(b, err)
		require.Equal(b, 195292, ans)
	}
}

func TestPart1(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		reader := strings.NewReader(input)
		ans, err := part1(reader)
		require.NoError(t, err)
		require.Equal(t, 24000, ans)
	})
	t.Run("actual", func(t *testing.T) {
		reader, err := os.Open("testdata/input.txt")
		require.NoError(t, err)
		ans, err := part1(reader)
		require.NoError(t, err)
		require.Equal(t, 66306, ans)
	})
}

func TestPart2(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		reader := strings.NewReader(input)
		ans, err := part2(reader)
		require.NoError(t, err)
		require.Equal(t, 45000, ans)
	})
	t.Run("actual", func(t *testing.T) {
		reader, err := os.Open("testdata/input.txt")
		require.NoError(t, err)
		ans, err := part2(reader)
		require.NoError(t, err)
		require.Equal(t, 195292, ans)
	})
}
