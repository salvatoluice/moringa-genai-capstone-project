package main

import (
	"testing"
)

func TestCalculator_Add(t *testing.T) {
	calc := NewCalculator()
	
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 5.0, 3.0, 8.0},
		{"negative numbers", -5.0, -3.0, -8.0},
		{"mixed numbers", 5.0, -3.0, 2.0},
		{"with decimals", 5.5, 3.3, 8.8},
		{"zero values", 0.0, 5.0, 5.0},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%.2f, %.2f) = %.2f; expected %.2f", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestCalculator_Subtract(t *testing.T) {
	calc := NewCalculator()
	
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 10.0, 3.0, 7.0},
		{"negative result", 3.0, 10.0, -7.0},
		{"negative numbers", -5.0, -3.0, -2.0},
		{"with decimals", 5.7, 2.2, 3.5},
		{"zero values", 5.0, 0.0, 5.0},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Subtract(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Subtract(%.2f, %.2f) = %.2f; expected %.2f", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestCalculator_Multiply(t *testing.T) {
	calc := NewCalculator()
	
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 4.0, 3.0, 12.0},
		{"negative numbers", -4.0, -3.0, 12.0},
		{"mixed signs", 4.0, -3.0, -12.0},
		{"with decimals", 2.5, 4.0, 10.0},
		{"with zero", 5.0, 0.0, 0.0},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Multiply(%.2f, %.2f) = %.2f; expected %.2f", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestCalculator_Divide(t *testing.T) {
	calc := NewCalculator()
	
	tests := []struct {
		name        string
		a, b        float64
		expected    float64
		expectError bool
	}{
		{"positive numbers", 10.0, 2.0, 5.0, false},
		{"negative numbers", -10.0, -2.0, 5.0, false},
		{"mixed signs", 10.0, -2.0, -5.0, false},
		{"with decimals", 7.5, 2.5, 3.0, false},
		{"division by zero", 5.0, 0.0, 0.0, true},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.Divide(tt.a, tt.b)
			
			if tt.expectError {
				if err == nil {
					t.Errorf("Divide(%.2f, %.2f) expected error but got none", tt.a, tt.b)
				}
			} else {
				if err != nil {
					t.Errorf("Divide(%.2f, %.2f) unexpected error: %v", tt.a, tt.b, err)
				}
				if result != tt.expected {
					t.Errorf("Divide(%.2f, %.2f) = %.2f; expected %.2f", tt.a, tt.b, result, tt.expected)
				}
			}
		})
	}
}

func TestCalculator_Power(t *testing.T) {
	calc := NewCalculator()
	
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"basic power", 2.0, 3.0, 8.0},
		{"power of 1", 5.0, 1.0, 5.0},
		{"power of 0", 5.0, 0.0, 1.0},
		{"base 1", 1.0, 5.0, 1.0},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Power(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Power(%.2f, %.2f) = %.2f; expected %.2f", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestCalculator_PerformOperation(t *testing.T) {
	calc := NewCalculator()
	
	tests := []struct {
		name        string
		a, b        float64
		operation   string
		expected    float64
		expectError bool
	}{
		{"addition with +", 5.0, 3.0, "+", 8.0, false},
		{"addition with add", 5.0, 3.0, "add", 8.0, false},
		{"subtraction", 10.0, 4.0, "-", 6.0, false},
		{"multiplication", 6.0, 7.0, "*", 42.0, false},
		{"division", 15.0, 3.0, "/", 5.0, false},
		{"power", 2.0, 4.0, "^", 16.0, false},
		{"invalid operation", 5.0, 3.0, "%", 0.0, true},
		{"division by zero", 5.0, 0.0, "/", 0.0, true},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.PerformOperation(tt.a, tt.b, tt.operation)
			
			if tt.expectError {
				if err == nil {
					t.Errorf("PerformOperation(%.2f, %.2f, %s) expected error but got none", tt.a, tt.b, tt.operation)
				}
			} else {
				if err != nil {
					t.Errorf("PerformOperation(%.2f, %.2f, %s) unexpected error: %v", tt.a, tt.b, tt.operation, err)
				}
				if result != tt.expected {
					t.Errorf("PerformOperation(%.2f, %.2f, %s) = %.2f; expected %.2f", tt.a, tt.b, tt.operation, result, tt.expected)
				}
			}
		})
	}
}