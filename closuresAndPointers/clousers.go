package main

import "fmt"


func counter() func() int{
	var count int = 0
	return  func() int{ //this one is a closure function
		count++
		return count
	}

}

func main(){
	fmt.Println("Closures called")

	incre := counter()
	fmt.Println(incre())
	fmt.Println(incre())
}