package dao

import (
	"github.com/zbnzl/paas-pod/internal/conf"
	"github.com/zbnzl/paas-pod/internal/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewMysqlDb(data *conf.Data) (*gorm.DB, error) {
	// 初始化MySQL连接和配置，返回连接对象
	db, err := gorm.Open(data.Database.Driver, data.Database.Source)
	if err != nil {
		panic("mysql connect failed")
	}

	// 设置Gorm的配置
	db.LogMode(true) // 如果需要打印SQL执行日志

	// 自动迁移模型到数据库中
	db.AutoMigrate(&model.Pod{}, &model.PodPort{}, &model.PodEnv{})

	return db, nil
}
