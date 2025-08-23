package main

import "fmt"

func changeNum(num *int){
	*num = 5
	fmt.Println("in function change ",*num)
}

func main(){
	num := 12
	changeNum(&num)
	fmt.Println("After change num: ",num)
}
