package main

import (
	"fmt"
	"sync"
)

//Race condition - When multiple processes modify the same the resources , when changes done by one processes can be overridden by another processes.

//Mutex - a synchronization tool used in concurrent programming to prevent race conditions.


type post struct{
	view int 
	mu sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup){
	defer func ()  {
		p.mu.Unlock()	
		wg.Done()
	}()

	p.mu.Lock()
	p.view += 1
}

func main(){
	fmt.Println("Mutex ")

	var wg sync.WaitGroup

	myPost := post{view:0}
	for i:=0;i<100;i++{
		wg.Add(1)
		go myPost.inc(&wg)

	}

	wg.Wait()
	fmt.Println(myPost.view)

}