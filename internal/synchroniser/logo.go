package synchroniser

import "fmt"

const logo = `___         _       _   ___           _
| _ \___  __| |_____| |_| _ ) ___  ___| |__
|  _/ _ \/ _| / / -_)  _| _ \/ _ \/ _ \ / /
|_|_\___/\__|_\_\___|\__|___/\___/\___/_\_\
|   \ _ _ ___ _ __| |__  _____ __
| |) | '_/ _ \ '_ \ '_ \/ _ \ \ /
|___/|_| \___/ .__/_.__/\___/_\_\    _
|   \ _____ _|_|__ _ | |___  __ _ __| |___ _ _
| |) / _ \ V  V / ' \| / _ \/ _' / _' / -_) '_|
|___/\___/\_/\_/|_||_|_\___/\__,_\__,_\___|_|

`

func (ds *DropboxSynchroniser) infoHeader() {
	fmt.Fprint(ds.output, logo)
	fmt.Println()
	fmt.Fprintf(ds.output, "Account: %s\n", ds.dropbox.AccountDisplayName())
	fmt.Fprintf(ds.output, "Email: %s\n", ds.dropbox.AccountEmail())
}