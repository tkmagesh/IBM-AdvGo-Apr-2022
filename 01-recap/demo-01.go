/* Functions */

package main

import (
	"errors"
	"fmt"
	"log"
)

var DivideByZeroError = errors.New("divide by zero error")

func main() {
	/* functions can return more than one result */
	q, r, e := divide(100, 7)
	if e != nil {
		log.Println(e)
	}
	fmt.Println(q, r)

	/* anonymous functions */
	func() {
		fmt.Println("anonymous function invoked")
	}()

	func(x, y int) {
		fmt.Println(x + y)
	}(100, 200)

	result := func(x, y int) int {
		return x + y
	}(100, 200)
	fmt.Println("Result = ", result)

	/* higher order functions */
	/*
		functions as data
		a. functions can be assigned as values to variables
		b. functions can be passed as arguments to other functions
		c. functions can be returned as return values
	*/

	//a. functions can be assigned as values to variables
	/*
		fn := func() {
			fmt.Println("fn invoked")
		}
	*/
	var fn func()
	fn = func() {
		fmt.Println("fn invoked")
	}
	fn()

	var multiply func(int, int) int
	multiply = func(x, y int) int {
		return x * y
	}
	fmt.Println(multiply(100, 200))

	//b. functions can be passed as arguments to other functions
	exec(fn)

	nos := []int{3, 1, 4, 2, 5}
	/*
		evenNos := filterEven(nos)
		fmt.Println("Even Nos = ", evenNos)

		oddNos := filterOdd(nos)
		fmt.Println("Odd Nos = ", oddNos)
	*/
	evenPredicate := func(no int) bool {
		return no%2 == 0
	}
	evenNos := filter(nos, evenPredicate)
	fmt.Println("Even Nos = ", evenNos)

	oddNos := filter(nos, func(no int) bool {
		return no%2 != 0
	})
	fmt.Println("Odd Nos = ", oddNos)

	//c. functions can be returned as return values
	/*
		add(100, 200)
		subtract(100, 200)
	*/
	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/
	/*
		logOperation(100, 200, add)
		logOperation(100, 200, subtract)
	*/
	logAdd := getLoggedOperation(add)
	logAdd(100, 200)

	logSubtract := getLoggedOperation(subtract)
	logSubtract(100, 200)

	/* closures */
	increment := getIncrement()
	fmt.Println(increment()) // => 1
	fmt.Println(increment()) // => 2
	//dummyFn()
	fmt.Println(increment()) // => 3
	fmt.Println(increment()) // => 4

	fmt.Println()
	for i := 0; i < 100; i++ {
		go func(no int) {
			fx(no)
		}(i)
	}

	var input string
	fmt.Scanln(&input)
}

func fx(no int) {
	fmt.Println("fx - ", no)
}

/* func dummyFn() {
	counter = 1000
} */

func getIncrement() func() int {
	var counter int
	increment := func() int {
		counter++
		return counter
	}
	return increment
}

/*
func logAdd(x, y int) {
	fmt.Println("invocation started")
	add(x, y)
	fmt.Println("invocation completed")
}

func logSubtract(x, y int) {
	fmt.Println("invocation started")
	subtract(x, y)
	fmt.Println("invocation completed")
}
*/

func logOperation(x, y int, oper func(int, int)) {
	fmt.Println("invocation started")
	oper(x, y)
	fmt.Println("invocation completed")
}

func getLoggedOperation(oper func(int, int)) func(int, int) {
	return func(x, y int) {
		fmt.Println("invocation started")
		oper(x, y)
		fmt.Println("invocation completed")
	}
}

/*
func divide(x, y int) (quotient, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
*/

/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

/*
func divide(x, y int) (int, int, error) {
	if y == 0 {
		return 0, 0, DivideByZeroError
	}
	quotient := x / y
	remainder := x % y
	return quotient, remainder, nil
}
*/

func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = DivideByZeroError
		return
	}
	quotient = x / y
	remainder = x % y
	return
}

func exec(fn func()) {
	fn()
}

func filterEven(nos []int) (evenNos []int) {
	for _, no := range nos {
		if no%2 == 0 {
			evenNos = append(evenNos, no)
		}
	}
	return
}

func filterOdd(nos []int) (oddNos []int) {
	for _, no := range nos {
		if no%2 != 0 {
			oddNos = append(oddNos, no)
		}
	}
	return
}

func filter(nos []int, predicate func(int) bool) (result []int) {
	for _, no := range nos {
		if predicate(no) {
			result = append(result, no)
		}
	}
	return
}

func add(x, y int) {
	fmt.Println("Add result = ", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract result = ", x-y)
}
