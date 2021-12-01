package dropbox

import dropboxLib "github.com/tj/go-dropbox"

type dropboxOption = func(client *Client)

func WithGoDropbox(libClient *dropboxLib.Client) dropboxOption {
	return func(client *Client) {
		account, err := libClient.Users.GetCurrentAccount()
		if err != nil {
			panic(err)
		}

		client.files = libClient.Files
		client.account = account
	}
}

func WithAccount(account *dropboxLib.GetCurrentAccountOutput) dropboxOption {
	return func(client *Client) {
		client.account = account
	}
}

func WithFiles(files dropboxFiles) dropboxOption {
	return func(client *Client) {
		client.files = files
	}
}
