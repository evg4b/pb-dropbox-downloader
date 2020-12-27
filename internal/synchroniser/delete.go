package synchroniser

import (
	"fmt"
	"log"
	"path"
)

func (db *DropboxSynchroniser) delete(folder string, files []string) error {
	for _, file := range files {
		err := db.files.DeleteFile(path.Join(folder, file))
		if err != nil {
			fmt.Fprintln(db.output, fmt.Sprintf("%s .... [filed]", file))
			log.Println(err)

			return err
		}

		fmt.Fprintln(db.output, fmt.Sprintf("%s .... [ok]", file))
	}

	return nil
}
