package main

import (
	"fmt"

	"github.com/Shivam-Kumar5071/go-learning/auth"
	"github.com/Shivam-Kumar5071/go-learning/user"
	"github.com/fatih/color"
)

//go mod tidy --> it is used as if any dependencies is not installed but you are using it then it is installed automatically and if there is any dependencies which is indirect then it will make it direct if u are using it.

//go get --> it is used to install anything.

func main(){
	auth.LoginWithCredentials("Shivam","Razorpay123@")
	session := auth.GetSession()
	fmt.Println(session)

	user := user.User{
		Name : "Shivam Kumar",
		Email : "shivamkumar@gmail.com",
		Username : "shivamkumar",

	}
	color.Red("User name : %s",user.Name)
	color.Green("User Email: %s",user.Email)
	color.Yellow("User username: %s",user.Username)


}