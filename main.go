package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("🧮 Welcome to Go Calculator!")
	fmt.Println("================================")
	
	calculator := NewCalculator()
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Perform calculation")
		fmt.Println("2. View supported operations")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice (1-3): ")
		
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())
		
		switch choice {
		case "1":
			performCalculation(calculator, scanner)
		case "2":
			showSupportedOperations(calculator)
		case "3":
			fmt.Println("Thank you for using Go Calculator! Goodbye! 👋")
			return
		default:
			fmt.Println("❌ Invalid choice. Please enter 1, 2, or 3.")
		}
	}
}

func performCalculation(calc *Calculator, scanner *bufio.Scanner) {
	// Get first number
	fmt.Print("Enter first number: ")
	scanner.Scan()
	firstInput := strings.TrimSpace(scanner.Text())
	
	num1, err := strconv.ParseFloat(firstInput, 64)
	if err != nil {
		fmt.Println("❌ Invalid number. Please enter a valid number.")
		return
	}
	
	// Get operation
	fmt.Print("Enter operation (+, -, *, /, ^): ")
	scanner.Scan()
	operation := strings.TrimSpace(scanner.Text())
	
	// Get second number
	fmt.Print("Enter second number: ")
	scanner.Scan()
	secondInput := strings.TrimSpace(scanner.Text())
	
	num2, err := strconv.ParseFloat(secondInput, 64)
	if err != nil {
		fmt.Println("❌ Invalid number. Please enter a valid number.")
		return
	}
	
	// Perform calculation
	result, err := calc.PerformOperation(num1, num2, operation)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		return
	}
	
	// Display result
	fmt.Printf("✅ Result: %.2f %s %.2f = %.2f\n", num1, operation, num2, result)
	
	// Ask if user wants to perform another calculation
	fmt.Print("\nWould you like to perform another calculation? (y/n): ")
	scanner.Scan()
	response := strings.ToLower(strings.TrimSpace(scanner.Text()))
	
	if response == "y" || response == "yes" {
		performCalculation(calc, scanner)
	}
}

func showSupportedOperations(calc *Calculator) {
	fmt.Println("\n📋 Supported Operations:")
	fmt.Println("+ or add      : Addition")
	fmt.Println("- or subtract : Subtraction")
	fmt.Println("* or multiply : Multiplication")
	fmt.Println("/ or divide   : Division")
	fmt.Println("^ or power    : Power (exponentiation)")
	fmt.Println("\nNote: Division by zero is not allowed and will return an error.")
}

// Helper function to validate user input (bonus feature)
func validateInput(input string) (float64, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return 0, fmt.Errorf("input cannot be empty")
	}
	
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number format")
	}
	
	return number, nil
}