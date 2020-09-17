// 自动生成模板Information
package model

import (
	"github.com/jinzhu/gorm"
)

// 如果含有time.Time 请自行import time包
type Information struct {
      gorm.Model
      Name  string `json:"name" form:"name" gorm:"column:name;comment:'';type:varchar(20)"`
      Status  int `json:"status" form:"status" gorm:"column:status;comment:'';type:smallint"`
      Extra  string `json:"extra" form:"extra" gorm:"column:extra;comment:'';type:text"`
      Address  string `json:"address" form:"address" gorm:"column:address;comment:'';type:text"`
      License  string `json:"license" form:"license" gorm:"column:license;comment:'';type:varchar(20)"`
      InCharge  string `json:"inCharge" form:"inCharge" gorm:"column:in_charge;comment:'';type:varchar(20)"`
      IdentityNumber  string `json:"identityNumber" form:"identityNumber" gorm:"column:identity_number;comment:'';type:varchar(30)"`
      LicenseImg  string `json:"licenseImg" form:"licenseImg" gorm:"column:license_img;comment:'';type:text"`
      IdentityImg  string `json:"identityImg" form:"identityImg" gorm:"column:identity_img;comment:'';type:text"`
      ExtraImg  string `json:"extraImg" form:"extraImg" gorm:"column:extraImg;comment:'';type:text"` 
}


func (Information) TableName() string {
  return "information"
}
