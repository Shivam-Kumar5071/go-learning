package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Channels is a kind of pipe used to send and receive data

func processNum(numChan chan int){
	for num := range numChan{
	fmt.Println("Number is processed",num)
	time.Sleep(time.Second*1)
	}
}

func  main()  {
	numChan := make(chan int)

	go processNum(numChan)

	// numChan <- 5
	for {
		numChan <- rand.Intn(100)
	}
	// time.Sleep(time.Second*1)

	
}