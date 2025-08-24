package main

import "fmt"

func main(){
	//Range -> iterarting over data structure
    //range on slice
    sumArr := []int{1,2,3,4,5}
    sum := 0
    for i, num := range sumArr{
        fmt.Println(i)
        sum = sum + num
    }
    fmt.Println("Sum is : ",sum)

    //range on map

    rangeMp := map[string]string{"fname":"John","mname":"Bros","lname":"Doe"}
    for k,v := range rangeMp{
        fmt.Println(k,"->",v)
    }

    //range on string
    var company string = "Razorpay"
    for i,c:= range company{
        fmt.Println(i,c)
    }

}
