package models

import (

	"errors"
	"fmt"
	"apiserver/db"
	"strconv"
	"time"
)

var (
	UserList map[string]*User
)

func init() {
	UserList = make(map[string]*User)
	u := User{"user_11111", "astaxie", "11111", Profile{"male", 20, "Singapore", "astaxie@gmail.com","apple"}}
	UserList["user_11111"] = &u
}

type User struct {
	Id       string
	Username string
	Password string
	Profile  Profile
}

type Profile struct {
	Gender  string
	Age     int
	Address string
	Email   string
	Love	string
}


type UserInfo struct {
	UserId			int			`json:"userid"`				// 用户ID
	UserName		string		`json:"username"`			// 用户名称
	PassWord		string		`json:"password"`			// 登录密码
	Token			string 		`json:"token"`				// 令牌
	SessionKey		string		`json:"session_key"`		// 会话秘钥
	UserAge			int64		`json:"userage"`			// 年龄
	UserSex			string		`json:"usersex"`			// 性别
	Name			string		`json:"name"`				// 真实姓名
	Remain			int64		`json:"remain"`				// 余额
	Addr			string		`json:"addr"`				// 地址
	CardCount		int64		`json:"card_count"`			// 银行卡总数
	Status 			int										// 在线状态
	CreateTime		string									// 注册时间
	UpdateTime		string									// 更新时间
}
// 查询数据
func Query(sqlQuery string){
	 res := db.Exec(sqlQuery)
	 fmt.Print(res)

}


func AddUser(u User) string {
	u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	UserList[u.Id] = &u
	return u.Id
}

func GetUser(uid string) (u *User, err error) {
	if u, ok := UserList[uid]; ok {
		return u, nil
	}
	return nil, errors.New("User not exists")
}

func GetAllUsers() map[string]*User {
	return UserList
}

func UpdateUser(uid string, uu *User) (a *User, err error) {
	if u, ok := UserList[uid]; ok {
		if uu.Username != "" {
			u.Username = uu.Username
		}
		if uu.Password != "" {
			u.Password = uu.Password
		}
		if uu.Profile.Age != 0 {
			u.Profile.Age = uu.Profile.Age
		}
		if uu.Profile.Address != "" {
			u.Profile.Address = uu.Profile.Address
		}
		if uu.Profile.Gender != "" {
			u.Profile.Gender = uu.Profile.Gender
		}
		if uu.Profile.Email != "" {
			u.Profile.Email = uu.Profile.Email
		}
		return u, nil
	}
	return nil, errors.New("User Not Exist")
}

func Login(username, password string) bool {
	for _, u := range UserList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}
