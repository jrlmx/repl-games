package repl

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Command struct {
	Name        string
	Description string
	Hooks       bool
	Action      func(*Repl, ...string) error
}

type Repl struct {
	running         bool
	prompt          string
	Config          interface{}
	commands        map[string]Command
	runBeforeAction func(*Repl, ...string) error
	runAfterAction  func(*Repl, ...string) error
}

func NewRepl(prompt string, config interface{}, commands map[string]Command, root bool) *Repl {
	r := Repl{
		running:  false,
		prompt:   prompt,
		Config:   config,
		commands: commands,
		runBeforeAction: func(r *Repl, args ...string) error {
			return nil
		},
		runAfterAction: func(r *Repl, args ...string) error {
			return nil
		},
	}

	r.addHelpCommand()
	r.addExitCommand()

	if !root {
		r.addBackCommand()
	}

	return &r
}

func (r *Repl) AddCommand(Name string, cmd Command) {
	r.commands[Name] = cmd
}

func (r *Repl) SetRunBeforeAction(f func(*Repl, ...string) error) {
	r.runBeforeAction = f
}

func (r *Repl) SetRunAfterAction(f func(*Repl, ...string) error) {
	r.runAfterAction = f
}

func (r *Repl) stop() {
	r.running = false
}

func (r *Repl) Start() error {
	s := bufio.NewScanner(os.Stdin)

	if r.running {
		return errors.New("repl already running")
	}

	r.running = true

	for r.running {
		fmt.Print(r.prompt)

		s.Scan()

		input := cleanInput(s.Text())

		if len(input) <= 0 {
			continue
		}

		cmd, ok := r.commands[input[0]]

		if !ok {
			fmt.Println("invalid command")
			continue
		}

		args := []string{}

		if len(input) > 1 {
			args = input[1:]
		}

		err := r.runBeforeAction(r, args...)

		if err != nil {
			fmt.Println(err)
			continue
		}

		err = cmd.Action(r, args...)

		if err != nil {
			fmt.Println(err)
		}

		err = r.runAfterAction(r, args...)

		if err != nil {
			fmt.Println(err)
		}

		if !r.running {
			break
		}
	}

	return nil
}

/**
 * Add the help command to the REPL (Added by Default)
 * This is a built-in command that will display all available commands
 */
func (r *Repl) addHelpCommand() {
	r.AddCommand("help", Command{
		Name:        "help",
		Description: "display this help message",
		Hooks:       false,
		Action:      helpCommand,
	})
}

/**
 * Add the exit command to the REPL (Added by Default)
 * This is a built-in command that will exit the application
 */
func (r *Repl) addExitCommand() {
	r.AddCommand("exit", Command{
		Name:        "exit",
		Description: "exit the application",
		Hooks:       false,
		Action:      exitCommand,
	})
}

/**
 * Add the back command to the REPL (Added Optionally)
 * This is a built-in command that will break from the current repl loop
 */
func (r *Repl) addBackCommand() {
	r.AddCommand("back", Command{
		Name:        "back",
		Description: "return to the previous menu",
		Hooks:       false,
		Action:      backCommand,
	})
}

func cleanInput(input string) []string {
	return strings.Fields(strings.ToLower(input))
}
