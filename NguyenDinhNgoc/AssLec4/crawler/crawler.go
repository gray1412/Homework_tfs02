package crawler

import (
	"ngoc/storage"
	"sync"
)

type Crawler interface {
	CrawlFilm(sync.WaitGroup, string, map[int64]storage.Film)
	CrawlItem(sync.WaitGroup, string, map[int64]storage.Item)
}
