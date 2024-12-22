package calc

import (
	"errors"
	"strconv"
	"strings"
)

func Calc(expression string) (float64, error) {
	expression = strings.ReplaceAll(expression, " ", "")

	if !isValidParentheses(expression) {
		return 0, errors.New("invalid parentheses")
	}

	tokens := tokenize(expression)

	result, err := evaluate(tokens)
	if err != nil {
		return 0, errors.New("invalid parentheses")
	}

	return result, nil
}

func isValidParentheses(expression string) bool {
	stack := 0
	for _, char := range expression {
		switch char {
		case '(':
			stack++
		case ')':
			stack--
			if stack < 0 {
				return false
			}
		}
	}
	return stack == 0
}

type TokenType int

const (
	Number TokenType = iota
	Plus
	Minus
	Multiply
	Divide
	LeftParen
	RightParen
)

type Token struct {
	Type  TokenType
	Value string
}

func tokenize(expression string) []Token {
	var tokens []Token
	var currentNumber []rune

	for _, char := range expression {
		switch char {
		case '+':
			if len(currentNumber) > 0 {
				tokens = append(tokens, Token{Type: Number, Value: string(currentNumber)})
				currentNumber = nil
			}
			tokens = append(tokens, Token{Type: Plus})
		case '-':
			if len(currentNumber) > 0 {
				tokens = append(tokens, Token{Type: Number, Value: string(currentNumber)})
				currentNumber = nil
			}
			tokens = append(tokens, Token{Type: Minus})
		case '*':
			if len(currentNumber) > 0 {
				tokens = append(tokens, Token{Type: Number, Value: string(currentNumber)})
				currentNumber = nil
			}
			tokens = append(tokens, Token{Type: Multiply})
		case '/':
			if len(currentNumber) > 0 {
				tokens = append(tokens, Token{Type: Number, Value: string(currentNumber)})
				currentNumber = nil
			}
			tokens = append(tokens, Token{Type: Divide})
		case '(':
			if len(currentNumber) > 0 {
				tokens = append(tokens, Token{Type: Number, Value: string(currentNumber)})
				currentNumber = nil
			}
			tokens = append(tokens, Token{Type: LeftParen})
		case ')':
			if len(currentNumber) > 0 {
				tokens = append(tokens, Token{Type: Number, Value: string(currentNumber)})
				currentNumber = nil
			}
			tokens = append(tokens, Token{Type: RightParen})
		default:
			currentNumber = append(currentNumber, char)
		}
	}

	if len(currentNumber) > 0 {
		tokens = append(tokens, Token{Type: Number, Value: string(currentNumber)})
	}

	return tokens
}

func evaluate(tokens []Token) (float64, error) {
	var values []float64
	var operators []Token

	for _, token := range tokens {
		switch token.Type {
		case Number:
			value, err := strconv.ParseFloat(token.Value, 64)
			if err != nil {
				return 0, err
			}
			values = append(values, value)
		case Plus, Minus:
			for len(operators) > 0 && (operators[len(operators)-1].Type == Plus || operators[len(operators)-1].Type == Minus || operators[len(operators)-1].Type == Multiply || operators[len(operators)-1].Type == Divide) {
				if len(values) < 2 {
					return 0, errors.New("invalid expression")
				}
				values = append(values[:len(values)-2], applyOperator(operators[len(operators)-1], values[len(values)-2], values[len(values)-1]))
				operators = operators[:len(operators)-1]
			}
			operators = append(operators, token)
		case Multiply, Divide:
			for len(operators) > 0 && (operators[len(operators)-1].Type == Multiply || operators[len(operators)-1].Type == Divide) {
				if len(values) < 2 {
					return 0, errors.New("invalid expression")
				}
				values = append(values[:len(values)-2], applyOperator(operators[len(operators)-1], values[len(values)-2], values[len(values)-1]))
				operators = operators[:len(operators)-1]
			}
			operators = append(operators, token)
		case LeftParen:
			operators = append(operators, token)
		case RightParen:
			for len(operators) > 0 && operators[len(operators)-1].Type != LeftParen {
				if len(values) < 2 {
					return 0, errors.New("invalid expression")
				}
				values = append(values[:len(values)-2], applyOperator(operators[len(operators)-1], values[len(values)-2], values[len(values)-1]))
				operators = operators[:len(operators)-1]
			}
			if len(operators) == 0 || operators[len(operators)-1].Type != LeftParen {
				return 0, errors.New("mismatched parentheses")
			}
			operators = operators[:len(operators)-1]
		}
	}

	for len(operators) > 0 {
		if len(values) < 2 {
			return 0, errors.New("invalid expression")
		}
		values = append(values[:len(values)-2], applyOperator(operators[len(operators)-1], values[len(values)-2], values[len(values)-1]))
		operators = operators[:len(operators)-1]
	}

	if len(values) != 1 {
		return 0, errors.New("invalid expression")
	}

	return values[0], nil
}

func applyOperator(operator Token, a, b float64) float64 {
	switch operator.Type {
	case Plus:
		return a + b
	case Minus:
		return a - b
	case Multiply:
		return a * b
	case Divide:
		return a / b
	default:
		return 0
	}
}
