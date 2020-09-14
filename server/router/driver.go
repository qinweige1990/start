package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDriverRouter(Router *gin.RouterGroup) {
	DriverRouter := Router.Group("driver").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		DriverRouter.POST("createDriver", v1.CreateDriver)   // 新建Driver
		DriverRouter.DELETE("deleteDriver", v1.DeleteDriver) // 删除Driver
		DriverRouter.DELETE("deleteDriverByIds", v1.DeleteDriverByIds) // 批量删除Driver
		DriverRouter.PUT("updateDriver", v1.UpdateDriver)    // 更新Driver
		DriverRouter.GET("findDriver", v1.FindDriver)        // 根据ID获取Driver
		DriverRouter.GET("getDriverList", v1.GetDriverList)  // 获取Driver列表
	}
}
