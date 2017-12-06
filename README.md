# gosearcher
gosearcher is a small golang package which provides a number of functions that make it easy to scrape a number of popular search engines including Google, Yandex and soon Bing. The package supports the use proxies and multi-page scraping and allows users to scrape multiple pages of search results.

## Example Usage - Google Scraping
```go
package main

import (
	"fmt"
	"github.com/EdmundMartin/gosearcher"
)

func main() {
	res, err := googleGrabber.GoogleScrape("Edmund Martin", "com", "en", nil, 1, 10, 10)
	if err == nil {
		for _, res := range res {
			fmt.Println(res)
		}
	}
}
```
### Parameters - Google Scraping
* searchTerm - string
* countryCode - string - Will return an error if country is not supported by Google. "com" - will use Google.com
* languageCode - string - The language used to search - in the format of an ISO 639-1 Code
* proxyString - empty interface - The proxy (string format) you wish to use for the particular scrape, or nil to scrape without a proxy
* pages - int - The number of pages you wish to scrape.
* count - int - The number of results per page - multiples of 10 up to 100.
* backoff - int - The time to wait in between scraping pages, if more than one page of results is being scraped.
## Example Usage - Yandex Scraping
```go
package main

import (
	"fmt"
	"github.com/EdmundMartin/gosearcher"
)

func main() {
	res, err := googleGrabber.YandexScrape("Привет меня зовут", "com", "10393", nil, 1, 30, 20)
	if err == nil {
		for _, res := range res {
			fmt.Println(res)
		}
	} else {
		fmt.Println(err)
	}
}
```
### Parameters - Yandex Scraping
* searchTerm - string
* countryCode - string - Will return an error if country is not supported by Yandex. "com" - will use Yandex.com
* location - empty interface - Yandex's location code, can be a string or will use Moscow as base if nil is based. Full list can be found [here](https://yandex.ru/yaca/geo.c2n).
* proxyString - empty interface - The proxy (string format) you wish to use for the particular scrape, or nil to scrape without a proxy
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
