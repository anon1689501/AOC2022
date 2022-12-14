package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day13() {
	inputText, err := os.Open("input/day13.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer inputText.Close()

	scanner := bufio.NewScanner(inputText)
	//scanner.Split(bufio.ScanRunes)
	lines := make([][]string, 0)

	pairs := 1
	for scanner.Scan() {

		line := scanner.Text()
		if strings.TrimSpace(scanner.Text()) == "" {
			pairs++
			//	fmt.Println("Pair:", pairs)
			continue
		}
		line = strings.ReplaceAll(line, "[", "[,")
		line = strings.ReplaceAll(line, "]", ",]")
		line = strings.ReplaceAll(line, ",,", ",")
		lines = append(lines, strings.Split(line, ","))

		//fmt.Println(len(lineSlice), lineSlice)

		//fmt.Println(line)
		// for i, v := range line {

		// }

	}
	linesB := make([][]string, len(lines))
	copy(linesB, lines)
	linesB = append(linesB, []string{"[", "[", "2", "]", "]"})
	linesB = append(linesB, []string{"[", "[", "6", "]", "]"})

	for i := 0; i < len(linesB)-1; i++ {
		for j := 0; j < len(linesB)-i-1; j++ {
			//fmt.Println(linesB[j], linesB[j+1])
			if !isOrdered(linesB[j], linesB[j+1]) {
				linesB[j], linesB[j+1] = linesB[j+1], linesB[j]
			}
		}
	}
	for _, v := range linesB {
		fmt.Println(v)
	}
	//fmt.Println(linesB)

	twoLocaion := 0
	sixLocation := 0

	for i, v := range linesB {
		if sliceEqual(v, []string{"[", "[", "2", "]", "]"}) {
			twoLocaion = i + 1
			break

		}
	}

	for i, v := range linesB {
		if sliceEqual(v, []string{"[", "[", "6", "]", "]"}) {
			sixLocation = i + 1
			break

		}
	}

	sum := 0
	for i := 0; i < len(lines); i += 2 {
		difference := 0
		ordered := false

		//fmt.Println("Pair:", i/2+1)
		for j := 0; j < len(lines[i]); {

			if j > len(lines[i+1]) { // stay in bounds of second string
				break
			}
			if difference > 0 { //difference found
				break

			}
			//fmt.Println(lines[i])
			if lines[i][j] == lines[i+1][j] { //if ([ or (num) or ]) match between the two rows
				j++ //do nothing and move to next
				//if the TOP is a list and BOTTOM is a number (and the number is alone)
			} else if lines[i][j] == "[" && isNumber(lines[i+1][j]) { //&& !isNumber(lines[i+1][j+1]) {
				lines[i+1] = sliceInsert(lines[i+1], j, "[")
				lines[i+1] = sliceInsert(lines[i+1], j+2, "]")

				//if the BOTTOM one is a list and TOP is a number (and the number is alone)
			} else if lines[i+1][j] == "[" && isNumber(lines[i][j]) { //&& !isNumber(lines[i][j+1]) {
				lines[i] = sliceInsert(lines[i], j, "[")
				lines[i] = sliceInsert(lines[i], j+2, "]")
				//if both are non matching number
			} else if isNumber(lines[i][j]) && isNumber(lines[i+1][j]) {
				topNum, _ := strconv.Atoi(lines[i][j])
				bottomNum, _ := strconv.Atoi(lines[i+1][j])
				if topNum < bottomNum {
					difference = j
					ordered = true

				} else if topNum > bottomNum {
					difference = j
				}

				j++
				//bottom ran out of numbers

				// } else if isNumber(lines[i][j]) && lines[i+1][j] == "[" {
				// 	difference = j
				// 	ordered = true
				// 	j++
			} else if isNumber(lines[i][j]) && !isNumber(lines[i+1][j]) { //bottom ran out of numbers
				difference = j
				j++
				// } else if lines[i][j] == "[" && isNumber(lines[i+1][j]) {
				// 	difference = j
				// 	j++
			} else if !isNumber(lines[i][j]) && isNumber(lines[i+1][j]) { //top ran out of numbers
				difference = j
				ordered = true
				j++

			} else if lines[i][j] == "[" && lines[i+1][j] == "]" {
				difference = j
				j++
			} else if lines[i][j] == "]" && lines[i+1][j] == "[" {
				difference = j
				ordered = true
				j++
			} else {
				j++
				fmt.Println("unknown", lines[i][j], lines[i+1][j])

			}

			//missed rule: if top [ and bot a number ordered is false
			//if top a number and bot  [ ordered is true

		}
		// fmt.Println(lines[i])
		// fmt.Println()

		//display line details

		// fmt.Println(lines[i])
		// fmt.Println(lines[i+1])

		// for i := 0; i < difference*2+1; i++ {
		// 	fmt.Print(" ")
		// }
		// fmt.Println("^")

		//end display line details

		//fmt.Println(difference, ordered)
		// fmt.Println(i/2+1, ":", ordered)
		if ordered {
			sum += i/2 + 1
		}
		//fmt.Println("")

	}
	fmt.Println(sum)
	fmt.Println(twoLocaion, sixLocation)
	fmt.Println(sixLocation * twoLocaion)

}
func sliceInsert(slice []string, index int, value string) []string {
	if len(slice) == index {
		return append(slice, value)
	}
	slice = append(slice[:index+1], slice[index:]...)
	slice[index] = value
	return slice
}

func isNumber(input string) bool {
	if _, err := strconv.Atoi(input); err == nil {
		return true
	}
	return false

}

// 6285 too low
// 6648 to high
func sliceEqual(left []string, right []string) bool {
	if len(left) != len(right) {
		return false
	}
	for i, v := range left {
		if v != right[i] {
			return false
		}
	}
	return true
}
func isOrdered(left []string, right []string) (ordered bool) {
	for j := 0; j < len(left); {
		if j >= len(right) { // stay in bounds of second string
			//fmt.Println("length issue")
			return false
		}
		//fmt.Println(left)
		if left[j] == right[j] { //if ([ or (num) or ]) match between the two rows
			j++ //do nothing and move to next
			//if the TOP is a list and BOTTOM is a number (and the number is alone)
		} else if left[j] == "[" && isNumber(right[j]) { //&& !isNumber(right[j+1]) {
			right = sliceInsert(right, j, "[")
			right = sliceInsert(right, j+2, "]")
			//if the BOTTOM one is a list and TOP is a number (and the number is alone)
		} else if right[j] == "[" && isNumber(left[j]) { //&& !isNumber(left[j+1]) {
			left = sliceInsert(left, j, "[")
			left = sliceInsert(left, j+2, "]")
			//if both are non matching number
		} else if isNumber(left[j]) && isNumber(right[j]) {
			topNum, _ := strconv.Atoi(left[j])
			bottomNum, _ := strconv.Atoi(right[j])
			if topNum < bottomNum {
				return true
			} else if topNum > bottomNum {
				return false
			}
		} else if isNumber(left[j]) && !isNumber(right[j]) { //bottom ran out of numbers
			return false
		} else if !isNumber(left[j]) && isNumber(right[j]) { //top ran out of numbers
			return true
		} else if left[j] == "[" && right[j] == "]" {
			return false
		} else if left[j] == "]" && right[j] == "[" {
			return true
		} else {
			fmt.Println("unknown", left[j], right[j])
			break
		}
	}
	return false

}

//part b 20350 too high
