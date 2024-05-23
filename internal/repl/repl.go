package repl

/**
 * This is a simple reusable REPL (Read-Eval-Print-Loop) package that can be used to create
 * interactive command-line interfaces. It is designed to be used in a modular way, where
 * the commands are defined in the calling package and passed to the REPL instance.

 * The REPL instance is created with a prompt, a configuration object, and a map of commands.

 * The configuration object is an interface{} that can be used to store any data that needs to be
 * shared between commands. The commands are defined as a map of strings to Command structs.

 * The Command struct contains the name of the command, a description, a flag to indicate if the
 * command should run hooks before and after the action, and the action function that will be
 * executed when the command is entered.

 * The REPL instance has a Start method that will start the REPL loop and wait for user input.
 * The input is parsed and matched to the available commands, and the corresponding action is
 * executed. The REPL loop can be stopped by calling the stop method.

 * The package also provides some built-in commands that will be automatically added to the REPL instance:
 * - help: display all available commands
 * - exit: exit the application
 * - clear: clear the terminal screen
 * - back: return to the previous menu (added if the root flag is false when creating the REPL instance)

 * The package also provides a way to define before and after hooks that will be executed before and after
 * each command action. This can be used to perform common tasks such as rendering the screen or updating
 * the configuration object.

 */

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
	Config          any
	commands        map[string]Command
	runBeforeAction func(*Repl, ...string) error
	runAfterAction  func(*Repl, ...string) error
}

func NewRepl(prompt string, config any, commands map[string]Command, root bool) *Repl {
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
	r.addClearCommand()

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

		if cmd.Hooks {
			err := r.runBeforeAction(r, args...)

			if err != nil {
				fmt.Println(err)
				continue
			}
		}

		err := cmd.Action(r, args...)

		if err != nil {
			fmt.Println(err)
		}

		if cmd.Hooks {
			err = r.runAfterAction(r, args...)

			if err != nil {
				fmt.Println(err)
			}
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

/**
 * Add the clear command to the REPL (Added by Default)
 * This is a built-in command that will clear the terminal screen
 */
func (r *Repl) addClearCommand() {
	r.AddCommand("clear", Command{
		Name:        "clear",
		Description: "clear the terminal screen",
		Hooks:       false,
		Action:      clearCommand,
	})
}

func cleanInput(input string) []string {
	return strings.Fields(strings.ToLower(input))
}
