package main

import (
	"fmt"
	"sync"
)

//GoRoutines are light weight thread and used for multi threading and used to run things concurrently.

func items(val int,w *sync.WaitGroup){
	//defer is keyword used as finally in other lang as defer will make the functions to run at last. 
	defer w.Done()
	fmt.Println("Dealing with the task ",val)
}

func main(){
	fmt.Println("Go Routines")
	var wg sync.WaitGroup

	for i:=1;i<=10;i++{
		wg.Add(1)
		go items(i,&wg)
	}

	wg.Wait()
	
}
