package main

import "fmt"

type bot interface{
	getGreeting() string
}

type englishBoat struct{}
type spainishBot struct{}

func main() {
	eb:=englishBoat{}
	sb:=spainishBot{}


	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot){
	fmt.Println(b.getGreeting())
}

// func printGreeting(eb englishBoat) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spainishBot){
// 	fmt.Println(sb.getGreeting())
// }

func (englishBoat) getGreeting() string {
	//very custome logic for generating english logic
	return "hi there"
}

func (spainishBot) getGreeting() string {
	return "hola!"
}