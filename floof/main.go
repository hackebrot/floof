package main

import (
	"fmt"
	"log"
)

func main() {
	userConfig, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(userConfig.Gist.Username)
	fmt.Println(userConfig.Gist.Token)
}
