package models

type RegionMD struct {
	ID          string `db:"id" json:"id,omitempty"`
	Name        string `db:"name" json:"name,omitempty"`
	Code        string `db:"code" json:"code,omitempty"`
	IsActive    int    `db:"is_active" json:"is_active,omitempty"`
	CreatedTime int32  `db:"created_time" json:"created_time,omitempty"` // seconds UTC
	UpdatedTime int32  `db:"updated_time" json:"updated_time,omitempty"` // seconds UTC
}
