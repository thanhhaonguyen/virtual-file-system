package models

type Folder struct {
	Id        int64  `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Name      string `gorm:"size:255" json:"name"`
	ParentId  int64  `json:"parent_id"`
	CreatedAt int64  `gorm:"autoCreateTime;not null" json:"created_at"`
}
