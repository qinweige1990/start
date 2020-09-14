package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateCarBoss
// @description   create a CarBoss
// @param     carBoss               model.CarBoss
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateCarBoss(carBoss model.CarBoss) (err error) {
	err = global.GVA_DB.Create(&carBoss).Error
	return err
}

// @title    DeleteCarBoss
// @description   delete a CarBoss
// @auth                     （2020/04/05  20:22）
// @param     carBoss               model.CarBoss
// @return                    error

func DeleteCarBoss(carBoss model.CarBoss) (err error) {
	err = global.GVA_DB.Delete(carBoss).Error
	return err
}

// @title    DeleteCarBossByIds
// @description   delete CarBosss
// @auth                     （2020/04/05  20:22）
// @param     carBoss               model.CarBoss
// @return                    error

func DeleteCarBossByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.CarBoss{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateCarBoss
// @description   update a CarBoss
// @param     carBoss          *model.CarBoss
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateCarBoss(carBoss *model.CarBoss) (err error) {
	err = global.GVA_DB.Save(carBoss).Error
	return err
}

// @title    GetCarBoss
// @description   get the info of a CarBoss
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    CarBoss        CarBoss

func GetCarBoss(id uint) (err error, carBoss model.CarBoss) {
	err = global.GVA_DB.Where("id = ?", id).First(&carBoss).Error
	return
}

// @title    GetCarBossInfoList
// @description   get CarBoss list by pagination, 分页获取CarBoss
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetCarBossInfoList(info request.CarBossSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.CarBoss{})
    var carBosss []model.CarBoss
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("`name` LIKE ?","%"+ info.Name+"%")
    }
    if info.Status != 0 {
        db = db.Where("`status` = ?",info.Status)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&carBosss).Error
	return err, carBosss, total
}