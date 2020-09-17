// 自动生成模板Order
package model

import (
	"github.com/jinzhu/gorm"
      "time"
)

// 如果含有time.Time 请自行import time包
type PreOrder struct {
      gorm.Model
      TransportId  string `json:"transportId" form:"transportId" gorm:"column:transport_id;comment:''"`
      DepartureTime  time.Time `json:"departureTime" form:"departureTime" gorm:"column:departure_time;comment:'';type:datetime"`
      Start  string `json:"start" form:"start" gorm:"column:start;comment:''"`
      End  string `json:"end" form:"end" gorm:"column:end;comment:'';type:text"`
      Good  string `json:"good" form:"good" gorm:"column:good;comment:''"`
      StartWeight  float64 `json:"startWeight" form:"startWeight" gorm:"column:start_weight;comment:'';type:float"`
      TotalFare  float64 `json:"totalFare" form:"totalFare" gorm:"column:total_fare;comment:'';type:float"`
      StartCustomer  string `json:"startCustomer" form:"startCustomer" gorm:"column:start_customer;comment:'';type:varchar(20)"`
      StartPhone  string `json:"startPhone" form:"startPhone" gorm:"column:start_phone;comment:'';type:varchar(20)"`
      EndCustomer  string `json:"endCustomer" form:"endCustomer" gorm:"column:end_customer;comment:'';type:varchar(20)"`
      EndPhone  string `json:"endPhone" form:"endPhone" gorm:"column:end_phone;comment:'';type:varchar(20)"`
      Agent  string `json:"agent" form:"agent" gorm:"column:agent;comment:'';type:text"`
      CreatedBy  string `json:"createBy" form:"createBy" gorm:"column:create_by;comment:'';type:varchar(20)"`
}


func (PreOrder) TableName() string {
  return "pre_order"
}
