# Hospital API

## Dependencies

- [BodyParser](https://github.com/Periyot/BodyParser)
- [go-sqlite3](https://github.com/mattn/go-sqlite3 "go-sqlite3")

## Execution Program

```
$ cd Hospital-API/src
```

```
$ go run hospital_api.go
```

## Documentation

### Doctor

#####  POST `/doctor/register`

###### Body

- Doctor_TC_NO
- Doctor_Password
- Doctor_Name
- Doctor_Last_Name
- Doctor_Age
- Doctor_Gender
- Doctor_Hes_Code
- Doctor_Phone_Number
- Doctor_Mail
- Doctor_Expertise

#####  POST `/doctor/login`

###### Body

- Doctor_TC_NO
- Doctor_Password

> Returns Doctor information if login is successful, returns an error message if login is unsuccessful

#####  POST `/doctor/update/password`

###### Body

- Doctor_TC_NO
- Doctor_New_Password

#####  POST `/doctor/update/name`

###### Body

- Doctor_TC_NO
- Doctor_New_Name

#####  POST `/doctor/update/lastname`

###### Body

- Doctor_TC_NO
- Doctor_New_Last_Name

#####  POST `/doctor/update/age`

###### Body

- Doctor_TC_NO
- Doctor_New_Age

#####  POST `/doctor/update/gender`

###### Body

- Doctor_TC_NO
- Doctor_New_Gender
