package entities

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

// Save .
func (*UserInfoAtomicService) Save(u *UserInfo) error {
    _, err := engine.Insert(u)
    checkErr(err)
    if err != nil {
      return err
    }
    return nil
}

// FindAll .
func (*UserInfoAtomicService) FindAll() []UserInfo {
    everyone := make([]UserInfo, 0)
    err := engine.Find(&everyone)
    checkErr(err)
    return everyone
}

// FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
    u := new(UserInfo)
    engine.ID(id).Get(u)
    return u
}

// ModifyInfoByID
func (*UserInfoAtomicService) ModifyInfoByID(u *UserInfo) int64 {
  affect, err := engine.Id(u.Uid).Update(u)
  checkErr(err)
  return affect
}

//DeleteAll
// func (*UserInfoAtomicService) DeleteAll() error {
//   dao := userInfoDao{mydb}
//   return dao.DeleteAll()
// }

// DeleteByID
func (*UserInfoAtomicService) DeleteByID(id int) int64 {
  u := new(UserInfo)
  affect, err := engine.ID(id).Delete(u)
  checkErr(err)
  return affect
}
