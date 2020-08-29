package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	instutationInfos := getinsitutationInfos("http://erdincuzun.cv.nku.edu.tr/")

	fmt.Println(instutationInfos)
}

// GetLatestBlogTitles gets the latest blog title headings from the url
// given and returns them as a list.
func getinsitutationInfos(url string) []string {

	// Get the HTML
	resp, err := http.Get(url)
	if err != nil {
		return make([]string, 30)
	}
	// Convert HTML into goquery document
	doc, err := goquery.NewDocumentFromReader(resp.Body)

	// Save each .post-title as a list

	infos := make([]string, 30)
	doc.Find("#anasayfa > div.panel.panel-info > div.panel-body.table-responsive > table > tbody > tr > td").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
		infos = strings.Split(s.Text(), ":")
	})
	return infos
}
