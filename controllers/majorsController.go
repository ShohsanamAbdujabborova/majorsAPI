package controllers

import (
	"API/models"
	"encoding/json"
	"net/http"
	"os"
)

func MajorsController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case ("GET"):
		getAllMajors(w, r)
	}
}
func getAllMajors(w http.ResponseWriter, r *http.Request) {
	var data []models.Major
	Bytes,_ := os.ReadFile("data/majors.json")
	json.Unmarshal(Bytes, &data)
	json.NewEncoder(w).Encode(data)
}
