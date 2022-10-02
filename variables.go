package main

import "fmt"

///declaring constant

const LoginToken string = "ghdhjdbd  hrjgr" ///L capital L made it a public variable

var jwtToken = 30000

func main() {
	fmt.Println("Variables")
	var username string = "Saksham"
	fmt.Println(username)
	fmt.Printf("It is of type: %T \n", username)

	var isLogged bool = true
	fmt.Println(isLogged)
	fmt.Printf("Type: %T ", isLogged)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Type: %T ", smallValue)

	var smallFloat float32 = 255.8943784738565
	fmt.Println(smallFloat)
	fmt.Printf("Type: %T ", smallFloat)

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type: %T \n", anotherVariable)

	///implicit type
	var website = "learn"
	fmt.Println(website)

	//no var style
	numberOfUser := 300000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)

}
