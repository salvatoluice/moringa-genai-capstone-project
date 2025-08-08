# Getting Started with Go Programming – A Beginner's Guide to Building a Calculator

## 1. Title & Objective

**Technology**: Go (Golang) Programming Language  
**Project**: Interactive Command-Line Calculator  
**Why Go?**: Go is an excellent choice for beginners because of its simple syntax, fast compilation, strong standard library, and excellent error handling. It's used by companies like Google, Docker, Kubernetes, and many startups for backend services.  
**End Goal**: Build a fully functional calculator that performs basic arithmetic operations (+, -, *, /, ^) with proper error handling and user-friendly interface.

## 2. Quick Summary of the Technology

**What is Go?**  
Go (also called Golang) is a programming language developed by Google in 2009. It's designed to be simple, efficient, and reliable for building software that scales.

**Where is it used?**  
- **Web servers and APIs** (companies like Uber, Netflix)
- **Cloud infrastructure** (Docker, Kubernetes)
- **Microservices** (Dropbox, SoundCloud)
- **Command-line tools** (GitHub CLI, Terraform)
- **Database systems** (CockroachDB, InfluxDB)

**Real-world example**: Docker, the popular containerization platform, is written entirely in Go. When you run `docker run`, you're using Go code!

## 3. System Requirements

**Operating System**: Linux, macOS, or Windows  
**Required Tools**:
- Go compiler (version 1.19 or higher)
- Text editor (VS Code recommended) or any IDE
- Terminal/Command prompt
- Git (optional, for version control)

**Hardware Requirements**:
- Minimum 2GB RAM
- 500MB free disk space
- Internet connection for downloading Go

## 4. Installation & Setup Instructions

### Step 1: Install Go

**For Windows:**
1. Visit https://golang.org/dl/
2. Download `go1.21.windows-amd64.msi`
3. Run the installer and follow default settings
4. Open Command Prompt and verify: `go version`

**For macOS:**
1. Visit https://golang.org/dl/
2. Download `go1.21.darwin-amd64.pkg`
3. Run the installer
4. Open Terminal and verify: `go version`

**For Linux:**
```bash
# Download and extract
wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz

# Add to PATH
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc

# Verify installation
go version
```

### Step 2: Setup Development Environment

**Install VS Code (Recommended)**:
1. Download from https://code.visualstudio.com/
2. Install Go extension by Google
3. Open VS Code and install Go tools when prompted

### Step 3: Create Project

```bash
# Create project directory
mkdir go-calculator
cd go-calculator

# Initialize Go module
go mod init go-calculator

# Create main files
touch main.go calculator.go calculator_test.go
```

## 5. Minimal Working Example

### What the Example Does
Our calculator performs basic arithmetic operations through an interactive command-line interface. Users can:
- Choose operations from a menu
- Enter two numbers
- Select an operation (+, -, *, /, ^)
- Get calculated results with error handling

### Core Calculator Functions (calculator.go)
```go
package main

import (
    "errors"
    "fmt"
)

type Calculator struct{}

func NewCalculator() *Calculator {
    return &Calculator{}
}

// Basic arithmetic operations
func (c *Calculator) Add(a, b float64) float64 {
    return a + b
}

func (c *Calculator) Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero is not allowed")
    }
    return a / b, nil
}

// Additional operations: Subtract, Multiply, Power...
```

### Main Program (main.go)
```go
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
    calculator := NewCalculator()
    scanner := bufio.NewScanner(os.Stdin)
    
    // Interactive menu loop
    for {
        fmt.Println("1. Perform calculation")
        fmt.Println("2. Exit")
        // Handle user input and calculations...
    }
}
```

### Expected Output
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
```

### Running the Program
```bash
# Method 1: Direct run
go run main.go calculator.go

