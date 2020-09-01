import service from '@/utils/request'

// @Tags CargoOrder
// @Summary 创建CargoOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CargoOrder true "创建CargoOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /order/createCargoOrder [post]
export const createCargoOrder = (data) => {
     return service({
         url: "/order/createCargoOrder",
         method: 'post',
         data
     })
 }


// @Tags CargoOrder
// @Summary 删除CargoOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CargoOrder true "删除CargoOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /order/deleteCargoOrder [delete]
 export const deleteCargoOrder = (data) => {
     return service({
         url: "/order/deleteCargoOrder",
         method: 'delete',
         data
     })
 }

// @Tags CargoOrder
// @Summary 删除CargoOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CargoOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /order/deleteCargoOrder [delete]
 export const deleteCargoOrderByIds = (data) => {
     return service({
         url: "/order/deleteCargoOrderByIds",
         method: 'delete',
         data
     })
 }

// @Tags CargoOrder
// @Summary 更新CargoOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CargoOrder true "更新CargoOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /order/updateCargoOrder [put]
 export const updateCargoOrder = (data) => {
     return service({
         url: "/order/updateCargoOrder",
         method: 'put',
         data
     })
 }


// @Tags CargoOrder
// @Summary 用id查询CargoOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CargoOrder true "用id查询CargoOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /order/findCargoOrder [get]
 export const findCargoOrder = (params) => {
     return service({
         url: "/order/findCargoOrder",
         method: 'get',
         params
     })
 }


// @Tags CargoOrder
// @Summary 分页获取CargoOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取CargoOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /order/getCargoOrderList [get]
 export const getCargoOrderList = (params) => {
     return service({
         url: "/order/getCargoOrderList",
         method: 'get',
         params
     })
 }