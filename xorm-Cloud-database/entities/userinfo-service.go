package entities

import (
	"fmt"
	"time"
)

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

// Save .
func (*UserInfoAtomicService) Save(u *UserInfo) error {
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	checkErr(err)

	_, err = session.Insert(u)

	if err == nil {
		session.Commit()
	} else {
		session.Rollback()
	}
	return nil
}

// FindAll .
func (*UserInfoAtomicService) FindAll() []UserInfo {
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	checkErr(err)

	ulist := make([]UserInfo, 0, 0)
	err = session.Find(&ulist)
	checkErr(err)

	return ulist
}

// FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	checkErr(err)

	user := new(UserInfo)
	has, err := session.Id(id).Get(user)
	checkErr(err)

	fmt.Println(has)
	return user
}

// UpdateByID .
func (*UserInfoAtomicService) UpdateByID(id int, username string, departname string) error {
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	checkErr(err)
	/*user := new(UserInfo)
	user.UserName = username
	user.DepartName = departname
	user.UpdatedAt = updated*/

	var updated time.Time
	t := time.Now()
	st, _ := time.LoadLocation("Asia/Shanghai")
	updated = t.In(st)
	fmt.Println("time:", st, updated)

	sql := "update `xorm_user_info` set user_name=?,depart_name=?,updated_at=? WHERE i_d=?"
	affected, err := engine.Exec(sql, username, departname, updated, id)

	if err == nil {
		fmt.Println("yes")
		session.Commit()
	} else {
		fmt.Printf(err.Error())
		session.Rollback()
	}

	fmt.Println(affected)
	return nil
}

// DeleteByID .
func (*UserInfoAtomicService) DeleteByID(id int) error {
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	checkErr(err)

	user := new(UserInfo)

	affected, err := session.Id(id).Delete(user)
	if err == nil {
		session.Commit()
	} else {
		session.Rollback()
	}

	fmt.Println(affected)
	return nil
}
