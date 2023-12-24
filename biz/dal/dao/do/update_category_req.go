package do

type UpdateCategoryReq struct {
	ID string `db:"id" json:"id,omitempty"`

	AppID    string `db:"app_id" json:"app_id,omitempty" validate:"required,uuid4"`
	Name     string `db:"name" json:"name,omitempty" validate:"required,gte=3,lte=128"`
	ParentID string `db:"parent_id" json:"parent_id,omitempty" validate:"required,uuid4"`
	Priority int32  `db:"priority" json:"priority,omitempty"`
	Icon     string `db:"icon" json:"icon,omitempty"`
	Status   bool   `db:"status" json:"status,omitempty"`
	IsPro    bool   `db:"is_pro" json:"is_pro,omitempty"`
	IsNew    bool   `db:"is_new" json:"is_new,omitempty"`

	CustomFields []*CustomFieldTable `db:"-" json:"custom_fields"`
}
