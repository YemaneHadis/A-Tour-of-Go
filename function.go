package main

import "fmt"

// named return vaule

func split(sum int) (x, y int) {
	x = sum / 2
	y = sum - x
	return
}

// go variable

func varaibleType(num int8) {
	fmt.Println("Go Varaible", num)
}

// function can return any nnumber of result

func swap(x, y string) (string, string) {

	return y, x
}

func addtion(x int, y int) int {
	return x + y
}

// variables

var java, c, python bool
var maxInt uint64 = 1<<64 - 1

func main() {

	varaibleType(127)
	var a string
	var b string
	a = "stringA"

	b = "StringB"
	fmt.Printf("Before Swap  A %s,  B %s", a, b)
	a, b = swap(a, b)
	fmt.Printf("After Swap A  %s,  B %s", a, b)
	fmt.Print(addtion(12, 21))
	fmt.Println("")
	fmt.Print(split(21))

	fmt.Println(maxInt)

}
