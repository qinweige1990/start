import service from '@/utils/request'

// @Tags Order
// @Summary 创建Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Order true "创建Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /order/createOrder [post]
export const createPreOrder = (data) => {
     return service({
         url: "/preOrder/createPreOrder",
         method: 'post',
         data
     })
 }


// @Tags Order
// @Summary 删除Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Order true "删除Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /order/deleteOrder [delete]
 export const deletePreOrder = (data) => {
     return service({
         url: "/preOrder/deletePreOrder",
         method: 'delete',
         data
     })
 }

// @Tags Order
// @Summary 删除Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /order/deleteOrder [delete]
 export const deletePreOrderByIds = (data) => {
     return service({
         url: "/preOrder/deletePreOrderByIds",
         method: 'delete',
         data
     })
 }

// @Tags Order
// @Summary 更新Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Order true "更新Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /order/updateOrder [put]
 export const updatePreOrder = (data) => {
     return service({
         url: "/preOrder/updatePreOrder",
         method: 'put',
         data
     })
 }


// @Tags Order
// @Summary 用id查询Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Order true "用id查询Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /order/findOrder [get]
 export const findPreOrder = (params) => {
     return service({
         url: "/preOrder/findPreOrder",
         method: 'get',
         params
     })
 }


// @Tags Order
// @Summary 分页获取Order列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Order列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /order/getOrderList [get]
 export const getPreOrderList = (params) => {
     return service({
         url: "/preOrder/getPreOrderList",
         method: 'get',
         params
     })
 }
