package categories_handler

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/models"
	"gitlab.volio.vn/tech/fountain/baselib/base"
	"gitlab.volio.vn/tech/fountain/baselib/v_net/v_api"
	"gitlab.volio.vn/tech/fountain/proto/v_proto"
)

func (api *CategoriesAPI) Gets(ctx *fiber.Ctx) error {
	id := ctx.Query("id")

	if id != "" {
		res := api.categoriesController.Get(id, models.KTypeCategory)
		if res == nil {
			return v_api.WriteError(ctx, v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_NOT_FOUND), "not found category"))
		}

		return v_api.WriteSuccess(ctx, res)
	}

	limit := base.StringToInt(ctx.Query("limit", "50"))
	if limit == 0 {
		limit = 50
	}

	offset := base.StringToInt(ctx.Query("offset", "0"))
	if offset < 0 {
		offset = 0
	}

	appID := ctx.Query("app_id")
	if appID != "" {
		if v, err := base.StringToBool(ctx.Query("only_root")); v && err == nil {
			res := api.categoriesController.GetByApp(appID, offset, limit, true)
			return v_api.WriteSuccess(ctx, res)
		}

		res := api.categoriesController.GetByApp(appID, offset, limit)
		return v_api.WriteSuccess(ctx, res)
	}

	parentID := ctx.Query("parent_id")
	if parentID != "" {
		if limit != 50 || offset != 0 {
			return v_api.WriteError(ctx, v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), "limit offset is not support with parent_id"))
		}
		res := api.categoriesController.GetByParent(parentID)
		return v_api.WriteSuccess(ctx, res)
	}

	if ctx.Query("only_root") != "" {
		err := v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), "only_root is only support with app_id")
		return v_api.WriteError(ctx, err)
	}

	res := api.categoriesController.Gets(offset, limit, models.KTypeCategory)

	return v_api.WriteSuccess(ctx, res)
}
