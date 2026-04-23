package main

import "fmt"

func main(){
//var declaration by var and must use type
// var whatToSay string;
// whatToSay="Hey var";
//shorthand of variable declaration and assignment ;here type is based upon what is the type of assigned value
whatToSay :="Hey new style"
printHello(whatToSay);
}

func printHello(var1 string){
	fmt.Println(var1);
}