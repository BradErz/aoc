package main

import (
	"bufio"
	"io"

	"github.com/samber/lo"
)

// Each line is an elf's backpack. The backpack contains two compartments
// split the line in half to get the contents of each compartment
// then find the common char in each compartment assign the priority and sum it
func part1(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var total int
	for scanner.Scan() {
		txt := scanner.Text()
		// split the items into their respective compartments
		compartment1 := txt[:len(txt)/2]
		compartment2 := txt[len(txt)/2:]

		var itemType rune

		// find the common value which gives us the itemType
		for _, item1 := range compartment1 {
			for _, item2 := range compartment2 {
				if item1 == item2 {
					itemType = item1
				}
			}
		}

		// we take the index of the rune inside the priorities and add 1 to get the priority
		total += lo.IndexOf(lo.LettersCharset, itemType) + 1
	}
	return total, nil
}

// Split the input lines into groups of three elves and find the common
// char between them. Assign that char a priority and sum it up.
func part2(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var total int

	groups := []string{}
	for scanner.Scan() {
		groups = append(groups, scanner.Text())
	}

	for _, group := range lo.Chunk(groups, 3) {
		var commonC rune
		// TODO: this works, however my brain is not yet awake enough
		// to solve this in a better way
		for _, elf1C := range group[0] {
			for _, elf2C := range group[1] {
				for _, elf3C := range group[2] {
					//nolint:gocritic // :cry:
					if elf1C == elf2C && elf1C == elf3C {
						commonC = elf1C
					}
				}
			}
		}

		total += lo.IndexOf(lo.LettersCharset, commonC) + 1
	}

	return total, nil
}
