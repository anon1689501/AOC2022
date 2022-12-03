package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Day2() {

	//read input

	inputText, err := os.Open("input/day2.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer inputText.Close()

	scanner := bufio.NewScanner(inputText)

	game := ""
	points := 0

	for scanner.Scan() {
		game = scanner.Text()
		if game[0] == 65 && game[2] == 88 || game[0] == 66 && game[2] == 89 || game[0] == 67 && game[2] == 90 {
			points += 3
		}
		if game[0] == 65 && game[2] == 89 || game[0] == 66 && game[2] == 90 || game[0] == 67 && game[2] == 88 {
			points += 6
		}
		points += int(game[2]) - 87

	}

	fmt.Println("Day 2 Part 1:", points)

}

func Day2p2() {
	//read input

	inputText, err := os.Open("input/day2.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer inputText.Close()

	scanner := bufio.NewScanner(inputText)

	game := ""
	points := 0

	for scanner.Scan() {
		game = scanner.Text()
		if game[2] == 89 {
			points += 3
		}
		if game[2] == 90 {
			points += 6
		}

		//rock conditions
		if game[0] == 65 && game[2] == 88 {
			points += 3
		}
		if game[0] == 65 && game[2] == 89 {
			points += 1
		}
		if game[0] == 65 && game[2] == 90 {
			points += 2
		}

		// paper conditions
		if game[0] == 66 && game[2] == 88 {
			points += 1
		}
		if game[0] == 66 && game[2] == 89 {
			points += 2
		}
		if game[0] == 66 && game[2] == 90 {
			points += 3
		}

		// scissor conditions
		if game[0] == 67 && game[2] == 88 {
			points += 2
		}
		if game[0] == 67 && game[2] == 89 {
			points += 3
		}
		if game[0] == 67 && game[2] == 90 {
			points += 1
		}
	}

	fmt.Println("Day 2 Part 2:", points)

}
