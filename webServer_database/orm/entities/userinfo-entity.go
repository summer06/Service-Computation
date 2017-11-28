package entities

import (
    "time"
    // "github.com/go-xorm/xorm"
)

// UserInfo .
type UserInfo struct {
    Uid        int   `xorm:"pk autoincr notnull unique"` //语义标签
    UserName   string `xorm:"notnull"`
    DepartName string
    CreateAt   *time.Time `xorm:"created"`
}
