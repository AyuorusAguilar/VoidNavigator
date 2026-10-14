package main

import (
	"fmt"

	keymonitor "github.com/AyuorusAguilar/VoidNavigator.git/internal/KeyMonitor"
)

func main() {
	fmt.Printf("Hello world!")
	var port chan string = make(chan string)
	go keymonitor.KeyMon(port)
	for {
		value := <- port
		fmt.Printf("The key \"%s\" has been pressed\n", value)
	}
}