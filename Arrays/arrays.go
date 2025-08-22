package main

import "fmt"

func main(){

    //arrays - for fixed size we use array,memory optimization,constant time access
    var nums [4]int
    nums[0] = 21
    fmt.Println(nums)
    fmt.Println(len(nums))

    arr := [3]int{10,20,30}
    fmt.Println(arr)

    //2d arrays
    arr2D := [2][2]int{{23,10},{90,87}}
    fmt.Println(arr2D)

}
