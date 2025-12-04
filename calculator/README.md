<h1> Simple Calculator in Go </h1>

This is a **basic calculator** implemented in **Go** that performs arithmetic operations with **two numbers only**. It supports **addition, subtraction, multiplication, and division** through a **menu-driven interface**.

---

## Features

- ✅ **Addition**
- ✅ **Subtraction**
- ✅ **Multiplication**
- ✅ **Division** (handles division by zero!)
- ✅ **Change numbers anytime**
- ✅ **Simple and interactive menu**

---

## How It Works

1. The program asks you to **enter two numbers** (`x` and `y`).
2. It shows a **menu of operations**:

```
1 - Sum | 2 - Subtraction | 3 - Multiplication | 4 - Division | 5 - Change Numbers | 0 - Exit
```

3. Choose an operation by typing the corresponding number.
4. The result is displayed immediately.
5. You can **change the numbers** anytime by selecting option `5`.
6. Exit the program by selecting `0`.

---

## Example Usage

```bash
Enter a number (x): 10
Enter a number (y): 5

Welcome! Choose an option:
1 - Sum | 2 - Subtraction | 3 - Multiplication | 4 - Division | 5 - Change Numbers | 0 - Exit
Enter an option: 1
The sum of 10.0 and 5.0 is 15.0

Enter an option: 4
The division of 10.0 by 5.0 is 2.0
```

---

## Notes

- **Division by zero is not allowed**. The program will show an error if you try to divide by 0.
- The calculator **only works with two numbers at a time**.
- Input validation is included to **handle invalid entries gracefully**.

---

## Run the Program

### **Option 1: Run with Go**

1. Clone the repository or copy the files.
2. Run the program:

```bash
go run calculator.go operations.go
```

3. Follow the interactive prompts and enjoy your calculations! 🎉

---

### **Option 2: Build and Run Executable**

1. Build the executable:

```bash
go build -o calculator_exe calculator.go operations.go
```

2. Run the executable:

```bash
./calculator_exe   # Linux / macOS
calculator_exe.exe # Windows
```

---

## Run Tests

The calculator logic is separated into **pure functions** (`Add`, `Subtract`, `Multiply`, `Divide`) and is fully testable.

1. Run tests with Go:

```bash
go test ./...
```

2. If you are using **Ginkgo/Gomega**:

```bash
go run github.com/onsi/ginkgo/v2/ginkgo -v
```

---

## Folder Structure

```
calculator/
├── calculator.go           # Interactive menu
├── operations.go           # Calculator functions (Add, Subtract, Multiply, Divide)
├── calculator_test.go      # Unit tests
├── calculator_suite_test.go# Ginkgo setup
└── README.md               # Project documentation
```
