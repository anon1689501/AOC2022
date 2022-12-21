package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Day17() {

	inputText, err := os.Open("input/day17.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer inputText.Close()

	scanner := bufio.NewScanner(inputText)
	scanner.Split(bufio.ScanRunes)

	gasDirection := make([]int, 0)

	for scanner.Scan() {
		if scanner.Text() == "<" {
			gasDirection = append(gasDirection, 1)
		} else {
			gasDirection = append(gasDirection, 2)
		}

	}
	fmt.Println(gasDirection)
}