# Method 2: Build then run
go build -o calculator
./calculator
```

## 6. AI Prompt Journal

### Prompt 1: Technology Overview
**Prompt Used**: "Explain Go programming language in simple terms - what it is, why it's popular, where it's used in real-world applications, and why it's good for beginners learning programming"

**AI Response Summary**: AI explained that Go is a statically typed, compiled language created by Google that emphasizes simplicity and efficiency. It's popular for backend services, cloud infrastructure, and has a gentle learning curve for beginners.

**Evaluation**: ⭐⭐⭐⭐⭐ Extremely helpful. Provided clear context and real-world examples that helped understand why Go is valuable to learn.

### Prompt 2: Project Structure
**Prompt Used**: "How should I structure a simple calculator project in Go? What files do I need and how should I organize the code?"

**AI Response Summary**: AI recommended separating concerns into multiple files: main.go for user interface, calculator.go for logic, and calculator_test.go for testing. Suggested using Go modules for dependency management.

**Evaluation**: ⭐⭐⭐⭐⭐ Perfect guidance. The modular structure made the code much more maintainable and professional.

### Prompt 3: Error Handling
**Prompt Used**: "How do I implement proper error handling in Go, specifically for a calculator that needs to handle division by zero and invalid input?"

**AI Response Summary**: AI showed Go's explicit error handling pattern using multiple return values, how to create custom errors, and input validation techniques using strconv.ParseFloat.

**Evaluation**: ⭐⭐⭐⭐⭐ Excellent. Error handling is crucial and AI provided comprehensive examples that made the calculator robust.

### Prompt 4: Testing
**Prompt Used**: "How do I write unit tests for Go functions? Show me how to test a calculator with different test cases including edge cases."

**AI Response Summary**: AI demonstrated Go's built-in testing package, table-driven tests, and how to test both success and error cases. Provided examples of testing edge cases like zero values and negative numbers.

**Evaluation**: ⭐⭐⭐⭐⭐ Outstanding. Testing knowledge gained here will be valuable for all future Go projects.

### Prompt 5: User Interface
**Prompt Used**: "Create a user-friendly command-line interface for a Go calculator that takes user input and displays results clearly"

**AI Response Summary**: AI showed how to use bufio.Scanner for input, string manipulation for validation, and formatting for clean output display with emojis and clear messages.

**Evaluation**: ⭐⭐⭐⭐⭐ Very helpful. Made the calculator much more professional and user-friendly.

## 7. Common Issues & Fixes

### Issue 1: "go: command not found"
**Problem**: Go not installed properly or not in PATH  
**Solution**: 
- Reinstall Go following official instructions
- Add Go to system PATH: `export PATH=$PATH:/usr/local/go/bin`
- Restart terminal/command prompt

### Issue 2: "package main is not a main package"
**Problem**: Incorrect package declaration  
**Solution**: Ensure first line of main.go is `package main` and you have a `main()` function

### Issue 3: "invalid character 'ï' looking for beginning of value"
**Problem**: File encoding issues (BOM in UTF-8)  
**Solution**: Save files as UTF-8 without BOM in your editor

### Issue 4: Division by zero not handled
**Problem**: Calculator crashes on division by zero  
**Solution**: 
```go
func (c *Calculator) Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero is not allowed")
    }
    return a / b, nil
}
```

### Issue 5: Tests failing
**Problem**: Test file not recognized  
**Solution**: Ensure test file ends with `_test.go` and functions start with `Test`

**References for Issues**:
- Stack Overflow Go section: https://stackoverflow.com/questions/tagged/go
- Go GitHub Issues: https://github.com/golang/go/issues
- Go Forum: https://forum.golangbridge.org/

## 8. References

### Official Documentation
- **Go Official Website**: https://golang.org/
- **Go Documentation**: https://golang.org/doc/
- **Go Tour (Interactive)**: https://tour.golang.org/
- **Effective Go**: https://golang.org/doc/effective_go.html

### Video Tutorials
- **Go Programming Tutorial**: https://www.youtube.com/watch?v=YS4e4q9oBaU
- **Learn Go in 12 Minutes**: https://www.youtube.com/watch?v=C8LgvuEBraI
- **Go Crash Course**: https://www.youtube.com/watch?v=SqrbIlUwR0U

### Helpful Blog Posts
- **A Tour of Go**: https://tour.golang.org/welcome/1
- **Go by Example**: https://gobyexample.com/
- **Go Documentation**: https://pkg.go.dev/
- **awesome-go**: https://github.com/avelino/awesome-go (curated list of Go resources)

### Books (Free)
- **The Go Programming Language Specification**: https://golang.org/ref/spec
- **An Introduction to Programming in Go**: http://www.golang-book.com/

---

**Total Development Time**: ~6-8 hours  
**Difficulty Level**: Beginner-Friendly  
**Key Skills Learned**: Go syntax, error handling, user input/output, testing, project structure  
**Next Steps**: Try building a web API, working with databases, or exploring Go's concurrency features (goroutines)