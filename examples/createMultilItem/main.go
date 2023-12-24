package main

import (
	"encoding/json"
	"fmt"

	"gitlab.volio.vn/tech/backend/store-pro/pkg/store"
	"gitlab.volio.vn/tech/backend/store-pro/pkg/store/do"
)

func main() {
	data := `{
		"category_id": "51",
		"items": [
			{
				"icon": "stores/store-5/2023/05/26/1685075158_Screenshot 2023-05-16 103057.png",
				"thumbnail": "stores/store-5/2023/05/26/thumbnails_1685075158_Screenshot 2023-05-16 103057.png",
				"name": "hi",
				"custom_fields": [
					{
						"custom_field_id": 96,
						"custom_field_value": "test"
					},
					{
						"custom_field_id": 94,
						"custom_field_value": "tetset"
					}
				]
			}
		]
	}`

	uploadDO := &do.CreateMultipleItemDO{}
	json.Unmarshal([]byte(data), uploadDO)

	store.InstallStoreClient()
	client := store.GetStoreClient()
	client.Token = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwOlwvXC9zdG9yZXMudm9saW8udm5cL2FwaVwvbG9naW4iLCJpYXQiOjE2ODQ4MjAzMzgsImV4cCI6MTY4ODQyMDMzOCwibmJmIjoxNjg0ODIwMzM4LCJqdGkiOiJnaDd5YkZoWXRoMUdROVcxIiwic3ViIjo0NywicHJ2IjoiMjNiZDVjODk0OWY2MDBhZGIzOWU3MDFjNDAwODcyZGI3YTU5NzZmNyJ9.x6i5w_e8CpA_4OZsjmrZ34ty8zPZu8h2HsEAbWmPXwo"
	res := client.CreateMutilpleItem(5, uploadDO)

	fmt.Println(res.Code)

}
