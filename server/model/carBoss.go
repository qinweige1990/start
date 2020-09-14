// 自动生成模板CarBoss
package model

import (
	"github.com/jinzhu/gorm"
)

// 如果含有time.Time 请自行import time包
type CarBoss struct {
      gorm.Model
      Name  string `json:"name" form:"name" gorm:"column:name;comment:'';type:varchar(20)"`
      Phone  string `json:"phone" form:"phone" gorm:"column:phone;comment:'';type:varchar(20)"`
      Identity  string `json:"identity" form:"identity" gorm:"column:identity;comment:'';type:varchar(20)"`
      Bank  string `json:"bank" form:"bank" gorm:"column:bank;comment:'';type:varchar(40)"`
      BankNumber  string `json:"bankNumber" form:"bankNumber" gorm:"column:bank_number;comment:'';type:varchar(40)"`
      Status  int `json:"status" form:"status" gorm:"column:status;comment:'';type:smallint"`
      Img  string `json:"img" form:"img" gorm:"column:img;comment:'';type:text"` 
}


func (CarBoss) TableName() string {
  return "car_boss"
}
