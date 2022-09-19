//go:build !menu
// +build !menu

package main

import (
	"fmt"

	"github.com/RicardoSantosSantana/dbTokenControl"
)

func main() {
 
	dbTokenControl.TruncateAllItems()
	MakeRefreshToken()
	ListSave()

	fmt.Println("Fim do processo")
}
