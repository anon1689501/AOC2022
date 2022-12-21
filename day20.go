package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Day20() {

	inputText, err := os.Open("input/day20.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer inputText.Close()

	scanner := bufio.NewScanner(inputText)
	numberMix := 10
	key := 811589153
	numbers := make([]num, 0)
	position := 0

	for scanner.Scan() {
		lineNumber, _ := strconv.Atoi(scanner.Text())

		numbers = append(numbers, num{position, lineNumber * key})
		position++

	}
	//fmt.Println(numbers)
	//count := 0
	printNum(numbers)
	//fmt.Println(len(numbers))

	for count := 0; count < numberMix; count++ {
		for i := 0; i < len(numbers); i++ {
			//fmt.Println("i", i)
			for j, v := range numbers {
				//fmt.Println(v.op)
				if v.op == i {
					numbers = numMove(numbers, j, calcPos(len(numbers), j, v.value))
					//printNum(numbers)
					break
				}
			}
		}
		fmt.Println(count + 1)
		printNum(numbers)
	}

	start := 0

	for i, v := range numbers {
		if v.value == 0 {
			start = i
			break
		}
	}

	//grove := []int{1000, 2000, 3000}

	a := numbers[(1000+start)%len(numbers)].value
	b := numbers[(2000+start)%len(numbers)].value
	c := numbers[(3000+start)%len(numbers)].value

	fmt.Println(a, b, c, a+b+c)

	//itemToMove = last move -1
}

func calcPos(len int, fromPos int, amount int) int {
	if amount == 0 {
		//fmt.Println("zero found at", fromPos)
		return fromPos

		// } else if amount < 0 { //negative case so moving left
		// 	if Abs(amount) > len { //if its a large number use % to account for times it would wrap around
		// 		amount %= len
		// 	}
		// 	if Abs(amount) > fromPos { //if its position would cause it to go to the back

		// 		pos := len + fromPos + amount - 1
		// 		if pos == 0 {
		// 			fmt.Println("0moving", amount, "to", len-1)
		// 			return len - 1
		// 		}
		// 		fmt.Println("1moving", amount, "to", pos)
		// 		return pos
		// 	}

		// } else { //positive input
		// 	if amount+fromPos > len {
		// 		amount %= len
		// 	}
		// if amount+fromPos > len { //casue it to go back to the front

		// 	pos := fromPos + amount - len + 1
		// 	if pos == 0 {
		// 		fmt.Println("0moving", amount, "to", len-1)
		// 		return len - 1
		// 	}
		// 	fmt.Println("2moving", amount, "to", pos)
		// 	return pos
		// }
	}

	offset := (amount + fromPos) % len
	loops := ((amount + fromPos) / len) % len

	pos := offset + loops

	if pos > len {
		pos = pos%len + 1
	}
	if pos == 0 {
		//fmt.Println("0moving", amount, "to", len-1)
		return len - 1
	}
	if pos < 0 {
		pos += len - 1
	}

	//fmt.Println("3moving", amount, "to", fromPos+amount)
	return pos
}

func printNum(numbers []num) {
	for _, v := range numbers {
		fmt.Print(v.value, " ")

	}
	fmt.Println("")
}

func numMove(numbers []num, fromPos int, toPos int) []num {
	moveNum := numbers[fromPos]
	return numInsert(numDel(numbers, fromPos), moveNum, toPos)
}

func numDel(numbers []num, index int) []num {
	return append(numbers[:index], numbers[index+1:]...)
}

func numInsert(numbers []num, value num, index int) []num {
	return append(numbers[:index], append([]num{value}, numbers[index:]...)...)
}

type num struct {
	op    int //original position
	value int
}

//not right -8347

//part b
//-25258546 not right
//1454367762176 not right
//1792 too low
