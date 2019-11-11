package models

// Model model
type Model struct {
	CreatedTime int64  `gorm:"index:idx_created_time;not null;type:bigint(13)"` // 创建时间
	UpdatedTime *int64 `gorm:"index:idx_updated_time;type:bigint(13)"`          // 更新时间
	DeletedTime *int64 `gorm:"index:idx_deleted_time;type:bigint(13)"`          // 删除时间
	Description string `gorm:"null;size:1000"`                                  // 资源类型描述
	Enabled     bool   `gorm:"index:idx_enabled;not null;default:1"`            // 启用
}
