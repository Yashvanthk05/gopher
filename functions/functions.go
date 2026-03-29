package functions

import "fmt"

func Helloworld(name string) string {
	return "Hello, " + name + "!"
}

func add(a, b int) int {
	return a + b
}

func swap(x, y string) (string, string) {
	return y, x
}

func variadic(nums ...int) int {
	fmt.Println("Numbers:", nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func subtract(x, y int) int {
	return x - y
}

func EvenOld(nums ...int) (int,int) {
	evencnt := 0
	oldcnt := 0
	for _,num := range nums{
		if num%2==0 { evencnt++ } else { oldcnt++ }
	} 
	return evencnt,oldcnt
}

// higher order functions or first class functions - passing functions as arguements to another function

func Func() {
	fmt.Println(Helloworld("Yashvanth"))
	fmt.Println("Sum:", add(3, 5))
	a, b := swap("Hello", "World")
	fmt.Println("Swapped:", a, b)
	fmt.Println("Variadic sum:", variadic(1, 2, 3, 4, 5))
	evencnt, oldcnt := EvenOld(1,2,3,4,5)
	fmt.Printf("Even Count: %d, Old Count: %d\n",evencnt,oldcnt)
}