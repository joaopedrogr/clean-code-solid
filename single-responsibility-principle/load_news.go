package single_responsibility_principle

import "net/url"

func LoadNewsFromFile(filename string) {
	//...
}

func LoadNewsFromWeb(website *url.URL) {
	//...
}
