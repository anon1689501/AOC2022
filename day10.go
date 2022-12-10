package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day10() {
	inputText, err := os.Open("input/day10.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer inputText.Close()

	numbers := make([]int, 0)

	scanner := bufio.NewScanner(inputText)

	for scanner.Scan() {

		input := strings.Fields(scanner.Text())
		if input[0] == "noop" {
			numbers = append(numbers, 0)

		} else {
			addNumber, _ := strconv.Atoi(input[1])
			numbers = append(numbers, 0)
			numbers = append(numbers, addNumber)
		}
	}

	start := 18
	sum := 1
	signalStrength := 0
	for i, val := range numbers {
		sum += val
		if i == start {
			//fmt.Println(sum)
			signalStrength += sum * 20
		}
		if i == start+40 {
			//fmt.Println(sum)
			signalStrength += sum * 60
		}
		if i == start+40*2 {
			//fmt.Println(sum)
			signalStrength += sum * 100
		}
		if i == start+40*3 {
			//fmt.Println(sum)
			signalStrength += sum * 140
		}
		if i == start+40*4 {
			//fmt.Println(sum)
			signalStrength += sum * 180
		}
		if i == start+40*5 {
			//fmt.Println(sum)
			signalStrength += sum * 220
		}

	}
	fmt.Println(signalStrength)
	register := 1
	display := ""
	for i, val := range numbers {

		//fmt.Println(i, register)
		//spite location is at register, register +1, register +2
		//CRT row location is at i
		if (i+1)%40 == register || (i+1)%40 == register+1 || (i+1)%40 == register+2 {
			display += "#"
		} else {
			display += "."
		}
		register += val
	}
	fmt.Println(display[0:40])
	fmt.Println(display[40:80])
	fmt.Println(display[80:120])
	fmt.Println(display[120:160])
	fmt.Println(display[160:200])
	fmt.Println(display[200:240])

}
