package crawler

import (
	"hieu/storage"
	"sync"
)

type Crawler interface {
	CrawlFilm(sync.WaitGroup, string, map[int64]storage.Movie)
	CrawlItem(sync.WaitGroup, string, map[int64]storage.Item)
}
