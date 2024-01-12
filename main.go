package main

import (
	"github.com/gocolly/colly"
)

func main() {
    println("Scraping Goodsmile Nendoroid Figure Data")
    
	url := "https://www.goodsmile.info/en/product/8819/Nendoroid+Zelda+Breath+of+the+Wild+Ver.html"
    // "https://www.goodsmile.info/en/product/15354/Nendoroid+Nino+Nakano+Wedding+Dress+Ver.html"
    // "https://www.goodsmile.info/en/product/15353/Nendoroid+Ninomae+Ina+nis.html"

    nendo := getNendoroidData(url)

    for i, n := range nendo {
        println(n[i])
    }
}

func getNendoroidData(url string) []string {
	userAgent := "Mozilla/5.0 (X11; Linux x86_64; rv:121.0) Gecko/20100101 Firefox/121.0"

	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", userAgent)
	})

    var nendoroids []string

    c.OnHTML(".detailBox > dl", func(e *colly.HTMLElement) {
        e.ForEach("dt", func(_ int, el *colly.HTMLElement) {
            text, _ := el.DOM.Html()
            nendoroids = append(nendoroids, text)
        })
    })
    defer c.OnHTMLDetach(".detailBox > dl")

	c.Visit(url)

	return nendoroids
}

func saveNendoroidData() {

}
