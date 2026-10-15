package console

import (
	"fmt"

	"github.com/AyuorusAguilar/VoidNavigator.git/internal/state"
)

type Screen struct {
	PrintScreen func(*state.State) string
	ActionsMap  map[string]func(*state.State)
}

var Screens = map[string]Screen{
	"Test": {
		PrintScreen: func(s *state.State) string {
			return fmt.Sprintf("Current config!: %v\n Use Y to increase delay\n Use H to decrease delay\n", s.Cfg)
		},
		ActionsMap: map[string]func(*state.State){
			"Y": func(s *state.State) {
				s.Cfg.Delay = s.Cfg.Delay + 1
			},
			"H": func(s *state.State) {
				s.Cfg.Delay = s.Cfg.Delay - 1
			},
		},
	},
}

func StartConsole(s *state.State, keymonchan chan string) {
	fmt.Printf("Starting console...\n")
	for {
		cls()
		fmt.Print(Screens[s.CurrentScreen].PrintScreen(s))
		fmt.Print("Esperando tecla\n")
		action := <- keymonchan
		fmt.Printf("La tecla fue: %s\n", action)
		f, ok := Screens[s.CurrentScreen].ActionsMap[action]
		fmt.Printf("Ok is %v\n", ok)
		if ok { f(s) }
	}
}

func cls() {fmt.Print("\033[2J\033[H")}