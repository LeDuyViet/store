package stores

import (
	"time"

	"github.com/google/uuid"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/dao/do"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/models"
	"gitlab.volio.vn/tech/fountain/proto/v_proto"
)

func (ctrl *StoresController) AdminGet(id string) interface{} {
	return ctrl.dao.storesDAO.GetStore(id)
}

func (ctrl *StoresController) AdminCreate(reqInterface interface{}) *v_proto.VolioRpcError {
	// Need gen new hash for AccessKey

	req := reqInterface.(*do.CreateStoreReq)

	if req == nil {
		return v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), "not found object to create")
	}

	timeNow := int32(time.Now().Unix())
	md := &models.StoreMD{
		ID:          uuid.NewString(),
		Name:        req.Name,
		Address:     req.Address,
		IsActive:    req.IsActive,
		AccessKey:   uuid.NewString(),
		CreatedTime: timeNow,
		UpdatedTime: timeNow,
	}

	req.ID = md.ID

	if err := ctrl.dao.storesDAO.InsertStore(md); err != nil {
		return v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_SAVE_DATA), err.Error())
	}

	return nil
}

func (ctrl *StoresController) AdminUpdate(reqInterface interface{}) *v_proto.VolioRpcError {
	req := reqInterface.(*do.UpdateStoreReq)

	if req == nil {
		return v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), "not found object to update")
	}

	md := &models.StoreMD{
		ID:          req.ID,
		Name:        req.Name,
		Address:     req.Address,
		IsActive:    req.IsActive,
		UpdatedTime: int32(time.Now().Unix()),
	}

	if err := ctrl.dao.storesDAO.UpdateStore(md); err != nil {
		return v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_SAVE_DATA), err.Error())
	}

	return nil
}

func (ctrl *StoresController) AdminDelete(id string) *v_proto.VolioRpcError {
	if err := ctrl.dao.storesDAO.DeleteStore(id); err != nil {
		return v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_SAVE_DATA), "can not delete store")
	}

	return nil
}
