package main

import (
	"errors"
	"fmt"

	"github.com/RicardoSantosSantana/apiTokenControl"
	"github.com/RicardoSantosSantana/dbTokenControl"
)

func MakeNewToken() {

	_token, err := apiTokenControl.New()
	checkErr(err)

	token := dbTokenControl.Token{
		Access_token:  _token.Access_token,
		Token_type:    _token.Token_type,
		Expires_in:    _token.Expires_in,
		Scope:         _token.Scope,
		User_id:       _token.User_id,
		Refresh_token: _token.Refresh_token,
	}

	errAdd := dbTokenControl.Token.Add(token)
	checkErr(errAdd)

	dbTokenControl.LogMessage("new token generated")

	fmt.Println(token)
}

func MakeRefreshToken() {

	token, err := dbActiveToken()

	if err != nil {
		dbTokenControl.LogMessage("Impossible generate refresh token, trying get new token")
		MakeNewToken()
		return
	}

	refreshToken, err := apiTokenControl.MakeRefreshToken(token)

	if err != nil {
		dbTokenControl.LogMessage("Impossible generate refresh token, trying get new token")
		MakeNewToken()
		return
	}

	newRefreshToken := dbTokenControl.Token{
		Access_token:  refreshToken.Access_token,
		Token_type:    refreshToken.Token_type,
		Expires_in:    refreshToken.Expires_in,
		Scope:         refreshToken.Scope,
		User_id:       refreshToken.User_id,
		Refresh_token: refreshToken.Refresh_token,
	}

	errAdd := dbTokenControl.Token.Add(newRefreshToken)
	checkErr(errAdd)

	dbTokenControl.LogMessage("new refresh token generated")
	fmt.Println(refreshToken)
}

func dbActiveToken() (apiTokenControl.Token, error) {

	_token, err := dbTokenControl.Active()

	if err != nil {
		err := errors.New("impossible generate refresh token, trying get new token")
		return apiTokenControl.Token{}, err
	}

	token := apiTokenControl.Token{
		Access_token:  _token.Access_token,
		Token_type:    _token.Token_type,
		Expires_in:    _token.Expires_in,
		Scope:         _token.Scope,
		User_id:       _token.User_id,
		Refresh_token: _token.Refresh_token,
	}

	return token, nil

}
