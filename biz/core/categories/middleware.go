package categories

import (
	"github.com/go-playground/validator/v10"
	"gitlab.volio.vn/tech/fountain/baselib/validor"
)

func (ctrl *CategoriesController) registerCustomValidorRule() {
	// thêm -rule ở cuối để nhấn mạnh đây là rule tự tạo
	validor.NewValidator().RegisterRule("exist-cate-rule", ctrl.CheckExistCate, "not found cate for this request")
}

func (ctrl *CategoriesController) CheckExistCate(fl validator.FieldLevel) bool {
	cateID := fl.Field().String()
	return ctrl.dao.categoriesDAO.GetCategory(cateID) != nil
}
