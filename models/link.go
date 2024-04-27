package models

import "time"

type Link struct {
	ID          int
	HashedUrl   string
	OriginalUrl string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewLink(hashedUrl string, originalUrl string) *Link {
	return &Link{
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		HashedUrl:   hashedUrl,
		OriginalUrl: originalUrl,
	}
}
