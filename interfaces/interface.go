package main

import "fmt"

type payment struct{
	gateway stripe
	gateway2 razorpay
}

func (p payment) makePayment(amount float32){
	// razorpayPaymentGw := razorpay{}
	// razorpayPaymentGw.pay(amount)
	// stripePaymenetGw := stripe{}
	// stripePaymenetGw.pay(2000)
	p.gateway.pay(amount)
	p.gateway2.pay(amount)

}

type razorpay struct{ 
}

func (r razorpay) pay(amount float32){
	fmt.Println("Making payment using razorpay",amount)
}

type stripe struct{}

func (s stripe) pay(amount float32){
	fmt.Println("making payment using stripe",amount)
}

func main(){
	stripePaymentGateway := stripe{}
	razorpayPaymentGw := razorpay{}
	newPayment := payment{
		gateway: stripePaymentGateway,
		gateway2: razorpayPaymentGw,
	}
	newPayment.makePayment(100)
}
