package lab2

import (
	. "gopkg.in/check.v1"
	"io"
	"os"
	"strings"
)

type HandlerSuite struct{}

var _ = Suite(&HandlerSuite{})

func (s *HandlerSuite) TestCorrectOutput(c *C) {
	testCase := Case{
		description: "Should write correct result to output",
		input:       "3 5 * 6 2 - ^",
		output:      "^ * 3 5 - 6 2",
		error:       "",
	}

	output, _ := captureComputeOutput(testCase.input)
	c.Assert(output, Equals, testCase.output, Commentf(testCase.description))
}

func (s *HandlerSuite) TestValidatingInput1(c *C) {
	testCase := Case{
		description: "Should return an error: incorrect character in the expression",
		input:       "3 5 & 6 /",
		output:      "",
		error:       "incorrect character in the expression",
	}

	_, err := captureComputeOutput(testCase.input)
	c.Assert(err, ErrorMatches, testCase.error, Commentf(testCase.description))
}

func (s *HandlerSuite) TestValidatingInput2(c *C) {
	testCase := Case{
		description: "Should return an error: invalid expression",
		input:       "3 5 * 6 2 * 6",
		output:      "",
		error:       "empty expression",
	}

	_, err := captureComputeOutput(testCase.input)
	c.Assert(err, ErrorMatches, testCase.error, Commentf(testCase.description))
}

func captureComputeOutput(input string) (string, error) {
	r, w, _ := os.Pipe()
	handler := ComputeHandler{
		Input:  strings.NewReader(input),
		Output: w,
	}

	err := handler.Compute()
	if err != nil {
		return "", err
	}
	w.Close()
	output, _ := io.ReadAll(r)

	return string(output[:len(output)-1]), nil
}
