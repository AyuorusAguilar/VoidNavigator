package main

import (
	"fmt"
	"os"

	"github.com/AyuorusAguilar/VoidNavigator.git/internal/keymonitor"
	"github.com/AyuorusAguilar/VoidNavigator.git/internal/state"
)

func main() {
	fmt.Printf("Hello world!")
	var port chan string = make(chan string)
	s, err := state.NewState()
	if err != nil {
		fmt.Printf("Error creating the state!: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Configuration: %v", s.Cfg)
	
	
	go keymonitor.KeyMon(port)
	for {
		value := <- port
		fmt.Printf("The key \"%s\" has been pressed\n", value)
	}
}