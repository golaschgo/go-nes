package games

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/golaschgo/go-nes/command"
)

// GamesDB := []GamesDBElement{}
var GamesDBSlice []GamesDBElement
var MyFavorites []GamesDBElement

var JSONFiles = map[string]string{
	"nes":       "nes_games_db.json",
	"favorites": "myfavorite_games_db.json",
}

func Init() {
	ReadJSONFileNes("./" + JSONFiles["nes"])
	ReadJSONFileMyFavorites("./" + JSONFiles["favorites"])
}

func Search() {
	// ReadDBFile()

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

func ReadJSONFileNes(file string) {
	data, err := ioutil.ReadFile(file)
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

func ReadJSONFileMyFavorites(file string) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal([]byte(data), &MyFavorites)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ReadJSONFile(file string, sliceDestiny []GamesDBElement) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal([]byte(data), &sliceDestiny)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func SaveFavoritesFile() {
	data, err := json.Marshal(MyFavorites)
	if err != nil {
		fmt.Println(err)
	}
	f, err := os.Create("./" + JSONFiles["favorites"])
	defer f.Close()
	//write directly into file
	f.Write([]byte(data))
	f.Close()
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
		SelectYourFavoriteGame(searchResult, len(searchResult))
	}
}

func PrintSearchGamesDBList(gameList []GamesDBElement) {
	for i, game := range gameList {
		// fmt.Println(color.GreenString("["+strconv.Itoa(i+1)+"]",game.Title)
		fmt.Println("[", strconv.Itoa(i+1), "]", game.Title)
	}
	fmt.Println("[ 0 ] -- Para volver")
}

func SelectYourFavoriteGame(gameList []GamesDBElement, resultsQty int) {
	opt := command.Prompt("Escribe el número de tu juego favorito ->")
	optInt, _ := strconv.Atoi(opt)
	if opt != "0" && optInt <= resultsQty {
		gameIndex, _ := strconv.Atoi(opt)
		if !isFavoriteGame(gameList[gameIndex-1].Title) {
			addFavoriteGame(gameList[gameIndex-1])
		}
	} else {
		return
	}
}

func SelectYourFavoriteGame2Delete(gameList []GamesDBElement) {
	opt := command.Prompt("Escribe el número del juego a eliminar de tu lista / 0 Para volver ->")
	if opt != "0" {
		gameIndex, _ := strconv.Atoi(opt)
		removeFavoriteGame(gameIndex)
	} else {
		return
	}
}

func addFavoriteGame(game GamesDBElement) {
	MyFavorites = append(MyFavorites, game)
}

func removeFavoriteGame(i int) {
	i = i - 1
	fmt.Println("Juego a eliminar", MyFavorites[i])
	if i > 0 {
		MyFavorites = append(MyFavorites[:i], MyFavorites[i+1:]...)
	} else {
		MyFavorites = append(MyFavorites[:i], MyFavorites[i+1:]...)
	}

}

func showFavoriteGameList() {
	command.ClearTerminal()
	fmt.Println("Tus favoritos son:")
	command.Separator2()
	for i, game := range MyFavorites {
		fmt.Println("[", strconv.Itoa(i+1), "]", game.Title)
	}
	fmt.Println("[ 0 ] -- Para volver")
	if len(MyFavorites) > 0 {
		SelectYourFavoriteGame2Delete(MyFavorites)
	}
}

func isFavoriteGame(gameTitle string) bool {
	for _, game := range MyFavorites {
		if strings.Contains(game.Title, gameTitle) {
			return true
		}
	}
	return false
}
