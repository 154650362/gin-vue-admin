package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/myPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/router/pkgTest"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Mypkg   myPkg.RouterGroup
	Pkgtest pkgTest.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
