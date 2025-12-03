<h1> Simple Calculator in Go </h1>

This is a **basic calculator** implemented in **Go** that can perform operations with **two numbers only**. It supports addition, subtraction, multiplication, and division, with a simple **menu-driven interface**.

Features

- ✅ **Addition**
- ✅ **Subtraction**
- ✅ **Multiplication**
- ✅ **Division** (handles division by zero!)
- ✅ **Change numbers anytime**
- ✅ **Simple and interactive menu**

How it works

1. The program asks you to **enter two numbers** (`x` and `y`).
2. Shows a **menu of options**:

```
1 - Sum | 2 - Subtraction | 3 - Multiplication | 4 - Division | 5 - Change Numbers | 0 - Exit
```

3. Choose an operation by typing the corresponding number.
4. The result is displayed immediately.
5. You can **change the numbers** anytime by selecting option `5`.
6. Exit the program by selecting `0`.

Example Usage

```bash
Tap a number(x): 10
Tap a number(y): 5

Welcome! Choose an option:
1 - Sum | 2 - Subtraction | 3 - Multiplication | 4 - Division | 5 - Change Numbers | 0 - Exit
Enter an option: 1
The sum for 10.0 and 5.0 is 15.0

Enter an option: 4
The division for 10.0 and 5.0 is 2.0
```

Notes

- **Division by zero is not allowed**. The program will show an error if you try to divide by 0.
- The calculator **only works with two numbers at a time**.
- Input validation is included to **handle wrong entries** gracefully.

Run the Program

1. Clone the repository or copy the file.
2. Run the program with Go:

```bash
go run calculator.go
```

3. Follow the interactive prompts and enjoy your calculations! 🎉
