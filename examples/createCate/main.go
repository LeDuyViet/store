package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"gitlab.volio.vn/tech/backend/store-pro/pkg/store"
	"gitlab.volio.vn/tech/backend/store-pro/pkg/store/do"
)

func main() {
	data := `{
		"is_pro": 0,
		"is_new": 0,
		"status": 1,
		"thumbnail": "stores/store-5/2023/05/26/thumbnails_1685091079_Screenshot 2023-04-06 134626.png",
		"icon": "stores/store-5/2023/05/26/1685091078_Screenshot 2023-04-06 134626.png",
		"name": "test",
		"module_id": 17,
		"custom_fields": []
	}`

	uploadDO := &do.CategoryDO{}
	json.Unmarshal([]byte(data), uploadDO)

	store.InstallStoreClient()
	client := store.GetStoreClient()
	client.Token = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwOlwvXC9zdG9yZXMudm9saW8udm5cL2FwaVwvbG9naW4iLCJpYXQiOjE2ODQ4MjAzMzgsImV4cCI6MTY4ODQyMDMzOCwibmJmIjoxNjg0ODIwMzM4LCJqdGkiOiJnaDd5YkZoWXRoMUdROVcxIiwic3ViIjo0NywicHJ2IjoiMjNiZDVjODk0OWY2MDBhZGIzOWU3MDFjNDAwODcyZGI3YTU5NzZmNyJ9.x6i5w_e8CpA_4OZsjmrZ34ty8zPZu8h2HsEAbWmPXwo"

	start := time.Now()
	defer func() {
		log.Printf("End create category, time: %v", time.Since(start))
	}()

	res := client.CreateCategory(5, uploadDO)

	fmt.Println(res.Code)

}
