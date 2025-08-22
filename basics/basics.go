package main

import "fmt"

func main(){
    fmt.Println("Hello, World")

    var name string = "Shivam"
    // fmt.Println(name)

    age := 16
    fmt.Println(name,age)

    if age >= 18{
        fmt.Println("Eligible")
    }else if age >=12{
        fmt.Println("Teenager")
    }else{
        fmt.Println("Not eligible")
    }

    //while loop
    // i:=1
    // for i<=10{
    //     fmt.Println(i)
    //     i = i+1
    // }

    // for loop
    for j:=10;j<=20;j++{
        fmt.Println(j)
    }
    //go does not have ternary operator
}
