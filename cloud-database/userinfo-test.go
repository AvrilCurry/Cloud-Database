package main

import (
	"Cloud-database/entities"
	"fmt"
	"time"
)

func main() {
	var Service = entities.UserInfoAtomicService{}
	var num, id int
	var username, departname string
	var created *time.Time

	for true {
		fmt.Scanln(&num)
		if num >= 6 {
			break
		}

		switch num {
		// Insert
		case 1:
			user := new(entities.UserInfo)
			fmt.Scanln(&username, &departname)
			user.UserName = username
			user.DepartName = departname
			Service.Save(entities.NewUserInfo(*user))
		// SearchAll
		case 2:
			list := Service.FindAll()
			for _, useritem := range list {
				fmt.Println(useritem)
			}
		// SearchByID
		case 3:
			fmt.Scanln(&id)
			user := Service.FindByID(id)
			fmt.Println(*user)
		// UpdateByID
		case 4:
			fmt.Scanln(&id, &username, &departname)
			t := time.Now()
			created = &t
			Service.UpdateByID(id, username, departname, created)
		// DeleteByID
		case 5:
			fmt.Scanln(&id)
			Service.DeleteByID(id)
		}

	}

}
