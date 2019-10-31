package models

// Model model
type Model struct {
	CreatedTime int64  `gorm:"not null;type:bigint(13)"`
	UpdatedTime *int64 `gorm:"type:bigint(13)"`
	DeletedTime *int64 `gorm:"index;type:bigint(13)"`
}
