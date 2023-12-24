package do

import (
	"fmt"

	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/models"
	"gitlab.volio.vn/tech/fountain/baselib/base"
)

type ModulesPublicDo struct {
	ID           int32                      `db:"id" json:"id"`
	AppID        int32                      `db:"app_id" json:"app_id"`
	Priority     int8                       `db:"priority" json:"priority"`
	Name         string                     `db:"name" json:"name"`
	Thumbnail    *string                    `db:"thumbnail" json:"thumbnail"`
	Icon         *string                    `db:"icon" json:"icon"`
	Status       int8                       `db:"status" json:"status"`
	IsPro        int8                       `db:"is_pro" json:"is_pro"`
	IsNew        int8                       `db:"is_new" json:"is_new"`
	CustomFields []*StoreCustomFieldTableDO `db:"custom_fields" json:"custom_fields"`
}

func CreateModulesPublicDo(m *models.ModuleMD, custom_fields []*StoreCustomFieldTableDO, IsMobile bool) *ModulesPublicDo {
	var (
		thumnail *string = nil
		icon     *string = nil
	)

	if m.Thumbnail != "" {
		thumb := ""
		if IsMobile {
			thumb = fmt.Sprintf("%s/%s", BASE_URL_STORE, m.Thumbnail)
		} else {
			thumb = m.Thumbnail
		}
		thumnail = &thumb
	}

	if m.Icon != "" {
		i := ""
		if IsMobile {
			i = fmt.Sprintf("%s/%s", BASE_URL_STORE, m.Icon)
		} else {
			i = m.Icon
		}
		icon = &i
	}

	return &ModulesPublicDo{
		ID:           m.DataMigratesMD.DataID,
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
