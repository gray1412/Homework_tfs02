package models

type Input struct {
	Op     string  `json:"op"`
	A      int     `json:"a"`
	B      int     `json:"b"`
	Result float64 `json:"result"`
}