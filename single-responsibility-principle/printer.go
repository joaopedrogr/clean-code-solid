package single_responsibility_principle

import (
	"io/ioutil"
)

func PrintNews(filename string, journal JournalClassic) {
	_ = ioutil.WriteFile(filename, []byte(journal.String()), 0644)
}
