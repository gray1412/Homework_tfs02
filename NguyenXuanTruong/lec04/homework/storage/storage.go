package storage

// Struct Film chứa dữ liệu crawl của IMDb
type Film struct {
	Name  string
	Image string
	Rate  string
	Url   string
}

// Struct HomeDecor chứa dữ liệu crawl của Shopbase
type HomeDecor struct {
	ID    uint
	Title string
	Price int
}

type HomeDecors struct {
	HomeDecors []HomeDecor
}
