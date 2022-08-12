package model

import (
	"example.com/m/v2/src/helper"
)

type Usr struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *Usr) Create() error {

	sq := "INSERT INTO usr (name, email, password) VALUES ($1, $2, $3) RETURNING id"
	err := helper.App.DB.QueryRow(sq, u.Name, u.Email, u.Password).Scan(&u.ID)
	if err != nil {
		return err
	}

	return nil
}

func (u *Usr) Update(id int64) error {

	sq := "UPDATE usr SET name = $1, password = $2, email = $3 WHERE id = $4"
	_, err := helper.App.DB.Exec(sq, u.Name, u.Password, u.Email, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *Usr) Delete(id int64) error {

	sq := "DELETE FROM usr WHERE id = $1"
	_, err := helper.App.DB.Exec(sq, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *Usr) Get(id int64) error {

	sq := "SELECT id, name, password, email FROM usr WHERE id = $1"
	err := helper.App.DB.QueryRow(sq, id).Scan(&u.ID, &u.Name, &u.Password, &u.Email)
	if err != nil {
		return err
	}
	return nil
}

func (u *Usr) GetAll() ([]Usr, error) {

	rows, err := helper.App.DB.Query("SELECT id,name,email FROM usr")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// An album slice to hold data from returned rows.
	var usrs []Usr

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var usr Usr
		if err := rows.Scan(&usr.ID, &usr.Name, &usr.Email); err != nil {
			return usrs, err
		}
		usrs = append(usrs, usr)
	}
	if err = rows.Err(); err != nil {
		return usrs, err
	}
	return usrs, nil
}
