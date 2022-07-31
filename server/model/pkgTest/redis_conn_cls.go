// 自动生成模板RedisConnCls
package pkgTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// RedisConnCls 结构体
type RedisConnCls struct {
      global.GVA_MODEL
      RedisClusterName  string `json:"redisClusterName" form:"redisClusterName" gorm:"column:redisclustername;comment:redisClusterName;size:50;"`
      IP  string `json:"IP" form:"IP" gorm:"column:ip;comment:IP;size:50;"`
      Port  string `json:"port" form:"port" gorm:"column:port;comment:端口;size:10;"`
      Pwd  string `json:"pwd" form:"pwd" gorm:"column:pwd;comment:password;size:10;"`
      Owner  string `json:"owner" form:"owner" gorm:"column:owner;comment:owner;size:10;"`
}


// TableName RedisConnCls 表名
func (RedisConnCls) TableName() string {
  return "redis_conn_cls"
}

