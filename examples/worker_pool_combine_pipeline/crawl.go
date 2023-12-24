package main

import (
	"context"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

// func CrawlTopics() []*Topic {
// 	c := colly.NewCollector()

// 	linkArr := []*Topic{}
// 	// Chỉ định hàm callback cho sự kiện OnHTML
// 	c.OnHTML("a[href].topicIndexChicklet ", func(e *colly.HTMLElement) {
// 		link := e.Attr("href")
// 		linkArr = append(linkArr, &Topic{Url: BaseURL + link, Name: removeAccent(e.Text)})
// 	})

// 	err := c.Visit(BaseURL + "/topics")
// 	if err != nil {
// 		log.Printf("Lỗi khi truy cập trang web: %s - err %s", "step1", err)
// 	}

// 	return linkArr
// }

func CrawlTopics() []*Topic {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
	)
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))

	// Khởi tạo context
	defer cancel()

	// Truy cập trang web
	err := chromedp.Run(ctx, chromedp.Navigate(BaseURL+"/topics"))
	if err != nil {
		log.Printf("CrawlTopics::Lỗi khi truy cập trang web: %s - err %s", "step1", err)
	}

	// Lấy nội dung HTML của trang web
	var html string
	err = chromedp.Run(ctx, chromedp.InnerHTML("html", &html, chromedp.ByQuery))
	if err != nil {
		log.Printf("CrawlTopics::Lỗi khi lấy nội dung HTML: %s - err %s", "step2", err)
	}

	// Phân tích cú pháp HTML bằng goquery
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Printf("CrawlTopics::Lỗi khi phân tích cú pháp HTML: %s - err %s", "step3", err)
	}

	// Lấy các phần tử HTML chứa các đường dẫn đến các chủ đề
	topics := []*Topic{}
	doc.Find(`a[href].topicIndexChicklet`).Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		name := removeAccent(s.Text())
		topics = append(topics, &Topic{Url: BaseURL + removeBaseUrl(link), Name: strings.TrimSpace(name)})
	})

	return topics
}

// func CrawlLastPageUrlQuote(urlQuoteList string) string {

// 	c := colly.NewCollector()

// 	var url string
// 	// Chỉ định hàm callback cho sự kiện OnHTML
// 	c.OnHTML("body > main > div.infScrollFooter > div.bq_s.hideInfScroll.bq_pageNumbersCont > ul > li:nth-last-child(2) > a", func(e *colly.HTMLElement) {
// 		link := e.Attr("href")

// 		url = link
// 	})

// 	// Truy cập vào trang web cần crawl
// 	err := c.Visit(urlQuoteList)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	return url
// }

func CrawlLastPageUrlQuote(urlQuoteList string) string {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
	)
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	// Truy cập trang web
	err := chromedp.Run(ctx, chromedp.Navigate(BaseURL+removeBaseUrl(urlQuoteList)))
	if err != nil {
		log.Printf("Lỗi khi truy cập trang web: %s - err %s", urlQuoteList, err)
	}

	// Lấy nội dung HTML của trang web
	var html string
	err = chromedp.Run(ctx, chromedp.InnerHTML("html", &html, chromedp.ByQuery))
	if err != nil {
		log.Printf("Lỗi khi lấy nội dung HTML: %s - err %s", urlQuoteList, err)
	}

	// Phân tích cú pháp HTML bằng goquery
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Printf("Lỗi khi phân tích cú pháp HTML: %s - err %s", urlQuoteList, err)
	}

	// Lấy URL của trang cuối cùng
	lastPageLink, _ := doc.Find("body > main > div.infScrollFooter > div.bq_s.hideInfScroll.bq_pageNumbersCont > ul > li:nth-last-child(2) > a").Attr("href")

	return BaseURL + lastPageLink
}

// func CrawlUrlQuotes(urlQuoteList string) []string {

// 	c := colly.NewCollector()

// 	linkArr := []string{}

// 	c.OnHTML("a[href].b-qt", func(e *colly.HTMLElement) {
// 		link := e.Attr("href")
// 		linkArr = append(linkArr, link)
// 		// fmt.Println(link)
// 	})

// 	// Truy cập vào trang web cần crawl
// 	err := c.Visit(urlQuoteList)
// 	if err != nil {
// 		log.Printf("Lỗi khi truy cập trang web: %s - err %s", urlQuoteList, err)
// 	}

// 	return linkArr
// }

