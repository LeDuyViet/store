package do

type CategoriesDO struct {
	ID        string `db:"id" json:"id,omitempty"`
	ParentID  string `db:"parent_id" json:"parent_id,omitempty"`
	Priority  int32  `db:"priority" json:"priority,omitempty"`
	Name      string `db:"name" json:"name,omitempty"`
	Thumbnail string `db:"thumbnail" json:"thumbnail,omitempty"`
	Icon      string `db:"icon" json:"icon,omitempty"`
	Status    bool   `db:"status" json:"status,omitempty"`
	IsPro     bool   `db:"is_pro" json:"is_pro,omitempty"`
	IsNew     bool   `db:"is_new" json:"is_new,omitempty"`

	CustomFields []*CustomFieldTable `db:"-" json:"custom_fields,omitempty"`
	Childrens    []*CategoriesDO     `db:"-" json:"childrens,omitempty"`
}
