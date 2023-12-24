package custom_fields

import (
	"time"

	"github.com/google/uuid"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/dao/do"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/models"
	"gitlab.volio.vn/tech/fountain/proto/v_proto"
)

func (ctrl *CustomFieldsController) GetData(id string) *models.CustomFieldDataMD {
	return ctrl.dao.customFieldTablesDAO.Get(id)
}

func (ctrl *CustomFieldsController) GetByIDsData(ids []string) []*models.CustomFieldDataMD {
	return ctrl.dao.customFieldTablesDAO.GetByIDs(ids)
}

func (ctrl *CustomFieldsController) GetManyData(customFields []*do.CustomFieldTable) []*models.CustomFieldDataMD {
	return ctrl.dao.customFieldTablesDAO.GetMany(customFields)
}

func (ctrl *CustomFieldsController) GetByTableableData(CustomFieldTableableID, CustomFieldTableableType string) []*models.CustomFieldDataMD {
	return ctrl.dao.customFieldTablesDAO.GetByTableable(CustomFieldTableableID, CustomFieldTableableType)
}

func (ctrl *CustomFieldsController) DeleteData(id ...string) *v_proto.VolioRpcError {
	if len(id) == 0 {
		return v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), "custom field table id is required")
	}

	if len(id) == 1 {
		if err := ctrl.dao.customFieldTablesDAO.Delete(id[0]); err != nil {
			return v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), "err when delete custom field")
		}
	} else {
		if err := ctrl.dao.customFieldTablesDAO.DeleteMultiple(id); err != nil {
			return v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), "err when delete custom field")
		}
	}

	return nil
}

func (ctrl *CustomFieldsController) InsertData(md *models.CustomFieldDataMD) *v_proto.VolioRpcError {
	if md.ID == "" {
		md.ID = uuid.NewString()
		timeNow := int32(time.Now().Unix())
		md.CreatedTime = timeNow
		md.UpdatedTime = timeNow
	}

	if err := ctrl.dao.customFieldTablesDAO.Insert(md); err != nil {
		return v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), "err when insert custom field")
	}

	return nil
}

func (ctrl *CustomFieldsController) InsertManyData(do []*do.CustomFieldTable) *v_proto.VolioRpcError {
	mds := make([]*models.CustomFieldDataMD, 0)
	timeNow := int32(time.Now().Unix())
	for i := 0; i < len(do); i++ {

		mds = append(mds, &models.CustomFieldDataMD{
			ID:            uuid.NewString(),
			CustomFieldID: do[i].CustomFieldID,
			Value:         do[i].CustomFieldValue,
			// CustomFieldTableableID:   do[i].CustomFieldTableableID,
			// CustomFieldTableableType: do[i].CustomFieldTableableType,
			// IsActive:                 true,
			CreatedTime: timeNow,
			UpdatedTime: timeNow,
		})
	}

	if err := ctrl.dao.customFieldTablesDAO.InsertMany(mds); err != nil {
		return v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), "err when insert custom field")
	}

	return nil
}
func (ctrl *CustomFieldsController) UpdateManyData(do []*do.CustomFieldTable) *v_proto.VolioRpcError {
	mds := make([]*models.CustomFieldDataMD, 0)
	timeNow := int32(time.Now().Unix())

	for i := 0; i < len(do); i++ {
		mds = append(mds, &models.CustomFieldDataMD{
			ID:            do[i].ID,
			CustomFieldID: do[i].CustomFieldID,
			Value:         do[i].CustomFieldValue,
			// CustomFieldTableableID:   do[i].CustomFieldTableableID,
			// CustomFieldTableableType: do[i].CustomFieldTableableType,
			// IsActive:                 true,
			UpdatedTime: timeNow,
		})
	}

	if err := ctrl.dao.customFieldTablesDAO.UpdateMany(mds); err != nil {
		return v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), "err when update custom field")
	}

	return nil
}

func (ctrl *CustomFieldsController) UpdateData(md *models.CustomFieldDataMD) error {
	if md.ID == "" {
		md.ID = uuid.NewString()
		timeNow := int32(time.Now().Unix())
		md.CreatedTime = timeNow
		md.UpdatedTime = timeNow
	}

	return ctrl.dao.customFieldTablesDAO.Update(md)
}
