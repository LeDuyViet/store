package items_handler

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.volio.vn/tech/backend/store-pro/pkg/store/store_api"
)

func (api *ItemsAPI) GetAll(ctx *fiber.Ctx) error {
	// storeId := ctx.Params("store_id")
	// storeIdInt, err := strconv.Atoi(storeId)
	// if err != nil {
	// 	return v_api.WriteError(ctx, v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), err.Error()))
	// }
	// appId := ctx.Params("app_id")
	// if appId == "" {
	// 	err := v_proto.NewRpcError(400, "app_id is required!")
	// 	return store_api.WriteError(ctx, err)
	// }

	// _, err = strconv.Atoi(appId) // appIdInt
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

	// moduleId := ctx.Query("module_id", "")
	// if moduleId == "" {
	// 	err := v_proto.NewRpcError(400, "module_id is required!")
	// 	return store_api.WriteError(ctx, err)
	// }
	// _, err = strconv.Atoi(appId) //moduleIDInt
	// if err != nil {
	// 	return v_api.WriteError(ctx, v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), err.Error()))
	// }

	// categoryId := ctx.Query("category_id", "")
	// if categoryId == "" {
	// 	err := v_proto.NewRpcError(400, "category_id is required!")
	// 	return store_api.WriteError(ctx, err)
	// }

	// categoryIdInt, err := strconv.Atoi(categoryId)
	// if err != nil {
	// 	return v_api.WriteError(ctx, v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), err.Error()))
	// }

	// region := ctx.Query("region", "")
	// if region == "" {
	// 	resp := store.GetStoreClient().Forward(ctx)
	// 	return ctx.Send(resp)
	// }

	// category := api.dataMigratesController.GetByDataMigrate(&models.DataMigratesMD{
	// 	StoreID:  int32(storeIdInt),
	// 	DataID:   int32(categoryIdInt),
	// 	DataType: models.KTypeCategory,
	// })

	// if category == nil {
	// 	res := v_proto.NewRpcError(400, fmt.Sprintf("category_id %v not exist!", categoryIdInt))
	// 	return v_api.WriteError(ctx, res)
	// }

	// res := api.itemsController.GetItemsbyCategoryIdForAdmin(category, perPage, page)

	return store_api.WriteSuccess(ctx, nil)
}
