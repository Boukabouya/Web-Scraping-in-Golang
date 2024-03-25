/*package main

import (
	"encoding/csv"
	"log"
	"os"
	"sync"

	"github.com/gocolly/colly"
)

type PokemonProduct struct {
	url, image, name, price string
}
*/
/*
func main() {
	var pokemonProducts []PokemonProduct
	var wg sync.WaitGroup
	wg.Add(1)

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})

	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		pokemonProduct := PokemonProduct{}

		pokemonProduct.url = e.ChildAttr("a", "href")
		pokemonProduct.image = e.ChildAttr("img", "src")
		pokemonProduct.name = e.ChildText("h2")
		pokemonProduct.price = e.ChildText(".price")

		pokemonProducts = append(pokemonProducts, pokemonProduct)
	})

	c.OnScraped(func(r *colly.Response) {
		defer wg.Done()
		fmt.Println(r.Request.URL, " scraped!")
	})

	c.Visit("https://scrapeme.live/shop/")

	wg.Wait()

	file, err := os.Create("products.csv")
	if err != nil {
		log.Fatalln("Failed to create output CSV file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{
		"url",
		"image",
		"name",
		"price",
	}
	writer.Write(headers)

	for _, pokemonProduct := range pokemonProducts {
		record := []string{
			pokemonProduct.url,
			pokemonProduct.image,
			pokemonProduct.name,
			pokemonProduct.price,
		}
		writer.Write(record)
	}
}
*/
//Advanced Techniques in Web Scraping with Golang

/*func main() {
	var pokemonProducts []PokemonProduct

	// the first pagination URL to scrape

	pagesToScrape := []string{
		"https://scrapeme.live/shop/page/1/",
		"https://scrapeme.live/shop/page/2/",
		// ...
		"https://scrapeme.live/shop/page/47/",
		"https://scrapeme.live/shop/page/48/",
	}


	// initializing a Colly collector
	c := colly.NewCollector(
		// turning on the asynchronous request mode in Colly
		colly.Async(true),
	)
	c.Limit(&colly.LimitRule{
		// limit the parallel requests to 4 request at a time
		Parallelism: 4,
	})

	// Setting a valid User-Agent header for Linux Debian and Brave browser
	c.UserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36 Brave/90.0.0.0"

	// Lock to synchronize access to pokemonProducts slice
	var mu sync.Mutex

	// Callback to handle pagination links
	c.OnHTML("a.page-numbers", func(e *colly.HTMLElement) {
		newPaginationLink := e.Attr("href")
		if !contains(pagesToScrape, newPaginationLink) {
			pagesToScrape = append(pagesToScrape, newPaginationLink)
			c.Visit(newPaginationLink)
		}
	})

	// Callback to extract product information
	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		pokemonProduct := PokemonProduct{
			url:   e.ChildAttr("a", "href"),
			image: e.ChildAttr("img", "src"),
			name:  e.ChildText("h2"),
			price: e.ChildText(".price"),
		}
		mu.Lock()
		pokemonProducts = append(pokemonProducts, pokemonProduct)
		mu.Unlock()
	})
		// registering all pages to scrape
	for _, pageToScrape := range pagesToScrape {
		c.Visit(pageToScrape)
	}

	// Wait for the collector to finish
	c.Wait()

	// Write the data to a CSV file
	file, err := os.Create("products_by_pages.csv")
	if err != nil {
		log.Fatalf("Failed to create output CSV file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{
		"url",
		"image",
		"name",
		"price",
	}
	writer.Write(headers)

	for _, pokemonProduct := range pokemonProducts {
		record := []string{
			pokemonProduct.url,
			pokemonProduct.image,
			pokemonProduct.name,
			pokemonProduct.price,
		}
		writer.Write(record)
	}
}

// Helper function to check if a string exists in a slice
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
*/
//Scraping Dynamic-Content Websites with a Headless Browser in Go
package main

import (
	"context"
	"encoding/csv"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"log"
	"os"
)

type PokemonProduct struct {
	url, image, name, price string
}

func main() {
	var pokemonProducts []PokemonProduct

	// initializing a chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// navigate to the target web page and select the HTML elements of interest
	var nodes []*cdp.Node
	chromedp.Run(ctx,
		chromedp.Navigate("https://scrapeme.live/shop"),
		chromedp.Nodes(".product", &nodes, chromedp.ByQueryAll),
	)

	// scraping data from each node
	var url, image, name, price string
	for _, node := range nodes {
		chromedp.Run(ctx,
			chromedp.AttributeValue("a", "href", &url, nil, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.AttributeValue("img", "src", &image, nil, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text("h2", &name, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text(".price", &price, chromedp.ByQuery, chromedp.FromNode(node)),
		)

		pokemonProduct := PokemonProduct{}

		pokemonProduct.url = url
		pokemonProduct.image = image
		pokemonProduct.name = name
		pokemonProduct.price = price

		pokemonProducts = append(pokemonProducts, pokemonProduct)
	}

	// export logic
	// Write the data to a CSV file
	file, err := os.Create("products_With_chromedp.csv")
	if err != nil {
		log.Fatalf("Failed to create output CSV file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{
		"url",
		"image",
		"name",
		"price",
	}
	writer.Write(headers)

	for _, pokemonProduct := range pokemonProducts {
		record := []string{
			pokemonProduct.url,
			pokemonProduct.image,
			pokemonProduct.name,
			pokemonProduct.price,
		}
		writer.Write(record)
	}

}
