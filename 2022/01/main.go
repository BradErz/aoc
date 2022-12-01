package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"

	"github.com/samber/lo"

	"github.com/sirupsen/logrus"
)

func main() {
	if err := part1(); err != nil {
		logrus.WithError(err).Fatal("got an error running part1")
	}
	if err := part2(); err != nil {
		logrus.WithError(err).Fatal("got an error running part2")
	}
}

// Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?
func part1() error {
	calories, err := getCalories()
	if err != nil {
		return err
	}
	logrus.Infof("part 1: %+v", lo.Max(calories))
	return nil
}

// Find the top three Elves carrying the most Calories. How many Calories are those Elves carrying in total?
func part2() error {
	calories, err := getCalories()
	if err != nil {
		return err
	}

	// sort the array in descending order
	slices.Sort(calories)

	total := 0
	// extract the last 3 elements of the array, aka top 3 calories and add them together
	for _, i := range calories[len(calories)-3:] {
		total += i
	}
	logrus.Infof("part 2: %d", total)
	return nil
}

func getCalories() ([]int, error) {
	data, err := os.ReadFile("2022/01/testdata/input.txt")
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	var calories []int
	for _, elf := range strings.Split(string(data), "\n\n") {
		total := 0
		for _, val := range strings.Split(elf, "\n") {
			cal, err := strconv.Atoi(val)
			if err != nil {
				continue
			}
			total += cal
		}

		calories = append(calories, total)
	}

	return calories, nil
}
