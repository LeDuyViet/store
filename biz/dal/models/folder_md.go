package models

type FolderMD struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	ParentID    string `json:"parent_id" db:"parent_id"`
	CreatedTime int32  `json:"created_time" db:"created_time"`
	UpdatedTime int32  `json:"updated_time" db:"updated_time"`
}
