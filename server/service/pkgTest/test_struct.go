package pkgTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    pkgTestReq "github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest/request"
)

type TestStructService struct {
}

// CreateTestStruct 创建TestStruct记录
// Author [piexlmax](https://github.com/piexlmax)
func (testStructService *TestStructService) CreateTestStruct(testStruct pkgTest.TestStruct) (err error) {
	err = global.GVA_DB.Create(&testStruct).Error
	return err
}

// DeleteTestStruct 删除TestStruct记录
// Author [piexlmax](https://github.com/piexlmax)
func (testStructService *TestStructService)DeleteTestStruct(testStruct pkgTest.TestStruct) (err error) {
	err = global.GVA_DB.Delete(&testStruct).Error
	return err
}

// DeleteTestStructByIds 批量删除TestStruct记录
// Author [piexlmax](https://github.com/piexlmax)
func (testStructService *TestStructService)DeleteTestStructByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]pkgTest.TestStruct{},"id in ?",ids.Ids).Error
	return err
}

// UpdateTestStruct 更新TestStruct记录
// Author [piexlmax](https://github.com/piexlmax)
func (testStructService *TestStructService)UpdateTestStruct(testStruct pkgTest.TestStruct) (err error) {
	err = global.GVA_DB.Save(&testStruct).Error
	return err
}

// GetTestStruct 根据id获取TestStruct记录
// Author [piexlmax](https://github.com/piexlmax)
func (testStructService *TestStructService)GetTestStruct(id uint) (testStruct pkgTest.TestStruct, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&testStruct).Error
	return
}

// GetTestStructInfoList 分页获取TestStruct记录
// Author [piexlmax](https://github.com/piexlmax)
func (testStructService *TestStructService)GetTestStructInfoList(info pkgTestReq.TestStructSearch) (list []pkgTest.TestStruct, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&pkgTest.TestStruct{})
    var testStructs []pkgTest.TestStruct
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&testStructs).Error
	return  testStructs, total, err
}
