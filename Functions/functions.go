package main

import (
	"fmt"
)

func add(a int,b int) int{
    return a + b
}

func getLanguages()(string,string,int){
    return "Mern","Django",78
}

func sumation(nums ...int) int { //interface is used when we need all the data types 
	total := 0
	for _, num := range nums {
		total = total + num
	}
	return total
}

func main(){


    //functions in Go
    res := add(10,40)
    fmt.Println(res)

    lang1,lang2 , int1 := getLanguages()
    fmt.Println(lang1,lang2,int1)

    //variadic function -> where we can pass multiple parameters means we can pass n number of functions
    //for eg :- Println()
    tot := sumation(1,2,3,4,5)
    fmt.Println(tot)


}

