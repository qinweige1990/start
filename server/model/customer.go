// 自动生成模板Customer
package model

import (
	"github.com/jinzhu/gorm"
)

// 如果含有time.Time 请自行import time包
type Customer struct {
      gorm.Model
      Name  string `json:"name" form:"name" gorm:"column:name;comment:'';type:varchar(20)"`
      Type  int `json:"type" form:"type" gorm:"column:type;comment:'';type:smallint"`
      Company  string `json:"company" form:"company" gorm:"column:company;comment:'';type:text"`
      ContactMan  string `json:"contactMan" form:"contactMan" gorm:"column:contact_man;comment:'';type:varchar(20)"`
      Phone  string `json:"phone" form:"phone" gorm:"column:phone;comment:'';type:varchar(20)"`
      Address  string `json:"address" form:"address" gorm:"column:address;comment:'';type:text"`
      Extra  string `json:"extra" form:"extra" gorm:"column:extra;comment:'';type:text"` 
}


func (Customer) TableName() string {
  return "customer"
}
