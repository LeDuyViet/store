package regions_handler

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/models"
	"gitlab.volio.vn/tech/backend/store-pro/pkg/store/store_api"
	"gitlab.volio.vn/tech/fountain/baselib/v_log"
	"gitlab.volio.vn/tech/fountain/baselib/v_net/v_api"
	"gitlab.volio.vn/tech/fountain/proto/v_proto"
)

func (api *RegionAPI) AddApp(ctx *fiber.Ctx) error {
	input := &models.AppRegionMD{}
	if err := ctx.BodyParser(input); err != nil {
		err := v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), err.Error())
		v_log.V(1).WithError(err).Errorf("RegionAPI::AddApp - Error: %+v", err)

		return store_api.WriteError(ctx, err)
	}

	if input.RegionID == "" {
		err := v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), "not found region_id for this request")
		v_log.V(1).WithError(err).Errorf("StoreAPI::Create - Error: %+v", err)

		return store_api.WriteError(ctx, err)
	}

	if input.AppID == "" {
		err := v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), "not found app_id for this request")
		v_log.V(1).WithError(err).Errorf("StoreAPI::Create - Error: %+v", err)

		return store_api.WriteError(ctx, err)
	}

	err := api.regionsController.AddApp(input)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("RegionAPI::Add - Error: %+v", err)

		return store_api.WriteError(ctx, err)
	}

	return v_api.WriteSuccess(ctx, input)
}
