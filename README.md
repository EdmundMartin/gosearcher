# googleGrabber
Golang library for scraping search results

## Example Usage - Google
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
