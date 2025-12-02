package main

import (
	"fmt"

	"github.com/Glebegor/hackathon-go-templates/bootstrap"
)

func main() {
	fmt.Println("Start of the service")

	config, err := bootstrap.NewConfig("./config/config.json", "dev")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Config loaded: %+v\n", config)
}
