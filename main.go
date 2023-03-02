package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Rule struct {
	CurrentState string `json:"currentState"`
	TapeSymbol   string `json:"tapeSymbol"`
	NextState    string `json:"nextState"`
	WriteSymbol  string `json:"writeSymbol"`
	Move         string `json:"move"`
}

type Machine struct {
	Tape         []string `json:"tape"`
	CurrentState string   `json:"initialState"`
	Rules        []Rule   `json:"rules"`
	Head         int      `json:"head"`
	HaltState    string   `json:"haltState"`
}

func NewMachine(tape []string, currentState string, rules []Rule, head int, haltState string) *Machine {
	return &Machine{
		Tape:         tape,
		CurrentState: currentState,
		Rules:        rules,
		Head:         head,
		HaltState:    haltState,
	}
}

func (m *Machine) PrintTape() {
	fmt.Printf("|")
	for i := 0; i < len(m.Tape); i++ {
		fmt.Printf("%s|", m.Tape[i])
	}
	fmt.Printf("\n%s^%s\n", strings.Repeat(" ", m.Head*2+1), m.CurrentState)
}

func (m *Machine) Run() error {
	if m.Head == -1 {
		m.Head = len(m.Tape) / 2
	}
	var steps uint32 = 0
	for {
		m.PrintTape()
		for _, v := range m.Rules {
			if v.CurrentState == m.CurrentState && v.TapeSymbol == m.Tape[m.Head] {
				m.CurrentState = v.NextState
				m.Tape[m.Head] = v.WriteSymbol
				switch v.Move {
				case "L":
					m.Head--
				case "R":
					m.Head++
				default:
					return errors.New("Unknown movement")
				}
				break
			}
		}
		steps++
		if m.CurrentState == m.HaltState {
			break
		}
	}
	fmt.Println("\nThe turing machine has halted\nFinal tape config is:")
	m.PrintTape()
	fmt.Printf("steps=%d", steps)
	return nil
}

type Config struct {
	Tape         string `json:"tape"`
	InitialState string `json:"initialState"`
	Rules        []Rule `json:"rules"`
	HeadPosition int    `json:"headPosition,omitempty"`
	HaltState    string `json:"haltState"`
}

func main() {
	filename := os.Args[1] + ".tm.json"
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Notice that turing machine program file name have to end with .tm.json")
		fmt.Println("But you don't have to include .tm.json to an argument")
		os.Exit(1)
	}
	var machineConfig Config
	err = json.Unmarshal(file, &machineConfig)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	machine := NewMachine(
		strings.Split(machineConfig.Tape, ""),
		machineConfig.InitialState,
		machineConfig.Rules,
		machineConfig.HeadPosition,
		machineConfig.HaltState,
	)
	fmt.Printf("\nTURING MACHINE \nProgram: %s\n\n", filename)
	if err := machine.Run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
