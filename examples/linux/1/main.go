package main

import (
	who "github.com/kyaxcorp/go-os-who"
	"log"
)

func main() {
	users, err := who.List()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(users)
}
