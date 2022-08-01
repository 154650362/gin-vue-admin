import service from '@/utils/request'

// @Tags RedisConnCls
// @Summary 创建RedisConnCls
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RedisConnCls true "创建RedisConnCls"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /RedisCls/createRedisConnCls [post]
export const createRedisConnCls = (data) => {
  return service({
    url: '/RedisCls/createRedisConnCls',
    method: 'post',
    data
  })
}

// @Tags RedisConnCls
// @Summary 删除RedisConnCls
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RedisConnCls true "删除RedisConnCls"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /RedisCls/deleteRedisConnCls [delete]
export const deleteRedisConnCls = (data) => {
  return service({
    url: '/RedisCls/deleteRedisConnCls',
    method: 'delete',
    data
  })
}

// @Tags RedisConnCls
// @Summary 删除RedisConnCls
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除RedisConnCls"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /RedisCls/deleteRedisConnCls [delete]
export const deleteRedisConnClsByIds = (data) => {
  return service({
    url: '/RedisCls/deleteRedisConnClsByIds',
    method: 'delete',
    data
  })
}

// @Tags RedisConnCls
// @Summary 更新RedisConnCls
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RedisConnCls true "更新RedisConnCls"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /RedisCls/updateRedisConnCls [put]
export const updateRedisConnCls = (data) => {
  return service({
    url: '/RedisCls/updateRedisConnCls',
    method: 'put',
    data
  })
}

// @Tags RedisConnCls
// @Summary 用id查询RedisConnCls
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.RedisConnCls true "用id查询RedisConnCls"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /RedisCls/findRedisConnCls [get]
export const findRedisConnCls = (params) => {
  return service({
    url: '/RedisCls/findRedisConnCls',
    method: 'get',
    params
  })
}

// @Tags RedisConnCls
// @Summary 分页获取RedisConnCls列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取RedisConnCls列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /RedisCls/getRedisConnClsList [get]
export const getRedisConnClsList = (params) => {
  return service({
    url: '/RedisCls/getRedisConnClsList',
    method: 'get',
    params
  })
}


export const testRedisConn = (data) => {
  return service({
    url: '/RedisCls/testRedisConn',
    method: 'post',
    data
  })
}