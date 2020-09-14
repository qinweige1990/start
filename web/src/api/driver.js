import service from '@/utils/request'

// @Tags Driver
// @Summary 创建Driver
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Driver true "创建Driver"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /driver/createDriver [post]
export const createDriver = (data) => {
     return service({
         url: "/driver/createDriver",
         method: 'post',
         data
     })
 }


// @Tags Driver
// @Summary 删除Driver
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Driver true "删除Driver"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /driver/deleteDriver [delete]
 export const deleteDriver = (data) => {
     return service({
         url: "/driver/deleteDriver",
         method: 'delete',
         data
     })
 }

// @Tags Driver
// @Summary 删除Driver
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Driver"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /driver/deleteDriver [delete]
 export const deleteDriverByIds = (data) => {
     return service({
         url: "/driver/deleteDriverByIds",
         method: 'delete',
         data
     })
 }

// @Tags Driver
// @Summary 更新Driver
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Driver true "更新Driver"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /driver/updateDriver [put]
 export const updateDriver = (data) => {
     return service({
         url: "/driver/updateDriver",
         method: 'put',
         data
     })
 }


// @Tags Driver
// @Summary 用id查询Driver
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Driver true "用id查询Driver"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /driver/findDriver [get]
 export const findDriver = (params) => {
     return service({
         url: "/driver/findDriver",
         method: 'get',
         params
     })
 }


// @Tags Driver
// @Summary 分页获取Driver列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Driver列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /driver/getDriverList [get]
 export const getDriverList = (params) => {
     return service({
         url: "/driver/getDriverList",
         method: 'get',
         params
     })
 }