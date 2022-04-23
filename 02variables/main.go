package main

import "fmt"

const LoginToken string = "ghabdfkjan" // public variable because of cap letter


func main() {
	var username string = "hitesh"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n", username)

	var trueorfalse bool = false
	fmt.Println(trueorfalse)
	fmt.Printf("Variable is of type : %T \n", trueorfalse)

	// implicit type
	var website = "learncodeonline.in"
	fmt.Println(website)

	// no var style
	numberOfUser := 30000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
}