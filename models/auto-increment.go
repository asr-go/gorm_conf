package models

// AutoIncrementModel model
type AutoIncrementModel struct {
	Model
	ID int `gorm:"primary_key,not null;AUTO_INCREMENT"` // ID
}
