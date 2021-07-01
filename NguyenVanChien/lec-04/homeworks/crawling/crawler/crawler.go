package crawler

import (
	"chien/storage"
	"sync"

	"go.uber.org/zap"
)

type Crawler interface {
	Crawl(wg sync.WaitGroup, url string, sugar *zap.SugaredLogger, m *map[int64]storage.Film)
}
