package admin

import (
	"fmt"
	"reflect"

	"gitlab.volio.vn/tech/backend/store-pro/biz/core"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/dao/do"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/models"
	"gitlab.volio.vn/tech/fountain/baselib/v_log"
	"gitlab.volio.vn/tech/fountain/proto/v_proto"
)

const (
	KCustomFieldTablesKey = "%s_%s_%s"
)

func createCustomfieldTablesKey(CustomFieldID, CustomFieldTableableID, CustomFieldTableableType string) string {
	if CustomFieldID == "" || CustomFieldTableableID == "" || CustomFieldTableableType == "" {
		v_log.V(1).Errorf("createCustomfieldTablesKey - Error: %+v", "Invalid CustomFieldID, CustomFieldTableableID or CustomFieldTableableType")
	}
	return fmt.Sprintf(KCustomFieldTablesKey, CustomFieldID, CustomFieldTableableID, CustomFieldTableableType)
}

func (ctrl *AdminController) validateCustomFields(input interface{}, modelType string) *v_proto.VolioRpcError {
	ctfRaw := reflect.ValueOf(input).Elem().FieldByName("CustomFields")
	var customFieldsDO []*do.CustomFieldTable

	if ctfRaw.IsValid() {
		customFieldsDO = ctfRaw.Interface().([]*do.CustomFieldTable)
	} else {
		return nil
	}

	mCustomFieldsCount := make(map[string]int)

	for i := 0; i < len(customFieldsDO); i++ {
		if customField := ctrl.customFieldsCallback.GetCustomField(customFieldsDO[i].CustomFieldID); customField == nil {
			return v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), "not found custom field")
		}

		key := createCustomfieldTablesKey(customFieldsDO[i].CustomFieldID, "any", modelType)
		mCustomFieldsCount[key]++
		if mCustomFieldsCount[key] > 1 {
			return v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), "duplicate custom field")
		}
	}

	return nil
}

func (ctrl *AdminController) updateOrCreateCustomfields(input interface{}, modelType string) *v_proto.VolioRpcError {
	ctfRaw := reflect.ValueOf(input).Elem().FieldByName("CustomFields")
	var customFieldTablesDO []*do.CustomFieldTable
	var inputID string

	if ctfRaw.IsValid() {
		customFieldTablesDO = ctfRaw.Interface().([]*do.CustomFieldTable)
	} else {
		return nil
	}

	idRaw := reflect.ValueOf(input).Elem().FieldByName("ID")
	if idRaw.IsValid() {
		inputID = idRaw.Interface().(string)
		if inputID == "" {
			panic("Invalid ID for create or update custom field table. Forgot set ID after create or update?")
		}
	} else {
		v_log.V(1).WithError(v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), "not found in this request"))
		return nil
	}

	toCreate := make([]*do.CustomFieldTable, 0)
	toUpdate := make([]*do.CustomFieldTable, 0)
	mCustomFields := make(map[string]*models.CustomFieldDataMD)

	for i := 0; i < len(customFieldTablesDO); i++ {
		customFieldTablesDO[i].CustomFieldTableableID = inputID
		customFieldTablesDO[i].CustomFieldTableableType = modelType
	}

	mds := ctrl.customFieldsCallback.GetByTableableData(inputID, modelType)

	for _, md := range mds {
		key := createCustomfieldTablesKey(md.CustomFieldID, "md.CustomFieldTableableID", "md.CustomFieldTableableType")
		mCustomFields[key] = md
	}

	for _, customField := range customFieldTablesDO {
		key := createCustomfieldTablesKey(customField.CustomFieldID, customField.CustomFieldTableableID, customField.CustomFieldTableableType)

		if v, ok := mCustomFields[key]; !ok {
			toCreate = append(toCreate, &do.CustomFieldTable{
				CustomFieldID:            customField.CustomFieldID,
				CustomFieldValue:         customField.CustomFieldValue,
				IsActive:                 customField.IsActive,
				CustomFieldTableableID:   inputID,
				CustomFieldTableableType: modelType,
			})

			delete(mCustomFields, key)
		} else {
			if v.Value != customField.CustomFieldValue {
				toUpdate = append(toUpdate, &do.CustomFieldTable{
					ID:                       v.ID,
					CustomFieldID:            customField.CustomFieldID,
					CustomFieldValue:         customField.CustomFieldValue,
					IsActive:                 customField.IsActive,
					CustomFieldTableableID:   inputID,
					CustomFieldTableableType: modelType,
				})
			}

			delete(mCustomFields, key)
		}
	}

	if len(toCreate) > 0 {
		if err := ctrl.customFieldsCallback.InsertManyData(toCreate); err != nil {
			v_log.V(1).WithError(err).Errorf("AdminController::createCustomfields - InsertMany - Error: %+v", err)
			return err
		}
	}

	if len(toUpdate) > 0 {
		if err := ctrl.customFieldsCallback.UpdateManyData(toUpdate); err != nil {
			v_log.V(1).WithError(err).Errorf("AdminController::createCustomfields - UpdateMany - Error: %+v", err)
			return err
		}
	}

	if len(mCustomFields) > 0 {
		for _, value := range mCustomFields {
			if err := ctrl.customFieldsCallback.DeleteData(value.ID); err != nil {
				v_log.V(1).WithError(err).Errorf("AdminController::createCustomfields - Update - Error: %+v", err)
				return err
			}
		}
	}

	return nil
}

func (ctrl *AdminController) deleteCustomFieldByModel(id, modelType string) *v_proto.VolioRpcError {
	mds := ctrl.customFieldsCallback.GetByTableableData(id, modelType)

	if len(mds) == 0 {
		return nil
	}

	ids := []string{}
	for _, md := range mds {
		ids = append(ids, md.ID)
	}

	if err := ctrl.customFieldsCallback.DeleteData(ids...); err != nil {
		v_log.V(1).WithError(err).Errorf("AdminController::deleteCustomFieldByModel - Error: %+v", err)
		return err
	}

	return nil
}

func (ctrl *AdminController) getController(modelType string) core.Admin {
	controller, ok := ctrl.adminControllers.Load(modelType)
	if !ok {
		panic("admin controller not found")
	}
	return controller.(core.Admin)
}
