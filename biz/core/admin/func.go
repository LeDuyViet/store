package admin

import (
	"gitlab.volio.vn/tech/fountain/baselib/v_log"
	"gitlab.volio.vn/tech/fountain/proto/v_proto"
)

func (ctrl *AdminController) Get(ID string, modelType string) interface{} {
	controller := ctrl.getController(modelType)

	res := controller.AdminGet(ID)

	// TODO: get custom fields

	return res
}

func (ctrl *AdminController) Gets(limit, offset int, modelType string) interface{} {
	controller := ctrl.getController(modelType)

	res := controller.AdminGets(limit, offset)

	// TODO: get custom fields

	return res
}

func (ctrl *AdminController) Create(input interface{}, modelType string) *v_proto.VolioRpcError {
	controller := ctrl.getController(modelType)

	if err := ctrl.validateCustomFields(input, modelType); err != nil {
		v_log.V(1).WithError(err).Errorf("AdminController::Create - Error: %+v", err)
		return err
	}

	if err := controller.AdminCreate(input); err != nil {
		v_log.V(1).WithError(err).Errorf("AdminController::Create - Error: %+v", err)
		return err
	}

	if err := ctrl.updateOrCreateCustomfields(input, modelType); err != nil {
		v_log.V(1).WithError(err).Errorf("AdminController::Create - Error: %+v", err)
		return err
	}

	// TODO: log

	return nil
}

func (ctrl *AdminController) Update(input interface{}, modelType string) *v_proto.VolioRpcError {
	controller := ctrl.getController(modelType)

	if err := ctrl.validateCustomFields(input, modelType); err != nil {
		v_log.V(1).WithError(err).Errorf("AdminController::Create - Error: %+v", err)
		return err
	}

	if err := controller.AdminUpdate(input); err != nil {
		v_log.V(1).WithError(err).Errorf("AdminController::Update - Error: %+v", err)
		return err
	}

	if err := ctrl.updateOrCreateCustomfields(input, modelType); err != nil {
		v_log.V(1).WithError(err).Errorf("AdminController::Update - Error: %+v", err)
		return err
	}

	return nil
}

func (ctrl *AdminController) Delete(id string, modelType string) *v_proto.VolioRpcError {
	controller := ctrl.getController(modelType)

	if err := ctrl.deleteCustomFieldByModel(id, modelType); err != nil {
		v_log.V(1).WithError(err).Errorf("AdminController::Delete - Error: %+v", err)
		return err
	}

	if err := controller.AdminDelete(id); err != nil {
		v_log.V(1).WithError(err).Errorf("AdminController::Create - Error: %+v", err)
		return err
	}

	return nil
}
