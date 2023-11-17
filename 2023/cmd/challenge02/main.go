package main

import (
	"codember/internal/challenge/challenge02"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("Running challenge02...")

	inputFile, err := os.Open("./cmd/challenge02/input")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(inputFile)

	fileContents, err := io.ReadAll(inputFile)
	if err != nil {
		panic(err)
	}

	outputFile, err := os.Create("./cmd/challenge02/output")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(outputFile)

	solution := challenge02.Run(strings.Trim(string(fileContents), "\n"))

	fmt.Println("Solution: ", solution)

	_, err = outputFile.WriteString(solution)
}
