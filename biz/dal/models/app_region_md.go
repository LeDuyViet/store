package models

type AppRegionMD struct {
	AppID       string `db:"app_id" json:"app_id,omitempty"`
	RegionID    string `db:"region_id" json:"region_id,omitempty"`
	IsActive    int    `db:"is_active" json:"is_active,omitempty"`
	CreatedTime int32  `db:"created_time" json:"created_time,omitempty"` // seconds UTC
	UpdatedTime int32  `db:"updated_time" json:"updated_time,omitempty"` // seconds UTC
}
