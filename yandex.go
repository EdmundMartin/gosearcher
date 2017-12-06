package googleGrabber

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func buildYandexUrls(searchTerm, country string, location interface{}, pages, count int) ([]string, error) {
	toScrape := []string{}
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "%20", -1)
	location = returnLocation(location)
	if yandexBase, found := yandexDomains[country]; found {
		for i := 0; i < pages; i++ {
			scrapeUrl := fmt.Sprintf("%s%s&lr=%s&numdoc=%d&p=%d", yandexBase, searchTerm, location, count, i)
			toScrape = append(toScrape, scrapeUrl)
		}
	} else {
		err := fmt.Errorf("country (%s) is currently not supported", country)
		return nil, err
	}
	return toScrape, nil
}

func returnLocation(location interface{}) string {

	switch v := location.(type){
	case string:
		return v
	default:
		return "1"
	}
}

func yandexResultParser(response *http.Response, rank int) ([]SearchResult, error) {
	var re = regexp.MustCompile(`\d{1,3}`)
	repl := "${1}"
	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		return nil, err
	}
	results := []SearchResult{}
	sel := doc.Find("li.serp-item")
	rank += 1
	for i := range sel.Nodes {
		item := sel.Eq(i)
		linkTag := item.Find("a")
		link, _ := linkTag.Attr("href")
		titleTag := item.Find("h2")
		descTag := item.Find("div.organic__content-wrapper")
		desc := descTag.Text()
		title := titleTag.Text()
		link = strings.Trim(link, " ")
		startswith := strings.HasPrefix(link, "//")
		if startswith != true {
			result := SearchResult{
				rank,
				link,
				re.ReplaceAllString(title, repl),
				desc,
			}
			results = append(results, result)
			rank += 1
		}
	}
	return results, err
}

func YandexScrape(searchTerm, country string, location, proxyString interface{}, pages, count, backoff int) ([]SearchResult, error) {
	results := []SearchResult{}
	yandexPages, err := buildYandexUrls(searchTerm, country, location, pages, count)
	if err != nil {
		return nil, err
	}
	for _, searchPage := range yandexPages {
		res, err := scrapeClientRequest(searchPage, proxyString)
		if strings.Contains(res.Request.URL.String(),"captcha"){
			return nil, fmt.Errorf("yandex served a captcha to your request")
		}
		if err != nil {
			return nil, err
		}
		data, err := yandexResultParser(res, len(results))
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
