package global

import (
	"gin-vue-admin/model/request"
	"github.com/gin-gonic/gin"
)

const GIN_CONTEXT_THREAD_LOCAL = "GIN_CONTEXT_THREAD_LOCAL"

func SetGinCtx(ctx gin.Context) {
	Set(GIN_CONTEXT_THREAD_LOCAL, ctx)
}

func GetGinCtx() *gin.Context {
	return Get(GIN_CONTEXT_THREAD_LOCAL).(*gin.Context)
}

func SetString(k, v string) {
	Set(k, v)
}

func GetString(k string) string {
	if v := Get(k); v != nil {
		if ret, ok := v.(string); ok {
			return ret
		}
	}
	return ""
}


func GetClaim(ctx *gin.Context) *request.CustomClaims {
	if claims, exist := ctx.Get("claims"); exist {
		return claims.(*request.CustomClaims)
	}
	return nil
}

