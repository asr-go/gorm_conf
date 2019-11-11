package models

// SnowflakeModel model
type SnowflakeModel struct {
	Model
	ID int64 `gorm:"primary_key;not null;size:20;AUTO_INCREMENT:false"` // ID
}
