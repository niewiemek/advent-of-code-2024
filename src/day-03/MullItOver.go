package main

import (
	"../helper"
	"fmt"
	"regexp"
	"strconv"
)

const (
	FileName                         = "puzzle"
	EnabledInstructionsFilterEnabled = true
)

var (
	mulInstructionFilter     = regexp.MustCompile(`mul\(\d+,\d+\)`)
	enabledInstructionFilter = regexp.MustCompile(`^.*?don't\(\)|do\(\).*?don't\(\)|do\(\).*?$`)
	paramFilter              = regexp.MustCompile(`\d+`)
)

func main() {
	sum := 0
	memoryDump := ""
	handler := func(memoryDumpLine string) {
		memoryDump += memoryDumpLine
	}
	helper.ReadFile("day-03/"+FileName+".txt", handler)

	if EnabledInstructionsFilterEnabled {
		enabledInstructions := enabledInstructionFilter.FindAllString(memoryDump, -1)
		for i, instructions := range enabledInstructions {
			fmt.Printf("\nFound:\n%s\n", instructions)
			sum = execute(instructions, sum)
			fmt.Printf("Processed line %d\n", i)
		}
	} else {
		sum = execute(memoryDump, sum)
	}

	fmt.Printf("\nInstructions result = %d", sum)

}

func execute(enabledInstructions string, sum int) int {
	mulInstructions := mulInstructionFilter.FindAll([]byte(enabledInstructions), -1)

	for _, mul := range mulInstructions {
		params := paramFilter.FindAll(mul, -1)
		val, _ := strconv.Atoi(string(params[0]))
		val2, _ := strconv.Atoi(string(params[1]))

		product := val * val2

		sum += product
	}
	return sum
}
