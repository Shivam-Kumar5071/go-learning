package main

import (
	"fmt"
	"time"
)

type order struct{
	id int
	amount float32
	status string
	createdAt time.Time

}

func (o *order) changeStatus(status string){
	o.status = status
}

func (o *order) getID() int{
	return o.id
}

//constructor -> called with the help of a function
func constructor(id int,amount float32,status string)* order{
	myOrder := order{
		id : id,
		amount: amount,
		status:status,
	}
	return &myOrder
}

func main(){

	obj := order{
		id:123,
		amount:50.45,
		status: "ordered",
	}
	obj2 := order{
		id:456,
		amount:900.09,
		status: "unpaid",
		createdAt: time.Now(),
	}
	obj2.status = "paid"
	obj.createdAt = time.Now()
	obj2.changeStatus("confirmed")
	fmt.Println("Obj 1",obj)
	fmt.Println("Obj 2",obj2)

	fmt.Println(obj2.getID())


	obj3 := constructor(100,89.12,"paid")
	fmt.Println(obj3)
}