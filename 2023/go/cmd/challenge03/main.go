package main

import (
	"codember/internal/challenge/challenge03"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("Running challenge03...")

	inputFile, err := os.Open("./cmd/challenge03/input")
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

	outputFile, err := os.Create("./cmd/challenge03/output")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(outputFile)

	solution := challenge03.Run(strings.Split(strings.Trim(string(fileContents), " "), "\n"), 42)

	fmt.Println("Solution:", solution)

	_, err = outputFile.WriteString(solution)
}
