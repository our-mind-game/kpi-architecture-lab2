package lab2

import (
	"fmt"
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type TestSuite struct{}

var _ = Suite(&TestSuite{})

type Case struct {
	description string
	input       string
	output      string
	error       string
}

func (s *TestSuite) TestPostfixToPrefix1(c *C) {
	testCase := Case{
		description: "Should convert expression with +,- operands",
		input:       "7 5 + 1 -",
		output:      "- + 7 5 1",
		error:       "",
	}
	res, err := PostfixToPrefix(testCase.input)
	if err != nil {
		c.Error(err)
		return
	}
	c.Assert(res, Equals, testCase.output, Commentf(testCase.description))
}

func (s *TestSuite) TestPostfixToPrefix2(c *C) {
	testCase := Case{
		description: "Should convert expression with +,- and / operands",
		input:       "12 6 / 4 + 2 -",
		output:      "- + / 12 6 4 2",
		error:       "",
	}
	res, err := PostfixToPrefix(testCase.input)
	if err != nil {
		c.Error(err)
		return
	}
	c.Assert(res, Equals, testCase.output, Commentf(testCase.description))
}

func (s *TestSuite) TestPostfixToPrefix3(c *C) {
	testCase := Case{
		description: "Should convert expression with +,-,* and / operands",
		input:       "9 4 5 7 8 + - / * 7 - 9 + 22 *",
		output:      "* + - * 9 / 4 - 5 + 7 8 7 9 22",
		error:       "",
	}
	res, err := PostfixToPrefix(testCase.input)
	if err != nil {
		c.Error(err)
		return
	}
	c.Assert(res, Equals, testCase.output, Commentf(testCase.description))
}

func (s *TestSuite) TestPostfixToPrefix4(c *C) {
	testCase := Case{
		description: "Should convert expression with all required operands",
		input:       "8 3 4 * ^ 1 2 + - 4 6 9 / + - 8 5 0 2 / ^ + * 30 17 19 - * /",
		output:      "/ * - - ^ 8 * 3 4 + 1 2 + 4 / 6 9 + 8 ^ 5 / 0 2 * 30 - 17 19",
		error:       "",
	}
	res, err := PostfixToPrefix(testCase.input)
	if err != nil {
		c.Error(err)
		return
	}
	c.Assert(res, Equals, testCase.output, Commentf(testCase.description))
}

func (s *TestSuite) TestPostfixToPrefix5(c *C) {
	testCase := Case{
		description: "Should return error empty expression",
		input:       " ",
		output:      "",
		error:       "empty expression",
	}
	_, err := PostfixToPrefix(testCase.input)

	c.Assert(err, ErrorMatches, testCase.error, Commentf(testCase.description))
}

func (s *TestSuite) TestPostfixToPrefix6(c *C) {
	testCase := Case{
		description: "Should return error invalid expression",
		input:       "6 2 1 +",
		output:      "",
		error:       "invalid expression",
	}
	_, err := PostfixToPrefix(testCase.input)

	c.Assert(err, ErrorMatches, testCase.error, Commentf(testCase.description))
}

func (s *TestSuite) TestPostfixToPrefix7(c *C) {
	testCase := Case{
		description: "Should return error invalid expression",
		input:       "7 3 - 5 / 4",
		output:      "",
		error:       "invalid expression",
	}
	_, err := PostfixToPrefix(testCase.input)

	c.Assert(err, ErrorMatches, testCase.error, Commentf(testCase.description))
}

func (s *TestSuite) TestPostfixToPrefix8(c *C) {
	testCase := Case{
		description: "Should return error incorrect character in the expression",
		input:       "4 6 8 * - 4 + 9 ?",
		output:      "",
		error:       "incorrect character in the expression",
	}
	_, err := PostfixToPrefix(testCase.input)

	c.Assert(err, ErrorMatches, testCase.error, Commentf(testCase.description))
}

func (s *TestSuite) TestPostfixToPrefix9(c *C) {
	testCase := Case{
		description: "Should return error incorrect character in the expression",
		input:       "1 3 7 / - 2 + 9 b - *",
		output:      "",
		error:       "incorrect character in the expression",
	}
	_, err := PostfixToPrefix(testCase.input)

	c.Assert(err, ErrorMatches, testCase.error, Commentf(testCase.description))
}

func ExamplePostfixToPrefix() {
	res, _ := PostfixToPrefix("5 8 * 9 ^")
	fmt.Println(res)

	// Output:
	// ^ * 5 8 9
}
