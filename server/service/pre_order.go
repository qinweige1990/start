package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateOrder
// @description   create a Order
// @param     order               model.Order
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateOrder(order model.Order) (err error) {
	err = global.GVA_DB.Create(&order).Error
	return err
}

// @title    DeleteOrder
// @description   delete a Order
// @auth                     （2020/04/05  20:22）
// @param     order               model.Order
// @return                    error

func DeleteOrder(order model.Order) (err error) {
	err = global.GVA_DB.Delete(order).Error
	return err
}

// @title    DeleteOrderByIds
// @description   delete Orders
// @auth                     （2020/04/05  20:22）
// @param     order               model.Order
// @return                    error

func DeleteOrderByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Order{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateOrder
// @description   update a Order
// @param     order          *model.Order
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateOrder(order *model.Order) (err error) {
	err = global.GVA_DB.Save(order).Error
	return err
}

// @title    GetOrder
// @description   get the info of a Order
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Order        Order

func GetOrder(id uint) (err error, order model.Order) {
	err = global.GVA_DB.Where("id = ?", id).First(&order).Error
	return
}

// @title    GetOrderInfoList
// @description   get Order list by pagination, 分页获取Order
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetOrderInfoList(info request.OrderSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Order{})
    var orders []model.Order
    // 如果有条件搜索 下方会自动创建搜索语句
    if !info.DepartureTime.IsZero() {
         db = db.Where("`departure_time` <> ?",info.DepartureTime)
    }
    if info.End != "" {
        db = db.Where("`end` LIKE ?","%"+ info.End+"%")
    }
    if info.Driver != "" {
        db = db.Where("`driver` LIKE ?","%"+ info.Driver+"%")
    }
    if info.DriverPhone != "" {
        db = db.Where("`driver_phone` LIKE ?","%"+ info.DriverPhone+"%")
    }
    if info.CarNumber != "" {
        db = db.Where("`car_number` LIKE ?","%"+ info.CarNumber+"%")
    }
    if info.StartCustomer != "" {
        db = db.Where("`start_customer` LIKE ?","%"+ info.StartCustomer+"%")
    }
    if info.StartPhone != "" {
        db = db.Where("`start_phone` LIKE ?","%"+ info.StartPhone+"%")
    }
    if info.EndCustomer != "" {
        db = db.Where("`end_customer` LIKE ?","%"+ info.EndCustomer+"%")
    }
    if info.EndPhone != "" {
        db = db.Where("`end_phone` LIKE ?","%"+ info.EndPhone+"%")
    }
    if info.TransportStatus != 0 {
        db = db.Where("`transport_status` = ?",info.TransportStatus)
    }
    if info.FareAuditStatus != 0 {
        db = db.Where("`fare_audit_status` = ?",info.FareAuditStatus)
    }
    if info.OrderAuditStatus != 0 {
        db = db.Where("`order_audit_status` = ?",info.OrderAuditStatus)
    }
    if info.FarePayStatus != 0 {
        db = db.Where("`fare_pay_status` = ?",info.FarePayStatus)
    }
    if info.DueFareStatus != 0 {
        db = db.Where("`due_fare_status` = ?",info.DueFareStatus)
    }
    if info.InvoiceStatus != 0 {
        db = db.Where("`invoice_status` = ?",info.InvoiceStatus)
    }
	err = db.Count(&total).Error
	err = db.Order("created_at DESC").Limit(limit).Offset(offset).Find(&orders).Error
	return err, orders, total
}
