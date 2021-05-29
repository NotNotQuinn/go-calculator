package calc

import (
	"fmt"
	"os"
	"strconv"
)

type Calculator struct {
	Fp *os.File
}

func NewCalculator(inStream *os.File) *Calculator {
	return &Calculator{Fp: inStream}
}

type IntAndError struct {
	Num float64
	Err error
}

func (c Calculator) ReadInput(results chan<- IntAndError) {
	for {
		var arg1, op, arg2 string

		fmt.Fscanln(c.Fp, &arg1, &op, &arg2)
		num1, err := strconv.ParseFloat(arg1, 64)
		if err != nil {
			results <- IntAndError{0, err}
			continue
		}
		num2, err := strconv.ParseFloat(arg2, 64)
		if err != nil {
			results <- IntAndError{0, err}
			continue
		}
		switch op {
		case "+":
			results <- IntAndError{num1 + num2, nil}
		case "-":
			results <- IntAndError{num1 - num2, nil}
		case "*":
			results <- IntAndError{num1 * num2, nil}
		case "/":
			results <- IntAndError{num1 / num2, nil}
		default:
			results <- IntAndError{0, fmt.Errorf("invalid operator \"%s\"", op)}
		}
	}
}
