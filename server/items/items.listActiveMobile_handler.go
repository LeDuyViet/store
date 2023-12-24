package items_handler

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.volio.vn/tech/backend/store-pro/pkg/store/store_api"
)

// https://store.volio.vn/api/v2/public/categories?app_id=2&per_page=3&page=1
func (api *ItemsAPI) ListActive(ctx *fiber.Ctx) error {
	// region := ctx.Query("region", "")
	// if region == "" {
	// 	resp := store.GetStoreClient().Forward(ctx, true)
	// 	return ctx.Send(resp)
	// }

	// categoryID := ctx.Query("category", "")
	// if categoryID == "" {
	// 	err := v_proto.NewRpcError(400, "category is required!")
	// 	return store_api.WriteError(ctx, err)
	// }

	// categoryIDInt, err := strconv.Atoi(categoryID)
	// if err != nil {
	// 	return v_api.WriteError(ctx, v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), err.Error()))
	// }

	// perPage, err := strconv.Atoi(ctx.Query("per_page", "15"))
	// if err != nil {
	// 	perPage = 15
	// }

	// page, err := strconv.Atoi(ctx.Query("page", "1"))
	// if err != nil {
	// 	page = 1
	// }

	// storeID := v_api.GetContextDataString(ctx, stores.KStoreIDKey)

	// category := api.dataMigratesController.GetByDataMigrate(&models.DataMigratesMD{
	// 	StoreID:  base.StringToInt32(storeID),
	// 	DataID:   int32(categoryIDInt),
	// 	DataType: models.KTypeCategory,
	// })

	// if category == nil {
	// 	go func() {
	// 		api.dataMigratesController.MigrateData(base.StringToInt32(storeID), int32(14))
	// 	}()

	// 	resp := store.GetStoreClient().Forward(ctx, true)
	// 	return ctx.Send(resp)
	// }
	// res := api.itemsController.ListActive(category, perPage, page)

	return store_api.WriteSuccess(ctx, nil)
}
