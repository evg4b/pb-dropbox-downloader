package synchroniser

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/c2h5oh/datasize"
)

const maxLingth = 52
const minSpace = 3

func min(n int, m int) int {
	if n > m {
		return m
	}

	return n
}

func (s *DropboxSynchroniser) printerThread(source chan printInfo) {
	for info := range source {
		if info.success {
			size := fmt.Sprintf("[%s]", datasize.ByteSize(info.size).HumanReadable())
			s.printf(format(info.name, size))
		} else {
			s.printf(format(info.name, "[fail]"))
		}
	}
}

func format(name, suffix string) string {
	suffixLen := utf8.RuneCountInString(suffix)
	maxNameLen := min(maxLingth-suffixLen-minSpace, utf8.RuneCountInString(name))
	base := strings.Repeat(".", maxLingth-maxNameLen-suffixLen)

	return substring(name, maxNameLen) + base + suffix
}
