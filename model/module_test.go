package model

import "testing"

func TestUserCreateOne(t *testing.T){
	res := UserCreateOne("hello", "123")
	if res == 0{
		t.Errorf("error")
	}
}