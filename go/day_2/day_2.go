package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("go/day_2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	intList, err := readIntegers(file)
	noun, verb := calcLoop(42, 0, intList, 19690720)
	fmt.Println("Noun = ", noun, " Verb = ", verb)
}

func runCalc(numbers []int) {
	for i := 0; i < len(numbers); i += 4 {
		if numbers[i] == 99 {
			break
		}
		if i%4 == 0 {
			operation := numbers[i]
			firstItem := numbers[numbers[i+1]]
			secondItem := numbers[numbers[i+2]]
			destination := numbers[i+3]
			numbers[destination] = operationMap(operation, firstItem, secondItem)
		}
	}
}

func calcLoop(noun int, verb int, integerList []int, desiredValue int) (int, int) {
	for x := noun; x >= 0; x = x - 1 {
		for h := verb; h < 99; h++ {
			if modifyCalc(x, h, integerList) == desiredValue {
				fmt.Println("Found!")
				return x, h
			}
		}
	}
	return 0, 0
}

func modifyCalc(first int, second int, integerList []int) int {
	newList := make([]int, len(integerList))
	copy(newList, integerList)
	newList[1] = first
	newList[2] = second
	runCalc(newList)
	return newList[0]
}

func operationMap(op int, a int, b int) int {
	if op == 1 {
		return a + b
	}

	if op == 2 {
		return a * b
	}

	return 0
}

func readIntegers(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	var result []int
	for scanner.Scan() {
		x := strings.Split(scanner.Text(), ",")
		for i := 0; i < len(x); i++ {
			integer, err := strconv.Atoi(x[i])
			if err != nil {
				log.Fatal(err)
			}
			result = append(result, integer)
		}
	}
	return result, scanner.Err()
}
