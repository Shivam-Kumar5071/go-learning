package main

import (
	"fmt"
	"io"
	"os"
)

func main(){
	
	f,err := os.Open("example.txt")

	if err != nil{
		//logging the error.
		panic(err)
	}
	
	defer f.Close()


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


	//for reading folders

	dir,err := os.Open("../")
	if err != nil{
		panic(err)
	}

	dirInfo,err := dir.ReadDir(-1)
	if err != nil{
		panic(err)
	}

	for _,dir := range dirInfo{
		fmt.Println("File name in the directory",dir.Name(),dir.IsDir())
	}

	defer dir.Close()


	//creation and writing of a file.
	fileNew,err := os.Create("example2.txt")
	
	if err != nil {
		panic(err)
	}

	defer fileNew.Close()


	fileNew.WriteString("Hi from Golang , and Hi developers ")

	fileNew.Truncate(0) // removes or clears file content.
	fileNew.Seek(0,0) //this will get the cursor to starting position.
	fileNew.WriteString("Welcome to Razorpay")//now this will overwrite the file.

	//for copying the contents of old file to new file we can use io.Copy()

	//source file - f -> example.txt

	dst,err := os.Create("example3.txt")

	if err != nil{
		panic(err)
	}

	defer dst.Close()
	f.Seek(0,0)
	copiedData ,err := io.Copy(dst,f)
	if err != nil{
		panic(err)
	}

	fmt.Println("Data copied successfully",copiedData)

	//deletion of a file

	// prr := os.Remove("example3.txt")
	// if prr != nil{
	// 	panic(prr)
	// }
	// fmt.Println("File deleted successfully")



}