package main

import (
	"fmt"
	"myapp/packageone"
)


var one="One"//global variable scoped to the package main
func main(){
	
    var one="This one variable is block level variable and from main function scope that shadows the global variable one"// block level variable that shadows the global variable one
	fmt.Println(one)
	myFunc()
	newString:=packageone.PublicVar//accessing the public variable from packageone
	fmt.Println(newString)
	
	// secondSring:=packageone.PrivateVar// trying to access the private variable from packageone will result in an error because it is not exported and only accessible within the packageone
	// fmt.Println(secondSring)
}

func myFunc(){
	// var one="Number one"// function scope variable that shadows the global variable one
	fmt.Println(one)
}