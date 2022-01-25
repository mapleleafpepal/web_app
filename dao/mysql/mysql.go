package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

var db *sqlx.DB

func InitDB() (err error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
	)

	db, err = sqlx.Connect("mysql", dsn)

	if err != nil {
		return err
	}

	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle"))
	db.SetMaxOpenConns(viper.GetInt("mysql.max_connect_count"))

	return
}

func Close() {
	db.Close()
}
