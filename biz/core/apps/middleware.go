package apps

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"gitlab.volio.vn/tech/fountain/baselib/validor"
)

func (ctrl *AppsController) registerCustomValidorRule() {
	// thêm -rule ở cuối để nhấn mạnh đây là rule tự tạo
	validorIns := validor.NewValidator()
	validorIns.RegisterRule("unique-package-name-rule", ctrl.checkUniquePackageName, "exist package name for this request")
	validorIns.RegisterRule("exist-app-rule", ctrl.checkExistApp, "not found app for this request")
}

func (ctrl *AppsController) checkUniquePackageName(fl validator.FieldLevel) bool {
	parentValue := reflect.ValueOf(fl.Parent().Interface())
	currentID := parentValue.FieldByName("ID").String()

	packageName := fl.Field().String()

	apps := ctrl.dao.appsDAO.GetByPackageName(packageName)
	// ignore id của bản ghi hiện tại
	if len(apps) == 1 && apps[0].ID == currentID {
		return true
	}

	return len(apps) == 0
}

func (ctrl *AppsController) checkExistApp(fl validator.FieldLevel) bool {
	appID := fl.Field().String()
	return ctrl.dao.appsDAO.GetApp(appID) != nil
}
