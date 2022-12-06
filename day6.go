package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Day6(uniqueLength int) {
	inputText, err := os.Open("input/day6.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer inputText.Close()

	scanner := bufio.NewScanner(inputText)
	scanner.Split(bufio.ScanRunes)

	count := 0
	startMarker := ""

	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		count++
		startMarker = startMarker + scanner.Text()
		if len(startMarker) > uniqueLength {
			startMarker = startMarker[1:]
			if isUnique(startMarker) {
				fmt.Println(count)
				break
			}
		}

	}

}

func isUnique(testString string) (result bool) {
	for i := 0; i < len(testString)-1; i++ {
		for j := i + 1; j < len(testString); j++ {
			if testString[i] == testString[j] {
				return false
			}
		}
	}
	return true

}
