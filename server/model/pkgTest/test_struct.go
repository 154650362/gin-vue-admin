// 自动生成模板TestStruct
package pkgTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// TestStruct 结构体
type TestStruct struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:name;size:200;"`
}


// TableName TestStruct 表名
func (TestStruct) TableName() string {
  return "test_struct"
}

