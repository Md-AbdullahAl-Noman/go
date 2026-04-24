package packageone

import "fmt"
var privateVar="I am private"
var PublicVar="I am public or exported"//the its PascalCase that makes it public and accessible from other packages

//same for functions if we have a function with a name that starts with an uppercase letter it is exported and can be accessed from other packages but if it starts with a lowercase letter it is unexported and can only be accessed within the package
func PublicFunc(){
	fmt.Println("This is a public function")
}

func privateFunc(){
	fmt.Println("This is a private function")
}