package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/tj/go-dropbox"
)

// https://www.dropbox.com/oauth2/authorize?client_id=srb8i7o71xkx08c&response_type=code

func main() {
	config := struct {
		AccessToken string `json:"access_token"`
		Folder      string `json:"folder"`
		DeleteFiles bool   `json:"delete_files"`
	}{}

	data, err := ioutil.ReadFile("./demo.json")
	if err != nil {
		panic(err)
	}

	json.Unmarshal(data, &config)

	dp := dropbox.New(dropbox.NewConfig(config.AccessToken))
	account, err := dp.Users.GetCurrentAccount()
	if err != nil {
		panic(err)
	}

	fmt.Println(account.Name.DisplayName)
	fmt.Println(account.Email)

}
