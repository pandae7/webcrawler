package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Function to fetch URL's HTML content
func fetchContent(urlStr string) (*goquery.Document, error) {
	httpClient := http.Client{Timeout: 30 * time.Second}
	resp, err := httpClient.Get(urlStr)
	if err != nil {
		return nil, err
	}
	println("response status Code: ", resp.Status)
	defer resp.Body.Close()
	return goquery.NewDocumentFromReader(resp.Body)
}

// Function to extract URLs and print them
func extractAndPrintURLs(baseURL string, doc *goquery.Document) {
	baseURLParsed, err := url.Parse(baseURL)
	if err != nil {
		fmt.Printf("Error parsing base URL: %v\n", err)
		return
	}
	println("debig point")
	println(doc.Html())
	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		if href, exists := s.Attr("href"); exists {
			absURL, err := baseURLParsed.Parse(href)
			if err != nil {
				fmt.Printf("Error parsing URL %s: %v\n", href, err)
				return
			}
			fmt.Println(absURL.String())
		}
	})
}

func main() {
	// seedURL := "https://quotes.toscrape.com"
	seedURL := "https://thebearhouse.com/sitemap.xml"

	doc, err := fetchContent(seedURL)
	if err != nil {
		fmt.Printf("Error fetching content: %v\n", err)
		return
	}

	fmt.Println("URLs found:")
	extractAndPrintURLs(seedURL, doc)
}
