package main

import (
	"bytes"
	"errors"
	"example.com/m/v2/src/api"
	"example.com/m/v2/src/helper"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestList(t *testing.T) {

	conInfo := helper.PgConnectionInfo{
		Host:     "127.0.0.1",
		Port:     5432,
		Database: "rollic",
		Username: "postgres",
		Password: "tayitkan",
		SSLMode:  "disable",
	}

	db, err := helper.NewPgSqlxDbHandle(conInfo, 10)
	if err != nil {
		errors.New("create db handle error.")
	}
	err = db.Ping()
	if err != nil {
		errors.New("ping db error.")
	}

	// Create Appplication Service
	err = helper.InitApp(db)
	if err != nil {
		errors.New("init app error.")
	}

	req, err := http.NewRequest("GET", "/api/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.UserList)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		if status == http.StatusBadRequest {
			expected := `{"error": "Bad request"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusForbidden {
			expected := `{"error": "User with that email already exists"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusInternalServerError {
			expected := `{"error": "Internal server error"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusNotFound {
			expected := `{"error": "User with that id does not exist"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	} else {
		expected := `[{"id":35,"name":"burak2","email":"testemail6@gmail.com","password":""},{"id":36,"name":"burak2","email":"testemail7@gmail.com","password":""},{"id":34,"name":"burak3","email":"testemail77@gmail.com","password":""}]
`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}
	// Check the response body is what we expect.

}

func TestCreate(t *testing.T) {

	conInfo := helper.PgConnectionInfo{
		Host:     "127.0.0.1",
		Port:     5432,
		Database: "rollic",
		Username: "postgres",
		Password: "tayitkan",
		SSLMode:  "disable",
	}

	db, err := helper.NewPgSqlxDbHandle(conInfo, 10)
	if err != nil {
		errors.New("create db handle error.")
	}
	err = db.Ping()
	if err != nil {
		errors.New("ping db error.")
	}

	// Create Appplication Service
	err = helper.InitApp(db)
	if err != nil {
		errors.New("init app error.")
	}

	var jsonStr = []byte(`{"name":"burak2","email":"testemail7@gmail.com","password":"testbrk"}`)

	req, err := http.NewRequest("POST", "/api/users", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.UserCreate)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		if status == http.StatusBadRequest {
			expected := `{"error": "Bad request"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusForbidden {
			expected := `{"error": "User with that email already exists"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusInternalServerError {
			expected := `{"error": "Internal server error"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusNotFound {
			expected := `{"error": "User with that id does not exist"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	} else {
		var id int64

		err = db.Get(&id, "SELECT id from usr order by id desc limit 1")
		if err != nil {
			errors.New("get id error.")
		}

		expected := fmt.Sprintf(`{"id":%d,"name":"burak2","email":"testemail7@gmail.com"}
`, id)
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}

}

func TestGet(t *testing.T) {

	conInfo := helper.PgConnectionInfo{
		Host:     "127.0.0.1",
		Port:     5432,
		Database: "rollic",
		Username: "postgres",
		Password: "tayitkan",
		SSLMode:  "disable",
	}

	db, err := helper.NewPgSqlxDbHandle(conInfo, 10)
	if err != nil {
		errors.New("create db handle error.")
	}
	err = db.Ping()
	if err != nil {
		errors.New("ping db error.")
	}

	// Create Appplication Service
	err = helper.InitApp(db)
	if err != nil {
		errors.New("init app error.")
	}

	req, err := http.NewRequest("GET", "/api/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "22")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.UserGet)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		if status == http.StatusBadRequest {
			expected := `{"error": "Bad request"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusForbidden {
			expected := `{"error": "User with that email already exists"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusInternalServerError {
			expected := `{"error": "Internal server error"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusNotFound {
			expected := `{"error": "User with that id does not exist"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	} else {
		expected := `{"id":22,"name":"burak","email":"testemail@gmail.com"}
`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}
	// Check the response body is what we expect.

}

func TestDelete(t *testing.T) {

	conInfo := helper.PgConnectionInfo{
		Host:     "127.0.0.1",
		Port:     5432,
		Database: "rollic",
		Username: "postgres",
		Password: "tayitkan",
		SSLMode:  "disable",
	}

	db, err := helper.NewPgSqlxDbHandle(conInfo, 10)
	if err != nil {
		errors.New("create db handle error.")
	}
	err = db.Ping()
	if err != nil {
		errors.New("ping db error.")
	}

	// Create Appplication Service
	err = helper.InitApp(db)
	if err != nil {
		errors.New("init app error.")
	}

	req, err := http.NewRequest("DELETE", "/api/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "33")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.UserDelete)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		if status == http.StatusBadRequest {
			expected := `{"error": "Bad request"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusForbidden {
			expected := `{"error": "User with that email already exists"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusInternalServerError {
			expected := `{"error": "Internal server error"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusNotFound {
			expected := `{"error": "User with that id does not exist"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

	} else {
		// Check the response body is what we expect.
		expected := `"ok"
`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}

}

func TestUpdate(t *testing.T) {

	conInfo := helper.PgConnectionInfo{
		Host:     "127.0.0.1",
		Port:     5432,
		Database: "rollic",
		Username: "postgres",
		Password: "tayitkan",
		SSLMode:  "disable",
	}

	db, err := helper.NewPgSqlxDbHandle(conInfo, 10)
	if err != nil {
		errors.New("create db handle error.")
	}
	err = db.Ping()
	if err != nil {
		errors.New("ping db error.")
	}

	// Create Appplication Service
	err = helper.InitApp(db)
	if err != nil {
		errors.New("init app error.")
	}

	var jsonStr = []byte(`{"name":"burak3","email":"testemail77@gmail.com","password":"testbrk"}`)

	req, err := http.NewRequest("PATCH", "/api/users", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "34")
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.UserUpdate)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		if status == http.StatusBadRequest {
			expected := `{"error": "Bad request"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusForbidden {
			expected := `{"error": "User with that email already exists"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusInternalServerError {
			expected := `{"error": "Internal server error"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusNotFound {
			expected := `{"error": "User with that id does not exist"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	} else {
		expected := fmt.Sprintf(`{"id":34,"name":"burak3","email":"testemail77@gmail.com"}
`)
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}

}
