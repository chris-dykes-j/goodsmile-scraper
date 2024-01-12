package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
    println("Scraping Goodsmile Nendoroid Figure Data")
    
	// url := "https://www.goodsmile.info/en/product/8819/Nendoroid+Zelda+Breath+of+the+Wild+Ver.html"
    url := "https://www.goodsmile.info/en/product/15354/Nendoroid+Nino+Nakano+Wedding+Dress+Ver.html"
    // "https://www.goodsmile.info/en/product/15353/Nendoroid+Ninomae+Ina+nis.html"

    nendo := getNendoroidData(url)

    for _, n := range nendo {
        fmt.Printf("%s: %s\n", n.key, n.value)
    }
}

type Tuple struct {
    key string
    value string
}

func getNendoroidData(url string) []Tuple {
	userAgent := "Mozilla/5.0 (X11; Linux x86_64; rv:121.0) Gecko/20100101 Firefox/121.0"

	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", userAgent)
	})

    var keys []string
    var values []string

    c.OnHTML(".detailBox > dl", func(e *colly.HTMLElement) {
        key := e.ChildText("dt")
        value := e.ChildText("dd")
        keys = append(keys, key)
        values = append(values, value)
    })
    defer c.OnHTMLDetach(".detailBox > dl")

	c.Visit(url)

    var nendoroids []Tuple
    for i, key := range keys {
        nendoroids = append(nendoroids, Tuple{key, values[i]})
    }

	return nendoroids
}

func saveNendoroidData() {

}
