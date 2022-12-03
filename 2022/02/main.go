package main

import (
	"bufio"
	"io"
	"strings"
)

var (
	inputOptions = map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
		"X": 0,
		"Y": 1,
		"Z": 2,
	}

	// Part 1 matrix
	// A X for rock 1
	// B Y for paper 2
	// C Z for scissors 3

	// 0 for lose
	// 3 for draw
	// 6 for win
	part1Matrix = [][]int{
		// R  P  S
		{4, 8, 3}, // R
		{1, 5, 9}, // P
		{7, 2, 6}, // S
	}

	// Part 2 matrix
	// A for rock 1
	// B for paper 2
	// C for scissors 3
	// X 0 for lose
	// Y 3 for draw
	// Z 6 for win
	part2Matrix = [][]int{
		//  L   D  W
		{3, 4, 8}, // rock
		{1, 5, 9}, // paper
		{2, 6, 7}, // scissors
	}
)

func part1(reader io.Reader) (int, error) {
	return getTotal(reader, part1Matrix), nil
}

func part2(reader io.Reader) (int, error) {
	return getTotal(reader, part2Matrix), nil
}

func getTotal(reader io.Reader, m [][]int) int {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var total int
	for scanner.Scan() {
		txt := scanner.Text()
		plays := strings.Split(txt, " ")

		theirPlay := inputOptions[plays[0]]
		our := inputOptions[plays[1]]

		total += m[theirPlay][our]
	}

	return total
}
