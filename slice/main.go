package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

func changeVal(arr []int, index uint, val int) {
	if index < uint(len(arr)) {
		arr[index] = val
	}
	fmt.Println("Array inside func", arr)
}
func main() {
	// slice is a growable array
	// slice is a reference type
	arr := []int{10, 11, 12} // This is shorthand declaration of slice ;
	fmt.Println("Array", arr)
	changeVal(arr, 2, 112)
	fmt.Println("Array after changed", arr)
	arr = []int{10, 11, 12, 13, 14, 15}
	fmt.Println("Array", arr)
	// There are three defined reference types in golang
	// 1- slice
	// 2- map
	// 3- channel

	// pointers , slices, maps, channels you can check them whethere they are nil or not
	var slice []int // declaring the slice

	// Employee emp = new Employee();

	//
	// Emoployee emp1;

	// public EmployeeDetails(Employee emp){
	//if(emp!==null){
	// emp.ProcessSalary();
	// }
	// }

	// Ok EmployeeDetails ed = new EmployeeDetails(emp);
	// Fail EmployeeDetails ed1 = new EmployeeDetails(emp1);

	fmt.Println("Type of a slice", reflect.TypeOf(slice))
	fmt.Println("Length of a slice", len(slice), "\nCapacity of a slice", cap(slice))
	if slice == nil {
		fmt.Println("yes slice is nil, yet to be instantiated")
	}
	slice = make([]int, 10) // make makes a slice provided by given type , length and capacity
	if slice == nil {
		fmt.Println("Slice is nill")
	} else {
		fmt.Println("Len ", len(slice), "\ncap ", cap(slice))
		fmt.Println("Slice", slice)
	}

	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Intn(200)
	}
	fmt.Println("Slice", slice)

	slice = append(slice, 1000)
	slice = append(slice, 1001)
	for i := 0; i < len(slice); i++ {
		fmt.Println("Index", i, "Value", slice[i])
	}
	fmt.Println("Len ", len(slice), "\ncap ", cap(slice))
}
