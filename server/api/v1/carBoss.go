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

// @Tags CarBoss
// @Summary 创建CarBoss
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CarBoss true "创建CarBoss"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /carBoss/createCarBoss [post]
func CreateCarBoss(c *gin.Context) {
	var carBoss model.CarBoss
	_ = c.ShouldBindJSON(&carBoss)
	err := service.CreateCarBoss(carBoss)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags CarBoss
// @Summary 删除CarBoss
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CarBoss true "删除CarBoss"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /carBoss/deleteCarBoss [delete]
func DeleteCarBoss(c *gin.Context) {
	var carBoss model.CarBoss
	_ = c.ShouldBindJSON(&carBoss)
	err := service.DeleteCarBoss(carBoss)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags CarBoss
// @Summary 批量删除CarBoss
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CarBoss"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /carBoss/deleteCarBossByIds [delete]
func DeleteCarBossByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteCarBossByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags CarBoss
// @Summary 更新CarBoss
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CarBoss true "更新CarBoss"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /carBoss/updateCarBoss [put]
func UpdateCarBoss(c *gin.Context) {
	var carBoss model.CarBoss
	_ = c.ShouldBindJSON(&carBoss)
	err := service.UpdateCarBoss(&carBoss)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags CarBoss
// @Summary 用id查询CarBoss
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CarBoss true "用id查询CarBoss"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /carBoss/findCarBoss [get]
func FindCarBoss(c *gin.Context) {
	var carBoss model.CarBoss
	_ = c.ShouldBindQuery(&carBoss)
	err, recarBoss := service.GetCarBoss(carBoss.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"recarBoss": recarBoss}, c)
	}
}

// @Tags CarBoss
// @Summary 分页获取CarBoss列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CarBossSearch true "分页获取CarBoss列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /carBoss/getCarBossList [get]
func GetCarBossList(c *gin.Context) {
	var pageInfo request.CarBossSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetCarBossInfoList(pageInfo)
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
