package main

import "fmt"

//normal function

// func PrintItems(arr []int){
// 	for i ,val := range arr{
// 		fmt.Println(i,val)
// 	}
// }

//generic function

func PrintItems[T int | string | bool](arr []T){
	for _ ,val := range arr{
		fmt.Println(val)
	}
}

func main(){

	nums := []int{10,20,30,40,50}
	names := []string{"Golang","C++","MERN"}
	boolean := []bool{true,false,true,false}
	PrintItems(names)
	PrintItems(nums)
	PrintItems(boolean)

}