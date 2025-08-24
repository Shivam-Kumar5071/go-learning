package main

import(
	"fmt"
	"slices"
)

func main(){

    //slices - imp - dynamic array

    //uninitialised slice is nil

    // var brr[] int
    // fmt.Println(brr)

    // var brr = make([]int, 0,5)
    brr := []int{}
    brr = append(brr, 23)
    brr = append(brr, 21)
    brr = append(brr, 40)
    brr = append(brr, 2)
    fmt.Println(cap(brr))
    fmt.Println(brr)


    // copy of a slice
    var crr = make([]int,0,5)
    crr = append(crr, 2)
    var drr = make([]int,len(crr))

    copy(drr,crr)

    fmt.Println(drr,crr)

    //slice operator

    var slicer = []int{1,2,3,4,5}
    fmt.Println(slicer[:3])
    fmt.Println(slicer[1:])
    fmt.Println(slicer[2:4])

    //slices package

    var nums1 = []int{1,2,3}
    var nums2 = []int{1,2,4}

    fmt.Println(slices.Equal(nums1,nums2))
    //if the key is not present in map then it would return 0 value

}