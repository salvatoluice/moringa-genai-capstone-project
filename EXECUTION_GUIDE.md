# 🚀 Complete Execution Guide - Go Calculator Project

## 📁 Project Structure (Verify you have all these files)

```
go-calculator/
├── main.go                    # Main program entry point
├── calculator.go              # Core calculator functions  
├── calculator_test.go         # Unit tests
├── go.mod                     # Go module definition
├── .gitignore                 # Git ignore rules
├── README.md                  # User documentation
├── TOOLKIT.md                 # Beginner's guide (main deliverable)
├── AI_PROMPT_JOURNAL.md       # AI learning documentation
└── EXECUTION_GUIDE.md         # This file
```

## ⚡ Quick Start (Run in 2 minutes)

1. **Open Terminal/Command Prompt**
2. **Navigate to project folder**: `cd go-calculator`
3. **Run the calculator**: `go run main.go calculator.go`
4. **Test it**: Enter numbers and operations as prompted

## 🧪 Step-by-Step Execution

### Step 1: Verify Go Installation
```bash
go version
# Expected output: go version go1.21.0 (or higher)
```

### Step 2: Initialize and Verify Project
```bash
# Ensure you're in the project directory
cd go-calculator

# Check that go.mod exists
cat go.mod

# Verify all source files are present
ls -la *.go
```

### Step 3: Run Unit Tests
```bash
# Run all tests
go test

# Expected output:
# PASS
# ok      go-calculator   0.002s

# Run tests with detailed output
go test -v

# Check test coverage
go test -cover
```

### Step 4: Run the Calculator
```bash
# Method 1: Direct run (recommended for development)
go run main.go calculator.go

# Method 2: Build then run (creates executable)
go build -o calculator
./calculator        # Linux/Mac
calculator.exe      # Windows
```

## 🎮 Testing the Calculator

### Example Test Session
```
🧮 Welcome to Go Calculator!
================================

Choose an option:
1. Perform calculation
2. View supported operations
3. Exit
Enter your choice (1-3): 1

Enter first number: 10
Enter operation (+, -, *, /, ^): +
Enter second number: 5
✅ Result: 10.00 + 5.00 = 15.00

Would you like to perform another calculation? (y/n): y

Enter first number: 12
Enter operation (+, -, *, /, ^): /
Enter second number: 0
❌ Error: division by zero is not allowed

Would you like to perform another calculation? (y/n): n

Choose an option:
1. Perform calculation
2. View supported operations
3. Exit
Enter your choice (1-3): 3
Thank you for using Go Calculator! Goodbye! 👋
```

### Test All Operations
1. **Addition**: `10 + 5 = 15`
2. **Subtraction**: `10 - 3 = 7`
3. **Multiplication**: `4 * 6 = 24`
4. **Division**: `15 / 3 = 5`
5. **Power**: `2 ^ 3 = 8`
6. **Error Case**: `10 / 0 = Error message`

## 🔧 Troubleshooting

### Problem: "go: command not found"
**Solution**: Go is not installed or not in PATH
```bash
# Check if Go is installed
which go    # Linux/Mac
where go    # Windows

# If not found, reinstall Go from https://golang.org/dl/
```

### Problem: "package main is not a main package"
**Solution**: Check file structure and package declarations
```bash
# Verify first line of main.go and calculator.go
head -1 main.go calculator.go
# Both should show: package main
```

### Problem: Tests failing
**Solution**: Run tests individually to identify issues
```bash
# Run specific test
go test -run TestCalculator_Add
go test -run TestCalculator_Divide
```

### Problem: "no Go files in /path"
**Solution**: Ensure you're in the correct directory
```bash
pwd                    # Check current directory
ls *.go               # List Go files
cd go-calculator       # Navigate to project if needed
```

## 📤 Submission Preparation

### Final Checklist
- [ ] All tests pass: `go test`
- [ ] Calculator runs without errors: `go run main.go calculator.go`
- [ ] All operations work (test each one)
- [ ] Documentation is complete
- [ ] GitHub repository created (optional)

### Create Submission Package
```bash
# Create ZIP file for submission (if not using GitHub)
zip -r go-calculator-submission.zip go-calculator/

# Or create tar.gz
tar -czf go-calculator-submission.tar.gz go-calculator/
```

### GitHub Submission (Recommended)
```bash
# Initialize git repository
git init

# Add all files
git add .

# Commit
git commit -m "Initial commit: Go Calculator Capstone Project"

# Add remote repository (replace with your GitHub repo URL)
git remote add origin https://github.com/yourusername/go-calculator.git

# Push to GitHub
git push -u origin main
```

## 🎯 Evaluation Criteria Checklist

### Clarity & Completeness of Documentation (30%)
- [ ] TOOLKIT.md is comprehensive and well-structured
- [ ] README.md provides clear installation and usage instructions
- [ ] AI_PROMPT_JOURNAL.md documents learning process
- [ ] Code is well-commented

### Use of GenAI for Learning (20%)
- [ ] AI_PROMPT_JOURNAL.md shows 10+ meaningful prompts
- [ ] Each prompt is evaluated and linked to learning outcomes
- [ ] Shows progression from basic to advanced concepts
- [ ] Demonstrates effective AI interaction

### Functionality of Example (20%)
- [ ] Calculator performs all basic operations (+, -, *, /, ^)
- [ ] Error handling works (division by zero, invalid input)
- [ ] User interface is intuitive and user-friendly
- [ ] Code follows Go best practices

### Testing & Iteration (20%)
- [ ] Comprehensive unit tests with multiple test cases
- [ ] All tests pass
- [ ] Edge cases are covered
- [ ] Code has been tested with peers

### Creativity in Chosen Technology (10%)
- [ ] Go was well-chosen for the project
- [ ] Additional features beyond basic requirements
- [ ] Professional code organization
- [ ] Good use of Go-specific features

## 🎉 Success Indicators

Your project is ready for submission when:
1. ✅ Calculator runs smoothly without crashes
2. ✅ All unit tests pass
3. ✅ Documentation is complete and professional
4. ✅ Error handling works properly
5. ✅ A peer can follow your README and run the project
6. ✅ AI learning process is well-documented

## 📞 Need Help?

If you encounter issues:
1. Check this troubleshooting section
2. Refer to the official Go documentation: https://golang.org/doc/
3. Search Stack Overflow for specific error messages
4. Ask a peer to test your project

**Congratulations on completing your Go Calculator Capstone Project! 🎊**