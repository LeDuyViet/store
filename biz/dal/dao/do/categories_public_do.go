package do

import (
	"fmt"

	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/models"
	"gitlab.volio.vn/tech/fountain/baselib/base"
)

type CategoriesPublicDo struct {
	ID           int32   `db:"id" json:"id"`
	ModuleID     int32   `db:"module_id" json:"module_id"`
	Priority     int8    `db:"priority" json:"priority"`
	Name         string  `db:"name" json:"name"`
	Thumbnail    *string `db:"thumbnail" json:"thumbnail"`
	Icon         *string `db:"icon" json:"icon"`
	Status       int8    `db:"status" json:"status"`
	IsPro        int8    `db:"is_pro" json:"is_pro"`
	IsNew        int8    `db:"is_new" json:"is_new"`
	CustomFields []*StoreCustomFieldTableDO
}

func CreateCategoriesPublicDo(m *models.CategoryMD, custom_fields []*StoreCustomFieldTableDO) *CategoriesPublicDo {
	var (
		thumnail *string = nil
		icon     *string = nil
	)
	if m.Thumbnail != "" {
		thumb := fmt.Sprintf("%s/%s", BASE_URL_STORE, m.Thumbnail)
		thumnail = &thumb
	}
	if m.Icon != "" {
		i := fmt.Sprintf("%s/%s", BASE_URL_STORE, m.Icon)
		icon = &i
	}
	// for _, custom_field := range custom_fields {
	// 	custom_field.IsActiveInt = base.BoolToInt8(custom_field.IsActive)
	// }
	return &CategoriesPublicDo{
		ID:           m.IdInt,
		ModuleID:     m.ModuleIDInt,
		Priority:     int8(m.Priority),
		Name:         m.Name,
		Thumbnail:    thumnail,
		Icon:         icon,
		Status:       base.BoolToInt8(m.Status),
		IsPro:        base.BoolToInt8(m.IsPro),
		IsNew:        base.BoolToInt8(m.IsNew),
		CustomFields: custom_fields,
	}
}
