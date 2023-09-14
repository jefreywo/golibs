package mysql

import (
	"os"
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestNewMysqlDB(t *testing.T) {
	// dblog, err := os.OpenFile("gorm.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	t.Fatalf("open gorm.log failed: %v\n", err)
	// }

	db, err := NewMysqlDB(&MysqlDBConfig{
		User:         "root",
		Password:     "12345",
		Host:         "127.0.0.1",
		Port:         3306,
		Dbname:       "test",
		MaxIdleConns: 5,
		MaxOpenConns: 80,

		LogWriter:     os.Stdout, // dblog
		Colorful:      true,
		SlowThreshold: time.Second * 2,
		LogLevel:      "info",
	})
	if err != nil {
		t.Fatalf("NewMysqlDB failed: %v\n", err)
	}
	var u JUser
	stmt := db.Session(&gorm.Session{DryRun: true}).First(&u, 1232).Statement
	t.Log(stmt.SQL.String())
}
