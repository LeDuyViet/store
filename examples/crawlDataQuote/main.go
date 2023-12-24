package main

import (
	"log"
	"runtime"

	"github.com/gocolly/colly/v2"
)

const (
	BaseURL = "https://www.brainyquote.com"
)

func main() {

	topicLinks := CrawlTopics()

	workerCount := runtime.NumCPU()
	topicCount := len(topicLinks)
	jobs := make(chan string, topicCount)
	results := make(chan []string, topicCount)
	faildChan := make(chan string, topicCount)

	for w := 0; w < workerCount; w++ {
		go GetListQuoteWorker(faildChan, jobs, results)
	}

	for _, url := range topicLinks {
		jobs <- url
	}

	close(jobs)
	// resultLen := len(result.Get("audio_url").Array())
	count := 0
	for result := range results {
		log.Println("result: ", result)
		count += len(result)
	}

	a := 0

	log.Println("count: ", count, a)

	// for linkFaild := range faildChan {
	// 	fmt.Println("link faild: ", linkFaild)
	// }
}

func CrawlTopics() []string {
	c := colly.NewCollector()

	linkArr := []string{}
	// Chỉ định hàm callback cho sự kiện OnHTML
	c.OnHTML("a[href].topicIndexChicklet ", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		linkArr = append(linkArr, link)
		// fmt.Println(link)
	})

	// Truy cập vào trang web cần crawl
	err := c.Visit(BaseURL + "/topics")
	if err != nil {
		log.Printf("Lỗi khi truy cập trang web: %s - err %s", BaseURL, err)
	}

	return linkArr
}

func GetListQuoteWorker(faildchan chan<- string, jobs <-chan string, results chan<- []string) {

	for link := range jobs {
		c := colly.NewCollector()

		linkArr := []string{}
		// Chỉ định hàm callback cho sự kiện OnHTML
		c.OnHTML("a[href].b-qt", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			linkArr = append(linkArr, link)
			// fmt.Println(link)
		})

		// Truy cập vào trang web cần crawl
		err := c.Visit(BaseURL + link)
		if err != nil {
			log.Printf("Lỗi khi truy cập trang web: %s - err %s", BaseURL, err)
		}

		results <- linkArr
	}

}
