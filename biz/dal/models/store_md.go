package models

import (
	"time"

	"github.com/google/uuid"
)

type StoreMD struct {
	ID          string `db:"id" json:"id,omitempty"`
	Name        string `db:"name" json:"name,omitempty"`
	Address     string `db:"address" json:"address,omitempty"`
	IsActive    bool   `db:"is_active" json:"is_active"`
	AccessKey   string `db:"access_key" json:"access_key,omitempty"`
	CreatedTime int32  `db:"created_time" json:"created_time,omitempty"` // seconds UTC
	UpdatedTime int32  `db:"updated_time" json:"updated_time,omitempty"` // seconds UTC

	// AppIdsRes []*AppIdsRes `db:"-" json:"app_ids,omitempty"`
}

type AppIdsRes struct {
	AppId   interface{} `json:"app_id,omitempty"`
	AppName string      `json:"app_name,omitempty"`
}

// fill id and created_time
func (store *StoreMD) FillIdAndCreatedTime() {
	store.ID = uuid.New().String()
	store.CreatedTime = int32(time.Now().Unix())
	store.UpdatedTime = store.CreatedTime
}
