package models

import (
	//"./mo"
	"fmt"
)

func init() {
	fmt.Println("abc")
}

type  User struct {
	Name string
	Age int64
}

func GetUserFactory()User  {
	//var a models.ApiModel

	//fmt.Println(a)
	return User{
		Name: "abc",
	}
}
