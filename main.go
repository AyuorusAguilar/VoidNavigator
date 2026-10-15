package main

import (
	"fmt"
	"os"

	"github.com/AyuorusAguilar/VoidNavigator.git/internal/console"
	"github.com/AyuorusAguilar/VoidNavigator.git/internal/keymonitor"
	"github.com/AyuorusAguilar/VoidNavigator.git/internal/state"
)

func main() {
	// State
	s, err := state.NewState()
	if err != nil {
		fmt.Printf("Error creating the state!: %v", err)
		os.Exit(1)
	}
	// Keymon
	var port chan string = make(chan string)
	go keymonitor.KeyMon(port)
	// Console
	console.StartConsole(&s, port)
}