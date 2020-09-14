import service from '@/utils/request'

// @Tags CarBoss
// @Summary 创建CarBoss
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CarBoss true "创建CarBoss"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /carBoss/createCarBoss [post]
export const createCarBoss = (data) => {
     return service({
         url: "/carBoss/createCarBoss",
         method: 'post',
         data
     })
 }


// @Tags CarBoss
// @Summary 删除CarBoss
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CarBoss true "删除CarBoss"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /carBoss/deleteCarBoss [delete]
 export const deleteCarBoss = (data) => {
     return service({
         url: "/carBoss/deleteCarBoss",
         method: 'delete',
         data
     })
 }

// @Tags CarBoss
// @Summary 删除CarBoss
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CarBoss"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /carBoss/deleteCarBoss [delete]
 export const deleteCarBossByIds = (data) => {
     return service({
         url: "/carBoss/deleteCarBossByIds",
         method: 'delete',
         data
     })
 }

// @Tags CarBoss
// @Summary 更新CarBoss
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CarBoss true "更新CarBoss"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /carBoss/updateCarBoss [put]
 export const updateCarBoss = (data) => {
     return service({
         url: "/carBoss/updateCarBoss",
         method: 'put',
         data
     })
 }


// @Tags CarBoss
// @Summary 用id查询CarBoss
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CarBoss true "用id查询CarBoss"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /carBoss/findCarBoss [get]
 export const findCarBoss = (params) => {
     return service({
         url: "/carBoss/findCarBoss",
         method: 'get',
         params
     })
 }


// @Tags CarBoss
// @Summary 分页获取CarBoss列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取CarBoss列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /carBoss/getCarBossList [get]
 export const getCarBossList = (params) => {
     return service({
         url: "/carBoss/getCarBossList",
         method: 'get',
         params
     })
 }