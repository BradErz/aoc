package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strconv"

	"golang.org/x/exp/slices"

	"github.com/samber/lo"

	"github.com/sirupsen/logrus"
)

func main() {
	f, err := os.Open("2022/01/testdata/input.txt")
	if err != nil {
		logrus.WithError(err).Fatal("failed to open file")
	}
	defer f.Close()

	var buf bytes.Buffer
	tee := io.TeeReader(f, &buf)

	p1Ans, err := part1(tee)
	if err != nil {
		logrus.WithError(err).Fatal("got an error running part1")
	}
	logrus.Infof("part 1: %d", p1Ans)

	p2Ans, err := part2(&buf)
	if err != nil {
		logrus.WithError(err).Fatal("got an error running part2")
	}
	logrus.Infof("part 2: %d", p2Ans)
}

// Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?
func part1(reader io.Reader) (int, error) {
	calories := getCalories(reader)
	return lo.Max(calories), nil
}

// Find the top three Elves carrying the most Calories. How many Calories are those Elves carrying in total?
func part2(reader io.Reader) (int, error) {
	calories := getCalories(reader)
	// sort the array in descending order
	slices.Sort(calories)

	total := 0
	// extract the last 3 elements of the array, aka top 3 calories and add them together
	for _, i := range calories[len(calories)-3:] {
		total += i
	}

	return total, nil
}

func getCalories(reader io.Reader) []int {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var (
		total    int
		calories []int
	)
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			// add a new total to the list
			calories = append(calories, total)
			total = 0
			continue
		}

		cal, err := strconv.Atoi(txt)
		if err != nil {
			continue
		}

		total += cal
	}

	// TODO: is there a better way of solving this?
	// basically if its the last line of the scanner
	// we need to append what is currently in the total
	calories = append(calories, total)
	return calories
}
