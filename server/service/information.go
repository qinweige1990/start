package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateInformation
// @description   create a Information
// @param     information               model.Information
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateInformation(information model.Information) (err error) {
	err = global.GVA_DB.Create(&information).Error
	return err
}

// @title    DeleteInformation
// @description   delete a Information
// @auth                     （2020/04/05  20:22）
// @param     information               model.Information
// @return                    error

func DeleteInformation(information model.Information) (err error) {
	err = global.GVA_DB.Delete(information).Error
	return err
}

// @title    DeleteInformationByIds
// @description   delete Informations
// @auth                     （2020/04/05  20:22）
// @param     information               model.Information
// @return                    error

func DeleteInformationByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Information{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateInformation
// @description   update a Information
// @param     information          *model.Information
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateInformation(information *model.Information) (err error) {
	err = global.GVA_DB.Save(information).Error
	return err
}

// @title    GetInformation
// @description   get the info of a Information
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Information        Information

func GetInformation(id uint) (err error, information model.Information) {
	err = global.GVA_DB.Where("id = ?", id).First(&information).Error
	return
}

// @title    GetInformationInfoList
// @description   get Information list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetInformationInfoList(info request.InformationSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Information{})
    var informations []model.Information
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&informations).Error
	return err, informations, total
}