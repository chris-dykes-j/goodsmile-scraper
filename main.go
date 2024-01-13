package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func main() {
	// url := "product/8819/Nendoroid+Zelda+Breath+of+the+Wild+Ver.html"
	url := "product/15354/"
	// url := "product/15353/Nendoroid+Ninomae+Ina+nis.html"

    figure := getFigure(url)
	saveFigureData(figure)

	fmt.Println(figure)
}

type FigureData struct {
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

type Figure struct {
    English FigureData `json:"en"`
    Japanese FigureData `json:"ja"`
    Chinese FigureData `json:"zh"`
}

func getFigure(url string) Figure {
    var figure Figure

    figure.English = getFigureData("https://www.goodsmile.info/en/", url)
    sleep()
    figure.Japanese = getFigureData("https://www.goodsmile.info/ja/", url)
    sleep()
    figure.Chinese = getFigureData("https://www.goodsmile.info/zh/", url)
    sleep()

    return figure
}

func getFigureData(baseUrl string, itemUrl string) FigureData {
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

	// Get images
	var imageLinks []string
	c.OnHTML(".itemImg", func(e *colly.HTMLElement) {
		img := e.Attr("src")
		imageLinks = append(imageLinks, img)
	})

	c.Visit(baseUrl + itemUrl)


	var figure FigureData

	// Add name
	figure.Name = strings.TrimSpace(name)
	// Add description
	figure.Description = strings.TrimSpace(desc)
	// Add itemLink
	figure.ItemLink = itemUrl
	// Add blogLink
	figure.BlogLink = blogLink
	// Add details
	for i, key := range keys {
		figure.Details = append(figure.Details, Details{key, values[i]})
	}

	// Save images
    if baseUrl == "https://www.goodsmile.info/en/" {
        // saveImages(imageLinks)
        fmt.Printf("Saving Images for %s", figure.Name)
    }

	return figure
}

func saveFigureData(figure Figure) {
	fileName := "test.jsonl"

	// open file to append
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// serialize nendo
	data, err := json.Marshal(figure)
	if err != nil {
		fmt.Println(err)
	}

	// add data to jsonl
	_, err = file.WriteString(string(data) + "\n")
	if err != nil {
		fmt.Println(err)
	}
}

func sleep() {
	randomNumber := rand.Float64()*(4-2) + 2
	time.Sleep(time.Duration(randomNumber) * time.Second)
}

/*
func saveImages(links []string) {
    userAgent := "Mozilla/5.0 (X11; Linux x86_64; rv:121.0) Gecko/20100101 Firefox/121.0"

    client := &http.Client{}

    for _, link := range links {
        req, _ := http.NewRequest("GET", link, nil)
        req.Header.Set("User-Agent", userAgent)
        res, _ := client.Do(req)

        imgPath := filepath
        imgFile, _ := os.Create(imgPath)

        io.Copy(imgFile, res.Body)
    }
}
*/
