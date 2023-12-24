package apps_handler

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.volio.vn/tech/fountain/baselib/base"
	"gitlab.volio.vn/tech/fountain/baselib/v_log"
	"gitlab.volio.vn/tech/fountain/baselib/v_net/v_api"
	"gitlab.volio.vn/tech/fountain/proto/v_proto"
)

func (api *AppsAPI) GetCategories(ctx *fiber.Ctx) error {
	appID := ctx.Params("id", "")
	if appID == "" {
		err := v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), "store id is required")
		v_log.V(1).WithError(err).Errorf("StoreAPI::Delete - Error: %+v", err)

		return v_api.WriteError(ctx, err)
	}

	offset := base.StringToInt(ctx.Query("offset", "0"))
	if offset < 0 {
		offset = 0
	}

	limit := base.StringToInt(ctx.Query("limit", "50"))
	if limit == 0 {
		limit = 50
	}

	res := api.appsController.GetCategories(appID, offset, limit)

	return v_api.WriteSuccess(ctx, res)
}
