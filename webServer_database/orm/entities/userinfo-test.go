package entities

import (
	"fmt"
)

func Userinfo_test() {
	Initial()
	// var user = new(UserInfo)
	// user.UserName = "orm_summer"
	// user.DepartName = "orm_shensi"
	// fmt.Println("the item we want to insert is")
	// fmt.Println(*user)
	// err := UserInfoService.Save(user)
	// if err != nil {
	// 	fmt.Println("error occur")
	// }
	user_in_database := UserInfoService.FindAll()
	// fmt.Println("the database after insertion:")
	// fmt.Println(user_in_database)
	// fmt.Println("we want to find item which id is 2")
	// user_id := UserInfoService.FindByID(2)
	// fmt.Println("return item is:")
	// fmt.Println(user_id)
	// fmt.Println("we want to modify item with username winter and departname bahao which id is 3")
	// u := new(UserInfo)
  // u.Uid = 3
  // u.UserName = "winter"
  // u.DepartName = "orm_shensiyuan"
	// UserInfoService.ModifyInfoByID(u)
	// user_in_database = UserInfoService.FindAll()
	fmt.Println("now the database is:")
	fmt.Println(user_in_database)
  UserInfoService.DeleteByID(2)
  user_in_database = UserInfoService.FindAll()
	fmt.Println("now the database is:")
	fmt.Println(user_in_database)

}
