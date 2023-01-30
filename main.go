package main

import (
	"github.com/golaschgo/go-nes/command"
	"github.com/golaschgo/go-nes/games"
)

func main() {

	for {
		command.ClearTerminal()
		opt := command.ShowMenu()
		switch opt {
		case "1":
			games.Search()
		case "2":
			games.ShowFavorites()
		case "3":
			println("Salir")
			return
		}
	}

}
