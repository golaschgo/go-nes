package main

import (
	"fmt"

	"github.com/golaschgo/go-nes/command"
	"github.com/golaschgo/go-nes/games"
)

func main() {

	games.Init()

	for {
		command.ClearTerminal()
		opt := command.ShowMenu()
		switch opt {
		case "1":
			games.Search()
		case "2":
			games.ShowFavorites()
		case "3":
			games.SaveFavoritesFile()
			fmt.Println("Guardando Favoritos antes de salir...")
			fmt.Println("Cerrando la aplicación")
			return
		}
	}

}
