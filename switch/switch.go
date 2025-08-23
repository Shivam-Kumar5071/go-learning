package main

import (
	"fmt"
	"time"
)

func main(){
    //switch statement - break statement is not in go as it is inbuilt.
    k:=30
    switch k{
        case 1:
            fmt.Println("one")
        case 2:
            fmt.Println("two")
        case 3:
            fmt.Println("three")
        default:
            fmt.Println("final statement")
    }

    //multiple switch conditions
    switch time.Now().Weekday(){
        case time.Sunday,time.Saturday:
            fmt.Println("Weekend")
        default:
            fmt.Println("Work day")
        
        
    }

    typeFinder := func(i interface {}){
        switch t:= i.(type){
            case int:
                fmt.Println("Integer")
            case string:
                fmt.Println("String")
            case bool:
                fmt.Println("Boolean")
            default:
                fmt.Println("other data type",t)
        }
    }

    typeFinder(true)

}
