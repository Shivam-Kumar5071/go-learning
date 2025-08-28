package auth

import "fmt"


//in this if you make functions first letter with capital then it is globally accessible and if you make functions with first letter small then it is only in the current directory.
func LoginWithCredentials(username string,password string){
	fmt.Println("User login with credentials with : ",username,password)

} 