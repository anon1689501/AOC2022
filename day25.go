package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func Day25() {

	inputText, err := os.Open("input/day25.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer inputText.Close()

	scanner := bufio.NewScanner(inputText)
	numbers := make([]int, 0)
	base := float64(5)

	line := 0
	for scanner.Scan() {
		numbers = append(numbers, 0)
		for i := 0; i < len(scanner.Text()); i++ {
			place := float64(len(scanner.Text()) - i - 1)
			placeVal := int(math.Pow(base, place))
			//fmt.Println("line:", line, "char:", string(scanner.Text()[i]), "place:", place, "place val:", placeVal)
			switch scanner.Text()[i] {
			case '1':
				numbers[line] += 1 * placeVal
			case '2':
				numbers[line] += 2 * placeVal
			case '-':
				numbers[line] += -1 * placeVal
			case '=':
				numbers[line] += -2 * placeVal
			case '0':
				//do nothing
			default:
				fmt.Println("unknown char:", scanner.Text()[i])

			}

		}
		line++
	}

	fmt.Println(numbers)

	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	fmt.Println(sum)

	fmt.Println(makeBaseFive(sum))

}

func makeBaseFive(input int) string {
	if input <= 0 {
		return ""
	}
	nextInput := (input + 2) / 5
	switch input % 5 {
	case 1:
		return makeBaseFive(nextInput) + "1"
	case 2:
		return makeBaseFive(nextInput) + "2"
	case 3:
		return makeBaseFive(nextInput) + "="
	case 4:
		return makeBaseFive(nextInput) + "-"
	case 0:
		return makeBaseFive(nextInput) + "0"
	}

	return ""

}

// func makeBaseFive(input int) string {

// 	base := float64(5)
// 	testPlace := 0
// 	for i := 0; ; i++ {
// 		testPlace = int(math.Pow(base, float64(i)))
// 		if input/int(testPlace) == 0 {
// 			testPlace = int(math.Pow(base, float64(i-1)))
// 			break
// 		}
// 	}
// 	result := input / testPlace
// 	switch result {
// 	case 1:
// 		return "1" + makeBaseFive(input/result)
// 	case 2:
// 		return "2" + makeBaseFive(input/result)
// 	}
// 	return ""
// }
