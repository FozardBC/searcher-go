package processor

import (
	"fmt"
	"log"
	"searcher/internal/crawler/spider"
	"searcher/internal/database"
	"searcher/internal/database/files"
	"searcher/internal/index"
)

type Proc struct {
	I  *index.Index
	S  *spider.Service
	Db database.IDatabase
}

func New() *Proc {
	p := Proc{
		I:  index.New(),
		S:  spider.New(),
		Db: db(),
	}

	return &p
}

func db() files.FilesDB {
	db, err := files.New()
	if err != nil {
		log.Fatal("can't make db: %w", err)
	}
	return *db
}

func (p *Proc) Save() {
	for s, i := range p.I.Words {

		str := fmt.Sprintf("%s:%v", s, i)
		p.Db.Write([]byte(str))

	}
}

func (p *Proc) FindUrls(t string) {

	urlsID := p.I.DocsID(t)

	for y, url := range urlsID {
		fmt.Printf("[%v]: %s\n", y, p.I.Docs[url].URL)
	}
}
