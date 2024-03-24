# Web-Scraping-in-Golang

Web scraping in Golang is a popular approach to automatically retrieve data from the web.

## part 01 How to Scrape a Website in Go

[Scrape](https://www.zenrows.com/blog/web-scraping-golang#how-to-web-scrape-in-go)

we will using  [ScrapeMe](https://scrapeme.live/shop/) website as an target for learn how to scrape a website in Go

Colly: A fast and flexible web scraping framework for Go

The program scrapes product information from a single page of a website.
 It uses the Colly library to make requests, handle responses, and extract data from HTML elements.
  The scraped data is then saved to a CSV file named products, csv in the project directory.

### Part 02 Advanced Techniques in Web Scraping with Golang

The program scrapes product information from multiple pages of a website.
It uses pagination links to navigate through the pages and extract the data.
The scraped data is then saved to a CSV file named products_by_pages,csv in the project directory parallel.
