import service from '@/utils/request'

// @Tags Information
// @Summary 创建Information
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Information true "创建Information"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /information/createInformation [post]
export const createInformation = (data) => {
     return service({
         url: "/information/createInformation",
         method: 'post',
         data
     })
 }


// @Tags Information
// @Summary 删除Information
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Information true "删除Information"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /information/deleteInformation [delete]
 export const deleteInformation = (data) => {
     return service({
         url: "/information/deleteInformation",
         method: 'delete',
         data
     })
 }

// @Tags Information
// @Summary 删除Information
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Information"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /information/deleteInformation [delete]
 export const deleteInformationByIds = (data) => {
     return service({
         url: "/information/deleteInformationByIds",
         method: 'delete',
         data
     })
 }

// @Tags Information
// @Summary 更新Information
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Information true "更新Information"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /information/updateInformation [put]
 export const updateInformation = (data) => {
     return service({
         url: "/information/updateInformation",
         method: 'put',
         data
     })
 }


// @Tags Information
// @Summary 用id查询Information
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Information true "用id查询Information"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /information/findInformation [get]
 export const findInformation = (params) => {
     return service({
         url: "/information/findInformation",
         method: 'get',
         params
     })
 }


// @Tags Information
// @Summary 分页获取Information列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Information列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /information/getInformationList [get]
 export const getInformationList = (params) => {
     return service({
         url: "/information/getInformationList",
         method: 'get',
         params
     })
 }