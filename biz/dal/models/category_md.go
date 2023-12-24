package models

type CategoryMD struct {
	ID          string `db:"id" json:"id,omitempty"`
	ParentID    string `db:"parent_id" json:"parent_id,omitempty"`
	Level       int8   `db:"level" json:"level,omitempty"`
	AppID       string `db:"app_id" json:"app_id,omitempty"`
	Priority    int32  `db:"priority" json:"priority,omitempty"`
	Name        string `db:"name" json:"name,omitempty"`
	Thumbnail   string `db:"thumbnail" json:"thumbnail,omitempty"`
	Icon        string `db:"icon" json:"icon,omitempty"`
	Status      bool   `db:"status" json:"status,omitempty"`
	IsPro       bool   `db:"is_pro" json:"is_pro,omitempty"`
	IsNew       bool   `db:"is_new" json:"is_new,omitempty"`
	CreatedTime int32  `db:"created_time" json:"created_time,omitempty"` // seconds UTC
	UpdatedTime int32  `db:"updated_time" json:"updated_time,omitempty"` // seconds UTC

	IdInt       int32 `db:"-" json:"-"`
	ModuleIDInt int32 `db:"-" json:"-"`

	DataMigratesMD *DataMigratesMD `db:"-" json:"data_migrate,omitempty"`

	CustomFields []*CustomFieldMD `db:"-" json:"custom_fields,omitempty"`
}
