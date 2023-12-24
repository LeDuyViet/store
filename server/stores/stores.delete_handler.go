package stores_handler

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/models"
	"gitlab.volio.vn/tech/fountain/baselib/v_log"
	"gitlab.volio.vn/tech/fountain/baselib/v_net/v_api"
	"gitlab.volio.vn/tech/fountain/proto/v_proto"
)

func (api *StoresAPI) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		err := v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), "store id is required")
		v_log.V(1).WithError(err).Errorf("StoreAPI::Delete - Error: %+v", err)

		return v_api.WriteError(ctx, err)
	}

	if err := api.storesController.Delete(id, models.KTypeStoreData); err != nil {
		v_log.V(1).WithError(err).Errorf("StoreAPI::Delete - Error: %+v", err)

		return v_api.WriteError(ctx, err)
	}

	v_log.V(3).Infof("StoreAPI::Delete - Reply: Oke")

	return v_api.WriteSuccess(ctx, nil)
}
