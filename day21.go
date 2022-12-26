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
out:
	for humn := 0; ; humn++ {
		fmt.Println(humn)
		numbersCopy := make(map[string]int)
		formulasCopy := make(map[string]formula)
		for i, v := range numbers {
			numbersCopy[i] = v
		}
		for i, v := range formulas {
			formulasCopy[i] = v
		}
		numbersCopy["humn"] = humn
		//fmt.Println(numbersCopy)
		//fmt.Println(formulasCopy)
		for {
			for i, v := range formulasCopy {
				val1, ok1 := numbersCopy[v.operand1]
				val2, ok2 := numbersCopy[v.operand2]
				if ok1 && ok2 {
					//fmt.Println(i)

					if i == "root" {
						//fmt.Println("humn found")
						if val1 == val2 {
							fmt.Println(val1, val2)
							fmt.Println(humn)
							os.Exit(3)
						}
						continue out

					}
					switch v.operator {
					case "+":
						numbersCopy[i] = val1 + val2
						delete(formulasCopy, i)
					case "-":
						numbersCopy[i] = val1 - val2
						delete(formulasCopy, i)
					case "*":
						numbersCopy[i] = val1 * val2
						delete(formulasCopy, i)
					case "/":
						numbersCopy[i] = val1 / val2
						delete(formulasCopy, i)
					default:
						fmt.Println("unknown operator", v.operator)
					}

				}
			}
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
