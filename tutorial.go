package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sep() {
	fmt.Println("---")
	fmt.Println()
}

func calcAge() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter the year you were born: ")
	scanner.Scan()
	input, err := strconv.ParseInt(scanner.Text(), 10, 64)

	if err != nil {
		fmt.Println("Error. Error. ", err)
		return
	}

	fmt.Printf("You will be %d years old in 2020\n", 2020-input)
}

func aritmetic() {
	var num1 float64 = 9
	var num2 int = 4
	answer := num1 / float64(num2)
	fmt.Printf("%g\n", answer)
}

func compare() {
	x := 5
	y := 6
	val := x <= y

	fmt.Printf("x: %v, y: %v, val: %v", x, y, val)
}

func conditionals() {
	fmt.Println("Before if")

	if false {
		fmt.Println("in if")
	}

	fmt.Println("after if")
}

func loop() {
	// go only has one looping construct

	// for { } // least amount of code

	x := 3
	for x < 5 {
		fmt.Println(x)
		x++
	}
	sep()

	for x := 3; x < 5; x++ {
		fmt.Println(x)
	}
	sep()

	for i := 0; i < 1000; i++ {
		if i != 0 && i%3 == 0 && i%7 == 0 && i%9 == 0 {
			fmt.Println(i)
		}

	}

}

func switchStatement() {
	ans := 10

	switch ans {
	case 0:
		fmt.Println("zero")
	case 1, -1:
		fmt.Println("one")
	case 10:
		fmt.Println("ten")
	default:
		fmt.Println("not a case")
	}

	// doesn't require a condition next to the keyword switch
	// same as switch true
	switch {
	// condition needed next to the keywork case
	case ans > 0:
		fmt.Println("greater than zero")
	default:
		fmt.Println("less than zero")
	}

	v1 := 30

	// switch evaluates from top to bottom, stopping when a case succeeds
	switch {
	case v1 > 0:
		fmt.Println("v1")
	case ans > 0:
		fmt.Println("ans")

	}

}

func main() {
	// calcAge()
	// aritmetic()
	// compare()
	// conditionals()
	// loop()
	switchStatement()
}
