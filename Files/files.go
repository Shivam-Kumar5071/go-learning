package main

import (
	"fmt"
	"os"
)

func main(){
	
	f,err := os.Open("example.txt")

	if err != nil{
		panic(err)
	}

	fileInfo , err := f.Stat()

	if err != nil{
		//this means that error is not nil , or error encountered then you have to log the error.
		panic(err)
	}

	fmt.Println(fileInfo.IsDir())
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Mode())
	fmt.Println(fileInfo.Size())



}