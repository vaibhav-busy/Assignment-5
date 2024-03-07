package main

import "fmt"

func divide(num1, num2 int) (result float32, err error) {
	defer func() {
		fmt.Println("Another defer1")
	}()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from division Zero panic.")
			err = fmt.Errorf("Divsion by Zero", r)
		}
	}()

	defer func() {
		fmt.Println("Another defer2")
	}()

	if num2 == 0 {
		fmt.Println("Panic Division by Zero")
		panic("Division by Zero")
	}

	res := float32(num1) / float32(num2)

	fmt.Println("Normal flow resume 1")

	return float32(res), nil

}

func main() {
	var num1, num2 int

	fmt.Println("Enter Num 1: ")
	_, err := fmt.Scanln(&num1)

	if err != nil {
		fmt.Println("Error in input1")
		return
	}

	fmt.Println("Enter Num 2: ")
	_, err2 := fmt.Scanln(&num2)

	if err2 != nil {
		fmt.Println("Error in input2")
		return
	}

	res, err3 := divide(num1, num2)

	fmt.Println("Normal flow resume 2")

	if err3 != nil {
		fmt.Println("Error is: ", err3)
		return
	}

	fmt.Println("Normal flow resume 3")

	fmt.Println("division result is : ", res)
}
