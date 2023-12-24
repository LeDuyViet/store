package main

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"gitlab.volio.vn/tech/backend/store-pro/pkg/store"
	"gitlab.volio.vn/tech/backend/store-pro/pkg/store/do"
)

func CreateCategory(job Job, resultCh chan interface{}) {
	client := store.GetStoreClient()

	res := client.CreateCategory(StoreID, job.Data.(*do.CategoryDO))
	if res.Code != 200 {
		log.Println("create category failed")
	}

	category := job.Data.(*do.CategoryDO)
	category.ID = res.Data.Data[0].ID
	category.Name = res.Data.Data[0].Name

	resultCh <- Job{ID: job.ID, Data: category}
}

func GetUrlQuotes(job Job, resultCh chan interface{}) {
	category := job.Data.(*do.CategoryDO)

	lastPage := CrawlLastPageUrlQuote(category.Url)

	result := CrawlUrlQuotes(category.Url)
	quoteDO := makeQuoteDO(result, category)

	resultCh <- Job{ID: job.ID, Data: quoteDO}

	if lastPage != "" {
		lastPageNum, err := ExtractNumberFromURL(lastPage)
		if err != nil {
			log.Println(err)
		}

		for i := 2; i <= lastPageNum; i++ {
			results := CrawlUrlQuotes(fmt.Sprintf("%s_%d", category.Url, i))
			quoteDO := makeQuoteDO(results, category)

			resultCh <- Job{ID: job.ID, Data: quoteDO}
		}
	}

}

var count int
var faildCount int
var successCount int

func CrawlQuote(job Job, resultCh chan interface{}, jobCh chan interface{}) {
	quotesDO := job.Data.([]*QuoteDO)
	quoteFaild := make([]*QuoteDO, 0)
	quotesResult := make([]*QuoteDO, 0)
	count += len(quotesDO)
	if job.IsFaild {
		log.Printf("===reCrawl job faild")
	}
	for _, quoteDO := range quotesDO {
		quoteRes := CrawlQuoteDetail(quoteDO.Url)
		if quoteRes.Content == "" {
			faildCount++
			quoteFaild = append(quoteFaild, quoteDO)
		} else {
			quoteRes.TopicID = quoteDO.TopicID
			quotesResult = append(quotesResult, quoteRes)
			successCount++
			log.Printf("quote: %s", quoteDO.Url)
		}
	}
	log.Printf("count: %d, faild: %d, success: %d", count, faildCount, successCount)
	if len(quoteFaild) > 0 {
		jobCh <- Job{ID: job.ID, Data: quoteFaild, IsFaild: true}
	}
	resultCh <- Job{ID: job.ID, Data: quotesResult}
}

// upload to store
func UploadQuote(job Job, resultCh chan interface{}, jobCh chan interface{}) {
	client := store.GetStoreClient()
	quotesDO := job.Data.([]*QuoteDO)
	uploadDo := &do.CreateMultipleItemDO{
		CategoryID: quotesDO[0].TopicID,
	}

	if job.IsFaild {
		log.Printf("===reCrawl job upload faild")
	}

	for _, quoteDO := range quotesDO {
		customFields := make([]*do.StoreCustomFieldTableDO, 0)
		customFields = append(customFields, &do.StoreCustomFieldTableDO{
			CustomFieldID:    AuthorCustomFieldID,
			CustomFieldValue: quoteDO.Author,
		})

		customFields = append(customFields, &do.StoreCustomFieldTableDO{
			CustomFieldID:    TagCustomFieldID,
			CustomFieldValue: strings.Join(quoteDO.Tag, "|"),
		})

		uploadDo.Items = append(uploadDo.Items, &do.ItemDO{
			Name:              quoteDO.Content,
			CustomFieldTables: customFields,
		})
	}

	storeLock.Lock()
	itemsFromStore := GetItems(uploadDo.CategoryID)
	storeLock.Unlock()
	items := uploadDo.Items

	uploadDo.Items = GetItemsNotInStore(items, itemsFromStore)

	storeLock.Lock()
	res := client.CreateMutilpleItem(StoreID, uploadDo)
	storeLock.Unlock()
	if res == nil {
		log.Printf("===upload faild quote: %s", uploadDo.Items)
		jobCh <- Job{ID: job.ID, Data: quotesDO, IsFaild: true}
	}
}

func main() {
	start := time.Now()
	defer func() {
		end := time.Since(start)
		fmt.Println("duration:", end)

	}()

	const numJobs = 1000

	store.InstallStoreClient()
	client := store.GetStoreClient()
	client.Token = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwOlwvXC9zdG9yZXMudm9saW8udm5cL2FwaVwvbG9naW4iLCJpYXQiOjE2ODQ4MjAzMzgsImV4cCI6MTY4ODQyMDMzOCwibmJmIjoxNjg0ODIwMzM4LCJqdGkiOiJnaDd5YkZoWXRoMUdROVcxIiwic3ViIjo0NywicHJ2IjoiMjNiZDVjODk0OWY2MDBhZGIzOWU3MDFjNDAwODcyZGI3YTU5NzZmNyJ9.x6i5w_e8CpA_4OZsjmrZ34ty8zPZu8h2HsEAbWmPXwo"

	// create category
	cateToCreate := GetCategories(client)

	jobCh1 := make(chan interface{}, numJobs)
	resultCh1 := make(chan interface{}, numJobs)
	resultCh2 := make(chan interface{}, numJobs)
	resultCh3 := make(chan interface{}, numJobs)
	resultCh4 := make(chan interface{}, numJobs)

	wp1 := &WorkerPool{
		NumWorkers: 1,
		JobCh:      jobCh1,
		ResultCh:   resultCh1,
		ProcessJob: func(job interface{}, resultCh chan interface{}, JobCh chan interface{}) {
			CreateCategory(job.(Job), resultCh)
		},
	}

	var allWg sync.WaitGroup
	allWg.Add(1)
	go func() {
		defer allWg.Done()

		wp1.Run()
	}()

	// init job to pool
	go func() {
		for i, topic := range cateToCreate {
			uploadDO := &do.CategoryDO{
				ID:       0,
				Name:     topic.Name,
				ModuleID: ModuleID,
				Status:   1,
				Url:      topic.Url,
			}

			jobCh1 <- Job{ID: i, Data: uploadDO}

		}

		close(jobCh1)
	}()

	wp2 := &WorkerPool{
		NumWorkers: 5,
		JobCh:      wp1.ResultCh,
		ResultCh:   resultCh2,
		ProcessJob: func(job interface{}, resultCh chan interface{}, JobCh chan interface{}) {
			GetUrlQuotes(job.(Job), resultCh)
		},
	}

	allWg.Add(1)
	go func() {
		defer allWg.Done()
		wp2.Run()
	}()

	wp3 := &WorkerPool{
		NumWorkers: 10,
		JobCh:      wp2.ResultCh,
		ResultCh:   resultCh3,
		ProcessJob: func(job interface{}, resultCh chan interface{}, JobCh chan interface{}) {
			CrawlQuote(job.(Job), resultCh, JobCh)
		},
	}

	allWg.Add(1)
	go func() {
		defer allWg.Done()
		wp3.Run()
	}()

	wp4 := &WorkerPool{
		NumWorkers: 5,
		JobCh:      wp3.ResultCh,
		ResultCh:   resultCh4,
		ProcessJob: func(job interface{}, resultCh chan interface{}, JobCh chan interface{}) {
			UploadQuote(job.(Job), resultCh, JobCh)
		},
	}

	allWg.Add(1)
	go func() {
		defer allWg.Done()
		wp4.Run()
	}()

	allWg.Wait()

}
