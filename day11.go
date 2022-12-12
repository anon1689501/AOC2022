package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day11() {
	inputText, err := os.Open("input/day11.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer inputText.Close()

	currentMonkey := 0

	items := make(map[int][]int)
	var testDivisor [8]int
	var operation [8]string
	var operationValue [8]int
	var falseMonkey [8]int
	var trueMonkey [8]int
	var itemsInspected [8]int
	monkeyCount := 0

	scanner := bufio.NewScanner(inputText)

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		input := strings.Fields(scanner.Text())
		if input[0] == "Monkey" {
			monkeyCount++
			currentMonkey, _ = strconv.Atoi(strings.Trim(input[1], ":"))

		} else if input[0] == "Starting" {
			itemCount := len(input)
			for i := 2; i < itemCount; i++ {
				tempItem, _ := strconv.Atoi(strings.Trim(input[i], ","))
				items[currentMonkey] = append(items[currentMonkey], tempItem)
			}

		} else if input[0] == "Operation:" {
			operation[currentMonkey] = input[4]
			if input[5] == "old" {
				operationValue[currentMonkey] = -1
			} else {
				operationValue[currentMonkey], _ = strconv.Atoi(input[5])
			}

		} else if input[0] == "Test:" {
			testDivisor[currentMonkey], _ = strconv.Atoi(input[3])

		} else if input[0] == "If" {
			if input[1] == "true:" {
				trueMonkey[currentMonkey], _ = strconv.Atoi(input[5])

			} else { //assume false
				falseMonkey[currentMonkey], _ = strconv.Atoi(input[5])
			}
		}
	}

	//process data
	for round := 0; round < 10000; round++ {
		for currentMonkey = 0; currentMonkey < monkeyCount; currentMonkey++ {
			//fmt.Println("Monkey", currentMonkey)
			for _, item := range items[currentMonkey] {
				itemsInspected[currentMonkey]++
				var newValue int
				newValue = 0
				//fmt.Println("inspects item lvl", item)
				if operationValue[currentMonkey] == -1 { //only time old value is used is for mult resulting in a square
					newValue = item * item
				} else if operation[currentMonkey] == "*" {
					newValue = item * operationValue[currentMonkey]
				} else {
					newValue = item + operationValue[currentMonkey]
				}
				//fmt.Println(newValue)
				//newValue /= 3 //uncomment line for part 1
				someVal := 2 * 7 * 11 * 19 * 3 * 5 * 17 * 13 //TO DO change based on input
				newValue %= someVal

				if newValue%testDivisor[currentMonkey] == 0 {
					items[trueMonkey[currentMonkey]] = append(items[trueMonkey[currentMonkey]], newValue)
				} else {
					items[falseMonkey[currentMonkey]] = append(items[falseMonkey[currentMonkey]], newValue)
				}
				items[currentMonkey] = items[currentMonkey][1:]
			}
		}

	}
	fmt.Println(itemsInspected)

}

//20709554856 too high
//880721776 too low
//2257694042 too low
//18170818354
