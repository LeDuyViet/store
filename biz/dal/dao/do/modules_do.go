package do

type ModulesDo struct {
	ID           string              `db:"id" json:"id"`
	AppID        string              `db:"app_id" json:"app_id"`
	Priority     int32               `db:"priority" json:"priority"`
	Name         string              `db:"name" json:"name"`
	Thumbnail    string              `db:"thumbnail" json:"thumbnail"`
	Icon         string              `db:"icon" json:"icon"`
	Status       bool                `db:"status" json:"status"`
	IsPro        bool                `db:"is_pro" json:"is_pro"`
	IsNew        bool                `db:"is_new" json:"is_new"`
	CustomFields []*CustomFieldTable `db:"custom_fields" json:"custom_fields"`
}
