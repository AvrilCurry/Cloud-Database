package service

import (
	"net/http"
	"strconv"

	"xorm-Cloud-database/entities"

	"github.com/unrolled/render"
)

func postUserInfoHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if len(req.Form["username"][0]) == 0 {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"Bad Input Without \"userid\"!"})
			return
		}
		u := entities.NewUserInfo(entities.UserInfo{UserName: req.Form["username"][0]})
		u.DepartName = req.Form["departname"][0]
		entities.UserInfoService.Save(u)
		formatter.JSON(w, http.StatusOK, u)
	}
}

func getUserInfoHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if len(req.Form["userid"][0]) != 0 {
			i, _ := strconv.ParseInt(req.Form["userid"][0], 10, 32)

			u := entities.UserInfoService.FindByID(int(i))
			formatter.JSON(w, http.StatusBadRequest, u)
			return
		}
		ulist := entities.UserInfoService.FindAll()
		formatter.JSON(w, http.StatusOK, ulist)
	}
}

// updateUserInfoHandler .
func updateUserInfoHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if len(req.Form["userid"][0]) != 0 {
			i, _ := strconv.ParseInt(req.Form["userid"][0], 10, 32)

			if len(req.Form["username"][0]) == 0 {
				formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"Bad Input Without \"username\"!"})
				return
			}

			var username = req.Form["username"][0]
			var departname = req.Form["departname"][0]

			err := entities.UserInfoService.UpdateByID(int(i), username, departname)

			if err == nil {
				formatter.JSON(w, http.StatusOK, "Updated Sucessfully")
				return
			}
		}
		formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"Bad Input Without \"userid\"!"})
		return
	}
}

// deleteUserInfoHandler .
func deleteUserInfoHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()

		if len(req.Form["userid"][0]) == 0 {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"Bad Input Without \"userid\"!"})
			return
		}

		i, _ := strconv.ParseInt(req.Form["userid"][0], 10, 32)
		err := entities.UserInfoService.DeleteByID(int(i))

		if err == nil {
			formatter.JSON(w, http.StatusOK, "Delete Successfully")
			return
		}
	}
}
