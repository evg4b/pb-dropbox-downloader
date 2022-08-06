// nolint: wrapcheck
package synchroniser

import (
	"fmt"
	"log"
	"os"
	"pb-dropbox-downloader/internal/utils"
)

func (s *DropboxSynchroniser) deleteFiles(folder string, files []os.FileInfo) error {
	for _, file := range files {
		fileName := file.Name()
		exist, err := s.storage.KeyExists(fileName)
		if err != nil {
			return err
		}

		if !exist {
			err := s.files.Remove(utils.JoinPath(folder, fileName))
			if err != nil {
				fmt.Fprintf(s.output, "%s .... [filed]\n", file)
				log.Println(err)

				return err
			}

			fmt.Fprintf(s.output, "%s .... [ok]\n", file)
			err = s.storage.Remove(fileName)
			if err != nil {
				fmt.Fprintf(s.output, "%s .... [filed]\n", file)
				log.Println(err)

				return err
			}
		}
	}

	return nil
}
