package factfibpack

import (
	"fmt"
)

func Fibonacci(n int) {
	a := 0
	b := 1
	fmt.Println(a)
	fmt.Println(b)
	count := 2
	for i := 0; i <= n-2; i++ {
		c := a + b
		fmt.Println(c)
		a = b
		b = c
		count = count + 1
	}

}
