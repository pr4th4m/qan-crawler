package crawler

import (
	"log"
	"net/url"
	"sync"
)

// Define defaults for crawler
const (
	DefaultURL   = "https://google.com"
	DefaultDepth = 1
)

var DefaultExclude = []string{"/", "#"}

// Crawler to crawl web pages
type Crawler struct {
	URL   string
	Depth int
	Web   *Web
	// NOTE: Advanced caching can be done
	// with redis or similar
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

	// WaitGroup for goroutines
	wg := sync.WaitGroup{}

	if depth <= 0 {
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
			return nil
		}

		links := cw.Web.Links(parsedURI.String())
		if len(links) > 0 {
			cw.Cacher[parsedURI.String()] = links
			log.Println(parsedURI.String())

			// Crawl pages concurrently with goroutines
			wg.Add(1)
			go func() {
				cw.crawl(parsedURI.Scheme, parsedURI.Host, links, depth-1)
				wg.Done()
			}()
		}
	}

	wg.Wait()
	return cw.Cacher
}