func CrawlUrlQuotes(urlQuoteList string) []string {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-web-security", true),
	)
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	// Truy cập trang web
	err := chromedp.Run(ctx, chromedp.Navigate(urlQuoteList))
	if err != nil {
		log.Printf("CrawlLastPageUrlQuote::Lỗi khi truy cập trang web: %s - err %s", urlQuoteList, err)
	}

	// Lấy nội dung HTML của trang web
	var html string
	err = chromedp.Run(ctx, chromedp.InnerHTML("html", &html, chromedp.ByQuery))
	if err != nil {
		log.Printf("CrawlLastPageUrlQuote::Lỗi khi lấy nội dung HTML: %s - err %s", urlQuoteList, err)
	}

	// Phân tích cú pháp HTML bằng goquery
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Printf("CrawlLastPageUrlQuote::Lỗi khi phân tích cú pháp HTML: %s - err %s", urlQuoteList, err)
	}

	// check if link contain BaseURL remove it

	// Lấy các đường dẫn của các bài viết trên trang web
	links := []string{}
	doc.Find("a[href].b-qt").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		links = append(links, BaseURL+removeBaseUrl(link))
	})

	return links
}

// func CrawlQuoteDetail(urlQuoteList string) *QuoteDO {

// 	do := &QuoteDO{}
// 	c := colly.NewCollector(
// 	// colly.Async(true),
// 	// colly.MaxDepth(1),
// 	// colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36"),
// 	// colly.CacheDir("./colly_cache"),
// 	// colly.AllowURLRevisit(),
// 	// colly.IgnoreRobotsTxt(),
// 	)

// 	// c.SetRequestTimeout(1 * time.Second)
// 	c.OnHTML("div.quoteContent p.b-qt", func(e *colly.HTMLElement) {
// 		do.Content = e.Text
// 		// log.Println("content:", content)
// 	})

// 	c.OnHTML("div.quoteContent p.bq_fq_a a", func(e *colly.HTMLElement) {
// 		do.Author = e.Text
// 	})

// 	c.OnHTML(".kw-box a", func(e *colly.HTMLElement) {
// 		do.Tag = append(do.Tag, e.Text)
// 	})

// 	// Truy cập vào trang web cần crawl
// 	err := c.Visit(BaseURL + urlQuoteList)
// 	if err != nil {
// 		do.Content = ""
// 		log.Printf("Lỗi khi truy cập trang web: %s - err %s", urlQuoteList, err)
// 	}

//		return do
//	}
func CrawlQuoteDetail(urlQuoteList string) *QuoteDO {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-web-security", true),
		chromedp.Flag("user-agent", getRandomUserAgent()),
	)
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	do := &QuoteDO{}

	// Truy cập trang web
	err := chromedp.Run(ctx, chromedp.Navigate(BaseURL+removeBaseUrl(urlQuoteList)))
	if err != nil {
		do.Content = ""
		log.Printf("CrawlQuoteDetail::Lỗi khi truy cập trang web: %s - err %s", urlQuoteList, err)
	}

	// Lấy nội dung HTML của trang web
	var html string
	err = chromedp.Run(ctx, chromedp.InnerHTML("html", &html, chromedp.ByQuery))
	if err != nil {
		log.Printf("CrawlQuoteDetail::Lỗi khi lấy nội dung HTML: %s - err %s", urlQuoteList, err)
	}

	// Phân tích cú pháp HTML bằng goquery
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Printf("CrawlQuoteDetail::Lỗi khi phân tích cú pháp HTML: %s - err %s", urlQuoteList, err)
	}

	// Lấy nội dung bài viết
	doc.Find("div.quoteContent p.b-qt").Each(func(i int, s *goquery.Selection) {
		do.Content = s.Text()
	})

	// Lấy tác giả của bài viết
	doc.Find("div.quoteContent p.bq_fq_a a").Each(func(i int, s *goquery.Selection) {
		do.Author = s.Text()
	})

	// Lấy các tag của bài viết
	doc.Find(".kw-box a").Each(func(i int, s *goquery.Selection) {
		do.Tag = append(do.Tag, s.Text())
	})

	return do
}

var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36 Edg/114.0.0.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36",
	"Mozilla/5.0 (Linux; Android 9; SM-G960F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Mobile Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.1.1 Safari/605.1.15",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 13_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.1.1 Mobile/15E148 Safari/604.1",
}

func getRandomUserAgent() string {
	rand.Seed(time.Now().UnixNano())
	return userAgents[rand.Intn(len(userAgents))]
}
