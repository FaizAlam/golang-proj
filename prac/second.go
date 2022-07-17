package main

import (
	"fmt"
)

func main() {

	//copy_arr()
	//looping()
	checkifExist()
	filtering()
}

func looping() {
	intArr := []int{10, 20, 5, 15, 30}
	loop1(intArr)
	loop2(intArr)
	loop3(intArr)
	loop4(intArr)
}

func loop1(arr []int) {
	println("Loop 1")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func loop2(arr []int) {
	println("Loop 2")
	for index, element := range arr {
		println(index, "=>", element)
	}
}

func loop3(arr []int) {
	println("Loop 3")
	for _, elem := range arr {
		println(elem)
	}
}

func loop4(arr []int) {
	println("Loop 4")
	j := 0
	for range arr {
		println(arr[j])
		j++
	}
}

func copy_arr() {

	intArr := [5]int{10, 20, 5, 15, 30}
	secArr := intArr  //data is passed by value
	thiArr := &intArr //data is passed by reference

	fmt.Printf("........Normal arr........\n")
	fmt.Printf("str arr: %v\n", intArr)
	fmt.Printf("str arr1: %v\n", secArr)
	fmt.Printf("str arr2: %v\n", thiArr)

	fmt.Printf("........secArr[0]=60........\n")
	secArr[0] = 60

	fmt.Printf("str arr: %v\n", intArr)
	fmt.Printf("str arr1: %v\n", secArr)
	fmt.Printf("str arr2: %v\n", thiArr)

	fmt.Printf("........thiArr[0]=60........\n")
	thiArr[0] = 60

	fmt.Printf("str arr: %v\n", intArr)
	fmt.Printf("str arr1: %v\n", secArr)
	fmt.Printf("str arr2: %v\n", thiArr)
}

func checkifExist() {
	arr := []string{"Hello", "Canada", "England"}
	fmt.Println(itemExists(arr, "Bombay"))
}

func itemExists(arr []string, item string) bool {
	ifExist := false

	for _, elem := range arr {
		if elem == item {
			ifExist = true
			break
		}
	}

	return ifExist
}
