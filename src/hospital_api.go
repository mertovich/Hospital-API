package main

import (
	"fmt"
	"hospital_api/src/Doctor"
	"log"
	"net/http"
	"os"
)

func main() {
	// doktor kayıt
	http.HandleFunc(`/doctor/register`, Doctor.DoctorRegisterHandler)
	http.HandleFunc(`/doctor/login`, Doctor.DoctorLoginHandler)
	http.HandleFunc(`/doctor/update/password`, Doctor.DoctorUpdatePasswordHandler)
	http.HandleFunc(`/doctor/update/name`, Doctor.DoctorUpdateNameHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
