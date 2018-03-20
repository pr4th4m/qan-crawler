package main

import (
	"fmt"
	qanCrawler "github.com/pratz/qan-crawler/crawler"
)

func main() {

	exclude := []string{"/", "#"}
	web := &qanCrawler.Web{Exclude: exclude}

	crawler := &qanCrawler.Crawler{
		URL:    "https://pratz.github.io",
		Depth:  2,
		Web:    web,
		Cacher: map[string][]string{},
	}
	fmt.Println(crawler.Start())
}
