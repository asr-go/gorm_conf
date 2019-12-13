package model

import (
	"time"

	"github.com/asr-go/snowflake"
)

// BaseModel 基础模型
type BaseModel struct {
	ID          snowflake.ID `gorm:"size:20;AUTO_INCREMENT:false"`             // ID
	CreatedAt   time.Time    `gorm:"index;not null;DEFAULT:CURRENT_TIMESTAMP"` // 创建时间
	UpdatedAt   time.Time    `gorm:"index;not null;DEFAULT:CURRENT_TIMESTAMP"` // 更新时间
	DeletedAt   *time.Time   `gorm:"index;"`                                   // 删除时间
	Description string       `gorm:"not null;size:1000;DEFAULT:''"`            // 资源类型描述
	Enabled     bool         `gorm:"index;not null;DEFAULT:1"`                 // 启用
}
