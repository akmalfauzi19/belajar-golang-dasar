package main

func factorial(n int) int {
	result := 1

	for i := n; i > 0; i-- {
		result *= i
	}

	return result
}

func factorialRecursive(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorialRecursive(n-1)
}

func main() {
	println(factorialRecursive(5))
	println(factorial(5))
}
