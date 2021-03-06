package DataManagerPatient

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Patient struct {
	Patient_TC_NO        int
	Patient_Name         string
	Patient_Last_Name    string
	Patient_Age          int
	Patient_Gender       string
	Patient_Hes_Code     string
	Patient_Phone_Number int
	Patient_Mail         string
}

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

func GetPatient(patientTcNo int) (Patient, string) {
	var patientNull string
	p := Patient{}
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		defer db.Close()
	}
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS patients(Patient_TC_NO integer not null primary key, Patient_Name text, Patient_Last_Name text, Patient_Age integer, Patient_Gender text, Patient_Hes_Code text, Patient_Phone_Number integer,Patient_Mail text);
	`
	db.Exec(sqlStmt)

	rows, err := db.Query("select Patient_TC_NO, Patient_Name, Patient_Last_Name, Patient_Age, Patient_Gender, Patient_Hes_Code, Patient_Phone_Number, Patient_Mail from patients")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var Patient_TC_NO int
		var Patient_Name string
		var Patient_Last_Name string
		var Patient_Age int
		var Patient_Gender string
		var Patient_Hes_Code string
		var Patient_Phone_Number int
		var Patient_Mail string
		err = rows.Scan(&Patient_TC_NO, &Patient_Name, &Patient_Last_Name, &Patient_Age, &Patient_Gender, &Patient_Hes_Code, &Patient_Phone_Number, &Patient_Mail)
		if err != nil {
			log.Fatal(err)
		}
		if patientTcNo == Patient_TC_NO {
			p.Patient_TC_NO = Patient_TC_NO
			p.Patient_Name = Patient_Name
			p.Patient_Last_Name = Patient_Last_Name
			p.Patient_Age = Patient_Age
			p.Patient_Gender = Patient_Gender
			p.Patient_Hes_Code = Patient_Hes_Code
			p.Patient_Phone_Number = Patient_Phone_Number
			p.Patient_Mail = Patient_Mail
			patientNull = ``
			break
		} else {
			patientNull = `Patient not found`
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return p, patientNull
}

func UpdateName(patientNewName string, patientTcNo int) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		defer db.Close()
	}

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS  patients(Patient_TC_NO integer not null primary key, Patient_Name text, Patient_Last_Name text, Patient_Age integer, Patient_Gender text, Patient_Hes_Code text, Patient_Phone_Number integer,Patient_Mail text);
	`
	db.Exec(sqlStmt)

	tx, _ := db.Begin()
	smtm, _ := tx.Prepare("UPDATE patients SET Patient_Name = ? WHERE Patient_TC_NO = ?")
	smtm.Exec(patientNewName, patientTcNo)
	tx.Commit()
}

func UpdateLastName(patientNewLastName string, patientTcNo int) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		defer db.Close()
	}

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS  patients(Patient_TC_NO integer not null primary key, Patient_Name text, Patient_Last_Name text, Patient_Age integer, Patient_Gender text, Patient_Hes_Code text, Patient_Phone_Number integer,Patient_Mail text);
	`
	db.Exec(sqlStmt)

	tx, _ := db.Begin()
	smtm, _ := tx.Prepare("UPDATE patients SET Patient_Last_Name = ? WHERE Patient_TC_NO = ?")
	smtm.Exec(patientNewLastName, patientTcNo)
	tx.Commit()
}

func UpdateAge(patientAge int, patientTcNo int) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		defer db.Close()
	}

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS  patients(Patient_TC_NO integer not null primary key, Patient_Name text, Patient_Last_Name text, Patient_Age integer, Patient_Gender text, Patient_Hes_Code text, Patient_Phone_Number integer,Patient_Mail text);
	`
	db.Exec(sqlStmt)

	tx, _ := db.Begin()
	smtm, _ := tx.Prepare("UPDATE patients SET Patient_Age = ? WHERE Patient_TC_NO = ?")
	smtm.Exec(patientAge, patientTcNo)
	tx.Commit()
}

func UpdateGender(patientGender string, patientTcNo int) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		defer db.Close()
	}

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS  patients(Patient_TC_NO integer not null primary key, Patient_Name text, Patient_Last_Name text, Patient_Age integer, Patient_Gender text, Patient_Hes_Code text, Patient_Phone_Number integer,Patient_Mail text);
	`
	db.Exec(sqlStmt)

	tx, _ := db.Begin()
	smtm, _ := tx.Prepare("UPDATE patients SET Patient_Gender = ? WHERE Patient_TC_NO = ?")
	smtm.Exec(patientGender, patientTcNo)
	tx.Commit()
}

func UpdateHesCode(patientHesCode string, patientTcNo int) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		defer db.Close()
	}

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS  patients(Patient_TC_NO integer not null primary key, Patient_Name text, Patient_Last_Name text, Patient_Age integer, Patient_Gender text, Patient_Hes_Code text, Patient_Phone_Number integer,Patient_Mail text);
	`
	db.Exec(sqlStmt)

	tx, _ := db.Begin()
	smtm, _ := tx.Prepare("UPDATE patients SET Patient_Hes_Code = ? WHERE Patient_TC_NO = ?")
	smtm.Exec(patientHesCode, patientTcNo)
	tx.Commit()
}

func UpdatePhoneNumber(patientPhoneNumber int, patientTcNo int) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		defer db.Close()
	}

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS  patients(Patient_TC_NO integer not null primary key, Patient_Name text, Patient_Last_Name text, Patient_Age integer, Patient_Gender text, Patient_Hes_Code text, Patient_Phone_Number integer,Patient_Mail text);
	`
	db.Exec(sqlStmt)

	tx, _ := db.Begin()
	smtm, _ := tx.Prepare("UPDATE patients SET Patient_Phone_Number = ? WHERE Patient_TC_NO = ?")
	smtm.Exec(patientPhoneNumber, patientTcNo)
	tx.Commit()
}

func UpdateMail(patientMail string, patientTcNo int) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		defer db.Close()
	}

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS  patients(Patient_TC_NO integer not null primary key, Patient_Name text, Patient_Last_Name text, Patient_Age integer, Patient_Gender text, Patient_Hes_Code text, Patient_Phone_Number integer,Patient_Mail text);
	`
	db.Exec(sqlStmt)

	tx, _ := db.Begin()
	smtm, _ := tx.Prepare("UPDATE patients SET Patient_Mail = ? WHERE Patient_TC_NO = ?")
	smtm.Exec(patientMail, patientTcNo)
	tx.Commit()
}
