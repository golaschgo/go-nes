package command

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

var Options []string = []string{"Buscar juego", "Ver lista de juegos", "Salir"}

func ShowWelcome() {
	ColorTitle := color.New(color.FgCyan).Add(color.Underline)
	textWelcome := "=== Seleccione una opción ==="
	ColorTitle.Println(textWelcome)
}

func ShowOptions() {
	ShowWelcome()
	for i, opt := range Options {
		fmt.Println(color.GreenString("- "+strconv.Itoa(i+1)), opt)
	}
}

func ShowMenu() string {
	ClearTerminal()
	ShowOptions()
	return Prompt("->")
}

func ClearTerminal() {
	cmd := exec.Command("cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Prompt(label string) string {
	var prompt string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		prompt, _ = r.ReadString('\n')
		if prompt != "" {
			break
		}
	}
	return strings.TrimSpace(prompt)

}

func Separator() {
	fmt.Println("===============")
}
func Separator2() {
	fmt.Println(color.GreenString("----------"))
}
