package main

import "fmt"

func main() {

	fmt.Println(fib(46))

}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	one := fib(n - 1)
	two := fib(n - 2)
	return one + two
}
