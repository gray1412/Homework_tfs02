package crawler

import (
	"sync"
)

type Crawler interface {
	CrawlFilm(sync.WaitGroup, string)
}
