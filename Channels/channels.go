package main

import (
	"fmt"
	"time"
	// "math/rand"
	// "time"
)

// Channels is a kind of pipe used to send and receive data

// func processNum(numChan chan int){
// 	for num := range numChan{
// 	fmt.Println("Number is processed",num)
// 	time.Sleep(time.Second*1)
// 	}
// }

// func sum(result chan int,num1 ,num2 int){
// 	numResult := num1 + num2
// 	result <- numResult
// }


func sendingMail(emailChan chan string,done chan bool){
	defer func(){ done <- true} ()

	for email := range emailChan{
		fmt.Println("sending email to : ",email)
		time.Sleep(time.Second)
	}
}

func  main()  {
	// numChan := make(chan int)

	// go processNum(numChan)

	// // numChan <- 5
	// for {
	// 	numChan <- rand.Intn(100)
	// }
	// time.Sleep(time.Second*1)

	//unbuffered channel - it focus on sender to wait untill the receiver is able to receive a message.it has capacity of zero.

	// sumRes := make(chan int)
	// go sum(sumRes,4,5)

	// res := <-sumRes
	// fmt.Println(res)


	//buffered channel - it doesn't wait for receiver to receive whereas it allows sender to send data without blocking because it has some capacity.

	// emailChan := make(chan string,100)

	// emailChan <- ("ex1@gmail.com")
	// emailChan <- ("ex2@gmail.com")

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)

	emailChan := make(chan string,20)
	done := make(chan bool)

	go sendingMail(emailChan,done)
	for i:=0; i<5; i++{
		emailChan <- fmt.Sprintf("%d@gmail.com",i)
	}

	fmt.Println("Done by sending the mail") 
	close(emailChan) //important closing the buffered channel.
	<-done


	chan1 := make(chan int)
	chan2 := make(chan string)

	go func(){
		chan1 <- 10
	}()

	go func(){
		chan2 <- "ping"
	}()

	for i:=0;i<2;i++{
		select{
		case chanVal1 := <- chan1:
			fmt.Println("Received data from chan1 ",chanVal1)
		case chanVal2 := <- chan2:
			fmt.Println("Received data from chan2 ",chanVal2)
		}

	}






	
}