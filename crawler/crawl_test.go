package crawler

import (
	"testing"
)

func TestCrawler(t *testing.T) {

	expected := []string{"/about"}

	web := &Web{Exclude: DefaultExclude}
	crawler := &Crawler{
		// TODO: This should be mock data
		URL:    "https://duckduckgo.com/",
		Depth:  1,
		Web:    web,
		Cacher: map[string][]string{},
	}
	pages := crawler.Start()

	if pages["https://duckduckgo.com/"][0] != expected[0] {
		t.Error("Excepted ", expected)
	}
}
