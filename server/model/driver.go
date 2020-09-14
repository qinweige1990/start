// 自动生成模板Driver
package model

import (
	"github.com/jinzhu/gorm"
)

// 如果含有time.Time 请自行import time包
type Driver struct {
      gorm.Model
      License  string `json:"license" form:"license" gorm:"column:license;comment:'';type:varchar(20)"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:'';type:varchar(20)"`
      Phone  string `json:"phone" form:"phone" gorm:"column:phone;comment:'';type:varchar(20)"`
      TransportNumber  string `json:"transportNumber" form:"transportNumber" gorm:"column:transport_number;comment:'';type:varchar(20)"`
      Loading  string `json:"loading" form:"loading" gorm:"column:loading;comment:'';type:varchar(20)"`
      Brand  string `json:"brand" form:"brand" gorm:"column:brand;comment:'';type:text"`
      Date  time.Time `json:"date" form:"date" gorm:"column:date;comment:'';type:data"`
      Identify  int `json:"identify" form:"identify" gorm:"column:identify;comment:'';type:smallint"`
      DriverReview  string `json:"driverReview" form:"driverReview" gorm:"column:driver_review;comment:'';type:varchar(20)"`
      CarReview  string `json:"carReview" form:"carReview" gorm:"column:car_review;comment:'';type:varchar(20)"`
      DriverLicenseImg  string `json:"driverLicenseImg" form:"driverLicenseImg" gorm:"column:driver_license_img;comment:'';type:text"`
      CarLicenseImg  string `json:"carLicenseImg" form:"carLicenseImg" gorm:"column:car_license_img;comment:'';type:text"` 
}


func (Driver) TableName() string {
  return "driver"
}
