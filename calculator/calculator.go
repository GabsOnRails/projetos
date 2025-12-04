// main.go
package main

import (
	"fmt"
)

func main() {
	var x, y float64
	// Start the calculator with initial numbers x and y
	fmt.Println(Calculator(x, y))
}

// ShowMenu displays the available options for the user to interact with the calculator.
func ShowMenu() {
	fmt.Println("\nWelcome! Choose an option:")

	fmt.Print(
		"1 - Sum | ",
		"2 - Subtraction | ",
		"3 - Multiplication | ",
		"4 - Division | ",
		"5 - Change Numbers | ",
		"0 - Exit\n",
		"Enter an option: ",
	)
}

// Calculator executes the main logic of the calculator, handling user input and
// performing the chosen arithmetic operation between two numbers.
func Calculator(x float64, y float64) string {
	var choice int
	// 'choice' stores the user's menu selection.

	// Prompt the user to input the initial numbers.
	fmt.Print("Enter a number (x): ")
	fmt.Scanln(&x)
	fmt.Print("Enter a number (y): ")
	fmt.Scanln(&y)

	// Main loop to repeatedly show the menu and execute operations.
	for {
		ShowMenu()
		_, err := fmt.Scanln(&choice)
		if err != nil {
			// Handle invalid input
			fmt.Println("Invalid input. Please enter a valid number!")
			var trash string
			fmt.Scanln(&trash) // Clear the input buffer
			continue
		}

		switch choice {
		case 0:
			// Exit the calculator
			return "See you!"
		case 1:
			// Perform addition using the Add() function
			sumResult := Add(x, y)
			fmt.Printf("The sum of %.1f and %.1f is %.1f\n", x, y, sumResult)
		case 2:
			// Perform subtraction using the Subtract() function
			subtractionResult := Subtract(x, y)
			fmt.Printf("The subtraction of %.1f and %.1f is %.1f\n", x, y, subtractionResult)
		case 3:
			// Perform multiplication using the Multiply() function
			multiplicationResult := Multiply(x, y)
			fmt.Printf("The multiplication of %.1f and %.1f is %.1f\n", x, y, multiplicationResult)
		case 4:
			// Perform division using the Divide() function
			divisionResult, err := Divide(x, y)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("The division of %.1f by %.1f is %.1f\n", x, y, divisionResult)
			}
		case 5:
			// Allow the user to change the input numbers
			fmt.Print("Enter a new number (x): ")
			fmt.Scanln(&x)
			fmt.Print("Enter a new number (y): ")
			fmt.Scanln(&y)
		default:
			// Handle invalid menu choices
			fmt.Print("Please choose a valid option between 0 and 5.\n")
		}
	}
}
