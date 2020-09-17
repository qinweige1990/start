package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitInformationRouter(Router *gin.RouterGroup) {
	InformationRouter := Router.Group("information").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		InformationRouter.POST("createInformation", v1.CreateInformation)   // 新建Information
		InformationRouter.DELETE("deleteInformation", v1.DeleteInformation) // 删除Information
		InformationRouter.DELETE("deleteInformationByIds", v1.DeleteInformationByIds) // 批量删除Information
		InformationRouter.PUT("updateInformation", v1.UpdateInformation)    // 更新Information
		InformationRouter.GET("findInformation", v1.FindInformation)        // 根据ID获取Information
		InformationRouter.GET("getInformationList", v1.GetInformationList)  // 获取Information列表
	}
}
