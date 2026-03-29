package main

import (
	"fmt"
)

func sum(nums ...int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum
}

func main(){
	// slice = reference to an underlying array
	// arrays - fixed size 
	// slices - dynamic size (can grow using append)
	// arr - array
	arr := [2]int{1,2}
	fmt.Println(arr)
	// arr = append(arr,3) will result in error since it is not a slice (it is fixed)
	// arr1 - slice
	arr1 := []string{"yashvanth","kavin","deepak"}
	fmt.Println(arr1)
	fmt.Println(arr1[1:])
	fmt.Println(len(arr1))
	arr1 = append(arr1,"abinav")
	fmt.Println(arr1)
	b := arr1[1:]
	b = append(b,"kishore")
	fmt.Println(b)
	fmt.Println(arr1)
	// slice reference arr and multiple slices can reference the same arr
	// slice values are not passed by value
	// make - built in function used to create and initialize slices, maps and channels
	// make(type, size, capacity); capacity - maximum space (optional)
	fmt.Println(cap(b))

	// go reallocation problem
	/*
	When you use append and the slice is full
	If capacity is not enough -> Go creates a new array
	Go may repeatedly allocate new memory, copy data again and again
	Solution is to use: make
	*/
	s := make([]int, 3, 10)
	s[0] = 1
	fmt.Println(s)
	/*
	“Go runtime doesn’t need to reallocate memory continuously” means: If you allocate enough capacity in advance, Go doesn’t need to repeatedly create new memory and copy data
	*/
	// cap - maximum length of the slice before reallocation into new underlying array
	// len - current length of the slice

	// variadic - variable number of arguements for a function (print, append are variadic function)
	fmt.Println(sum(1,2))
	fmt.Println(sum(2,3,4,5))
	// spread operator - allows to pass a slice to a variadic function
	nums := []int{5,6,7,8,9}
	nums2 := []int{10,11,12}
	// append(slice, ele1, ele2,....)
	nums = append(nums,nums2...)
	fmt.Println(sum(nums...))
}