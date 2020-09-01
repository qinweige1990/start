package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitCargoOrderRouter(Router *gin.RouterGroup) {
	CargoOrderRouter := Router.Group("order").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		CargoOrderRouter.POST("createCargoOrder", v1.CreateCargoOrder)   // 新建CargoOrder
		CargoOrderRouter.DELETE("deleteCargoOrder", v1.DeleteCargoOrder) // 删除CargoOrder
		CargoOrderRouter.DELETE("deleteCargoOrderByIds", v1.DeleteCargoOrderByIds) // 批量删除CargoOrder
		CargoOrderRouter.PUT("updateCargoOrder", v1.UpdateCargoOrder)    // 更新CargoOrder
		CargoOrderRouter.GET("findCargoOrder", v1.FindCargoOrder)        // 根据ID获取CargoOrder
		CargoOrderRouter.GET("getCargoOrderList", v1.GetCargoOrderList)  // 获取CargoOrder列表
	}
}
