package model

import "example.com/m/v2/src/helper"

type Usr struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *Usr) Create() error {
	db := helper.ConnectDb()

	sq := "INSERT INTO usr (name, email, password) VALUES ($1, $2) RETURNING id"
	_, err := db.Exec(sq, u.Name, u.Password)
	if err != nil {
		return err
	}
	return nil
}

func (u *Usr) Update() error {
	db := helper.ConnectDb()

	sq := "UPDATE usr SET name = $1, password = $2 WHERE id = $3"
	_, err := db.Exec(sq, u.Name, u.Password, u.ID)
	if err != nil {
		return err
	}
	return nil
}

func (u *Usr) Delete() error {
	db := helper.ConnectDb()

	sq := "DELETE FROM usr WHERE id = $1"
	_, err := db.Exec(sq, u.ID)
	if err != nil {
		return err
	}
	return nil
}

func (u *Usr) Get() error {
	db := helper.ConnectDb()

	sq := "SELECT id, name, password FROM usr WHERE id = $1"
	err := db.QueryRow(sq, u.ID).Scan(&u.ID, &u.Name, &u.Password)
	if err != nil {
		return err
	}
	return nil
}

func (u *Usr) GetAll() ([]Usr, error) {

	db := helper.ConnectDb()

	rows, err := db.Query("SELECT id,name,email,password FROM usr")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// An album slice to hold data from returned rows.
	var usrs []Usr

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var usr Usr
		if err := rows.Scan(&usr.ID, &usr.Name, &usr.Email,
			&usr.Password); err != nil {
			return usrs, err
		}
		usrs = append(usrs, usr)
	}
	if err = rows.Err(); err != nil {
		return usrs, err
	}
	return usrs, nil
}
