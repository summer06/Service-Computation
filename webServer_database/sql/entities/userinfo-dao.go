package entities

type userInfoDao DaoSource

var userInfoInsertStmt = "INSERT userinfo SET username=?,departname=?,created=?"

// Save .
func (dao *userInfoDao) Save(u *UserInfo) error {
    stmt, err := dao.Prepare(userInfoInsertStmt)
    checkErr(err)
    defer stmt.Close()

    res, err := stmt.Exec(u.UserName, u.DepartName, u.CreateAt)
    checkErr(err)
    if err != nil {
        return err
    }
    id, err := res.LastInsertId()
    if err != nil {
        return err
    }
    u.UID = int(id)
    return nil
}

var userInfoQueryAll = "SELECT * FROM userinfo"
var userInfoQueryByID = "SELECT * FROM userinfo where uid = ?"

// FindAll .
func (dao *userInfoDao) FindAll() []UserInfo {
    rows, err := dao.Query(userInfoQueryAll)
    checkErr(err)
    defer rows.Close()

    ulist := make([]UserInfo, 0, 0)
    for rows.Next() {
        u := UserInfo{}
        err := rows.Scan(&u.UID, &u.UserName, &u.DepartName, &u.CreateAt)
        checkErr(err)
        ulist = append(ulist, u)
    }
    return ulist
}

// FindByID .
func (dao *userInfoDao) FindByID(id int) *UserInfo {
    stmt, err := dao.Prepare(userInfoQueryByID)
    checkErr(err)
    defer stmt.Close()

    row := stmt.QueryRow(id)
    u := UserInfo{}
    err = row.Scan(&u.UID, &u.UserName, &u.DepartName, &u.CreateAt)
    checkErr(err)

    return &u
}

var modifyInfoByID = "UPDATE userinfo SET username=?, departname=? WHERE uid=?"

// modifyInfoByID
func (dao *userInfoDao) ModifyInfoByID (u *UserInfo) int64 {
  stmt, err := dao.Prepare(modifyInfoByID)
  checkErr(err)
  defer stmt.Close()

  res, err := stmt.Exec(u.UserName, u.DepartName, u.UID)
  checkErr(err)

  affect, err := res.RowsAffected()
  checkErr(err)

  return affect
}

var deleteAll = "DELETE FROM userinfo"
var deleteInfoByID = "DELETE FROM userinfo WHERE uid = ?"

// DelteAll
func (dao *userInfoDao) DeleteAll() error {
  stmt, err := dao.Prepare(deleteAll)
  checkErr(err)
  defer stmt.Close()

  _, err = stmt.Exec()
  checkErr(err)
  if err != nil {
    return err
  }
  return nil
}

// DeleteByID
func (dao *userInfoDao) DeleteByID(id int) int64 {
  stmt, err := dao.Prepare(deleteInfoByID)
  checkErr(err)
  defer stmt.Close()

  res, err := stmt.Exec(id)
  checkErr(err)

  affect, err := res.RowsAffected()
  checkErr(err)

  return affect
}
