package scalculator

/*
import (
	"strconv"
)

func Calc(expression string) float64 {
	slicedExpression := convert(expression)
	slicedExpression = prioritizeMulDiv(slicedExpression)

	result := execSlice(0.0, slicedExpression, 0, len(slicedExpression)-1)

	return result
}

func prioritizeMulDiv(expression []string) []string {

	for i := 0; i < len(expression); i++ {
		if expression[i] == "*" || expression[i] == "/" {
			dst := make([]string, len(expression[i+2:])-1)
			copy(dst, expression[i+2:])
			expression = append(expression[:i-1], "(", expression[i-1], expression[i], expression[i+1], ")")
			expression = append(expression, dst...)
			i += 1
		}
	}
	return expression
}

func execSlice(result float64, expression []string, bg, end int) float64 {
	if bg > len(expression)-1 {
		return result
	}

	for i := bg; i <= end; i++ {
		switch expression[i] {
		case "+":
			f, err := strconv.ParseFloat(expression[i+1], 64)
			if err != nil {
				//parethesis
			}
			result += f
			i++
		case "-":
			f, err := strconv.ParseFloat(expression[i+1], 64)
			if err != nil {
				//parethesis
			}
			result -= f
			i++
		case "/":
			f, err := strconv.ParseFloat(expression[i+1], 64)
			if err != nil {
				//parethesis
			}
			result /= f
			i++
		case "*":
			f, err := strconv.ParseFloat(expression[i+1], 64)
			if err != nil {
				//parethesis
			}
			result *= f
			i++
		case "(":

		case ")":

		case "":

		default:
			f, _ := strconv.ParseFloat(expression[bg], 64)
			result = f
		}
	}

	return result
}

func convert(exp string) []string {
	sliceOfExpressions := []string{}
	readingNumber := false
	for i := 0; i < len(exp); i++ {
		s := exp[i]
		switch s {
		case '+':
			readingNumber = false
			sliceOfExpressions = append(sliceOfExpressions, "+")
		case '-':
			readingNumber = false
			sliceOfExpressions = append(sliceOfExpressions, "-")
		case '/':
			readingNumber = false
			sliceOfExpressions = append(sliceOfExpressions, "/")
		case '*':
			readingNumber = false
			sliceOfExpressions = append(sliceOfExpressions, "*")
		case '(':
			readingNumber = false
			sliceOfExpressions = append(sliceOfExpressions, "(")
		case ')':
			readingNumber = false
			sliceOfExpressions = append(sliceOfExpressions, ")")
		case ' ':

		default:
			if !readingNumber {
				sliceOfExpressions = append(sliceOfExpressions, "")
				readingNumber = true
			}
			sliceOfExpressions[len(sliceOfExpressions)-1] += string(exp[i])
		}
	}
	return sliceOfExpressions
}
*/
