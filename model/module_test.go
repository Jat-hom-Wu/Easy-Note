package model

<<<<<<< HEAD
import (
	"testing"
)

=======
import "testing"
>>>>>>> v1

func TestUserCreateOne(t *testing.T){
	res := UserCreateOne("hello", "123")
	if res == 0{
		t.Errorf("error")
	}
<<<<<<< HEAD
}

func TestCreateDataList(t *testing.T){
	todo := &ToDo{
		User:"xiaoming",
		Title:"test",
		Status:true,
	}
	err := CreateDataList(todo)
	if err == 0{
		t.Errorf("error")
	}
=======
>>>>>>> v1
}