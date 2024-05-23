# A complilation of CLI-Games

## Project Description

A complilation of Cli-Games written in Golang

Todo:

- [] Seperate the Cli implementation from the Minesweeper package
- [] Clean up the Render function
- [] Code more games (Battleship, Tic Tac Toe, etc...)

Ideas & Stretch Goals:

- [] Build out a seperate web api and serve the games in a web application

## 

## Minesweeper

The minesweeper package provides a simple implementation of the classic game Minesweeper.

The package provides a Game struct that represents the state of a Minesweeper game, as well as
functions to create a new game, hit a square, flag a square, and draw the game board.

The Game struct contains the minefield, which is a 2D array of Squares, and the gamestate, which
indicates whether the game is in progress, won, or lost.

The package also provides a StartCommand function that creates a new REPL instance and starts the
Minesweeper game. The REPL instance is configured with a set of commands that allow the player to
interact with the game, such as starting a new game, hitting a square, flagging a square, and drawing
the game board.

The package uses the repl package to implement the REPL functionality, which allows the player to
enter commands and interact with the game through the command-line interface.


## The Internal/Repl Package

This is a simple reusable REPL (Read-Eval-Print-Loop) package that can be used to create
interactive command-line interfaces. It is designed to be used in a modular way, where
the commands are defined in the calling package and passed to the REPL instance.

The REPL instance is created with a prompt, a configuration object, and a map of commands.

The configuration object is an interface\{\} that can be used to store any data that needs to be
shared between commands. The commands are defined as a map of strings to Command structs.

The Command struct contains the name of the command, a description, a flag to indicate if the
command should run hooks before and after the action, and the action function that will be
executed when the command is entered.

The REPL instance has a Start method that will start the REPL loop and wait for user input.
The input is parsed and matched to the available commands, and the corresponding action is
executed. The REPL loop can be stopped by calling the stop method.

The package also provides some built-in commands that will be automatically added to the REPL instance:
- help: display all available commands
- exit: exit the application
- clear: clear the terminal screen
- back: return to the previous menu (added if the root flag is false when creating the REPL instance)

The package also provides a way to define before and after hooks that will be executed before and after
each command action. This can be used to perform common tasks such as rendering the screen or updating
the configuration object.