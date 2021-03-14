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

func arrays() {
	var arr [5]int

	arr[0] = 100
	arr[4] = 900

	fmt.Println(arr)
	fmt.Println(arr[0])

	arr1 := [3]int{4, 5, 6}
	fmt.Println(arr1)

	sum := 0

	for i := 0; i < len(arr1); i++ {
		sum += arr1[i]
	}
	fmt.Printf("sum of the elements in arr1 = %d\n", sum)

	// multidimensional array
	arr2D := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr2D)
	fmt.Println(arr2D[0][1])
}

func slices() {
	// slices are their own data type. Works with arrays.

	var x [5]int = [5]int{1, 2, 3, 4, 5}

	// slice that contains all the elements from the array
	// var s []int = x[:]

	// slice that contains the elements in position 1 and 2. value of 2 and 3
	// var s []int = x[1:3]

	// slice that contains the elements in position 1 through 3. value of 2,3, and 4
	var s []int = x[1:4]
	fmt.Printf("values: %v len: %d capacity %d \n", s, len(s), cap(s))

	var a []int = []int{5, 6, 7, 8, 9}
	a = append(a, 10)
	fmt.Println(a)

	b := make([]int, 5)
	fmt.Println(b)
}

func ranges() {
	var a []int = []int{1, 3, 4, 56, 7, 12, 4, 6, 6}

	for i := 0; i < len(a); i++ {
		fmt.Printf("%d: %d\n", i, a[i])
	}

	sep()

	for i, element := range a {
		fmt.Printf("%d: %d\n", i, element)
	}

	sep()

	for _, element := range a {
		fmt.Printf("%d\n", element)
	}

	sep()

	for i, element := range a {
		for j, element2 := range a {
			if element == element2 && i > j {
				fmt.Println(element)
			}
		}
	}

	sep()

	for i, element := range a {
		for j := i + 1; j < len(a); j++ {
			element2 := a[j]
			if element2 == element {
				fmt.Println(element)
			}
		}
	}
}

func maps() {
	var mp map[string]int = map[string]int{
		"apple":  5,
		"pear":   6,
		"organe": 9,
	}
	fmt.Println(mp["apple"])
	mp["apple"] = 900
	mp["tim"] = 500
	delete(mp, "apple")

	val, ok := mp["pear"]
	fmt.Println(val, ok)

	fmt.Println("map and len - ", mp, len(mp))

	emptyMap := make(map[string]int)
	fmt.Println(emptyMap)
}

func exampleReturns(x, y int) (z1, z2 int) {
	defer fmt.Println("hello")
	z1 = x + y
	z2 = x - y
	defer fmt.Println("before")
	return
}

func advfunc(x int) {
	fmt.Println("hello!", x)
}

// creating a function within a function
func anony() {
	test := func(x int) int {
		fmt.Println("hello", x)
		return x
	}
	test(4)

	t2 := func(x int) int {
		return x * -1
	}(8)
	fmt.Println(t2)

	test2(test)

	test3 := func(x int) int {
		return x * 7
	}
	test2(test3)

	returnFunc("hello")()
}

func test2(myFunc func(int) int) {
	fmt.Println(myFunc(7))
}

func returnFunc(x string) func() {
	return func() { fmt.Println(x) }
}

func mutableAndImmutable() {
	// Immutable
	// seems to store the value in memory
	x := 5
	y := x
	y = 7
	fmt.Println("values:", x, y)

	// slice and map is a mutable type
	// seems to store the memory address
	var s []int = []int{3, 4, 5}
	t := s
	t[0] = 100
	s[1] = 200
	fmt.Println("slice and map:", s, t)

	var a [2]int = [2]int{3, 4}
	b := a
	b[0] = 100
	fmt.Println("array:", a, b)

	var c []int = []int{1, 2, 3}
	fmt.Println("before changeFirst:", c)
	changeFrist(c)
	fmt.Println("after changeFirst:", c)
}

func changeFrist(slice []int) {
	slice[0] = 1000
}

func pointersAndDeference() {
	x := 7
	fmt.Println("memory address:", &x, "with a value of", x)
	y := &x
	fmt.Println("memory address:", &y, "with a value of", y)
	*y = 8
	fmt.Println("memory address:", &y, "with a value of", y)
	fmt.Println("memory address:", &x, "with a value of", x)

	toChange := "hello"
	fmt.Println("before toChange:", toChange)
	changeValue(&toChange)
	fmt.Println("after toChange:", toChange)

	toChange2 := "hello"
	fmt.Println("before toChange2:", toChange2)
	toChange2 = changeValue2(toChange2)
	fmt.Println("after toChange2:", toChange2)

}

func changeValue(str *string) {
	*str = "changed!"
}

func changeValue2(str string) string {
	str = "changed!"
	return str
}

func structsAndCustomTypes() {
	var p1 Point = Point{1, 2}
	p2 := Point{-5, 7}
	p3 := Point{x: 3}
	p4 := &Point{y: 6}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(p4)
	changeX(p4)
	fmt.Println(p4)

	c1 := Circle{4.56, &Point{4, 5}}
	fmt.Println(c1, c1.center, c1.center.x)

	c2 := Circle2{5.87, &Point{2, 4}}
	fmt.Println(c2, c2.x)
}

type Point struct {
	x int32
	y int32
}

func changeX(pt *Point) {
	pt.x = 100
}

type Circle struct {
	radius float64
	center *Point
}

type Circle2 struct {
	radius float64
	*Point
}

func structMethods() {
	s1 := Student{"Tim", []int{70, 90, 80, 85}, 19}
	fmt.Println(s1.getAge())
	s1.setAge(7)
	fmt.Println(s1)
	fmt.Println(s1.getAverageGrade())
	fmt.Println(s1.getMaxGrade())
}

type Student struct {
	name   string
	grades []int
	age    int
}

func (s Student) getAge() int {
	return s.age
}

func (s *Student) setAge(age int) {
	s.age = age
}

func (s Student) getAverageGrade() float32 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float32(sum) / float32(len(s.grades))
}

func (s *Student) getMaxGrade() int {
	curMax := 0

	for _, v := range s.grades {
		if curMax < v {
			curMax = v
		}
	}

	return curMax
}

func main() {
	// calcAge()
	// aritmetic()
	// compare()
	// conditionals()
	// loop()
	// switchStatement()
	// arrays()
	// slices()
	// ranges()
	// maps()
	// ans1, ans2 := exampleReturns(4, 5)
	// fmt.Println(ans1, ans2)

	// assigns a function to a variable
	// x := advfunc
	// x(5)
	// anony()

	// mutableAndImmutable()
	// pointersAndDeference()
	// structsAndCustomTypes()
	structMethods()
}
