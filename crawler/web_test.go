package crawler

import (
	"testing"
)

func TestLinks(t *testing.T) {

	expected := []string{"/about"}

	web := &Web{
		Exclude: DefaultExclude,
	}

	// TODO: This should be mock data
	links := web.Links("https://duckduckgo.com/")

	if links[0] != expected[0] {
		t.Error("Excepted ", expected)
	}

}
