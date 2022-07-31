package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type RedisConnClsSearch struct{
    pkgTest.RedisConnCls
    request.PageInfo
}
