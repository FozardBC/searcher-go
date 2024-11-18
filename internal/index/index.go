package index

import "searcher/internal/crawler"

type Index struct {
	Words map[string][]int
	Docs  []crawler.Document
}

func New() *Index {
	return &Index{Words: make(map[string][]int)}
}

func (i *Index) AddWord(word string, docId int) {
	i.Words[word] = append(i.Words[word], docId)
}

func (i *Index) DocsID(word string) []int {
	return i.Words[word]
}
