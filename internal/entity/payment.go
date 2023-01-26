package entity

import "time"

type Payment struct {
	ID          int
	Src         int
	Dst         int
	Sum         float64
	CreatedAt   *time.Time
	ProcessedAt *time.Time
	Status      int
	Description string
}
