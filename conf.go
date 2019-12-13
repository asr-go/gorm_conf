package conf

import (
	"time"

	"github.com/jinzhu/gorm"

	"github.com/asr-go/gorm_conf/logging"

	//
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// NewDatabase 新建数据库
func NewDatabase(dialect string, debug bool) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", dialect)
	if err != nil {
		return nil, err
	}

	db.LogMode(debug)
	db.SingularTable(true)

	//空闲
	db.DB().SetMaxIdleConns(50)

	//打开
	db.DB().SetMaxOpenConns(100)

	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	db.SetLogger(logging.New())

	return db, nil
}
