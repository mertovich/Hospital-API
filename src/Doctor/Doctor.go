package Doctor

import (
	"encoding/json"
	"fmt"
	"hospital_api/src/DataManagerDoctor"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/Periyot/BodyParser"
)

func DoctorRegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Fatal Access-Control-Allow-Origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.URL.Path != "/doctor/register" {
		http.NotFound(w, r)
		return
	}

	bodyByte, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyByte)

	maps := BodyParser.Parser(bodyString)
	mapsJson, _ := json.Marshal(maps)

	doctorTcNo := maps[`Doctor_TC_NO`]
	doctorPassword := maps[`Doctor_Password`]
	doctorName := maps[`Doctor_Name`]
	doctorLastName := maps[`Doctor_Last_Name`]
	doctorAge := maps[`Doctor_Age`]
	doctorGender := maps[`Doctor_Gender`]
	doctorHesCode := maps[`Doctor_Hes_Code`]
	doctorPhoneNumber := maps[`Doctor_Phone_Number`]
	doctorMail := maps[`Doctor_Mail`]
	doctorExpertise := maps[`Doctor_Expertise`]

	doctorMail = strings.ReplaceAll(doctorMail, `%40`, `@`)
	doctorExpertise = strings.ReplaceAll(doctorExpertise, `%20`, ` `)

	doctorTcNoInt, _ := strconv.Atoi(doctorTcNo)
	doctorAgeInt, _ := strconv.Atoi(doctorAge)
	doctorPhoneNumberInt, _ := strconv.Atoi(doctorPhoneNumber)

	DataManagerDoctor.DataAdd(doctorTcNoInt, doctorPassword, doctorName, doctorLastName, doctorAgeInt, doctorGender, doctorHesCode, doctorPhoneNumberInt, doctorMail, doctorExpertise)
	_, errors := fmt.Fprint(w, string(mapsJson))
	if errors != nil {
		w.WriteHeader(http.StatusInternalServerError)

	}
}

func DoctorLoginHandler(w http.ResponseWriter, r *http.Request) {
	// Fatal Access-Control-Allow-Origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.URL.Path != "/doctor/login" {
		http.NotFound(w, r)
		return
	}
	bodyByte, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyByte)

	maps := BodyParser.Parser(bodyString)

	doctorTcNo := maps[`Doctor_TC_NO`]
	doctorPassword := maps[`Doctor_Password`]
	doctorTcNoInt, _ := strconv.Atoi(doctorTcNo)

	user, err := DataManagerDoctor.Login(doctorTcNoInt, doctorPassword)
	if err == `Login Failed` {
		_, errors := fmt.Fprint(w, err)
		if errors != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		userJSON, _ := json.Marshal(user)
		_, errors := fmt.Fprint(w, string(userJSON))
		if errors != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

}

