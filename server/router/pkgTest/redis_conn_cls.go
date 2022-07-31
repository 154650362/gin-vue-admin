package pkgTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RedisConnClsRouter struct {
}

// InitRedisConnClsRouter 初始化 RedisConnCls 路由信息
func (s *RedisConnClsRouter) InitRedisConnClsRouter(Router *gin.RouterGroup) {
	RedisClsRouter := Router.Group("RedisCls").Use(middleware.OperationRecord())
	RedisClsRouterWithoutRecord := Router.Group("RedisCls")
	var RedisClsApi = v1.ApiGroupApp.PkgtestApiGroup.RedisConnClsApi
	{
		RedisClsRouter.POST("createRedisConnCls", RedisClsApi.CreateRedisConnCls)   // 新建RedisConnCls
		RedisClsRouter.DELETE("deleteRedisConnCls", RedisClsApi.DeleteRedisConnCls) // 删除RedisConnCls
		RedisClsRouter.DELETE("deleteRedisConnClsByIds", RedisClsApi.DeleteRedisConnClsByIds) // 批量删除RedisConnCls
		RedisClsRouter.PUT("updateRedisConnCls", RedisClsApi.UpdateRedisConnCls)    // 更新RedisConnCls
	}
	{
		RedisClsRouterWithoutRecord.GET("findRedisConnCls", RedisClsApi.FindRedisConnCls)        // 根据ID获取RedisConnCls
		RedisClsRouterWithoutRecord.GET("getRedisConnClsList", RedisClsApi.GetRedisConnClsList)  // 获取RedisConnCls列表
	}
}
