package state

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type State struct {
	CurrentLevel int
	RunType      string
	Cfg          config
	Sequences    map[string]level
}

type config struct {
	Delay           float32 `json:"delay"`
	Press_duration  float32 `json:"press_duration"`
	Ignore_branches bool    `json:"ignore_branches"`
}

type level struct {
	MainSequence string   `json:"Main Sequence"`
	Branches     []branch `json:"Branches"`
}

type branch struct {
	Info           string `json:"Info / Branch Condition"`
	BranchSequence string `json:"Branch Sequence"`
	ExitTo         string `json:"Exit to"`
}

func NewState() (State, error){
	// config
	configDir := filepath.Join("resources","config.json")
	var conf = config{
		Delay: 0.2,
		Press_duration: 0.2,
		Ignore_branches: false,
	}
	err := OpenParseJSONFile(configDir, &conf)
	if err != nil {
		return State{}, err
	}

	// sequences

	s := State{
		CurrentLevel: 0,
		RunType: "",
		Cfg: conf,
		Sequences: make(map[string]level),
	}
	return s, nil
}

func OpenParseJSONFile(dir string, v any) error {
	file, err := os.OpenFile(dir, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("Error opening file: %v", err)
	}
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("Error reading file: %v", err)
	}
	json.Unmarshal(fileBytes, v)
	return nil
}

func LoadState()  {
	
}