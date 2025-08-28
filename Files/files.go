package main

import (
	"fmt"
	"os"
)

func main(){
	
	f,err := os.Open("example.txt")

	if err != nil{
		//logging the error.
		panic(err)
	}

	fileInfo , err := f.Stat()//this will all the statical data about file.

	if err != nil{
		//this means that error is not nil , or error encountered then you have to log the error.
		panic(err)
	}

	fmt.Println(fileInfo.IsDir())
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Mode())
	fmt.Println(fileInfo.Size())

	
	// buff := make([]byte,fileInfo.Size())
	// d,err := f.Read(buff)
	// if err != nil{
	// 	panic(err)
	// }

	// for i:=0 ; i<len(buff);i++{
	// 	fmt.Println("data : ",d,string(buff[i]))
	// }

	buf := make([]byte,fileInfo.Size()) //making of a buffer which is of array of bytes which stores information about the file i.e what is written like this.
	c,err := f.Read(buf)
	if err != nil {
		panic(err)
	}

	for i:=0;i<len(buf);i++{
		fmt.Println("data ",c,string(buf[i]))
	}


	//another way to read file and shortest way.
	data , err := os.ReadFile("example.txt")
	
	if err != nil{
		panic(err)
	}
	fmt.Println(string(data))

	defer f.Close()



}