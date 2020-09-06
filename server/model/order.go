// 自动生成模板Order
package model

import (
	"github.com/jinzhu/gorm"
)

// 如果含有time.Time 请自行import time包
type Order struct {
      gorm.Model
      TransportId  string `json:"transportId" form:"transportId" gorm:"column:transport_id;comment:''"`
      DepartureTime  time.Time `json:"departureTime" form:"departureTime" gorm:"column:departure_time;comment:'';type:datetime"`
      Start  string `json:"start" form:"start" gorm:"column:start;comment:''"`
      End  string `json:"end" form:"end" gorm:"column:end;comment:'';type:text"`
      Good  string `json:"good" form:"good" gorm:"column:good;comment:''"`
      Driver  string `json:"driver" form:"driver" gorm:"column:driver;comment:'';type:varchar(20)"`
      DriverPhone  string `json:"driverPhone" form:"driverPhone" gorm:"column:driver_phone;comment:'';type:varchar(20)"`
      CarNumber  string `json:"carNumber" form:"carNumber" gorm:"column:car_number;comment:'';type:varchar(20)"`
      StartWeight  float64 `json:"startWeight" form:"startWeight" gorm:"column:start_weight;comment:'';type:float"`
      EndWeight  float64 `json:"endWeight" form:"endWeight" gorm:"column:end_weight;comment:'';type:float"`
      TotalFare  float64 `json:"totalFare" form:"totalFare" gorm:"column:total_fare;comment:'';type:float"`
      PreFare  float64 `json:"preFare" form:"preFare" gorm:"column:pre_fare;comment:'';type:float"`
      PaidFare  float64 `json:"paidFare" form:"paidFare" gorm:"column:paid_fare;comment:'';type:float"`
      DueFare  float64 `json:"dueFare" form:"dueFare" gorm:"column:due_fare;comment:'';type:float"`
      StartCustomer  string `json:"startCustomer" form:"startCustomer" gorm:"column:start_customer;comment:'';type:varchar(20)"`
      StartPhone  string `json:"startPhone" form:"startPhone" gorm:"column:start_phone;comment:'';type:varchar(20)"`
      EndCustomer  string `json:"endCustomer" form:"endCustomer" gorm:"column:end_customer;comment:'';type:varchar(20)"`
      EndPhone  string `json:"endPhone" form:"endPhone" gorm:"column:end_phone;comment:'';type:varchar(20)"`
      Station  string `json:"station" form:"station" gorm:"column:station;comment:'';type:text"`
      StationPhone  string `json:"stationPhone" form:"stationPhone" gorm:"column:station_phone;comment:'';type:varchar(20)"`
      PlanId  string `json:"planId" form:"planId" gorm:"column:plan_id`;comment:'';type:varchar(50)"`
      OrderId  string `json:"orderId" form:"orderId" gorm:"column:order_id;comment:'';type:varchar(50)"`
      Agent  string `json:"agent" form:"agent" gorm:"column:agent;comment:'';type:text"`
      TransportStatus  int `json:"transportStatus" form:"transportStatus" gorm:"column:transport_status;comment:'';type:smallint"`
      FareAuditStatus  int `json:"fareAuditStatus" form:"fareAuditStatus" gorm:"column:fare_audit_status;comment:'';type:smallint"`
      OrderAuditStatus  int `json:"orderAuditStatus" form:"orderAuditStatus" gorm:"column:order_audit_status;comment:'';type:smallint"`
      FarePayStatus  int `json:"farePayStatus" form:"farePayStatus" gorm:"column:fare_pay_status;comment:'';type:smallint"`
      DueFareStatus  int `json:"dueFareStatus" form:"dueFareStatus" gorm:"column:due_fare_status;comment:'';type:smallint"`
      InvoiceStatus  int `json:"invoiceStatus" form:"invoiceStatus" gorm:"column:invoice_status;comment:'';type:smallint"`
      CreatedBy  string `json:"createBy" form:"createBy" gorm:"column:create_by;comment:'';type:varchar(20)"` 
}


func (Order) TableName() string {
  return "order"
}
