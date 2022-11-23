package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strings"
)

var DBClient *gorm.DB

func init() {
	Setup()
}
func Setup() {
	d, err := InitDBClient(&MysqlConfig{
		Addr:       "localhost:3306",
		User:       "root",
		Passwd:     "123456",
		DB:         "test",
		Charset:    "utf8mb4",
		Options:    "parseTime=false",
		TimeoutSec: 1000,
		Tracing:    true,
	})
	if err != nil {
		panic(err)
	}
	DBClient = d
}

type MysqlConfig struct {
	Addr         string `toml:"addr" json:"addr" validate:"hostname_port" long:"addr"`
	User         string `toml:"user" json:"user" long:"user" description:"mysql user"`
	Passwd       string `toml:"passwd" json:"passwd" long:"passwd" description:"mysql passwd"`
	DB           string `toml:"db" json:"db" validate:"required" long:"db" description:"mysql database name"`
	MaxOpenCount int    `toml:"max_open_count" json:"max_open_count" validate:"required" long:"max_open_count"`
	MaxIdleCount int    `toml:"max_idle_count" json:"max_idle_count" validate:"required" long:"max_idle_count"`
	Charset      string `toml:"charset" json:"charset" long:"charset"`
	TimeoutSec   int    `toml:"timeout_sec" json:"timeout_sec" long:"timeout_sec"`
	Options      string `toml:"options" json:"options" long:"options"`
	Tracing      bool   `toml:"tracing" json:"tracing" long:"tracing"`
}

const MysqlDefaultCharset = "utf8mb4"

// InitDBClient InitDBClient
func InitDBClient(cfg *MysqlConfig) (*gorm.DB, error) {
	if cfg.Charset == "" {
		cfg.Charset = MysqlDefaultCharset
	}

	dsn := fmt.Sprintf("%s:%s@"+"tcp(%s)/%s?charset=%s",
		cfg.User, cfg.Passwd, cfg.Addr, cfg.DB, cfg.Charset)
	if cfg.TimeoutSec > 0 {
		// timeout in seconds has "s"
		dsn += fmt.Sprintf("&timeout=%ds", cfg.TimeoutSec)
	}
	if !strings.Contains(cfg.Options, "parseTime=") {
		dsn += "&parseTime=true"
	}
	if !strings.Contains(cfg.Options, "loc=") {
		dsn += "&loc=Local"
	}
	// other options
	if cfg.Options != "" {
		dsn += "&" + strings.Trim(cfg.Options, "&")
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if cfg.MaxIdleCount > 0 {
		sqlDB.SetMaxIdleConns(cfg.MaxIdleCount)
	}
	if cfg.MaxOpenCount > 0 {
		sqlDB.SetMaxOpenConns(cfg.MaxOpenCount)
	}
	return db, nil
}
