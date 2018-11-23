package main

import (
	"fmt"
	"gopkg.in/AlecAivazis/survey.v1"
	"log"
)

func main() {
	toolans := struct {
		Tool string `survey:"tool-select"`
	}{}
	fmt.Println("*****************")
	fmt.Println("* File-Maker-GO *")
	fmt.Println("*****************")
	fmt.Println("Note: For Sign & Encrypt, LZ Tools and others are available as different programs.")
	err := survey.Ask(MainMenu, &toolans)
	if err != nil {
		log.Panic(err)
	}

}
