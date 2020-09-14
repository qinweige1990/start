package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitCarBossRouter(Router *gin.RouterGroup) {
	CarBossRouter := Router.Group("carBoss").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		CarBossRouter.POST("createCarBoss", v1.CreateCarBoss)   // 新建CarBoss
		CarBossRouter.DELETE("deleteCarBoss", v1.DeleteCarBoss) // 删除CarBoss
		CarBossRouter.DELETE("deleteCarBossByIds", v1.DeleteCarBossByIds) // 批量删除CarBoss
		CarBossRouter.PUT("updateCarBoss", v1.UpdateCarBoss)    // 更新CarBoss
		CarBossRouter.GET("findCarBoss", v1.FindCarBoss)        // 根据ID获取CarBoss
		CarBossRouter.GET("getCarBossList", v1.GetCarBossList)  // 获取CarBoss列表
	}
}
