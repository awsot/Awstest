package models

import (
  "github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"
)

//var c *gorm.DB
type User struct {
    UserID       int     `gorm:"primary_key,unique_index,AUTO_INCREMENT"`
    UserName     string  `gorm:"type:varchar(100);unique_index`
    Password     string  `gorm:"size:255"`
}

func ConnectDB()  *gorm.DB {
  db, err := gorm.Open("mysql", "test:pass@/test?charset=utf8mb4&parseTime=True&loc=Local")
  if err != nil {
    panic("连接数据库失败")
  }
 db.DB().Ping()
  return db
}

func init()  {
  ConnectDB()
}
