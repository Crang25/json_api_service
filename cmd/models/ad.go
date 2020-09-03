package models

import "time"

// AdList ...
type AdList struct {
	Ads []Ad `json:"ads"`
}

// Ad ...
type Ad struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	PhotoLinks  []string `json:"photo_links"`
	Price       float64  `json:"price"`
	Date        time.Time
}
