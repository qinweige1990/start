// 自动生成模板CargoOrder
package model

import (
	"github.com/jinzhu/gorm"
)

// 如果含有time.Time 请自行import time包
type CargoOrder struct {
      gorm.Model
      Start  string `json:"start" form:"start" gorm:"column:start;comment:'';type:text(100)"`
      End  string `json:"end" form:"end" gorm:"column:end;comment:'';type:text(100)"` 
}


func (CargoOrder) TableName() string {
  return "cargo_order"
}
