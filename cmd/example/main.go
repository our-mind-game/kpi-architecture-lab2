package main

import (
	"errors"
	"flag"
	"fmt"
	lab2 "github.com/our-mind-game/kpi-architecture-lab2"
	"io"
	"os"
	"strings"
)

var (
	inputExpression = flag.String("e", "", "Expression to convert")
	inputFile       = flag.String("f", "", "Input expression from file")
	outputFile      = flag.String("o", "", "File to store the result of converting the expression")
)

func main() {
	flag.Parse()
	checkFlags(*inputExpression, *inputFile)

	var input io.Reader
	var output io.Writer

	if *inputExpression != "" {
		input = strings.NewReader(*inputExpression)
	} else {
		file, err := os.Open(*inputFile)
		exitIfError(err)
		defer file.Close()
		input = file
	}

	if *outputFile != "" {
		file, err := os.Create(*outputFile)
		exitIfError(err)
		defer file.Close()
		output = file
	} else {
		output = os.Stdout
	}

	handler := &lab2.ComputeHandler{
		Input:  input,
		Output: output,
	}
	err := handler.Compute()
	exitIfError(err)
}

func checkFlags(expressionFlag string, fileFlag string) {
	if expressionFlag == "" && fileFlag == "" {
		err := errors.New("unspecified input")
		exitIfError(err)
	}
	if expressionFlag != "" && fileFlag != "" {
		err := errors.New("both -e and -f flags cannot be specified")
		exitIfError(err)
	}
}

func exitIfError(error error) {
	if error != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", error)
		os.Exit(1)
	}
}
