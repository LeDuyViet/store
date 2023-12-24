package main

import (
	"log"

	"gitlab.volio.vn/tech/backend/store-pro/pkg/store"
)

func GetCategories(client *store.Client) []*Topic {
	topics := CrawlTopics()
	popularTopics := []string{"Love", "Happiness", "Success", "Friendship", "Health", "Family", "Education", "Motivational", "Wisdom", "Money", "Time", "Trust", "Beauty", "Communication", "Independence", "Hope", "Change", "Inspirational", "Respect", "Peace", "Relationship", "Parenting", "Technology", "Nature", "Music", "Movies", "Sports", "Travel", "Food", "Courage", "Science", "Art", "Humor", "Forgiveness", "Faith", "Leadership", "Dreams"}
	// popularTopics := []string{"Love"}

	storeCategories := client.GetCategoriesPrivate(StoreID, ModuleID)
	cateNames, err := GetNames(storeCategories)
	if err != nil {
		log.Println(err)
	}

	topicNames, err := GetNames(topics)
	if err != nil {
		log.Println(err)
	}

	cateToCreate := findMissingElements(cateNames, topicNames)

	cateToCreate = FilterByList(cateToCreate, popularTopics)

	// filter cateToCreate in popularTopics

	return MapTopics(cateToCreate, topics)
}
