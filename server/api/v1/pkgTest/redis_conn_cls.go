package pkgTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    pkgTestReq "github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type RedisConnClsApi struct {
}

var RedisClsService = service.ServiceGroupApp.PkgtestServiceGroup.RedisConnClsService


// CreateRedisConnCls 创建RedisConnCls
// @Tags RedisConnCls
// @Summary 创建RedisConnCls
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgTest.RedisConnCls true "创建RedisConnCls"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /RedisCls/createRedisConnCls [post]
func (RedisClsApi *RedisConnClsApi) CreateRedisConnCls(c *gin.Context) {
	var RedisCls pkgTest.RedisConnCls
	_ = c.ShouldBindJSON(&RedisCls)
    verify := utils.Rules{
        "RedisClusterName":{utils.NotEmpty()},
        "IP":{utils.NotEmpty()},
        "Port":{utils.NotEmpty()},
    }
	if err := utils.Verify(RedisCls, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := RedisClsService.CreateRedisConnCls(RedisCls); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteRedisConnCls 删除RedisConnCls
// @Tags RedisConnCls
// @Summary 删除RedisConnCls
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgTest.RedisConnCls true "删除RedisConnCls"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /RedisCls/deleteRedisConnCls [delete]
func (RedisClsApi *RedisConnClsApi) DeleteRedisConnCls(c *gin.Context) {
	var RedisCls pkgTest.RedisConnCls
	_ = c.ShouldBindJSON(&RedisCls)
	if err := RedisClsService.DeleteRedisConnCls(RedisCls); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteRedisConnClsByIds 批量删除RedisConnCls
// @Tags RedisConnCls
// @Summary 批量删除RedisConnCls
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除RedisConnCls"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /RedisCls/deleteRedisConnClsByIds [delete]
func (RedisClsApi *RedisConnClsApi) DeleteRedisConnClsByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := RedisClsService.DeleteRedisConnClsByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRedisConnCls 更新RedisConnCls
// @Tags RedisConnCls
// @Summary 更新RedisConnCls
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgTest.RedisConnCls true "更新RedisConnCls"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /RedisCls/updateRedisConnCls [put]
func (RedisClsApi *RedisConnClsApi) UpdateRedisConnCls(c *gin.Context) {
	var RedisCls pkgTest.RedisConnCls
	_ = c.ShouldBindJSON(&RedisCls)
      verify := utils.Rules{
          "RedisClusterName":{utils.NotEmpty()},
          "IP":{utils.NotEmpty()},
          "Port":{utils.NotEmpty()},
      }
    if err := utils.Verify(RedisCls, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := RedisClsService.UpdateRedisConnCls(RedisCls); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindRedisConnCls 用id查询RedisConnCls
// @Tags RedisConnCls
// @Summary 用id查询RedisConnCls
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query pkgTest.RedisConnCls true "用id查询RedisConnCls"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /RedisCls/findRedisConnCls [get]
func (RedisClsApi *RedisConnClsApi) FindRedisConnCls(c *gin.Context) {
	var RedisCls pkgTest.RedisConnCls
	_ = c.ShouldBindQuery(&RedisCls)
	if reRedisCls, err := RedisClsService.GetRedisConnCls(RedisCls.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reRedisCls": reRedisCls}, c)
	}
}

// GetRedisConnClsList 分页获取RedisConnCls列表
// @Tags RedisConnCls
// @Summary 分页获取RedisConnCls列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query pkgTestReq.RedisConnClsSearch true "分页获取RedisConnCls列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /RedisCls/getRedisConnClsList [get]
func (RedisClsApi *RedisConnClsApi) GetRedisConnClsList(c *gin.Context) {
	var pageInfo pkgTestReq.RedisConnClsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := RedisClsService.GetRedisConnClsInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
