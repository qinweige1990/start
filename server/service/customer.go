package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateCustomer
// @description   create a Customer
// @param     customer               model.Customer
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateCustomer(customer model.Customer) (err error) {
	err = global.GVA_DB.Create(&customer).Error
	return err
}

// @title    DeleteCustomer
// @description   delete a Customer
// @auth                     （2020/04/05  20:22）
// @param     customer               model.Customer
// @return                    error

func DeleteCustomer(customer model.Customer) (err error) {
	err = global.GVA_DB.Delete(customer).Error
	return err
}

// @title    DeleteCustomerByIds
// @description   delete Customers
// @auth                     （2020/04/05  20:22）
// @param     customer               model.Customer
// @return                    error

func DeleteCustomerByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Customer{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateCustomer
// @description   update a Customer
// @param     customer          *model.Customer
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateCustomer(customer *model.Customer) (err error) {
	err = global.GVA_DB.Save(customer).Error
	return err
}

// @title    GetCustomer
// @description   get the info of a Customer
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Customer        Customer

func GetCustomer(id uint) (err error, customer model.Customer) {
	err = global.GVA_DB.Where("id = ?", id).First(&customer).Error
	return
}

// @title    GetCustomerInfoList
// @description   get Customer list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetCustomerInfoList(info request.CustomerSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Customer{})
    var customers []model.Customer
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.Type != 0 {
        db = db.Where("type = ?",info.Type)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&customers).Error
	return err, customers, total
}