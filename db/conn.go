package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"strconv"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

const asyncTaskBacklog = 128

var (
	Database *xorm.Engine
	chWrite  chan interface{}
	chUpdate chan interface{}
)

func DbStartUp() func() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		strconv.Itoa(viper.GetInt("mysql.port")),
		viper.GetString("mysql.dbname"),
		viper.GetString("mysql.charset"))

	db, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		fmt.Println("mysql start err")
		panic(err)
	}
	Database = db
	Database.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
	Database.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))

	Database.SetLogLevel(log.LogLevel(viper.GetInt("mysql.log_level")))
	Database.ShowSQL(viper.GetBool("mysql.show_sql"))
	chWrite = make(chan interface{}, asyncTaskBacklog)
	chUpdate = make(chan interface{}, asyncTaskBacklog)
	envInit()

	closer := func() {
		close(chWrite)
		close(chUpdate)
		Database.Close()
		fmt.Println("db stop!!!")
	}

	return closer
}

func envInit() {
	// async task
	go func() {
		for {
			select {
			case t, ok := <-chWrite:
				if !ok {
					return
				}

				if _, err := Database.Insert(t); err != nil {
					fmt.Println("db insert error:", err)
				}

			case t, ok := <-chUpdate:
				if !ok {
					return
				}

				if _, err := Database.Update(t); err != nil {
					fmt.Println("db update error:", err)
				}
			}
		}
	}()

	// 定时ping数据库, 保持连接池连接
	go func() {
		ticker := time.NewTicker(time.Minute * 5)
		for {
			select {
			case <-ticker.C:
				Database.Ping()
			}
		}
	}()
}
