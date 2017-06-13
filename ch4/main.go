package main

import (
	"fmt"
)

func sliceT() {
	fmt.Println("----------sliceT---------------")
	//slice := make([]string, 5)
	//slice2 := make([]string, 3, 5)

	slice3 := []string{"Red", "Green", "Blue", "Yellow", "Pink"}

	newSlice := slice3[1:3]
	newSlice2 := slice3[2:3:4]

	//newSlice3[1] = "KKK"
	newSlice = append(newSlice, "Purple")
	newSlice = append(newSlice, "Black")
	newSlice = append(newSlice, "White")
	newSlice = append(newSlice, "Brown")
	for i, it := range newSlice {
		fmt.Printf("index %d val %s\n", i, it)
	}

	fmt.Println("----------newSlice2---------------")
	for i, it := range newSlice2 {
		fmt.Printf("index %d val %s\n", i, it)
	}
	slice4 := []int{10, 20, 30}

	fmt.Println("----------newSlice2 index--------------")
	for index := 2; index < len(newSlice); index++ {
		fmt.Printf("index %d val %s\n", index, newSlice[index])
	}

	slice4[2] = 25

	//slice5 := []string{99: ""}

	//var slice6 []int

	//slice7 := make([]int, 0)

	//slice8 := []int{}

}

func arrayT() {
	fmt.Println("----------arrayT---------------")
	array := []int{10, 20, 30, 40, 50}
	array2 := [5]int{1: 10, 2: 20}

	array[2] = 35
	array2[3] = 88

	array3 := [5]*int{0: new(int), 1: new(int)}

	*array3[0] = 10
	*array3[1] = 20

	var array4 [5]string

	array5 := [5]string{"red", "blue", "green", "yellow", "pink"}

	array4 = array5

	for _, a := range array4 {
		fmt.Println(a)
	}
}

func mapT() {
	fmt.Println("----------mapT---------------")
	//dict := make(map[string]int)

	//dict2 := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}

	colors := map[string]string{}

	colors["Red"] = "#da1337"
	colors["Orange"] = "#e95a22"

	value, exists := colors["Orange"]
	if exists {
		fmt.Println(value)
	}

}

func main() {
	arrayT()

	sliceT()

	mapT()
}
