package main

import (
	"log"

	"github.com/kc1116/go-speak2me/translate"
)

func main() {
	text, err := translate.TransText("es", "Hello World")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(text)

}
