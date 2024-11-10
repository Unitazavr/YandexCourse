package main

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

var (
	invalidInput           = errors.New("invalid input")
	devidedByZero          = errors.New("devided by zero")
	invalidBracketsInInput = errors.New("brackets amount is invalid")
)

func ValidateInput(expression string) bool {
	operations := []rune("/*+-")
	openCount, closeCount := 0, 0
	for i, r := range expression {
		if strings.ContainsRune("(", r) {
			openCount++
			continue
		} else if strings.ContainsRune(")", r) {
			closeCount++
			continue
		} else if i == len(expression)-1 && slices.Contains(operations, r) ||
			!slices.Contains(operations, r) && !unicode.IsDigit(r) {
			return false
		}
	}
	for i := 0; i < len(expression)-2; i++ {
		el1, el2 := expression[i], expression[i+1]
		if slices.Contains(operations, rune(el1)) && slices.Contains(operations, rune(el2)) {
			return false
		}
	}
	if openCount != closeCount || strings.EqualFold("", expression) {
		return false
	}
	return true
}

func Calculate(expression []string) (float64, error) {
	operations := []rune("/*-+")
	for _, p := range operations {
		op := string(p)
		for slices.Contains(expression, op) {
			pos := slices.Index(expression, op)
			var result float64
			firstNumber, secondNumber := expression[pos-1], expression[pos+1]
			firstNumberFloat, _ := strconv.ParseFloat(firstNumber, 64)
			secondNumberFloat, _ := strconv.ParseFloat(secondNumber, 64)
			switch op {
			case "/":
				if secondNumber == "0" {
					return 0, devidedByZero
				}
				result = firstNumberFloat / secondNumberFloat
				break
			case "*":
				result = firstNumberFloat * secondNumberFloat
				break
			case "+":
				result = firstNumberFloat + secondNumberFloat
				break
			case "-":
				result = firstNumberFloat - secondNumberFloat
				break
			}
			stringresult := strconv.FormatFloat(result, 'f', -1, 64)
			expression[pos] = stringresult
			expression = slices.Delete(expression, pos-1, pos+2)
			expression = slices.Insert(expression, pos-1, stringresult)
		}
	}

	return strconv.ParseFloat(strings.Join(expression, ""), 64)
}

// BracketsIndex возвращает индекс первой открывающей скобки и закрывающей ее скобки.
//
// Если скобок нет, то возвращает 0, 0, nil
// Если возникает ошибка, то возвращает -1, -1, error
func BracketsIndex(expression string) (int, int, error) {
	counter := -1
	pos1 := strings.Index(expression, "(")
	if pos1 == -1 {
		return 0, 0, nil
	}
	for i := pos1 + 1; i < len(expression); i++ {
		switch string(expression[i]) {
		case "(":
			counter--
			break
		case ")":
			counter++
			break
		}
		if counter == 0 {
			return pos1, i, nil
		}
	}
	return -1, -1, invalidBracketsInInput
}
func RecursiveBracketCalculator(expression string) (string, error) {
	pos1, pos2, err := BracketsIndex(expression)
	if err != nil {
		return "", err
	}
	if pos1 == 0 && pos2 == 0 {
		tokens := tokenize(expression)
		result, err := Calculate(tokens)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%f", result), err
	} else {
		exp, err := RecursiveBracketCalculator(expression[pos1+1 : pos2])
		if err != nil {
			return "", err
		}
		expression = expression[:pos1] + exp + expression[pos2+1:]
		return RecursiveBracketCalculator(expression)
	}
}
func Calc(expression string) (float64, error) {
	if !ValidateInput(expression) {
		return 0, invalidInput
	}
	result, err := RecursiveBracketCalculator(expression)
	if err != nil {
		return 0, err
	}
	result1, err := strconv.ParseFloat(result, 64)
	if err != nil {
		return 0, err
	}
	return result1, nil
}
func tokenize(expression string) []string {
	var tokens []string
	var number strings.Builder
	for _, ch := range expression {
		switch {
		case unicode.IsDigit(ch) || ch == '.':
			number.WriteRune(ch)
		case ch == '+' || ch == '-' || ch == '*' || ch == '/' || ch == '(' || ch == ')':
			if number.Len() > 0 {
				tokens = append(tokens, number.String())
				number.Reset()
			}
			tokens = append(tokens, string(ch))
		}
	}
	if number.Len() > 0 {
		tokens = append(tokens, number.String())
	}
	return tokens
}

// Main здесь создан исключительно для тестов!!!
func main() {
	answer, err := Calc("2:2")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(answer)
}
