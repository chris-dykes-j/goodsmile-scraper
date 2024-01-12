package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
    println("Scraping Goodsmile Nendoroid Figure Data")
    
	url := "https://www.goodsmile.info/en/product/8819/Nendoroid+Zelda+Breath+of+the+Wild+Ver.html"
    // url2 := "https://www.goodsmile.info/en/product/15354/Nendoroid+Nino+Nakano+Wedding+Dress+Ver.html"
    // url3 := "https://www.goodsmile.info/en/product/15353/Nendoroid+Ninomae+Ina+nis.html"

    nendo := getNendoroidData(url)
    saveNendoroidData(nendo)

    fmt.Println(nendo)
}

type Nendoroid struct {
    Name        string    `json:"name"`
    Description string    `json:"description"`
    ItemLink    string    `json:"itemLink"`
    BlogLink    string    `json:"blogLink"`
    Details     []Details `json:"details"`
}

type Details struct {
    Key   string `json:"key"`
    Value string `json:"value"`
}

func getNendoroidData(url string) Nendoroid {
	userAgent := "Mozilla/5.0 (X11; Linux x86_64; rv:121.0) Gecko/20100101 Firefox/121.0"

	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", userAgent)
	})

    // Get name
    var name string
    c.OnHTML(".title", func(e *colly.HTMLElement) {
        name = strings.TrimSpace(e.Text)
    })
    defer c.OnHTMLDetach(".title")

    // Get description
    var desc string
    c.OnHTML(".description", func(e *colly.HTMLElement) {
        desc = strings.TrimSpace(e.Text)
    })
    defer c.OnHTMLDetach(".description")

    // Get blogLinks
    var blogLink string
    c.OnHTML("#bloglink", func(e *colly.HTMLElement) {
        blogLink = e.ChildAttr("a", "href")
    })
    defer c.OnHTMLDetach("#blogLink")

    // Get details
    var keys []string
    var values []string
    c.OnHTML(".detailBox > dl", func(e *colly.HTMLElement) {
        e.ForEach("dt", func(_ int, el *colly.HTMLElement) {
            key := el.Text
            keys = append(keys, strings.TrimSpace(key))
        })
        e.ForEach("dd", func(_ int, el *colly.HTMLElement) {
            value := el.Text
            values = append(values, strings.TrimSpace(value))
        })
    })
    defer c.OnHTMLDetach(".detailBox > dl")

	c.Visit(url)

    var nendoroid Nendoroid

    // Add name
    nendoroid.Name = name
    // Add description
    nendoroid.Description = desc
    // Add itemLink
    nendoroid.ItemLink = url
    // Add blogLink
    nendoroid.BlogLink = blogLink
    // Add details
    for i, key := range keys {
        nendoroid.Details = append(nendoroid.Details, Details{key, values[i]})
    }

	return nendoroid 
}

func saveNendoroidData(nendo Nendoroid) {
    // create file if not exists (jsonl)

    // open file to append

    // serialize nendo

    // add nendo to file
}
