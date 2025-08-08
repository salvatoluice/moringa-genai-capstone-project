package main

import (
	"errors"
	"fmt"
)

// Calculator represents our calculator with basic operations
type Calculator struct{}

// NewCalculator creates a new calculator instance
func NewCalculator() *Calculator {
	return &Calculator{}
}

// Add performs addition of two numbers
func (c *Calculator) Add(a, b float64) float64 {
	return a + b
}

// Subtract performs subtraction of two numbers
func (c *Calculator) Subtract(a, b float64) float64 {
	return a - b
}

// Multiply performs multiplication of two numbers
func (c *Calculator) Multiply(a, b float64) float64 {
	return a * b
}

// Divide performs division of two numbers with error handling for division by zero
func (c *Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

// Power calculates a to the power of b (bonus feature)
func (c *Calculator) Power(a, b float64) float64 {
	result := 1.0
	for i := 0; i < int(b); i++ {
		result *= a
	}
	return result
}

// PerformOperation executes the specified operation
func (c *Calculator) PerformOperation(a, b float64, operation string) (float64, error) {
	switch operation {
	case "+", "add":
		return c.Add(a, b), nil
	case "-", "subtract":
		return c.Subtract(a, b), nil
	case "*", "multiply":
		return c.Multiply(a, b), nil
	case "/", "divide":
		return c.Divide(a, b)
	case "^", "power":
		return c.Power(a, b), nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", operation)
	}
}

// GetSupportedOperations returns a list of supported operations
func (c *Calculator) GetSupportedOperations() []string {
	return []string{"+", "-", "*", "/", "^"}
}