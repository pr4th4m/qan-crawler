package crawler

import (
	"log"
	"net/url"
)

// Crawler to crawl web pages
type Crawler struct {
	URL   string
	Depth int
	Web   *Web
	// TODO: This can be redis store or similar
	Cacher map[string][]string
}

// Start crawling web pages
func (cw *Crawler) Start() map[string][]string {

	url, err := url.Parse(cw.URL)
	if err != nil {
		log.Println(err)
	}
	return cw.crawl(url.Scheme, url.Host,
		[]string{url.String()}, cw.Depth)
}

// crawl web pages
func (cw *Crawler) crawl(schema, host string,
	urls []string, depth int) map[string][]string {

	if depth <= 0 {
		log.Println("Crawling completed")
		return nil
	}

	for _, uri := range urls {

		// Url is parsed to get schema and host
		// so that it is available for next pages
		// to crawl
		parsedURI, err := url.Parse(uri)
		if err != nil {
			log.Println(err)
		}
		if parsedURI.Host == "" {
			parsedURI.Host = host
		}
		if parsedURI.Scheme == "" {
			parsedURI.Scheme = schema
		}

		// Check in cache if url already exist
		if _, ok := cw.Cacher[parsedURI.String()]; ok {
			continue
		}

		links := cw.Web.Links(parsedURI.String())
		cw.crawl(parsedURI.Scheme, parsedURI.Host, links, depth-1)
		cw.Cacher[parsedURI.String()] = links
	}

	return cw.Cacher
}
