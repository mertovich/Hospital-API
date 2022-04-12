package DataManagerPatient

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func DataAdd(Patient_TC_NO int, Patient_Name string, Patient_Last_Name string, Patient_Age int, Patient_Gender string, Patient_Hes_Code string, Patient_Phone_Number int, Patient_Mail string) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		defer db.Close()
	}

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS  patients(Patient_TC_NO integer not null primary key, Patient_Name text, Patient_Last_Name text, Patient_Age integer, Patient_Gender text, Patient_Hes_Code text, Patient_Phone_Number integer,Patient_Mail text);
	`
	db.Exec(sqlStmt)

	tx, _ := db.Begin()

	stmt, err := tx.Prepare("insert into patients(Patient_TC_NO, Patient_Name, Patient_Last_Name, Patient_Age, Patient_Gender, Patient_Hes_Code, Patient_Phone_Number, Patient_Mail) values(?, ?, ?, ?, ?, ?, ?, ?)")

	stmt.Exec(Patient_TC_NO, Patient_Name, Patient_Last_Name, Patient_Age, Patient_Gender, Patient_Hes_Code, Patient_Phone_Number, Patient_Mail)

	tx.Commit()
	defer db.Close()
}
