package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day5() {

	inputText, err := os.Open("input/day5.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer inputText.Close()

	scanner := bufio.NewScanner(inputText)
	count := 0

	stacks := make(map[int]string)

	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		count++
		if count < 9 {
			chars := []rune(scanner.Text())

			//fmt.Println(string(chars))

			stacks[1] += string(chars[1])
			stacks[2] += string(chars[5])
			stacks[3] += string(chars[9])
			stacks[4] += string(chars[13])
			stacks[5] += string(chars[17])
			stacks[6] += string(chars[21])
			stacks[7] += string(chars[25])
			stacks[8] += string(chars[29])
			stacks[9] += string(chars[33])
			stacks[1] = strings.Trim(stacks[1], " ")
			stacks[2] = strings.Trim(stacks[2], " ")
			stacks[3] = strings.Trim(stacks[3], " ")
			stacks[4] = strings.Trim(stacks[4], " ")
			stacks[5] = strings.Trim(stacks[5], " ")
			stacks[6] = strings.Trim(stacks[6], " ")
			stacks[7] = strings.Trim(stacks[7], " ")
			stacks[8] = strings.Trim(stacks[8], " ")
			stacks[9] = strings.Trim(stacks[9], " ")
		} else if count < 11 {
			continue
		} else {
			command := strings.Fields(scanner.Text())
			number, _ := strconv.Atoi(command[1])
			fromStack, _ := strconv.Atoi(command[3])
			toStack, _ := strconv.Atoi(command[5])

			for i := 0; i < number; i++ {
				stacks[toStack] = string(stacks[fromStack][0]) + stacks[toStack]
				stacks[fromStack] = stacks[fromStack][1:]
			}

		}

	}
	fmt.Println(stacks)

}

func Day5b() {

	inputText, err := os.Open("input/day5.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer inputText.Close()

	scanner := bufio.NewScanner(inputText)
	count := 0

	stacks := make(map[int]string)

	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		count++
		if count < 9 {
			chars := []rune(scanner.Text())

			//fmt.Println(string(chars))

			stacks[1] += string(chars[1])
			stacks[2] += string(chars[5])
			stacks[3] += string(chars[9])
			stacks[4] += string(chars[13])
			stacks[5] += string(chars[17])
			stacks[6] += string(chars[21])
			stacks[7] += string(chars[25])
			stacks[8] += string(chars[29])
			stacks[9] += string(chars[33])
			stacks[1] = strings.Trim(stacks[1], " ")
			stacks[2] = strings.Trim(stacks[2], " ")
			stacks[3] = strings.Trim(stacks[3], " ")
			stacks[4] = strings.Trim(stacks[4], " ")
			stacks[5] = strings.Trim(stacks[5], " ")
			stacks[6] = strings.Trim(stacks[6], " ")
			stacks[7] = strings.Trim(stacks[7], " ")
			stacks[8] = strings.Trim(stacks[8], " ")
			stacks[9] = strings.Trim(stacks[9], " ")
		} else if count < 11 {
			continue
		} else {
			command := strings.Fields(scanner.Text())
			number, _ := strconv.Atoi(command[1])
			fromStack, _ := strconv.Atoi(command[3])
			toStack, _ := strconv.Atoi(command[5])

			stacks[toStack] = string(stacks[fromStack][:number]) + stacks[toStack]
			stacks[fromStack] = stacks[fromStack][number:]

		}
		fmt.Println(stacks)
	}

}
