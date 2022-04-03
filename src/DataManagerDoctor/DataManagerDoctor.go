package DataManagerDoctor

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Doctor struct {
	Doctor_TC_NO        int
	Doctor_Password     string
	Doctor_Name         string
	Doctor_Last_Name    string
	Doctor_Age          int
	Doctor_Gender       string
	Doctor_Hes_Code     string
	Doctor_Phone_Number int
	Doctor_Mail         string
	Doctor_Expertise    string
}

func DataAdd(Doctor_TC_NO int, Doctor_Password string, Doctor_Name string, Doctor_Last_Name string, Doctor_Age int, Doctor_Gender string, Doctor_Hes_Code string, Doctor_Phone_Number int, Doctor_Mail string, Dcotor_Expertise string) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		defer db.Close()
	}

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS  doctors(Doctor_TC_NO integer not null primary key, Doctor_Password text, Doctor_Name text, Doctor_Last_Name text, Doctor_Age integer, Doctor_Gender text, Doctor_Hes_Code text, Doctor_Phone_Number integer,Doctor_Mail text, Doctor_Expertise text);
	`
	db.Exec(sqlStmt)

	tx, _ := db.Begin()

	stmt, err := tx.Prepare("insert into doctors(Doctor_TC_NO, Doctor_Password, Doctor_Name, Doctor_Last_Name, Doctor_Age, Doctor_Gender, Doctor_Hes_Code, Doctor_Phone_Number, Doctor_Mail, Doctor_Expertise) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")

	stmt.Exec(Doctor_TC_NO, Doctor_Password, Doctor_Name, Doctor_Last_Name, Doctor_Age, Doctor_Gender, Doctor_Hes_Code, Doctor_Phone_Number, Doctor_Mail, Dcotor_Expertise)

	tx.Commit()
	defer db.Close()
}

func Login(Doctor_TC_NO int, Doctor_Password string) (Doctor, string) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		defer db.Close()
	}

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS  doctors(Doctor_TC_NO integer not null primary key, Doctor_Password text, Doctor_Name text, Doctor_Last_Name text, Doctor_Age integer, Doctor_Gender text, Doctor_Hes_Code text, Doctor_Phone_Number integer,Doctor_Mail text, Doctor_Expertise text);
	`
	db.Exec(sqlStmt)

	tx, _ := db.Begin()

	rows, err := tx.Query(`select Doctor_TC_NO, Doctor_Password, Doctor_Name, Doctor_Last_Name, Doctor_Age, Doctor_Gender, Doctor_Hes_Code, Doctor_Phone_Number, Doctor_Mail, Doctor_Expertise from doctors`)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var d Doctor
	var loginControl string
	for rows.Next() {
		var doctorTcNo int
		var doctorPassword string
		var Doctor_Name string
		var Doctor_Last_Name string
		var Doctor_Age int
		var Doctor_Gender string
		var Doctor_Hes_Code string
		var Doctor_Phone_Number int
		var Doctor_Mail string
		var Doctor_Expertise string
		err = rows.Scan(&doctorTcNo, &doctorPassword, &Doctor_Name, &Doctor_Last_Name, &Doctor_Age, &Doctor_Gender, &Doctor_Hes_Code, &Doctor_Phone_Number, &Doctor_Mail, &Doctor_Expertise)
		if err != nil {
			panic(err.Error())
		}
		if doctorTcNo == Doctor_TC_NO && doctorPassword == Doctor_Password {
			fmt.Println("Login Successful")
			d = Doctor{Doctor_TC_NO, doctorPassword, Doctor_Name, Doctor_Last_Name, Doctor_Age, Doctor_Gender, Doctor_Hes_Code, Doctor_Phone_Number, Doctor_Mail, Doctor_Expertise}
			loginControl = ``
			break
		} else {
			fmt.Println("Login Failed")
			loginControl = `Login Failed`
		}
	}
	tx.Commit()
	defer db.Close()
	return d, loginControl
}
