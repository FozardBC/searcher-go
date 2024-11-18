package app

import (
	"fmt"
	"searcher/internal/processor"
	"strings"
)

func Start() {

	p := processor.New()

	//word := "intl"
	urls := []string{"https://google.com"}
	var err error

	for _, url := range urls {
		p.I.Docs, err = p.S.Scan(url, 2)
		if err != nil {
			fmt.Print(err)
		}

		splitFunc := func(r rune) bool {
			return r == '/' || r == '?' || r == '&' || r == '#' || r == ' ' || r == '.'
		}

		for _, doc := range p.I.Docs {
			words := strings.FieldsFunc(doc.URL, splitFunc)

			for _, w := range words {
				p.I.AddWord(w, doc.ID)
			}
		}

	}

	p.Save()

}
