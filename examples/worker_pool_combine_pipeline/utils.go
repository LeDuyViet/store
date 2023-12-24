package main

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/rainycape/unidecode"
	"gitlab.volio.vn/tech/backend/store-pro/pkg/store"
	"gitlab.volio.vn/tech/backend/store-pro/pkg/store/do"
)

const (
	BaseURL             = "https://www.brainyquote.com"
	StoreID             = 6
	ModuleID            = 7
	AuthorCustomFieldID = 18
	TagCustomFieldID    = 19
)

var storeLock sync.Mutex

// Worker represents a worker that processes jobs
type Worker struct {
	ID     int
	JobCh  chan interface{}
	Result chan interface{}
}

// Job represents a unit of work
type Job struct {
	ID      int
	Data    interface{}
	IsFaild bool
}

type Topic struct {
	Url  string
	Name string
}

type QuoteDO struct {
	TopicID int32
	Content string
	Author  string
	Tag     []string
	Url     string
}

// remove all accents \n
// remove all accents and special characters
func removeAccent(s string) string {
	// convert Unicode characters to ASCII
	s = unidecode.Unidecode(s)
	// replace all non-alphanumeric characters with dashes
	s = regexp.MustCompile("[^a-zA-Z0-9]+").ReplaceAllString(s, "-")
	// remove leading and trailing dashes
	s = regexp.MustCompile("^-|-$").ReplaceAllString(s, "")
	return s
}

func removeBaseUrl(url string) string {
	return strings.Replace(url, BaseURL, "", 1)
}

// tìm phần tử thuộc arr2 nhưng không có trong arr1
func findMissingElements(arr1 []string, arr2 []string) []string {
	found := make(map[string]bool)
	missing := []string{}
	for _, num := range arr1 {
		found[num] = true
	}
	for _, num := range arr2 {
		if !found[num] {
			missing = append(missing, num)
		}
	}
	return missing
}

func GetItemsNotInStore(items []*do.ItemDO, itemsFromStore []*do.ItemDO) []*do.ItemDO {
	// Tạo một slice mới để lưu trữ các phần tử không xuất hiện trong itemsFromStore
	result := make([]*do.ItemDO, 0)
	// Tạo một map để lưu trữ tên của các phần tử trong itemsFromStore
	nameMap := make(map[string]bool)
	for _, item := range itemsFromStore {
		nameMap[item.Name] = true
	}
	// Lặp qua các phần tử trong items và kiểm tra xem tên của phần tử đó có xuất hiện trong itemsFromStore không
	// Nếu tên không xuất hiện, thêm phần tử đó vào slice kết quả
	for _, item := range items {
		if !nameMap[item.Name] {
			result = append(result, item)
		}
	}
	return result
}

func findMissingElementsByBits(arr1 []int, arr2 []int) []int {
	maxVal := 0
	for _, num := range arr1 {
		if num > maxVal {
			maxVal = num
		}
	}
	bitset := make([]byte, (maxVal/8)+1)
	for _, num := range arr1 {
		idx := num / 8
		offset := uint(num % 8)
		bitset[idx] |= 1 << offset
	}
	missing := []int{}
	for _, num := range arr2 {
		idx := num / 8
		offset := uint(num % 8)
		if (bitset[idx] & (1 << offset)) == 0 {
			missing = append(missing, num)
		}
	}
	return missing
}

func findMissingElementsString(arr1 []string, arr2 []string) []string {
	// Tạo một bitset có độ dài bằng với số lượng ký tự ASCII.
	bitset := make([]uint64, 4)
	// Đặt các bit tương ứng với các ký tự trong arr1 thành 1.
	for _, str := range arr1 {
		for _, ch := range str {
			idx := ch / 64
			offset := uint(ch % 64)
			bitset[idx] |= 1 << offset
		}
	}
	// Tìm các phần tử của arr2 không thuộc arr1.
	missing := []string{}
	for _, str := range arr2 {
		found := true
		for _, ch := range str {
			idx := ch / 64
			offset := uint(ch % 64)
			if (bitset[idx] & (1 << offset)) != 0 {
				found = false
				break
			}
		}
		if found {
			missing = append(missing, str)
		}
	}
	return missing
}

func GetNames1(items interface{}) ([]string, error) {
	value := reflect.ValueOf(items)
	if value.Kind() != reflect.Slice {
		return nil, errors.New("input is not a slice")
	}
	names := make([]string, value.Len())
	for i := 0; i < value.Len(); i++ {
		item := value.Index(i)
		nameField := item.FieldByName("Name")
		if !nameField.IsValid() {
			return nil, errors.New("item does not have a Name field")
		}
		name, ok := nameField.Interface().(string)
		if !ok {
			return nil, errors.New("Name field is not a string")
		}
		names[i] = name
	}
	return names, nil
}

func GetNames(items interface{}) ([]string, error) {
	value := reflect.ValueOf(items)
	if value.Kind() != reflect.Slice {
		return nil, errors.New("input is not a slice")
	}
	if value.Type().Elem().Kind() == reflect.Ptr {
		// slice of pointers, dereference them
		dereferenced := reflect.MakeSlice(reflect.SliceOf(value.Type().Elem().Elem()), value.Len(), value.Len())
		for i := 0; i < value.Len(); i++ {
			dereferenced.Index(i).Set(value.Index(i).Elem())
		}
		value = dereferenced
	}
	names := make([]string, value.Len())
	for i := 0; i < value.Len(); i++ {
		item := value.Index(i)
		nameField := item.FieldByName("Name")
		if !nameField.IsValid() {
			return nil, errors.New("item does not have a Name field")
		}
		name, ok := nameField.Interface().(string)
		if !ok {
			return nil, errors.New("Name field is not a string")
		}
		names[i] = name
	}
	return names, nil
}

func ExtractNumberFromURL(url string) (int, error) {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(url, -1)
	if len(matches) == 0 {
		return 0, fmt.Errorf("No numbers found in URL")
	}
	num, err := strconv.Atoi(matches[len(matches)-1])
	if err != nil {
		return 0, fmt.Errorf("Failed to convert number to int")
	}
	return num, nil
}

func MapTopics(cateToCreate []string, topics []*Topic) []*Topic {
	mappedTopics := []*Topic{}

	for _, cate := range cateToCreate {
		for _, topic := range topics {
			if topic.Name == cate {
				mappedTopic := &Topic{
					Url:  topic.Url,
					Name: topic.Name,
				}
				mappedTopics = append(mappedTopics, mappedTopic)
				break
			}
		}
	}

	return mappedTopics
}

// lấy phần tử của arr có trong toFilters
func FilterByList(arr []string, toFilters []string) []string {
	toFiltersMap := make(map[string]bool)
	for _, v := range toFilters {
		toFiltersMap[v] = true
	}

	result := []string{}
	for _, v := range arr {
		if toFiltersMap[v] {
			result = append(result, v)
		}
	}

	return result
}

func makeQuoteDO(cateCrawl []string, category *do.CategoryDO) []*QuoteDO {
	quoteDO := make([]*QuoteDO, 0)
	for i := 0; i < len(cateCrawl); i++ {
		quoteDO = append(quoteDO, &QuoteDO{
			TopicID: category.ID,
			Url:     cateCrawl[i],
		})
	}
	return quoteDO
}

// get items

var items []*do.ItemDO

func GetItems(categoryID int32) []*do.ItemDO {
	if items != nil {
		return items
	}
	items := make([]*do.ItemDO, 0)
	client := store.GetStoreClient()
	items = client.GetItemsPrivate(StoreID, categoryID)

	return items
}
