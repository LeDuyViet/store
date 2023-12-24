package models

import "fmt"

const (
	KTypeStoreData        = "STORE"
	KTypeAppData          = "APP"
	KTypeModule           = "MODULE"
	KTypeCategory         = "CATEGORY"
	KTypeItem             = "ITEM"
	KTypeCustomField      = "CUSTOM_FIELD"
	KTypeCustomFieldTable = "CUSTOM_FIELD_TABLE"
)

type DataMigratesMD struct {
	StoreID     int32  `db:"store_id" json:"store_id,omitempty"`
	MigrationID string `db:"migration_id" json:"migration_id,omitempty"`
	DataID      int32  `db:"data_id" json:"data_id,omitempty"`
	DataType    string `db:"data_type" json:"data_type,omitempty"`
	IsSync      bool   `db:"is_sync" json:"is_sync,omitempty"`
	CreatedTime int32  `db:"created_time" json:"created_time,omitempty"` // seconds UTC
	UpdatedTime int32  `db:"updated_time" json:"updated_time,omitempty"` // seconds UTC
}

func CreateDataMigrageKey(id string) string {
	return fmt.Sprintf("%s:%s", "DataMigratesMD", id)
}

func CreateDataMigrageKeyWithParam(dataMigrate *DataMigratesMD) string {
	return fmt.Sprintf("%d:%d:%s", dataMigrate.StoreID, dataMigrate.DataID, dataMigrate.DataType)
}
