# googleGrabber
Golang library for scraping search results. Currently, the library supports scraping results from Google and allows the caller to define a number of important variables.

## Example Usage - Google Scraping
```go
package main

import (
	"fmt"
	"github.com/EdmundMartin/googleGrabber"
)

func main() {
	res, err := googleGrabber.GoogleScrape("Edmund Martin", "com", "en", "", 1, 10, 10)
	if err == nil {
		for _, res := range res {
			fmt.Println(res)
		}
	}
}
```
### Parameters - Google Scraping
* searchTerm - string
* countryCode - string - Will return an error if country is not supported by Google.
* languageCode - string - The language used to search - in the format of an ISO 639-1 Code
* proxyString - string - The proxy you wish to use for the particular scrape, an empty string to scrape without a proxy
* pages - int - The number of pages you wish to scrape.
* count - int - The number of results per page - multiples of 10 up to 100.
* backoff - int - The time to wait in between scraping pages, if more than one page of results is being scraped.
## Result Format
```go
type SearchResult struct {
	ResultRank int
	ResultURL string
	ResultTitle string
	ResultDesc string
}
```
All supported search engines return a slice of SearchResult. This struct contains the rank, url, title and description of the particular result in question.
