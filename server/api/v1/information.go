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

// @Tags Information
// @Summary 创建Information
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Information true "创建Information"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /information/createInformation [post]
func CreateInformation(c *gin.Context) {
	var information model.Information
	_ = c.ShouldBindJSON(&information)
	err := service.CreateInformation(information)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Information
// @Summary 删除Information
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Information true "删除Information"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /information/deleteInformation [delete]
func DeleteInformation(c *gin.Context) {
	var information model.Information
	_ = c.ShouldBindJSON(&information)
	err := service.DeleteInformation(information)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Information
// @Summary 批量删除Information
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Information"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /information/deleteInformationByIds [delete]
func DeleteInformationByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteInformationByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Information
// @Summary 更新Information
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Information true "更新Information"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /information/updateInformation [put]
func UpdateInformation(c *gin.Context) {
	var information model.Information
	_ = c.ShouldBindJSON(&information)
	err := service.UpdateInformation(&information)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Information
// @Summary 用id查询Information
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Information true "用id查询Information"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /information/findInformation [get]
func FindInformation(c *gin.Context) {
	var information model.Information
	_ = c.ShouldBindQuery(&information)
	err, reinformation := service.GetInformation(information.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"reinformation": reinformation}, c)
	}
}

// @Tags Information
// @Summary 分页获取Information列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.InformationSearch true "分页获取Information列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /information/getInformationList [get]
func GetInformationList(c *gin.Context) {
	var pageInfo request.InformationSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetInformationInfoList(pageInfo)
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
