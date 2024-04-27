package services

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/j-monteiro/url-shortener/models"
)

type ShortenUrlService struct {
}

func (s *ShortenUrlService) CreateShortenedLink(url string) *models.Link {
	shortUrl := s.shortenUrl(url)

	link := models.NewLink(shortUrl, url)

	return link
}

func (s *ShortenUrlService) shortenUrl(url string) string {
	hasher := sha256.New()

	hasher.Write([]byte(url))

	hash := hasher.Sum(nil)

	return hex.EncodeToString(hash)
}
