package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt= "and Press Enter when ready."

func main(){
	//seeding random number generator	so that we get different random numbers each time we run the program as computer follow formulas not randomness
	rand.New(rand.NewSource(time.Now().UnixNano()))//now everytime we run the program we get different random numbers
	var firstNumber=rand.Intn(9)+2
	var secondNumber=rand.Intn(9)+2
	var subtraction=rand.Intn(9)+2

	var answer=firstNumber*secondNumber-subtraction
	playThegame(firstNumber,secondNumber,subtraction,answer)
	

}


func playThegame(firstNumber,secondNumber,subtraction,answer int){
reader:=bufio.NewReader(os.Stdin)
	//display a welcome /instruction
	fmt.Println("Guess the number game")
	fmt.Println("---------------------")
	fmt.Println("")
	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')
	//take them through the games
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber)
	reader.ReadString('\n')

	fmt.Println("Now divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')
	fmt.Println("Now subtract",subtraction)
	//give them answer
	
	fmt.Println("The answer is", answer)
}