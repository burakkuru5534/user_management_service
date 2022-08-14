package api

import (
	"encoding/json"
	"example.com/m/v2/src/helper"
	"example.com/m/v2/src/model"
	"github.com/Shyp/go-dberror"
	_ "github.com/letsencrypt/boulder/db"
	"net/http"
)

func UserCreate(w http.ResponseWriter, r *http.Request) {

	var usr model.Usr

	err := helper.BodyToJsonReq(r, &usr)
	if err != nil {
		http.Error(w, "{\"error\": \"Bad request\"}", http.StatusBadRequest)
		return
	}

	err = usr.Create()
	if err != nil {
		dberr := dberror.GetError(err)
		switch e := dberr.(type) {
		case *dberror.Error:
			if e.Code == "23505" {
				http.Error(w, "{\"error\": \"User with that email already exists\"}", http.StatusForbidden)
				return
			}
		}

		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}

	respBody := struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}{
		ID:    usr.ID,
		Name:  usr.Name,
		Email: usr.Email,
	}

	json.NewEncoder(w).Encode(respBody)

}

func UserUpdate(w http.ResponseWriter, r *http.Request) {

	var usr model.Usr

	//id := helper.StrToInt64(chi.URLParam(r, "id"))
	id := helper.StrToInt64(r.URL.Query().Get("id"))

	isExists, err := helper.CheckIfUserExists(id)
	if err != nil {
		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}

	if !isExists {
		http.Error(w, "{\"error\": \"User with that id does not exist\"}", http.StatusNotFound)
		return
	}

	err = helper.BodyToJsonReq(r, &usr)
	if err != nil {
		http.Error(w, "{\"error\": \"Bad request\"}", http.StatusBadRequest)
		return
	}

	err = usr.Update(id)
	if err != nil {
		dberr := dberror.GetError(err)
		switch e := dberr.(type) {
		case *dberror.Error:
			if e.Code == "23505" {
				http.Error(w, "{\"error\": \"User with that email already exists\"}", http.StatusForbidden)
				return
			}
		}

		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}

	respBody := struct {
		ID    int64  `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}{
		ID:    id,
		Name:  usr.Name,
		Email: usr.Email,
	}
	json.NewEncoder(w).Encode(respBody)

}

func UserDelete(w http.ResponseWriter, r *http.Request) {

	var usr model.Usr

	//id := helper.StrToInt64(chi.URLParam(r, "id"))
	id := helper.StrToInt64(r.URL.Query().Get("id"))

	isExists, err := helper.CheckIfUserExists(id)
	if err != nil {
		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}

	if !isExists {
		http.Error(w, "{\"error\": \"User with that id does not exist\"}", http.StatusNotFound)
		return
	}

	err = usr.Delete(id)
	if err != nil {
		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("ok")

}

func UserGet(w http.ResponseWriter, r *http.Request) {

	var usr model.Usr

	id := helper.StrToInt64(r.URL.Query().Get("id"))
	//id := helper.StrToInt64(chi.URLParam(r, "id"))

	isExists, err := helper.CheckIfUserExists(id)
	if err != nil {
		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}

	if !isExists {
		http.Error(w, "{\"error\": \"User with that id does not exist\"}", http.StatusNotFound)
		return
	}

	err = usr.Get(id)
	if err != nil {
		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}

	respBody := struct {
		ID    int64  `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}{
		ID:    id,
		Name:  usr.Name,
		Email: usr.Email,
	}
	json.NewEncoder(w).Encode(respBody)

}

func UserList(w http.ResponseWriter, r *http.Request) {

	var usr model.Usr

	usrList, err := usr.GetAll()
	if err != nil {
		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(usrList)

}
