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

// @Tags Driver
// @Summary 创建Driver
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Driver true "创建Driver"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /driver/createDriver [post]
func CreateDriver(c *gin.Context) {
	var driver model.Driver
	_ = c.ShouldBindJSON(&driver)
	err := service.CreateDriver(driver)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Driver
// @Summary 删除Driver
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Driver true "删除Driver"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /driver/deleteDriver [delete]
func DeleteDriver(c *gin.Context) {
	var driver model.Driver
	_ = c.ShouldBindJSON(&driver)
	err := service.DeleteDriver(driver)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Driver
// @Summary 批量删除Driver
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Driver"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /driver/deleteDriverByIds [delete]
func DeleteDriverByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteDriverByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Driver
// @Summary 更新Driver
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Driver true "更新Driver"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /driver/updateDriver [put]
func UpdateDriver(c *gin.Context) {
	var driver model.Driver
	_ = c.ShouldBindJSON(&driver)
	err := service.UpdateDriver(&driver)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Driver
// @Summary 用id查询Driver
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Driver true "用id查询Driver"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /driver/findDriver [get]
func FindDriver(c *gin.Context) {
	var driver model.Driver
	_ = c.ShouldBindQuery(&driver)
	err, redriver := service.GetDriver(driver.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"redriver": redriver}, c)
	}
}

// @Tags Driver
// @Summary 分页获取Driver列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DriverSearch true "分页获取Driver列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /driver/getDriverList [get]
func GetDriverList(c *gin.Context) {
	var pageInfo request.DriverSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetDriverInfoList(pageInfo)
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
