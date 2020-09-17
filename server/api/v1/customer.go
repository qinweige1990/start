package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)

// @Tags Customer
// @Summary 创建Customer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Customer true "创建Customer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/createCustomer [post]
func CreateCustomer(c *gin.Context) {
	var customer model.Customer
	_ = c.ShouldBindJSON(&customer)
	err := service.CreateCustomer(customer)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Customer
// @Summary 删除Customer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Customer true "删除Customer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /customer/deleteCustomer [delete]
func DeleteCustomer(c *gin.Context) {
	var customer model.Customer
	_ = c.ShouldBindJSON(&customer)
	err := service.DeleteCustomer(customer)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Customer
// @Summary 批量删除Customer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Customer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /customer/deleteCustomerByIds [delete]
func DeleteCustomerByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteCustomerByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Customer
// @Summary 更新Customer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Customer true "更新Customer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /customer/updateCustomer [put]
func UpdateCustomer(c *gin.Context) {
	var customer model.Customer
	_ = c.ShouldBindJSON(&customer)
	err := service.UpdateCustomer(&customer)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Customer
// @Summary 用id查询Customer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Customer true "用id查询Customer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /customer/findCustomer [get]
func FindCustomer(c *gin.Context) {
	var customer model.Customer
	_ = c.ShouldBindQuery(&customer)
	err, recustomer := service.GetCustomer(customer.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"recustomer": recustomer}, c)
	}
}

// @Tags Customer
// @Summary 分页获取Customer列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CustomerSearch true "分页获取Customer列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/getCustomerList [get]
func GetCustomerList(c *gin.Context) {
	var pageInfo request.CustomerSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetCustomerInfoList(pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}
