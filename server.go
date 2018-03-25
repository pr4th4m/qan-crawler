package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	qanCrawler "github.com/pratz/qan-crawler/crawler"
	"log"
	"net/http"
	"strconv"
)

const (
	apiVersion = "/api/v1"
)

func main() {
	var router = mux.NewRouter()
	endpoint := fmt.Sprintf("%s/crawl", apiVersion)
	router.HandleFunc(endpoint, crawl).Methods(http.MethodGet)

	// NOTE: We can accept host/port from flags as well
	address := "0.0.0.0:8080"
	log.Println("Server running on", address)
	log.Fatal(http.ListenAndServe(address, router))
}

func crawl(res http.ResponseWriter, req *http.Request) {

	// API filtering support
	queryParams := req.URL.Query()
	var url string
	if queryParams.Get("url") == "" {
		url = qanCrawler.DefaultURL
	} else {
		url = queryParams.Get("url")
	}

	var depth int
	if queryParams.Get("depth") == "" {
		depth = qanCrawler.DefaultDepth
	} else {
		depthParam, _ := strconv.Atoi(queryParams.Get("depth"))
		depth = depthParam
	}

	exclude := qanCrawler.DefaultExclude
	if queryParams.Get("exclude") != "" {
		exclude = append(exclude, queryParams["exclude"]...)
	}

	// Define crawler and start it
	web := &qanCrawler.Web{Exclude: exclude}
	crawler := &qanCrawler.Crawler{
		URL:    url,
		Depth:  depth,
		Web:    web,
		Cacher: map[string][]string{},
	}
	pages := crawler.Start()
	RenderJSON(res, http.StatusOK, pages)
}

// RenderJSON as rest response
func RenderJSON(w http.ResponseWriter, status int, obj interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	// NOTE: Uncomment to display api response
	// log.Println("Render response", obj)
	return json.NewEncoder(w).Encode(obj)
}
