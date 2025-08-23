package main

import (
	"fmt"
	"time"
)


type customer struct{
	name string 
	phone string
}


type payment struct{

	id int
	amount float32
	status string
	receivedAt time.Time
	customer
}

func (p *payment) setPayment(amount float32){
	p.amount = amount
}

func (p *payment) getPayment()float32{
	return p.amount
}

func newPayment(id int ,amount float32,status string) *payment{
	
	newCust := customer{
		name: "Mark",
		phone: "0123456789",
	}
	
	pay1 := payment{
		id : id,
		amount : amount,
		status: status,
		customer: newCust,
	}
	return &pay1
}



func main(){

	pay1 := newPayment(123,4000,"Paid")
	pay1.receivedAt = time.Now()
	pay1.status = "paid"

	pay1.setPayment(3000)
	fmt.Println(pay1.getPayment())
	pay1.customer.name = "John Doe"

	fmt.Println(pay1)

}