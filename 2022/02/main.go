package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"

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

	//p2Ans, err := part2(&buf)
	//if err != nil {
	//	logrus.WithError(err).Fatal("got an error running part2")
	//}
	//logrus.Infof("part 2: %d", p2Ans)
}

// Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock.
// If both players choose the same shape, the round instead ends in a draw.

// A X for rock 1
// B Y for paper 2
// C Z for scissors 3

// 0 for lose
// 3 for draw
// 6 for win

var (
	thing = map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	action = map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}
)

type Ciccio struct {
	TheirPlay string
	MyPlay    string
	Result    int
}

func part1(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	ciccio := []Ciccio{
		{TheirPlay: "A", MyPlay: "X", Result: 3},
		{TheirPlay: "A", MyPlay: "Y", Result: 6},
		{TheirPlay: "A", MyPlay: "Z", Result: 0},
		{TheirPlay: "B", MyPlay: "X", Result: 0},
		{TheirPlay: "B", MyPlay: "Y", Result: 3},
		{TheirPlay: "B", MyPlay: "Z", Result: 6},
		{TheirPlay: "C", MyPlay: "X", Result: 6},
		{TheirPlay: "C", MyPlay: "Y", Result: 0},
		{TheirPlay: "C", MyPlay: "Z", Result: 3},
	}

	var total int
	for scanner.Scan() {
		txt := scanner.Text()
		plays := strings.Split(txt, " ")
		theirPlay := plays[0]
		myPlay := plays[1]

		result := thing[myPlay]
		for _, c := range ciccio {
			if theirPlay == c.TheirPlay && myPlay == c.MyPlay {
				result += c.Result
			}
		}
		total += result
	}

	return total, nil
}

type Ciccio2 struct {
	TheirPlay       string
	ExpectedOutcome int
	Result          int
}

func part2(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	//ciccio := []Ciccio2{
	//	{TheirPlay: "A", ExpectedOutcome: 0, Result: 3},
	//	{TheirPlay: "B", ExpectedOutcome: 3, Result: 6},
	//	{TheirPlay: "C", ExpectedOutcome: 6, Result: 0},
	//}

	var total int
	for scanner.Scan() {
		txt := scanner.Text()
		plays := strings.Split(txt, " ")
		theirPlay := plays[0]
		expectedOutcome := plays[1]

		result := action[expectedOutcome]
		if expectedOutcome == "Y" {
			if theirPlay == "A" {
				result += 1
			}
			if theirPlay == "B" {
				result += 2
			}
			if theirPlay == "C" {
				result += 3
			}
		}

		if expectedOutcome == "X" {
			if theirPlay == "A" {
				result += 3
			}
			if theirPlay == "B" {
				result += 1
			}
			if theirPlay == "C" {
				result += 2
			}
		}

		if expectedOutcome == "Z" {
			if theirPlay == "A" {
				result += 2
			}
			if theirPlay == "B" {
				result += 3
			}
			if theirPlay == "C" {
				result += 1
			}
		}

		//for _, c := range ciccio {
		//
		//
		//	if theirPlay == c.TheirPlay && expectedOutcome == c.MyPlay {
		//		result += c.Result
		//	}
		//}
		total += result
	}

	return total, nil
}

func part1v2(reader io.Reader) (int, error) {
	thing2 := map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
		"X": 0,
		"Y": 1,
		"Z": 2,
	}

	matrix := [][]int{
		0: {4, 8, 3},
		1: {1, 5, 9},
		2: {7, 2, 6},
	}

	total := 0
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		txt := scanner.Text()
		plays := strings.Split(txt, " ")

		theirPlay := thing2[plays[0]]
		ourPlay := thing2[plays[1]]

		total += matrix[theirPlay][ourPlay]
	}

	return total, nil
}
