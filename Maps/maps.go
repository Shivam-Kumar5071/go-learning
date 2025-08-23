package main

import "fmt"

func main(){
	mp := make(map[int]string)

    mp[1] = "Golang"
    mp[2] = "Backend Dev"
    mp[3] = "Online Course"
    fmt.Println(mp)
    // delete(mp, 2);
    //clear(mp) -> clears the map
    fmt.Println(mp)

    //another way to create a map

    m := map[string]int{"id":123,"quantity":3,"price":1224}
    fmt.Println(m)

    val,ok := m["id"]

    fmt.Println(val)
    if ok{
        fmt.Println("ok")
    }else{
        fmt.Println("not ok")
    }
    

    // fmt.Println(maps.Equal(m,mp))

}
