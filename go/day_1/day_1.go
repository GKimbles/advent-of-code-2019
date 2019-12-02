package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("go/day_1/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	totalWeight := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		totalWeight += FuelWeight(i)
		totalWeight += recursiveFuelWeight(FuelWeight(i))
	}

	fmt.Println(totalWeight)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func FuelWeight(weight int) int {
	return (weight / 3) - 2
}

func recursiveFuelWeight(weight int) int {
	fuelFuelWeight := 0
	for FuelWeight(weight) >= 0 {
		fuelFuelWeight += FuelWeight(weight)
		weight = FuelWeight(weight)
		fmt.Println(weight)
	}
	return fuelFuelWeight
}