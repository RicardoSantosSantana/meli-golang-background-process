package main

import (
	"encoding/json"
	"fmt"

	"github.com/RicardoSantosSantana/apiTokenControl"
	"github.com/RicardoSantosSantana/dbTokenControl"
)

func init() {
	dbTokenControl.StringConnection = StringConnection()
	apiTokenControl.InitialParams = InitialParams()
}

func SaveProduct(itemId string) {

	items, err := GetItemsDetails(itemId)
	checkErr(err)

	var p []apiTokenControl.Pictures = items.Pictures
	picture, errPicture := json.Marshal(p)
	checkErr(errPicture)

	var s []apiTokenControl.SalesTerms = items.Sale_terms
	salesTerms, errSalesTerm := json.Marshal(s)
	checkErr(errSalesTerm)

	des, errDescription := apiTokenControl.GetItemsDescription(itemId)
	checkErr(errDescription)

	var d apiTokenControl.Description = des
	description, errDescription := json.Marshal(d)
	checkErr(errDescription)

	dbItem := dbTokenControl.Items{
		Id:                 items.Id,
		Title:              items.Title,
		Site_id:            items.Site_id,
		Subtitle:           items.Subtitle,
		Seller_id:          items.Seller_id,
		Category_id:        items.Category_id,
		Official_store_id:  items.Official_store_id,
		Price:              items.Price,
		Base_price:         items.Base_price,
		Original_price:     items.Original_price,
		Currency_id:        items.Currency_id,
		Initial_quantity:   items.Initial_quantity,
		Available_quantity: items.Available_quantity,
		Sold_quantity:      items.Sold_quantity,
		Sale_terms:         string(salesTerms),
		Buying_mode:        items.Buying_mode,
		Listing_type_id:    items.Listing_type_id,
		Start_time:         items.Start_time,
		Stop_time:          items.Stop_time,
		Condition:          items.Condition,
		Permalink:          items.Permalink,
		Thumbnail_id:       items.Thumbnail_id,
		Thumbnail:          items.Thumbnail,
		Secure_thumbnail:   items.Secure_thumbnail,
		Status:             items.Status,
		Warranty:           items.Warranty,
		Catalog_product_id: items.Catalog_product_id,
		Domain_id:          items.Domain_id,
		Health:             items.Health,
		Pictures:           string(picture),
		Description:        string(description),
	}

	erroOnSave := dbTokenControl.Save(dbItem)

	if erroOnSave != nil {
		panic(erroOnSave)
	}
}
func ListSave() {

	ids, err := GetItemsIds()
	checkErr(err)

	for i := 0; i < len(ids); i++ {
		SaveProduct(ids[i])
	}
}

func main() {

	// MakeRefreshToken()
	// _token, err := dbTokenControl.Active()
	// checkErr(err)

	// token := apiTokenControl.Token{
	// 	Access_token:  _token.Access_token,
	// 	Token_type:    _token.Token_type,
	// 	Expires_in:    _token.Expires_in,
	// 	Scope:         _token.Scope,
	// 	User_id:       _token.User_id,
	// 	Refresh_token: _token.Refresh_token,
	// }
	// 	itemId := "MLB2750276798"
	// 	items, err := GetItemsDetails(itemId)

	// //	checkErr(err)
	//	itens, err := apiTokenControl.Token.GetProductItemsIds(token)
	// [MLB2767979721 MLB2750258500 MLB2761143241 MLB2750276798]

	//fmt.Println(items.Sale_terms)

	//Limpo todos os dados, pois não tem regra de negócio de update
	//dbTokenControl.TruncateAllItems()

	//ListSave()
	// MakeRefreshToken()
	// _token, err := dbTokenControl.Active()
	// checkErr(err)

	//Limpo todos os dados, pois não tem regra de negócio de update
	//dbTokenControl.TruncateAllItems()

	//fmt.Println(dbTokenControl.MakeSelectStatement("products"))
	dbTokenControl.TruncateAllItems()
	MakeRefreshToken()
	ListSave()

	fmt.Println("Fim do processo")
}
