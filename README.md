# GORM Conf

提供GORM常用配置

## 快速上手

``` go
package main

import(
  "github.com/jinzhu/gorm"
  "github.com/sirupsen/logrus"
  conf "github.com/surh-go/gorm_conf"

  //
  _ "github.com/jinzhu/gorm/dialects/mysql"

  //
  _ "github.com/surh97/logrus_config"
)

func main() {
  db, err := gorm.Open("mysql", connString)
  // Error
  if err != nil {
    logrus.Panic("连接数据库不成功", err)
  }
  conf.Conf(db)
}
```
