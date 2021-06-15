package storage

import "go.uber.org/zap"

//hàm tạo file chứa logger
func NewFileLogger(filepath string) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	if filepath != "" {
		cfg.OutputPaths = []string{
			filepath,
		}
	}
	return cfg.Build()
}
