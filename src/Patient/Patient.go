package Patient

import (
	"encoding/json"
	"fmt"
	"hospital_api/src/DataManagerPatient"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/Periyot/BodyParser"
)

func PatientRegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Fatal Access-Control-Allow-Origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.URL.Path != "/patient/register" {
		http.NotFound(w, r)
		return
	}

	bodyByte, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyByte)

	maps := BodyParser.Parser(bodyString)
	mapsJson, _ := json.Marshal(maps)

	patientTcNo := maps[`Patient_TC_NO`]
	patientName := maps[`Patient_Name`]
	patientLastName := maps[`Patient_Last_Name`]
	patientAge := maps[`Patient_Age`]
	patientGender := maps[`Patient_Gender`]
	patientHesCode := maps[`Patient_Hes_Code`]
	patientPhoneNumber := maps[`Patient_Phone_Number`]
	patientMail := maps[`Patient_Mail`]

	patientMail = strings.ReplaceAll(patientMail, `%40`, `@`)

	patientTcNoInt, _ := strconv.Atoi(patientTcNo)
	patientAgeInt, _ := strconv.Atoi(patientAge)
	patientPhoneNumberInt, _ := strconv.Atoi(patientPhoneNumber)

	DataManagerPatient.DataAdd(patientTcNoInt, patientName, patientLastName, patientAgeInt, patientGender, patientHesCode, patientPhoneNumberInt, patientMail)
	_, errors := fmt.Fprint(w, string(mapsJson))
	if errors != nil {
		w.WriteHeader(http.StatusInternalServerError)

	}
}

func PatientInfoHandler(w http.ResponseWriter, r *http.Request) {
	// Fatal Access-Control-Allow-Origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.URL.Path != "/patient/info" {
		http.NotFound(w, r)
		return
	}

	bodyByte, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyByte)

	maps := BodyParser.Parser(bodyString)

	patientTcNo := maps[`Patient_TC_NO`]
	patientTcNoInt, _ := strconv.Atoi(patientTcNo)

	patient, err := DataManagerPatient.GetPatient(patientTcNoInt)
	if err != `` {
		_, errors := fmt.Fprint(w, string(err))
		if errors != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		patientJSON, _ := json.Marshal(patient)

		_, errors := fmt.Fprint(w, string(patientJSON))
		if errors != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
