package main

import (
	"fmt"
	"bufio"
	"os"
	"errors"
	"strconv"
)

func main() {
	//First description for how to use calculator
	introduction()

	//Get First Operation to enter for loop
	firstChoice := bufio.NewScanner(os.Stdin)

	//If first Operation is exist, calculator will work forever
	for firstChoice.Scan() {

		//Controlling Chosen Operation if it is proper
		operation, err := selectOperation(firstChoice.Text())
		if err != nil {
			fmt.Println(err)

			//Describe how to use calculator again
			introduction()

			//Return for loop
			continue
		}

		// Get first and second numbers
		firstNumber, secondNumber := inputNumbers()

		// Calculate them!
		result := calculateResult(firstNumber, secondNumber, operation)

		//Print result
		fmt.Println("Result:", result)

		//Describe how to use calculator again
		introduction()
	}
}

// Function of The Description of Calculator
func introduction() {
	fmt.Println("\n----- Choose the Operation You Want to Perform -----")
	fmt.Println("----- + (Addition) ------")
	fmt.Println("----- - (Subtraction) ------")
	fmt.Println("----- / (Division) ------")
	fmt.Println("----- * (Multiplication) ------")
	fmt.Print("> ")
}

func selectOperation(choice string) (string, error) {
	switch choice {
	case "+":
		return "+", nil
	case "-":
		return "-", nil
	case "/":
		return "/", nil
	case "*":
		return "*", nil
	default:
		return "", errors.New("! Invalid Operation Selection !")
	}
}

// Function of The Getting Numbers
func inputNumbers() (int, int) {
	var firstInt, secondInt int

	fmt.Print("First Number > ")
	first := bufio.NewScanner(os.Stdin)
	for first.Scan() {
		firstInt_, firstErr := strconv.Atoi(first.Text())
		if firstErr != nil {
			fmt.Println("Invalid input. Please enter a valid integer.")
			fmt.Print("> ")
			continue
		}
		firstInt = firstInt_
		break
	}

	fmt.Print("Second Number > ")
	second := bufio.NewScanner(os.Stdin)
	for second.Scan() {
		secondInt_, secondErr := strconv.Atoi(second.Text())
		if secondErr != nil {
			fmt.Println("Invalid input. Please enter a valid integer.")
			fmt.Print("> ")
			continue
		}
		secondInt = secondInt_
		break
	}

	return firstInt, secondInt
}

func calculateResult(first int, second int, operation string) (int) {
	switch operation {
	case "+":
		return first + second
	case "-":
		return first - second
	case "/":
		return first / second
	case "*":
		return first * second
	default:
		return 0
	}
}