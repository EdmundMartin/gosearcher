package gosearcher

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func builBingUrls(searchTerm, country string, pages, count int) ([]string, error) {
	toScrape := []string{}
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	if countryCode, found := bingDomains[country]; found {
		for i := 0; i < pages; i++ {
			first := firstParameter(i, count)
			scrapeURL := fmt.Sprintf("https://bing.com/search?q=%s&first=%d&count=%d%s", searchTerm, first, count, countryCode)
			toScrape = append(toScrape, scrapeURL)
		}
	} else {
		err := fmt.Errorf("country (%s) is currently not supported", country)
		return nil, err
	}
	return toScrape, nil
}

func firstParameter(number, count int) int {
	if number == 0 {
		return number + 1
	}
	return number*count + 1
}

func bingResultParser(response *http.Response, rank int) ([]SearchResult, error) {
	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		return nil, err
	}
	results := []SearchResult{}
	sel := doc.Find("li.b_algo")
	rank++
	for i := range sel.Nodes {
		item := sel.Eq(i)
		linkTag := item.Find("a")
		link, _ := linkTag.Attr("href")
		titleTag := item.Find("h2")
		descTag := item.Find("div.b_caption p")
		desc := descTag.Text()
		title := titleTag.Text()
		link = strings.Trim(link, " ")
		if link != "" && link != "#" && !strings.HasPrefix(link, "/") {
			result := SearchResult{
				rank,
				link,
				title,
				desc,
			}
			results = append(results, result)
			rank++
		}
	}
	return results, err
}

// BingScrape scrapes bing.com with desired parameters
func BingScrape(searchTerm, country string, proxyString interface{}, pages, count, backoff int) ([]SearchResult, error) {
	results := []SearchResult{}
	bingPages, err := builBingUrls(searchTerm, country, pages, count)
	if err != nil {
		return nil, err
	}
	for _, page := range bingPages {
		rank := len(results)
		res, err := scrapeClientRequest(page, proxyString)
		if err != nil {
			return nil, err
		}
		data, err := bingResultParser(res, rank)
		if err != nil {
			return nil, err
		}
		for _, result := range data {
			results = append(results, result)
		}
		time.Sleep(time.Duration(backoff) * time.Second)
	}
	return results, nil
}
