package calculator

import (
	"fmt"
	"strconv"
	"strings"
)

func parse(args []string) (string, float64, float64, error) {
	if len(args) != 3 {
		return "", 0, 0, fmt.Errorf("error: <operand> <operator> <operand>\n")
	}

	var err error

	a, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return "", 0, 0, fmt.Errorf("error: wrong argument '%v'\n", args[0])
	}

	b, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		return "", 0, 0, fmt.Errorf("error: wrong argument '%v'\n", args[2])
	}

	opt := args[1]
	if !strings.ContainsAny(opt, "+-*/") {
		return "", 0, 0, fmt.Errorf("error: invalid operator '%v'\n", args[1])
	}

	return opt, a, b, nil
}

func calculate(opt string, a, b float64) (float64, error) {
	switch opt {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("error: cannot divide by zero\n")
		}
		return a / b, nil
	}
	return 0, fmt.Errorf("error: wrong operator\n")
}

func Run(args []string) {
	args = args[1:]
	opt, a, b, err := parse(args)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	res, err := calculate(opt, a, b)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	fmt.Printf("Result: %v %v %v = %v\n", a, opt, b, res)
}
