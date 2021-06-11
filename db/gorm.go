package db

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MysqlDBConfig struct {
	User         string // 必填
	Password     string // 必填
	Host         string // 必填
	Port         int    // 必填
	Dbname       string // 必填
	MaxIdleConns int    // 必填
	MaxOpenConns int    // 必填

	LogWriter     io.Writer     // 选填 日志输出方式，不传默认os.Stdout
	Colorful      bool          // 选填 是否禁用彩色打印
	SlowThreshold time.Duration // 选填 慢sql阈值
	LogLevel      string        // 选填 db日志级别: silent,info,error,warn
}

func NewMysqlDB(conf *MysqlDBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=16s&readTimeout=3s&writeTimeout=5s",
		conf.User, conf.Password, conf.Host, conf.Port, conf.Dbname)

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: initDBLogger(conf),
	})
	if err != nil {
		return nil, err
	}

	cDB, err := DB.DB()
	if err != nil {
		return nil, err
	}
	err = cDB.Ping()
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	cDB.SetMaxIdleConns(conf.MaxIdleConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	cDB.SetMaxOpenConns(conf.MaxOpenConns)

	return DB, nil
}

var logLevelMap = map[string]logger.LogLevel{
	"silent": logger.Silent,
	"info":   logger.Info,
	"error":  logger.Error,
	"warn":   logger.Warn,
}

func initDBLogger(conf *MysqlDBConfig) logger.Interface {
	slow := conf.SlowThreshold
	if slow == 0 {
		slow = time.Second
	}

	w := conf.LogWriter
	if conf.LogWriter == nil {
		w = os.Stdout
	}

	level, ok := logLevelMap[conf.LogLevel]
	if !ok {
		level = logger.Info
	}

	dbLogger := logger.New(
		log.New(w, time.Now().Format("2006-01-02 15:04:05")+" ", 0), // io writer
		logger.Config{
			SlowThreshold: slow,
			LogLevel:      level,
			Colorful:      conf.Colorful,
		})
	return dbLogger
}
