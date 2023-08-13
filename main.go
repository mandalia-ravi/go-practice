package main

import "fmt"

func main() {
	// learning the most important line
	fmt.Println("Hello World!!!")

	// variable declaration
	var x int = 5
	y := 10
	const pi = 3.14

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("pi:", pi)

	var age int = 25
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Pi: %.2f\n", pi)

	// if else
	if age >= 25 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}

	result := add(5, 3)
	fmt.Println("Sum:", result)

	// difference between using var and :=
	/*
		Variables declared with var at the package level have a package-level scope.
		Variables declared with var within a function have a function-level scope.
		Variables declared with := have a block-level scope, which means they are only accessible within the block where they are declared.
	*/

	printNumbers()
	ArrAndSlice()
	maps()
	printOddEven(21)
	print(50)
	copyMapValues()

	fmt.Println(reverseString("ravi"))
	fmt.Println(reverseStringUsingSlice("ravi"))

	input := "racecar"
	if isPalindrome(input) {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not a palindrome")
	}

	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println("Maximum Subarray Sum:", maxSubarraySum(nums))
}

// add 2 numbers
func add(a, b int) int {
	return a + b
}

// print numbers 1 to 10
func printNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

// Arrays and slices
func ArrAndSlice() {
	var nums [5]int
	nums[0] = 10
	nums[1] = 20

	fruits := []string{"apple", "banana", "peach"}

	var veggies []string

	fmt.Println("Array: ", nums)
	fmt.Println("Slice: ", fruits)
	fmt.Println("Veggies: ", veggies)
}

func maps() {
	ages := map[string]int{
		"Alice": 20,
		"Bob":   22,
	}

	var countString map[string]int

	ages["Charlie"] = 26

	fmt.Println("Ages:", ages)
	fmt.Println(countString)
	x := "Bob"
	fmt.Println(ages["Alice"])
	fmt.Println(ages[x])

	if age, ok := ages["Ravi"]; ok {
		fmt.Println("Charlie's age", age)
	} else {
		fmt.Println("Charlie not found in the map.")
	}
	delete(ages, "Bob")
	fmt.Println("Ages:", ages)

	for name, age := range ages {
		fmt.Printf("%s %d\n", name, age)
	}
}

func printOddEven(x int) {
	if x%2 == 0 {
		fmt.Printf("Number %d is even\n", x)
		return
	}
	fmt.Printf("Number %d is odd\n", x)
}

func print(x int) {
	for i := 0; i <= x; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func copyMapValues() {
	original := map[string]int{"a": 1, "b": 2}

	copyMap := make(map[string]int)
	for key, value := range original {
		copyMap[key] = value
	}

	copyMap["a"] = 42
	fmt.Println(copyMap["a"])
	fmt.Println(original["a"])
}

func reverseString(s string) (result string) {
	for _, i := range s {
		result = string(i) + result
	}
	return
}

// reverse using string slice
func reverseStringUsingSlice(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// palindrome check
func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func maxSubarraySum(nums []int) int {
	maxSum, currentSum := nums[0], nums[0]

	for i := range nums {
		currentSum = max(i, currentSum+i)
		maxSum = max(maxSum, currentSum)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
