package entities

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

// Save .
func (*UserInfoAtomicService) Save(u *UserInfo) error {
    tx, err := mydb.Begin()
    checkErr(err)

    dao := userInfoDao{tx}
    err = dao.Save(u)

    if err == nil {
        tx.Commit()
    } else {
        tx.Rollback()
    }
    return nil
}

// FindAll .
func (*UserInfoAtomicService) FindAll() []UserInfo {
    dao := userInfoDao{mydb}
    return dao.FindAll()
}

// FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
    dao := userInfoDao{mydb}
    return dao.FindByID(id)
}

// ModifyInfoByID
func (*UserInfoAtomicService) ModifyInfoByID(u *UserInfo) int64 {
  dao := userInfoDao{mydb}
  return dao.ModifyInfoByID(u)
}

//DeleteAll
func (*UserInfoAtomicService) DeleteAll() error {
  dao := userInfoDao{mydb}
  return dao.DeleteAll()
}

// DeleteByID
func (*UserInfoAtomicService) DeleteByID(id int) int64 {
  dao := userInfoDao{mydb}
  return dao.DeleteByID(id)
}
