# Web Scraping in Golang

Web scraping in Golang is a popular approach to automatically retrieve data from the web.

## Part 01: How to Scrape a Website in Go

[Scrape Tutorial](https://www.zenrows.com/blog/web-scraping-golang#how-to-web-scrape-in-go)

In this tutorial, we will use the [ScrapeMe](https://scrapeme.live/shop/) website as a target to learn how to scrape a website in Go.

### Tools Used

- Colly: A fast and flexible web scraping framework for Go.

### Description

The program scrapes product information from a single page of a website. It uses the Colly library to make requests, handle responses, and extract data from HTML elements. The scraped data is then saved to a CSV file named `products.csv` in the project directory.

## Part 02: Advanced Techniques in Web Scraping with Golang

The program scrapes product information from multiple pages of a website. It uses pagination links to navigate through the pages and extract the data. The scraped data is then saved to a CSV file named `products_by_pages.csv` in the project directory.

## Part 03: Scraping Dynamic-Content Websites with a Headless Browser in Go

- chromedp: A tool for crawling dynamic-content websites and interacting with a web page in a browser as a real user would.

### Description2

With chromedp, you can crawl dynamic-content websites and interact with a web page in a browser as a real user would. This also means that your script is less likely to be detected as a bot, so chromedp makes it easy to scrape a web page without getting blocked. On the contrary, Colly is limited to static-content websites and doesn't offer browser's capabilities such as executing JavaScript or rendering CSS.
