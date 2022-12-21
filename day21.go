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
	fmt.Println(numbers)
	fmt.Println(formulas)

	for {
		for i, v := range formulas {
			val1, ok1 := numbers[v.operand1]
			val2, ok2 := numbers[v.operand2]
			if ok1 && ok2 {
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
		rootVal, rootOk := numbers["root"]
		if rootOk {
			fmt.Println(rootVal)
			break
		}
	}

}

type formula struct {
	operand1 string
	operator string
	operand2 string
}
