package initialize

import (
	"fmt"
	"mall-demo/internal/model"
	"time"

	"gorm.io/plugin/dbresolver"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

// InitDB 初始化数据库
func InitDB() {
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	database := viper.GetString("datasource.database")
	charset := viper.GetString("datasource.charset")

	// 主从数据库 主数据库负责读写 从数据库负责读
	masterDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", username, password, host, port, database, charset)
	slaveDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", username, password, host, port, database, charset)

	var ormLogger logger.Interface
	if gin.Mode() == "debug" { // 开发模式
		// 如果当前为调试模式, ormLogger将会被赋值为一个logger.Interface类型的日志对象，并设置日志级别为Info级别
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		// 如果不是调试模式，ormLogger将被赋值为一个默认的日志对象，日志级别为默认级别
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.Open(masterDSN), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 关闭表名自动变为复数
		},
	})
	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)                  // 最大空闲连接池
	sqlDB.SetMaxOpenConns(100)                 // 最大连接数
	sqlDB.SetConnMaxLifetime(time.Second * 30) // 最大连接时间
	DB = db

	// 主从数据库配置 读写分离
	err2 := DB.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(masterDSN)},                       // 写
		Replicas: []gorm.Dialector{mysql.Open(masterDSN), mysql.Open(slaveDSN)}, // 读
		// sources/replicas 负载均衡策略
		Policy: dbresolver.RandomPolicy{},
	}))
	if err2 != nil {
		panic(err2)
	}

	fmt.Println("database connect success")

	migration() // 自动迁移
}

// 自动迁移建表
func migration() {
	// 创建表的同时进行表属性配置
	err := DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(
		&model.Address{},
		&model.Admin{},
		&model.Carousel{},
		&model.Cart{},
		&model.Category{},
		&model.Favorite{},
		&model.Notice{},
		&model.Order{},
		&model.Product{},
		&model.ProductImg{},
		&model.User{},
		&model.SecKillProduct{},
		&model.SecKillOrder{},
	)
	if err != nil {
		panic(err)
	}
}

// GetDB 获取数据库连接
func GetDB() *gorm.DB {
	return DB
}
