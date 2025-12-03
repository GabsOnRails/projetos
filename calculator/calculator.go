// Operations package is used to do basic operations (+ - * /) with two numbers only.

package main

import (
	"fmt"
)

func main() {
	var x, y float64
	fmt.Println(Calculator(x, y))

}

// ShowMenu function is used to show available options for user.
func ShowMenu() {
	fmt.Println("Welcome! Choose an option:")

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

// Calculator func is used to do math with numbers.
func Calculator(x float64, y float64) string {
	var choice int
	//Choice var is used for archived the user choice.

	// Recording numbers for operations.
	fmt.Print("Tap a number(x): ")
	fmt.Scanln(&x)
	fmt.Print("Tap a number(y): ")
	fmt.Scanln(&y)

	// Loop beggin.
	for {
		ShowMenu()
		_, err := fmt.Scanln(&choice)
		if err != nil {
			// Error message
			fmt.Println("Wrong input. Please, type a valid number!")
			// Trash var can clean the buffer for new value for input
			var trash string
			fmt.Scanln(&trash)

			// continue loop
			continue
		}

		switch choice {
		case 0:
			// This ir for exit.
			return "See you!"
		case 1:
			// This case if for sum.
			sumResult := x + y
			fmt.Printf("The sum for %.1f and %.1f is %.1f\n", x, y, sumResult)
		case 2:
			// This case if for subtraction.
			SubtractionResult := x - y
			fmt.Printf("The Subtraction for %.1f and %.1f is %.1f\n", x, y, SubtractionResult)
		case 3:
			// This case if for multiple.
			multipleResult := x * y
			fmt.Printf("The multiple for %.1f and %.1f is %.1f\n", x, y, multipleResult)
		case 4:
			// This case if for divide and treat a error.
			if y == 0 {
				fmt.Print("Error: Division by zero is not allowed\n")
			} else {
				divisionResult := x / y
				fmt.Printf("The division for %.1f and %.1f is %.1f\n", x, y, divisionResult)
			}
		case 5:
			// This case is for change number if user wants.
			fmt.Print("Tap a number(x): ")
			fmt.Scanln(&x)
			fmt.Print("Tap a number(y): ")
			fmt.Scanln(&y)
		default:
			fmt.Print("Choose a option between 0 and 5, please.\n")

		}
	}

}
