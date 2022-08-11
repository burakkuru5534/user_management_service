package api

import (
	"example.com/m/v2/src/helper"
	"example.com/m/v2/src/model"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func UserCreate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var usr model.Usr

		err := helper.BodyToJsonReq(r, &usr)
		if err != nil {
			http.Error(w, "body to json request error.", http.StatusBadRequest)
			return
		}

		err = usr.Create()
		if err != nil {
			http.Error(w, "create user error.", http.StatusInternalServerError)
			return
		}

	})
}

func UserUpdate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		id := helper.StrToInt64(chi.URLParam(r, "id"))

	})
}

func UserGet() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Get User"))
	})
}

func UserList() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("List Users"))
	})
}

func UserDelete() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Delete User"))
	})
}
