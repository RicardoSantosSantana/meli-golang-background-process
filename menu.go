package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/RicardoSantosSantana/apiTokenControl"
	"github.com/RicardoSantosSantana/dbTokenControl"
)

func clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()

}

func opcaoInvalida(msg string) {

	fmt.Println("_/﹋\\_")
	fmt.Println("(҂`_`)")
	fmt.Println("<,︻╦╤─ ҉ ------ - - - - - - - - " + msg)
	fmt.Println("/﹋\\")

}
func optionsMenu() int {
	fmt.Println(" ")
	fmt.Println(" Menu ")
	fmt.Println("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")

	fmt.Println("  1 - Exit")
	fmt.Println("  2 - Access Token")
	fmt.Println("  3 - List Product")
	fmt.Println("  4 - Delete all Produts on Meli")

	fmt.Println("┗━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
	fmt.Println("╭👀  ┃ ")
	fmt.Println("┫╰╯┣╯  ")
	fmt.Print("╰┳┳╯")
	var option int
	fmt.Scanln(&option)
	return option
}
func showAcessToken() {
	clear()
	token, err := dbTokenControl.Active()
	if err != nil {
		opcaoInvalida(err.Error())
	}

	fmt.Println(" ")
	fmt.Print(token.Access_token)
	fmt.Println(" ")
}
func deleteAllitems() {
	token, err := dbActiveToken()

	if err != nil {
		opcaoInvalida(err.Error())
	}

	_, errDeleted := apiTokenControl.DeleteAllItems(token)

	if errDeleted != nil {
		opcaoInvalida(errDeleted.Error())
	}
}
func listProduct() {
	clear()
	ids, err := GetItemsIds()

	if err != nil {
		opcaoInvalida(err.Error())
	}

	fmt.Println("Count: ", len(ids))
	fmt.Println("List: ", ids)
}
func Menu() {

	clear()
	for {
		var option = optionsMenu()

		if option == 1 {
			clear()
			break
		}

		if option == 2 {
			showAcessToken()

		} else if option == 4 {
			deleteAllitems()

		} else if option == 3 {

			listProduct()

		} else {
			clear()
			opcaoInvalida("opção inválida")
		}

	}

}