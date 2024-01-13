package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func main() {
	urls := []string{
		"product/8819/",
		"product/15354/",
		"product/15353/",
		"product/15397",
	}
	fmt.Println("Begin Scraping")

	for _, url := range urls {
		figure := getFigure(url)
		saveFigureData(figure)
		fmt.Println(figure)
        sleep()
	}

	fmt.Println("Complete")
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
	English  FigureData `json:"en"`
	Japanese FigureData `json:"ja"`
	Chinese  FigureData `json:"zh"`
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
		saveImages(imageLinks, figure.Name)
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

func saveImages(links []string, figureName string) {
	userAgent := "Mozilla/5.0 (X11; Linux x86_64; rv:121.0) Gecko/20100101 Firefox/121.0"
	client := &http.Client{}

	for _, link := range links {
        req, err := http.NewRequest("GET", "https:" + link, nil)
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Set("User-Agent", userAgent)
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
		}

        imgPath := getDir(figureName)
		imgFile, err := os.Create(imgPath + "/" + path.Base(link)) 
		if err != nil {
			log.Fatal(err)
		}

		io.Copy(imgFile, res.Body)
	}
}

func getDir(figureName string) string {
	root := "/extra/nendoroid"
    dir := filepath.Join(root, figureName)
	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Made %s directory\n", dir)
	}
	return dir 
}
