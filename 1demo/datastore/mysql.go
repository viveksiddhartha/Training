package datastore

import (
	"log"

	"Training/1demo/model"
)

func InsertDB(usr *model.User) error {

	m := DBCONN()

	tx, err := m.Begin()
	if err != err {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("INSERT into user(UserId,FirstName,LastName) value(?,?,?)")
	if err != err {
		log.Fatal(err)
	}

	_, err = stmt.Exec(usr.UserID, usr.FirstName, usr.LastName)
	if err != err {
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != err {
		log.Fatal(err)
	}

	return nil

}

func GetDBUsers(usr *model.User) ([]model.User, error) {
	m := DBCONN()

	GetUserResult := make([]model.User, 0)

	stmt, err := m.Prepare("SELECT UserId,FirstName,LastName FROM USER where UserID=?")
	if err != err {
		log.Fatal(err)
	}

	row, err := stmt.Query(usr.UserID)
	if err != err {
		log.Fatal(err)
	}

	for row.Next() {
		u := model.User{}

		err := row.Scan(&u.UserID, &u.FirstName, &u.LastName)
		if err != err {
			log.Fatal(err)
		}
		GetUserResult = append(GetUserResult, u)
	}

	return GetUserResult, nil

}

func UpdateDB(usr *model.User) error {
	m := DBCONN()

	tx, err := m.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("UPDATE user SET FirstName=? ,LastName=? WHERE UserID=?")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(usr.FirstName, usr.LastName, usr.UserID)
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func DeleteDBUsers(usr *model.User) error {
	m := DBCONN()

	tx, err := m.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("DELETE from USER WHERE UserID=?")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(usr.UserID)
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
