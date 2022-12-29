package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day21() {

	inputText, err := os.Open("input/day21.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer inputText.Close()

	scanner := bufio.NewScanner(inputText)

	numbers := make(map[string]int)
	formulas := make(map[string]formula)

	for scanner.Scan() {
		line := strings.ReplaceAll(scanner.Text(), ":", "")
		fields := strings.Fields(line)
		if len(fields) == 2 {
			val, _ := strconv.Atoi(fields[1])
			numbers[fields[0]] = val

		} else {
			formulas[fields[0]] = formula{fields[1], fields[2], fields[3]}
		}
	}
	//fmt.Println(numbers)
	//fmt.Println(formulas)
	delete(numbers, "humn")
	formulasCount := len(formulas)

	for {
		for i, v := range formulas {
			val1, ok1 := numbers[v.operand1]
			val2, ok2 := numbers[v.operand2]
			if ok1 && ok2 {
				//fmt.Println(i)
				switch v.operator {
				case "+":
					numbers[i] = val1 + val2
					delete(formulas, i)
				case "-":
					numbers[i] = val1 - val2
					delete(formulas, i)
				case "*":
					numbers[i] = val1 * val2
					delete(formulas, i)
				case "/":
					numbers[i] = val1 / val2
					delete(formulas, i)
				default:
					fmt.Println("unknown operator", v.operator)
				}

			}
		}
		if len(formulas) < formulasCount {
			formulasCount = len(formulas)
			fmt.Println(formulasCount)
		} else {
			break
		}
	}
	fmt.Println(numbers["czdp"])

	nextVal := 0
	nextVar := ""

	rootLeft, rootLeftOk := numbers[formulas["root"].operand1]
	rootRight, rootRightOk := numbers[formulas["root"].operand2]

	if rootLeftOk {
		nextVal = rootLeft
		nextVar = formulas["root"].operand2
	} else if rootRightOk {
		nextVal = rootRight
		nextVar = formulas["root"].operand1
	} else {
		fmt.Println("no good root val found")
	}

	fmt.Println(nextVal, nextVar)

	for {
		val1, ok1 := numbers[formulas[nextVar].operand1]
		val2, ok2 := numbers[formulas[nextVar].operand2]
		if ok1 && ok2 {
			fmt.Println("both vals found for", nextVar)
		}
		if ok1 { //left val has number
			switch formulas[nextVar].operator {
			case "+": //nextVal = val1 + nextVar
				nextVal = nextVal - val1
			case "-": //nextVal = val1 - nextVar
				nextVal = -1 * (nextVal - val1)
			case "*": //nextVal = val1 * nextVar
				nextVal = nextVal / val1
			case "/": //nextVal = val1 / nextVar
				nextVal = val1 / nextVal
			default:
				fmt.Println("unknown operator", formulas[nextVar].operator)
			}
			nextVar = formulas[nextVar].operand2

		} else if ok2 { //right has a number
			switch formulas[nextVar].operator {
			case "+": //nextVal = nextVar + val2
				nextVal = nextVal - val2
			case "-": //nextVal = nextVar - val2
				nextVal = nextVal + val2
			case "*": //nextVal = nextVar * val2
				nextVal = nextVal / val2
			case "/": //nextVal = nextVar / val2
				nextVal = val2 * nextVal
			default:
				fmt.Println("unknown operator", formulas[nextVar].operator)
			}
			nextVar = formulas[nextVar].operand1

		} else {
			fmt.Println("neither val found", nextVar)
		}
		if nextVar == "humn" {
			fmt.Println(nextVal)
			break
		}
		// rootVal, rootOk := numbers["root"]
		// if rootOk {
		// 	fmt.Println(rootVal)
		// 	break
		// }

	}

}

type formula struct {
	operand1 string
	operator string
	operand2 string
}

//part2 215 too low
