package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Day1() {
	//read input
	inputText, err := os.Open("input/day1.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer inputText.Close()

	scanner := bufio.NewScanner(inputText)

	caloriesPerElf := make([]int, 1)
	elf := 0

	//sum calories per elf

	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			caloriesPerElf = append(caloriesPerElf, 0)
			//fmt.Println(caloriesPerElf[elf])
			elf++
			continue

		}
		calories, _ := strconv.Atoi(scanner.Text())
		caloriesPerElf[elf] += calories

	}

	//Day 1 part 1

	fmt.Println("Day 1 Part 1:", FindMax(caloriesPerElf))

	//Day 1 part 2

	threeFatestElvesSum := 0
	for i := 0; i < 3; i++ {
		threeFatestElvesSum += FindMax(caloriesPerElf)
		caloriesPerElf = Remove(caloriesPerElf, FindMaxIndex(caloriesPerElf))
	}
	fmt.Println("Day 1 Part 2:", threeFatestElvesSum)

}
func FindMax(intSlice []int) (max int) {
	max = 0
	for _, v := range intSlice {
		if v > max {
			max = v
		}
	}
	return max
}
func FindMaxIndex(intSlice []int) (index int) {
	maxSize := 0
	index = 0
	for i, v := range intSlice {
		if v > maxSize {
			maxSize = v
			index = i
		}
	}
	return index
}

func Remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
