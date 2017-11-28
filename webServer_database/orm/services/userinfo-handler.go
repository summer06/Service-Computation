package service

import (
	"net/http"
	"strconv"

	"webServer_database/orm/entities"

	"github.com/unrolled/render"
)

func postUserInfoHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if len(req.Form["username"][0]) == 0 {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"Bad Input!"})
			return
		}
		// u := entities.NewUserInfo(entities.UserInfo{UserName: req.Form["username"][0]})
		u := new(entities.UserInfo)
		u.UserName = req.Form["username"][0]
		u.DepartName = req.Form["departname"][0]
		entities.UserInfoService.Save(u)
		formatter.JSON(w, http.StatusOK, u)
	}
}

func modifyuserinfoHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if len(req.Form["uid"][0]) == 0 {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"You must have an uid"})
			return
		}
		u := new(entities.UserInfo)
		u.UserName = req.Form["username"][0]
		u.DepartName = req.Form["departname"][0]
		var err error
		u.Uid, err = strconv.Atoi(req.Form["uid"][0])
		if err != nil {
			panic(err)
		}
		affected := entities.UserInfoService.ModifyInfoByID(u)
		formatter.JSON(w, http.StatusOK, affected)
	}
}

func deleteUserInfoHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if len(req.Form["uid"][0]) == 0 {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"You must have an uid"})
			return
		}
		id, err := strconv.Atoi(req.Form["uid"][0])
		if err != nil {
			panic(err)
		}
		affected := entities.UserInfoService.DeleteByID(id)
		formatter.JSON(w, http.StatusOK, affected)
	}
}

func deleteAllUserInfoHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		affected := entities.UserInfoService.DeleteAll()
		formatter.JSON(w, http.StatusOK, affected)
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
