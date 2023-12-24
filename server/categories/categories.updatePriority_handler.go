package categories_handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/models"
	"gitlab.volio.vn/tech/backend/store-pro/pkg/store"
	"gitlab.volio.vn/tech/backend/store-pro/pkg/store/store_api"
	"gitlab.volio.vn/tech/fountain/baselib/v_net/v_api"
	"gitlab.volio.vn/tech/fountain/proto/v_proto"
)

func (api *CategoriesAPI) UpdatePriority(ctx *fiber.Ctx) error {
	region := ctx.Query("region", "")

	payload := struct {
		Data []*models.CategoryMD `json:"data"`
	}{}

	payload.Data = make([]*models.CategoryMD, 0)

	if region == "" {
		resp := store.GetStoreClient().Forward(ctx)
		return ctx.Send(resp)
	}

	if err := ctx.BodyParser(&payload); err != nil {
		return v_api.WriteError(ctx, v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), err.Error()))
	}
	priorities := make(map[int32]bool, 0)
	for _, category := range payload.Data {
		if !api.categoriesController.CheckCategoryExist(category.ID) {
			return v_api.WriteError(ctx, v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), fmt.Sprintf("Category %v invalid ", category.ID)))
		}
		if ok := priorities[category.Priority]; ok == true {
			return v_api.WriteError(ctx, v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), fmt.Sprintf("Duplicate priorty value is %v ", category.Priority)))
		} else {
			priorities[category.Priority] = true
		}
	}
	err := api.categoriesController.UpdatePriority(payload.Data)

	if err != nil {
		// err := fmt.Errorf("modules.UpdatePriority: %v")
		return store_api.WriteError(ctx, err)
	}

	return store_api.WriteSuccess(ctx, nil)
}
