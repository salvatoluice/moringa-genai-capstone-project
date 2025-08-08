# 🧮 Go Calculator

A simple, interactive command-line calculator built with Go programming language. This project demonstrates basic Go programming concepts including functions, error handling, user input/output, and unit testing.

## 🚀 Features

- **Basic Arithmetic Operations**: Addition, Subtraction, Multiplication, Division
- **Advanced Operations**: Power/Exponentiation
- **Error Handling**: Division by zero protection and input validation
- **Interactive CLI**: User-friendly command-line interface
- **Unit Tests**: Comprehensive test coverage for all calculator functions
- **Input Validation**: Handles invalid inputs gracefully

## 📋 Prerequisites

Before running this calculator, make sure you have:

- **Go Programming Language** (version 1.19 or higher)
  - Download from: https://golang.org/dl/
  - Verify installation: `go version`

## 🛠️ Installation & Setup

1. **Clone or Download the Project**
   ```bash
   git clone <repository-url>
   cd go-calculator
   ```

2. **Initialize Go Module** (if not already done)
   ```bash
   go mod init go-calculator
   ```

3. **Verify Installation**
   ```bash
   go version
   ```

## ▶️ How to Run

### Method 1: Direct Run
```bash
go run main.go calculator.go
```

### Method 2: Build and Execute
```bash
# Build the executable
go build -o calculator

# Run the executable
./calculator        # On Linux/Mac
calculator.exe      # On Windows
```

## 🧪 Running Tests

Execute all unit tests:
```bash
go test
```

Run tests with verbose output:
```bash
go test -v
```

Check test coverage:
```bash
go test -cover
```

## 📖 How to Use

1. **Start the Calculator**
   - Run the program using one of the methods above
   - You'll see the welcome message and main menu

2. **Choose an Option**
   - `1`: Perform calculation
   - `2`: View supported operations
   - `3`: Exit

3. **Perform Calculations**
   - Enter first number
   - Enter operation (+, -, *, /, ^)
   - Enter second number
   - View the result

4. **Example Usage**
   ```
   Enter first number: 10
   Enter operation (+, -, *, /, ^): +
   Enter second number: 5
   ✅ Result: 10.00 + 5.00 = 15.00
   ```

## 🔧 Supported Operations

| Operation | Symbol | Description | Example |
|-----------|--------|-------------|---------|
| Addition | `+` or `add` | Adds two numbers | 5 + 3 = 8 |
| Subtraction | `-` or `subtract` | Subtracts second from first | 5 - 3 = 2 |
| Multiplication | `*` or `multiply` | Multiplies two numbers | 5 * 3 = 15 |
| Division | `/` or `divide` | Divides first by second | 6 / 3 = 2 |
| Power | `^` or `power` | Raises first to power of second | 2 ^ 3 = 8 |

## 🚨 Error Handling

The calculator handles various error cases:

- **Division by Zero**: Returns error message "division by zero is not allowed"
- **Invalid Numbers**: Prompts user to enter valid numeric values
- **Unsupported Operations**: Shows error for operations not supported
- **Empty Input**: Validates and handles empty input fields

## 📁 Project Structure

```
go-calculator/
├── main.go              # Main program with user interface
├── calculator.go        # Core calculator functions
├── calculator_test.go   # Unit tests
├── README.md           # This file
├── TOOLKIT.md          # Beginner's guide documentation
├── AI_PROMPT_JOURNAL.md # AI learning process documentation
├── go.mod              # Go module file
└── .gitignore          # Git ignore rules
```

## 🧪 Test Coverage

The project includes comprehensive unit tests covering:

- ✅ All basic arithmetic operations
- ✅ Error handling scenarios
- ✅ Edge cases (zero values, negative numbers, decimals)
- ✅ Operation validation
- ✅ Input validation

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/new-feature`)
3. Commit your changes (`git commit -am 'Add new feature'`)
4. Push to the branch (`git push origin feature/new-feature`)
5. Create a Pull Request

## 📝 License

This project is open source and available under the [MIT License](LICENSE).

## 👨‍💻 Author

Created as part of Moringa AI Capstone Project - demonstrating Go programming fundamentals with AI-assisted learning.

## 🙏 Acknowledgments

- Go Programming Language Team
- Moringa School AI Program
- AI assistance for learning and development guidance