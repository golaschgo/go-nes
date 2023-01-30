package games

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/golaschgo/go-nes/command"
)

// GamesDB := []GamesDBElement{}
var GamesDBSlice []GamesDBElement
var MyFavorites []GamesDBElement

func Search() {
	ReadDBFile()
	command.Separator()
	for {
		fmt.Println("Escribe un nombre para buscar / o ", color.YellowString("exit"), "para salir")
		suboption := command.Prompt("->")
		fmt.Println("Buscando :", color.GreenString(suboption))
		SearchGamesDB(suboption)
		if suboption == "exit" {
			return
		}
	}
}

func ShowFavorites() {
	showFavoriteGameList()
	// seleccionar para eliminar
}

func ReadDBFile() {
	data, err := ioutil.ReadFile("./nes_games_db.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal([]byte(data), &GamesDBSlice)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func SearchGamesDB(s string) {
	var searchResult []GamesDBElement

	fmt.Println("Buscando por ", s)

	for _, game := range GamesDBSlice {
		if strings.Contains(game.Title, s) {
			searchResult = append(searchResult, game)
		}
	}
	if len(searchResult) > 0 {
		PrintSearchGamesDBList(searchResult)
		SelectYourFavoriteGame(searchResult)
	}
}

func PrintSearchGamesDBList(gameList []GamesDBElement) {
	for i, game := range gameList {
		// fmt.Println(color.GreenString("["+strconv.Itoa(i+1)+"]",game.Title)
		fmt.Println("[", strconv.Itoa(i+1), "]", game.Title)
	}
	fmt.Println("[ 0 ] -- Para volver")
}

func SelectYourFavoriteGame(gameList []GamesDBElement) {
	opt := command.Prompt("Escribe el número de tu juego favorito ->")
	if opt != "0" {
		gameIndex, _ := strconv.Atoi(opt)
		if !isFavoriteGame(gameList[gameIndex-1].Title) {
			addFavoriteGame(gameList[gameIndex-1])
		}
	}
}

func addFavoriteGame(game GamesDBElement) {
	MyFavorites = append(MyFavorites, game)
}

func removeFavoriteGame(game GamesDBElement) {
	// remove
}

func showFavoriteGameList() {
	command.ClearTerminal()
	fmt.Println("Tus favoritos son:")
	command.Separator2()
	for i, game := range MyFavorites {
		fmt.Println("[", strconv.Itoa(i+1), "]", game.Title)
	}
	fmt.Println("[ 0 ] -- Para volver")
}

func isFavoriteGame(gameTitle string) bool {
	for _, game := range MyFavorites {
		if strings.Contains(game.Title, gameTitle) {
			return true
		}
	}
	return false
}
