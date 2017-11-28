package entities

import (
	"fmt"
)

func Userinfo_test() {
	Initial()
	var testuser = UserInfo{UserName: "summer"}
	user := NewUserInfo(testuser)
	user.DepartName = "shensi"
	fmt.Println("the item we want to insert is")
	fmt.Println(*user)
	err := UserInfoService.Save(user)
	if err != nil {
		fmt.Println("error occur")
	}
	user_in_database := UserInfoService.FindAll()
	fmt.Println("the database after insertion:")
	fmt.Println(user_in_database)
	fmt.Println("we want to find item which id is 7")
	user_id := UserInfoService.FindByID(7)
	fmt.Println("return item is:")
	fmt.Println(user_id)
	fmt.Println("we want to modify item with username winter and departname bahao which id is 6")
	u := NewUserInfo(UserInfo{UID: 6, UserName: "winter", DepartName: "bahao"})
	UserInfoService.ModifyInfoByID(u)
	user_in_database = UserInfoService.FindAll()
	fmt.Println("now the database is:")
	fmt.Println(user_in_database)
}
