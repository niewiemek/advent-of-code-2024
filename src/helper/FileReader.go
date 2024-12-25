package helper

import (
	"bufio"
	"os"
)

type LineHandler func(line string)

func ReadFile(filename string, handler LineHandler) {
	input, err := os.Open("src/" + filename)
	HandleErr(err)

	defer input.Close()

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		handler(scanner.Text())
	}
}
