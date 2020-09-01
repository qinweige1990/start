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

// @Tags CargoOrder
// @Summary 创建CargoOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CargoOrder true "创建CargoOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /order/createCargoOrder [post]
func CreateCargoOrder(c *gin.Context) {
	var order model.CargoOrder
	_ = c.ShouldBindJSON(&order)
	err := service.CreateCargoOrder(order)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags CargoOrder
// @Summary 删除CargoOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CargoOrder true "删除CargoOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /order/deleteCargoOrder [delete]
func DeleteCargoOrder(c *gin.Context) {
	var order model.CargoOrder
	_ = c.ShouldBindJSON(&order)
	err := service.DeleteCargoOrder(order)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags CargoOrder
// @Summary 批量删除CargoOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CargoOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /order/deleteCargoOrderByIds [delete]
func DeleteCargoOrderByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteCargoOrderByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags CargoOrder
// @Summary 更新CargoOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CargoOrder true "更新CargoOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /order/updateCargoOrder [put]
func UpdateCargoOrder(c *gin.Context) {
	var order model.CargoOrder
	_ = c.ShouldBindJSON(&order)
	err := service.UpdateCargoOrder(&order)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags CargoOrder
// @Summary 用id查询CargoOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CargoOrder true "用id查询CargoOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /order/findCargoOrder [get]
func FindCargoOrder(c *gin.Context) {
	var order model.CargoOrder
	_ = c.ShouldBindQuery(&order)
	err, reorder := service.GetCargoOrder(order.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"reorder": reorder}, c)
	}
}

// @Tags CargoOrder
// @Summary 分页获取CargoOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CargoOrderSearch true "分页获取CargoOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /order/getCargoOrderList [get]
func GetCargoOrderList(c *gin.Context) {
	var pageInfo request.CargoOrderSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetCargoOrderInfoList(pageInfo)
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
