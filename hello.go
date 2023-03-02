package main

import  "fmt"
import "rsc.io/quote"

func hello() {
	fmt.Println("Hello World!")
}

func getQuote() {
	fmt.Println(quote.Go())
}

func main() {
	hello()
	getQuote()
}