func DoctorUpdatePasswordHandler(w http.ResponseWriter, r *http.Request) {
	// Fatal Access-Control-Allow-Origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.URL.Path != "/doctor/update/password" {
		http.NotFound(w, r)
		return
	}
	bodyByte, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyByte)

	maps := BodyParser.Parser(bodyString)

	doctorTcNo := maps[`Doctor_TC_NO`]
	doctorNewPassword := maps[`Doctor_New_Password`]
	doctorTcNoInt, _ := strconv.Atoi(doctorTcNo)
	DataManagerDoctor.UpdatePassword(doctorNewPassword, doctorTcNoInt)
	_, errors := fmt.Fprint(w, `success`)
	if errors != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func DoctorUpdateNameHandler(w http.ResponseWriter, r *http.Request) {
	// Fatal Access-Control-Allow-Origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.URL.Path != "/doctor/update/name" {
		http.NotFound(w, r)
		return
	}
	bodyByte, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyByte)

	maps := BodyParser.Parser(bodyString)

	doctorTcNo := maps[`Doctor_TC_NO`]
	doctorNewName := maps[`Doctor_New_Name`]
	doctorTcNoInt, _ := strconv.Atoi(doctorTcNo)
	DataManagerDoctor.UpdateName(doctorNewName, doctorTcNoInt)
	_, errors := fmt.Fprint(w, `success`)
	if errors != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func DoctorUpdateLastNameHandler(w http.ResponseWriter, r *http.Request) {
	// Fatal Access-Control-Allow-Origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.URL.Path != "/doctor/update/lastname" {
		http.NotFound(w, r)
		return
	}
	bodyByte, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyByte)

	maps := BodyParser.Parser(bodyString)

	doctorTcNo := maps[`Doctor_TC_NO`]
	doctorNewLastName := maps[`Doctor_New_Last_Name`]
	doctorTcNoInt, _ := strconv.Atoi(doctorTcNo)
	DataManagerDoctor.UpdateLastName(doctorNewLastName, doctorTcNoInt)
	_, errors := fmt.Fprint(w, `success`)
	if errors != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func DoctorUpdateAgeHandler(w http.ResponseWriter, r *http.Request) {
	// Fatal Access-Control-Allow-Origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.URL.Path != "/doctor/update/age" {
		http.NotFound(w, r)
		return
	}
	bodyByte, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyByte)

	maps := BodyParser.Parser(bodyString)

	doctorTcNo := maps[`Doctor_TC_NO`]
	doctorNewAge := maps[`Doctor_New_Age`]
	doctorTcNoInt, _ := strconv.Atoi(doctorTcNo)
	doctorNewAgeInt, _ := strconv.Atoi(doctorNewAge)
	DataManagerDoctor.UpdateAge(doctorNewAgeInt, doctorTcNoInt)
	_, errors := fmt.Fprint(w, `success`)
	if errors != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func DoctorUpdateGenderHandler(w http.ResponseWriter, r *http.Request) {
	// Fatal Access-Control-Allow-Origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.URL.Path != "/doctor/update/gender" {
		http.NotFound(w, r)
		return
	}
	bodyByte, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyByte)

	maps := BodyParser.Parser(bodyString)

	doctorTcNo := maps[`Doctor_TC_NO`]
	doctorNewGender := maps[`Doctor_New_Gender`]
	doctorTcNoInt, _ := strconv.Atoi(doctorTcNo)

	DataManagerDoctor.UpdateGender(doctorNewGender, doctorTcNoInt)
	_, errors := fmt.Fprint(w, `success`)
	if errors != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func DoctorUpdateHesCodeHandler(w http.ResponseWriter, r *http.Request) {
	// Fatal Access-Control-Allow-Origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.URL.Path != "/doctor/update/hescode" {
		http.NotFound(w, r)
		return
	}
	bodyByte, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyByte)

	maps := BodyParser.Parser(bodyString)

	doctorTcNo := maps[`Doctor_TC_NO`]
	doctorNewHesCode := maps[`Doctor_New_Hes_Code`]
	doctorTcNoInt, _ := strconv.Atoi(doctorTcNo)

	DataManagerDoctor.UpdateHesCode(doctorNewHesCode, doctorTcNoInt)
	_, errors := fmt.Fprint(w, `success`)
	if errors != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func DoctorUpdatePhoneNumberHandler(w http.ResponseWriter, r *http.Request) {
	// Fatal Access-Control-Allow-Origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.URL.Path != "/doctor/update/phone" {
		http.NotFound(w, r)
		return
	}
	bodyByte, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyByte)

	maps := BodyParser.Parser(bodyString)

	doctorTcNo := maps[`Doctor_TC_NO`]
	doctorPhoneNumber := maps[`Doctor_New_Phone_Number`]
	doctorTcNoInt, _ := strconv.Atoi(doctorTcNo)
	doctorPhoneNumberInt, _ := strconv.Atoi(doctorPhoneNumber)

	DataManagerDoctor.UpdatePhoneNumber(doctorPhoneNumberInt, doctorTcNoInt)
	_, errors := fmt.Fprint(w, `success`)
	if errors != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func DoctorUpdateMailHandler(w http.ResponseWriter, r *http.Request) {
	// Fatal Access-Control-Allow-Origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.URL.Path != "/doctor/update/mail" {
		http.NotFound(w, r)
		return
	}
	bodyByte, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyByte)

	maps := BodyParser.Parser(bodyString)

	doctorTcNo := maps[`Doctor_TC_NO`]
	doctorNewMail := maps[`Doctor_New_Mail`]
	doctorNewMail = strings.ReplaceAll(doctorNewMail, `%40`, `@`)
	doctorTcNoInt, _ := strconv.Atoi(doctorTcNo)

	DataManagerDoctor.UpdateMail(doctorNewMail, doctorTcNoInt)
	_, errors := fmt.Fprint(w, `success`)
	if errors != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
