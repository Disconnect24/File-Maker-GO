package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("*****************")
	fmt.Println("* File-Maker-GO *")
	fmt.Println("*****************")

	fmt.Println("What would you like to do today? \n" +
		"1. Sign and Encrypt")
	//reading an integer
	var choice int
	_, err := fmt.Scan(&choice)
	if err != nil {
		log.Panic("oops", err)
	}
	switch choice {
	case 1:
		fmt.Println("placeholder")
	}
}
