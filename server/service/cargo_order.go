package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateCargoOrder
// @description   create a CargoOrder
// @param     order               model.CargoOrder
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateCargoOrder(order model.CargoOrder) (err error) {
	err = global.GVA_DB.Create(&order).Error
	return err
}

// @title    DeleteCargoOrder
// @description   delete a CargoOrder
// @auth                     （2020/04/05  20:22）
// @param     order               model.CargoOrder
// @return                    error

func DeleteCargoOrder(order model.CargoOrder) (err error) {
	err = global.GVA_DB.Delete(order).Error
	return err
}

// @title    DeleteCargoOrderByIds
// @description   delete CargoOrders
// @auth                     （2020/04/05  20:22）
// @param     order               model.CargoOrder
// @return                    error

func DeleteCargoOrderByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.CargoOrder{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateCargoOrder
// @description   update a CargoOrder
// @param     order          *model.CargoOrder
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateCargoOrder(order *model.CargoOrder) (err error) {
	err = global.GVA_DB.Save(order).Error
	return err
}

// @title    GetCargoOrder
// @description   get the info of a CargoOrder
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    CargoOrder        CargoOrder

func GetCargoOrder(id uint) (err error, order model.CargoOrder) {
	err = global.GVA_DB.Where("id = ?", id).First(&order).Error
	return
}

// @title    GetCargoOrderInfoList
// @description   get CargoOrder list by pagination, 分页获取CargoOrder
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetCargoOrderInfoList(info request.CargoOrderSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.CargoOrder{})
    var orders []model.CargoOrder
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Start != "" {
        db = db.Where("`start` LIKE ?","%"+ info.Start+"%")
    }
    if info.End != "" {
        db = db.Where("`end` LIKE ?","%"+ info.End+"%")
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&orders).Error
	return err, orders, total
}