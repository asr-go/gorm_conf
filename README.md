# GORM Conf

提供GORM常用配置

## 快速上手

``` go
package main

import(
  "github.com/jinzhu/gorm"
  gorm_conf "github.com/asr-go/gorm_conf"

  //
  _ "github.com/jinzhu/gorm/dialects/mysql"

)

func main() {
  db, err := gorm.Open("mysql", connString)
  // Error
  if err != nil {
    panic("连接数据库不成功", err)
  }
  gorm_conf.Conf(db)
}
```
