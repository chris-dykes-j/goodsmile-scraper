package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
    println("Scraping Goodsmile Nendoroid Figure Data")
    
	url := "https://www.goodsmile.info/en/product/8819/Nendoroid+Zelda+Breath+of+the+Wild+Ver.html"
    // url := "https://www.goodsmile.info/en/product/15354/Nendoroid+Nino+Nakano+Wedding+Dress+Ver.html"
    // url := "https://www.goodsmile.info/en/product/15353/Nendoroid+Ninomae+Ina+nis.html"

    nendo := getNendoroidData(url)

    for _, n := range nendo {
        fmt.Printf("%s: %s\n", n.key, n.value)
    }
}

type Details struct {
    key string
    value string
}

func getNendoroidData(url string) []Details {
	userAgent := "Mozilla/5.0 (X11; Linux x86_64; rv:121.0) Gecko/20100101 Firefox/121.0"

	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", userAgent)
	})

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

    var nendoroids []Details
    for i, key := range keys {
        nendoroids = append(nendoroids, Details{key, values[i]})
    }

	return nendoroids
}

func saveNendoroidData() {

}
