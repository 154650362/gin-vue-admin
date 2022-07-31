package pkgTest

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    pkgTestReq "github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest/request"
	"github.com/go-redis/redis/v8"
)

type RedisConnClsService struct {
}

// CreateRedisConnCls 创建RedisConnCls记录
// Author [piexlmax](https://github.com/piexlmax)
func (RedisClsService *RedisConnClsService) CreateRedisConnCls(RedisCls pkgTest.RedisConnCls) (err error) {
	err = global.GVA_DB.Create(&RedisCls).Error
	return err
}

// DeleteRedisConnCls 删除RedisConnCls记录
// Author [piexlmax](https://github.com/piexlmax)
func (RedisClsService *RedisConnClsService)DeleteRedisConnCls(RedisCls pkgTest.RedisConnCls) (err error) {
	err = global.GVA_DB.Delete(&RedisCls).Error
	return err
}

// DeleteRedisConnClsByIds 批量删除RedisConnCls记录
// Author [piexlmax](https://github.com/piexlmax)
func (RedisClsService *RedisConnClsService)DeleteRedisConnClsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]pkgTest.RedisConnCls{},"id in ?",ids.Ids).Error
	return err
}

// UpdateRedisConnCls 更新RedisConnCls记录
// Author [piexlmax](https://github.com/piexlmax)
func (RedisClsService *RedisConnClsService)UpdateRedisConnCls(RedisCls pkgTest.RedisConnCls) (err error) {
	err = global.GVA_DB.Save(&RedisCls).Error
	return err
}

// GetRedisConnCls 根据id获取RedisConnCls记录
// Author [piexlmax](https://github.com/piexlmax)
func (RedisClsService *RedisConnClsService)GetRedisConnCls(id uint) (RedisCls pkgTest.RedisConnCls, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&RedisCls).Error
	return
}

// GetRedisConnClsInfoList 分页获取RedisConnCls记录
// Author [piexlmax](https://github.com/piexlmax)
func (RedisClsService *RedisConnClsService)GetRedisConnClsInfoList(info pkgTestReq.RedisConnClsSearch) (list []pkgTest.RedisConnCls, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&pkgTest.RedisConnCls{})
    var RedisClss []pkgTest.RedisConnCls
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.RedisClusterName != "" {
        db = db.Where("redisclustername LIKE ?","%"+ info.RedisClusterName+"%")
    }
    if info.IP != "" {
        db = db.Where("ip LIKE ?","%"+ info.IP+"%")
    }
    if info.Owner != "" {
        db = db.Where("owner = ?",info.Owner)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&RedisClss).Error
	return  RedisClss, total, err
}


// 测试连接是否通，standalone版本
func (RedisClsService *RedisConnClsService)TestRedisConn(RedisCls pkgTest.RedisConnCls) (err error) {
	//err = global.GVA_DB.Delete(&RedisCls).Error
	addr := RedisCls.IP + ":" + RedisCls.Port
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: RedisCls.Pwd, // no password set
	})
	_, err = client.Ping(context.Background()).Result()
	return err
}

// 测试连接是否通，版本 redis cluster
// todo 密码需要测试下
func (RedisClsService *RedisConnClsService)TestRedisConnCls(RedisCls pkgTest.RedisConnCls) (err error) {
	//err = global.GVA_DB.Delete(&RedisCls).Error
	addr := RedisCls.IP + ":" + RedisCls.Port
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:     []string{addr},
		Password: RedisCls.Pwd,
	})
	_, err = client.Ping(context.Background()).Result()
	return err
}