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

		findItemType := func(c1, c2 string) rune {
			// find the common value which gives us the itemType
			for _, item1 := range c1 {
				for _, item2 := range c2 {
					if item1 == item2 {
						return item1
					}
				}
			}
			return 0
		}
		itemType := findItemType(compartment1, compartment2)

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

		findCommonChar := func(group []string) rune {
			commonChars := map[rune]struct{}{}
			commonCharsIntersection := map[rune]struct{}{}
			for _, v := range group[0] {
				if _, ok := commonChars[v]; ok {
					continue
				}
				commonChars[v] = struct{}{}
			}

			// search in the 2nd group for common chars
			for _, v := range group[1] {
				if _, ok := commonChars[v]; ok {
					commonCharsIntersection[v] = struct{}{}
				}
				continue
			}

			// search commonCharsIntersection in the 3rd group and early return
			for _, v := range group[2] {
				if _, ok := commonCharsIntersection[v]; ok {
					return v
				}
				continue
			}

			return 0
		}
		commonC := findCommonChar(group)
		total += lo.IndexOf(lo.LettersCharset, commonC) + 1
	}

	return total, nil
}
