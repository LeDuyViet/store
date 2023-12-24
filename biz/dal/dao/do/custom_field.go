package do

type CustomFieldTable struct {
	ID                       string `db:"id" json:"id,omitempty"`
	CustomFieldID            string `db:"custom_field_id" json:"custom_field_id,omitempty"`
	CustomFieldValue         string `db:"value" json:"custom_field_value,omitempty"`
	IsActive                 bool   `db:"is_active" json:"is_active,omitempty"`
	IsActiveInt              int8   `db:"-" json:"is_active,omitempty"`
	CustomFieldTableableID   string `db:"custom_field_tableable_id" json:"custom_field_tableable_id,omitempty"`
	CustomFieldTableableType string `db:"custom_field_tableable_type" json:"custom_field_tableable_type,omitempty"`
	Name                     string `db:"name" json:"name,omitempty"`
	Type                     int8   `db:"type" json:"type,omitempty"`
}

// sử dụng do cho các app sử dụng store cũ
type StoreCustomFieldTableDO struct {
	ID               int32  `db:"id" json:"id"`
	CustomFieldID    int32  `db:"custom_field_id" json:"custom_field_id"`
	CustomFieldValue string `db:"value" json:"custom_field_value"`
	IsActive         int8   `db:"-" json:"is_active"`
	Name             string `db:"name" json:"name"`
	Type             int8   `db:"type" json:"type"`
}
