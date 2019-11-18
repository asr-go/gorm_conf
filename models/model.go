package models

import "database/sql"

// Model model
type Model struct {
	CreatedTime int64         `gorm:"index:idx_created_time;not null;type:bigint(13)"` // 创建时间
	UpdatedTime sql.NullInt64 `gorm:"index:idx_updated_time;type:bigint(13)"`          // 更新时间
	DeletedTime sql.NullInt64 `gorm:"index:idx_deleted_time;type:bigint(13)"`          // 删除时间
	Description string        `gorm:"not null;size:1000;default:''"`                   // 资源类型描述
	Enabled     bool          `gorm:"index:idx_enabled;not null;default:1"`            // 启用
}
