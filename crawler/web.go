package crawler

import (
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
)

// Web - To get web pages and links
type Web struct {
	Exclude []string // exclude urls
}

// Links - Get all links for web page
func (web *Web) Links(url string) []string {

	// Get page content
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return []string{}
	}
	defer resp.Body.Close()

	// Return all links for web page
	return web.getLinks(resp.Body)
}

// getLinks - Get all links for web page
func (web *Web) getLinks(body io.Reader) []string {

	linkIndex := []string{}
	tokeniser := html.NewTokenizer(body)

	for {
		tokenType := tokeniser.Next()

		switch tokenType {

		case html.ErrorToken:
			return linkIndex

		case html.StartTagToken, html.EndTagToken:
			token := tokeniser.Token()

			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {

						// Control exclusion of urls
						if in(web.Exclude, attr.Val) {
							continue
						}

						// Store urls in slice
						if !in(linkIndex, attr.Val) {
							linkIndex = append(linkIndex, attr.Val)
						}

					}
				}
			}
		}
	}
}

// in - check for duplicate in slice
func in(slice []string, value string) bool {
	for _, item := range slice {
		if value == item {
			return true
		}
	}
	return false
}
