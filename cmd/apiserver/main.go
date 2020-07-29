package main

import (
	"alphatest/internal/apiserver"
	"fmt"
	"log"
)

func main() {
	config := apiserver.NewConfig()
	fmt.Println("Starting server")
	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
