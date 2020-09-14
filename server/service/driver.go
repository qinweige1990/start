package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateDriver
// @description   create a Driver
// @param     driver               model.Driver
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateDriver(driver model.Driver) (err error) {
	err = global.GVA_DB.Create(&driver).Error
	return err
}

// @title    DeleteDriver
// @description   delete a Driver
// @auth                     （2020/04/05  20:22）
// @param     driver               model.Driver
// @return                    error

func DeleteDriver(driver model.Driver) (err error) {
	err = global.GVA_DB.Delete(driver).Error
	return err
}

// @title    DeleteDriverByIds
// @description   delete Drivers
// @auth                     （2020/04/05  20:22）
// @param     driver               model.Driver
// @return                    error

func DeleteDriverByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Driver{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateDriver
// @description   update a Driver
// @param     driver          *model.Driver
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateDriver(driver *model.Driver) (err error) {
	err = global.GVA_DB.Save(driver).Error
	return err
}

// @title    GetDriver
// @description   get the info of a Driver
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Driver        Driver

func GetDriver(id uint) (err error, driver model.Driver) {
	err = global.GVA_DB.Where("id = ?", id).First(&driver).Error
	return
}

// @title    GetDriverInfoList
// @description   get Driver list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetDriverInfoList(info request.DriverSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Driver{})
    var drivers []model.Driver
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.License != "" {
        db = db.Where("license LIKE ?","%"+ info.License+"%")
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.DriverReview != "" {
        db = db.Where("driver_review = ?",info.DriverReview)
    }
    if info.CarReview != "" {
        db = db.Where("car_review = ?",info.CarReview)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&drivers).Error
	return err, drivers, total
}