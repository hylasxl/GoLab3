package main

import (
	"fmt"
	"math"
)

type Employee struct {
	name     string
	age      int
	position string
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height float64
}

type Car struct {
	name string
}

type Bike struct {
	name string
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Vehicle interface {
	Start()
	Stop()
}

func (circle Circle) Area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func (car Car) Start() {
	fmt.Println(car.name + " start...")
}

func (car Car) Stop() {
	fmt.Println(car.name + " stop...")
}

func (bike Bike) Start() {
	fmt.Println(bike.name + " start...")
}

func (bike Bike) Stop() {
	fmt.Println(bike.name + " stop...")
}

func (rect Rectangle) Area() float64 {
	return rect.width * rect.height
}

func (circle Circle) Perimeter() float64 {
	return 2 * math.Pi * circle.radius
}

func (rect Rectangle) Perimeter() float64 {
	return 2 * rect.width * rect.height
}

func getHighestAge(employee []Employee) string {
	result := ""
	highestAge := 0
	for _, e := range employee {
		if e.age > highestAge {
			highestAge = e.age
			result = e.name
		}
	}
	return result
}

func getCircumference(circle Circle) float64 {
	return 2 * circle.radius * math.Pi
}

func getArea(circle Circle) float64 {
	return math.Pi * circle.radius * circle.radius
}

func swapValue(x *int, y *int) (swapedX int, swapedY int) {
	temp := x
	x = y
	y = temp
	return *x, *y
}

func increaseValue(x *int) int {
	return *x + 1
}

func factorial(x int) int {
	if x == 1 || x == 0 {
		return 1
	}
	return x + factorial(x-1)
}

func printFibonacci(n int, current int, a int, b int) {
	if current < n {
		fmt.Print(a)
		if current < n-1 {
			fmt.Print(", ")
		}
		printFibonacci(n, current+1, b, a+b)
	}
}

func divideAbyB(a int, b int) int {
	if b == 0 {
		panic("Lỗi: không thể chia cho 0")
	}
	return a / b
}

func handleDefer() {
	defer fmt.Println("Goodbye!")
	fmt.Print("Hello!")
	panic("Có lỗi xảy ra!")

}

func variadicSum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func variadicHighest(nums ...int) int {
	highest := 0
	for _, num := range nums {
		if num > highest {
			highest = num
		}
	}
	return highest
}

func main() {
	employees := []Employee{
		{"Alice", 30, "Engineer"},
		{"Bob", 45, "Manager"},
		{"Charlie", 40, "Designer"},
	}

	circle := Circle{float64(10)}
	rectangle := Rectangle{float64(10), float64(10)}
	car := Car{"G63"}
	bike := Bike{"Martin"}
	x := 3
	y := 4
	fmt.Println("Struct 1: " + getHighestAge(employees))
	fmt.Println("Struct 2 Circumference: " + fmt.Sprintf("%.2f", getCircumference(circle)))
	fmt.Println("Struct 2 Area: " + fmt.Sprintf("%.2f", getArea(circle)))
	fmt.Println("\n")
	fmt.Println("Interface 1 Circle Area: " + fmt.Sprintf("%.2f", circle.Area()))
	fmt.Println("Interface 1 Circle Perimeter: " + fmt.Sprintf("%.2f", circle.Perimeter()))
	fmt.Println("Interface 1 Rectangle Area: " + fmt.Sprintf("%.2f", rectangle.Area()))
	fmt.Println("Interface 1 Rectangle Perimeter: " + fmt.Sprintf("%.2f", rectangle.Perimeter()))
	car.Start()
	car.Stop()
	bike.Start()
	bike.Stop()
	fmt.Println("\n")
	fmt.Print("Pointer 1:")
	fmt.Println(swapValue(&x, &y))
	fmt.Print("Pointer 2:")
	fmt.Print(increaseValue(&x))
	fmt.Println("\n")
	fmt.Println("Recursion:")
	fmt.Println("Factorial: " + fmt.Sprintf("%d", factorial(x)))
	fmt.Println("Fibonacci: " + fmt.Sprintf("%d", factorial(x)))
	printFibonacci(10, 0, 0, 1)
	fmt.Println()
	fmt.Println("Defer, Panic, Recover: ")
	//divideAbyB(5, 0)
	fmt.Println()
	//handleDefer()
	fmt.Println()
	fmt.Println("Variadic")
	fmt.Println("Sum: " + fmt.Sprintf("%d", variadicSum(3, 6, 7, 8, 4, 6)))
	fmt.Println("Highest: " + fmt.Sprintf("%d", variadicHighest(3, 6, 7, 8, 4, 6)))
}
