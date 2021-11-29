package mocks

import (
	dropboxLib "github.com/tj/go-dropbox"
)

var Account = dropboxLib.GetCurrentAccountOutput{
	Name: struct {
		GivenName    string "json:\"given_name\""
		Surname      string "json:\"surname\""
		FamiliarName string "json:\"familiar_name\""
		DisplayName  string "json:\"display_name\""
	}{
		DisplayName: "Display Name",
	},
	Email: "Email@mail.com",
}
