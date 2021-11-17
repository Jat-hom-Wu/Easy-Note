package model

import (
	"fmt"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"errors"
)

var DB *gorm.DB

type ToDo struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

type User struct{
	Name string `gorm:"primaryKey"`
	Password string
}


func CreateDataList(todo *ToDo) int{
	DB.AutoMigrate(&ToDo{})
	// err := DB.Create(&todo).Error
	err := DB.Exec("insert into to_dos(title,status) values(?,?)",todo.Title,todo.Status)

	if	err.Error != nil{
		fmt.Println("create data error")
		fmt.Println("error: ", err.Error)
		return 0
	}else{
		return 1
	}
}

func GetAllData() []ToDo{
	var allData []ToDo
	r := DB.Raw("select * from to_dos order by id desc").Scan(&allData)
	fmt.Println("ok")
	if r.Error != nil{
		fmt.Println("get all data failed")
	}
	return allData
}

func UpdateOneData(_id int, _status bool){
	var strStatus string
	if _status{
		strStatus = "true"
	}else{
		strStatus = "false"
	}
	dsn := fmt.Sprintf("update to_dos set status = %s where id = %d", strStatus,_id)
	r := DB.Exec(dsn)
	if r.Error != nil{
		fmt.Println("update one data error")
	}
}

func DeleteOneData(_id int){
	dsn := fmt.Sprintf("delete from to_dos where id = %d", _id)
	r := DB.Exec(dsn)
	if r.Error != nil{
		fmt.Println("delete one data error")
	}
}

func FindOneData(_id int) ToDo{
	var result ToDo
	dsn := fmt.Sprintf("select * from to_dos where id = %d", _id)
	r := DB.Raw(dsn).Scan(&result)
	if r.Error != nil{
		fmt.Println("find one data error")
	}
	return result
}

func UserCreateOne(name string, password string) int{
	DB.AutoMigrate(&User{})
	err := DB.Exec("insert into users(name,password) values(?,?)",name,password)
	if	err.Error != nil{
		fmt.Println("user table create data error")
		fmt.Println("error: ", err.Error)
		return 0
	}else{
		return 1
	}
}

func UserFind(name string) (int,User){
	var userData User
	dsn := fmt.Sprintf("select * from users where name = %s", name)
	r := DB.Raw(dsn).Scan(&userData)
	isNil := errors.Is(r.Error,gorm.ErrRecordNotFound)	//判断是否为空
	if r.Error != nil {
		fmt.Println("fine name in users error")
	}
	if isNil{
		fmt.Println("no this user")
		return 1,userData
	}
	return 0,userData

}
