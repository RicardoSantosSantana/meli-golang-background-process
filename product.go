package main

import "github.com/RicardoSantosSantana/apiTokenControl"

func GetItemsIds() ([]string, error) {

	token, err := dbActiveToken()
	checkErr(err)

	items, err := apiTokenControl.Token.GetProductItemsIds(token)
	checkErr(err)

	return items, nil
}

func GetItemsDetails(itemId string) (apiTokenControl.Items, error) {

	items, err := apiTokenControl.GetItemsDetails(itemId)
	checkErr(err)

	return items, nil

}

func GetItemsDescription(itemId string) (apiTokenControl.Description, error) {

	items, err := apiTokenControl.GetItemsDescription(itemId)
	checkErr(err)

	return items, nil

}
