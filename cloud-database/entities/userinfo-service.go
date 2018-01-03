package entities

import "time"

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

// UpdateByID .
func (*UserInfoAtomicService) UpdateByID(id int, username string, departname string, created *time.Time) error {
	tx, err := mydb.Begin()
	checkErr(err)

	dao := userInfoDao{tx}
	err = dao.UpdateByID(id, username, departname, created)

	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}
	return nil
}

// DeleteByID .
func (*UserInfoAtomicService) DeleteByID(id int) error {
	tx, err := mydb.Begin()
	checkErr(err)

	dao := userInfoDao{tx}
	err = dao.DeleteByID(id)

	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}
	return nil
}
