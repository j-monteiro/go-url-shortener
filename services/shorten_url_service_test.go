package services

import (
	"testing"
)

func TestCreateShortenedLink(t *testing.T) {
	shortenUrlService := ShortenUrlService{}
	link := shortenUrlService.CreateShortenedLink("https://www.google.com")

	if link.OriginalUrl != "https://www.google.com" {
		t.Errorf("Expected link.OriginalUrl to be 'https://www.google.com', got %s", link.OriginalUrl)
	}

	if link.HashedUrl == "" {
		t.Errorf("Expected link.HashedUrl to be not empty, got %s", link.HashedUrl)
	}
}
