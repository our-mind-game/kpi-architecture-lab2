package lab2

import (
	"errors"
	"strings"
	"unicode"
)

func isOperator(char string) bool {
	operators := [5]string{"+", "-", "*", "/", "^"}
	for _, operator := range operators {
		if operator == char {
			return true
		}
	}

	return false
}

func pop(slice []string) (string, []string) {
	if len(slice) == 0 {
		return "", slice
	}

	lastElem := slice[len(slice)-1]
	poppedSlice := slice[:len(slice)-1]

	return lastElem, poppedSlice
}

// TODO: document this function.
// PrefixToPostfix converts
func PostfixToPrefix(expression string) (string, error) {
	expression = strings.Trim(expression, " ")
	if len(expression) == 0 {
		return "", errors.New("empty expression")
	}

	var stack []string
	expressionChars := strings.Fields(expression)

	for _, char := range expressionChars {
		runeChar := []rune(char)[0]
		if unicode.IsDigit(runeChar) {
			stack = append(stack, char)
			continue
		}
		if !isOperator(char) {
			return "", errors.New("incorrect character in the expression")
		}

		var secondOperand string
		secondOperand, stack = pop(stack)
		var firstOperand string
		firstOperand, stack = pop(stack)

		if len(secondOperand) < 1 || len(firstOperand) < 1 {
			return "", errors.New("invalid expression")
		}

		prefixExpression := char + " " + firstOperand + " " + secondOperand
		stack = append(stack, prefixExpression)
	}

	if len(stack) != 1 {
		return "", errors.New("invalid expression")
	}

	return stack[0], nil
}
