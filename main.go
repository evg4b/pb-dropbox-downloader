package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"

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

	files, err := dp.Files.ListFolder(&dropbox.ListFolderInput{
		Path:             "/",
		Recursive:        true,
		IncludeMediaInfo: true,
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(account.Name.DisplayName)
	fmt.Println(account.Email)

	data2 := map[string]string{}
	for _, file := range files.Entries {
		// fmt.Println(file.ContentHash)
		// fmt.Println(file.Name)
		fmt.Println(filepath.Clean(file.PathLower))
		fmt.Println(file.PathLower)
		data2[file.ContentHash] = file.Name
	}

	// for _, file := range pocketbook.GetFiles("E:/system") {
	// 	fmt.Println(file)
	// }

	// encoded, err := binary.Marshal(data)
	// ioutil.WriteFile("demo.bin", encoded, 0666)

}

func runDonwloader(chanel chan string, client *dropbox.Client) {

	for i := 0; i < 3; i++ {
		go func() {
			for file := range chanel {
				output, err := client.Files.Download(&dropbox.DownloadInput{
					Path: file,
				})
				if err != nil {
					continue
				}
				io.Copy(nil, output.Body)
			}
		}()
	}
}
