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

	sq := "INSERT INTO usr (name, email, password) VALUES ($1, $2, $3) RETURNING id"
	err := db.QueryRow(sq, u.Name, u.Email, u.Password).Scan(&u.ID)
	if err != nil {
		return err
	}
	defer db.Close()

	return nil
}

func (u *Usr) Update(id int64) error {
	db := helper.ConnectDb()

	sq := "UPDATE usr SET name = $1, password = $2, email = $3 WHERE id = $4"
	_, err := db.Exec(sq, u.Name, u.Password, u.Email, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *Usr) Delete(id int64) error {
	db := helper.ConnectDb()

	sq := "DELETE FROM usr WHERE id = $1"
	_, err := db.Exec(sq, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *Usr) Get(id int64) error {
	db := helper.ConnectDb()

	sq := "SELECT id, name, password, email FROM usr WHERE id = $1"
	err := db.QueryRow(sq, id).Scan(&u.ID, &u.Name, &u.Password, &u.Email)
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
