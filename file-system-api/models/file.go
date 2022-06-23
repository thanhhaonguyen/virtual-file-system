package models

type File struct {
	Id        int64  `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Name      string `gorm:"size:255" json:"name"`
	Data      string `json:"data"`
	FolderId  int64  `json:"folder_id"`
	CreatedAt int64  `gorm:"autoCreateTime;not null" json:"created_at"`
}
