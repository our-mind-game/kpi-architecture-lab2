package lab2

import (
	"bufio"
	"fmt"
	"io"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	scanner := bufio.NewScanner(ch.Input)
	scanner.Scan()
	expression := scanner.Text()

	res, err := PostfixToPrefix(expression)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(ch.Output, res)
	if err != nil {
		return fmt.Errorf("failed to write output: %v", err)
	}

	return nil
}
