package synchroniser

import (
	"fmt"
	"strings"
)

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

func (s *DropboxSynchroniser) infoHeader() {
	logoLength := 46
	versionLine := strings.Repeat(" ", logoLength)
	versionSuffix := fmt.Sprintf("version: %s", s.version)
	versionPreffix := versionLine[:logoLength-len(versionSuffix)]

	fmt.Fprint(s.output, logo)
	fmt.Fprintln(s.output, versionPreffix, versionSuffix)
	fmt.Println()
	fmt.Fprintf(s.output, "Account: %s\n", s.dropbox.AccountDisplayName())
	fmt.Fprintf(s.output, "Email: %s\n", s.dropbox.AccountEmail())
	fmt.Println()
}
