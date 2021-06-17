package crawler

import (
	"sync"
	"tfs-02/crawl/storage"
)

type Crawler interface {
	CrawlFilm(sync.WaitGroup, string, map[int64]storage.Film)
	CrawlItem(sync.WaitGroup, string, map[int64]storage.Item)
	CrawlYouTube(sync.WaitGroup, string, map[int64]storage.YouTube)
}
