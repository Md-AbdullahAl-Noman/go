package main

import (
	"fmt"
	"scope-example-app/packageone"
)

var myVar="Main package variable"
func main(){

	var blockVar="Block variable inside main function"
	fmt.Println(myVar)
	fmt.Println(blockVar)
	fmt.Println(packageone.PackageVar)
	packageone.PrintMe()

}
