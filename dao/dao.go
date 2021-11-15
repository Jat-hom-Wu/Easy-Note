package dao

import "github.com/jinzhu/gorm"
import _ "github.com/jinzhu/gorm/dialects/mysql"
import "fmt"

type MysqlSetting struct{
	User string
	Password string
	Host string
	Port string
	DBname string
}

func InitMysql(cfg MysqlSetting) (*gorm.DB, error){
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBname)
	fmt.Println(dsn)
	DB,err := gorm.Open("mysql",dsn)
	if err != nil{
		fmt.Println("gorm open mysql failed")
	}else{
		fmt.Println("gorm open mysql succeed")
	}
	return DB,err
}