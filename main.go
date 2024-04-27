package main

import (
	"fmt"

	"github.com/j-monteiro/url-shortener/services"
)

func main() {
	s := services.ShortenUrlService{}
	link := s.CreateShortenedLink("https://www.google.com")
	fmt.Println(link)
}
