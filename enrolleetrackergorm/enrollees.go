package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Enrollee represents the model for an enrollee
// Default table name will be `enrollees`
type Enrollee struct {
	// gorm.Model
	EnrolleeID    uint        `json:"enrolleeId" gorm:"primary_key"`
	LastName      string      `json:"lastName"`
	FirstName     string      `json:"FirstName"`
	MiddleInitial string      `json:"middleInitial"`
	BirthDate     string      `json:"birthDate"`
	Sex           string      `json:"sex"`
	PhoneNumber   string      `json:"phoneNumber"`
	ActiveStatus  bool        `json:"activeStatus"`
	Dependents    []Dependent `json:"dependents" gorm:"foreignkey:EnrolleeID"`
}

// Dependent represents the model for an item in the enrollee
type Dependent struct {
	// gorm.Model
	DependentID   uint   `json:"dependentId" gorm:"primary_key"`
	LastName      string `json:"lastName"`
	FirstName     string `json:"FirstName"`
	MiddleInitial string `json:"middleInitial"`
	BirthDate     string `json:"birthDate"`
	Sex           string `json:"sex"`
	EnrolleeID    uint   `json:"-"`
}

var db *gorm.DB

func initDB() {
	var err error
	dataSourceName := "root:5filas@tcp(localhost:3306)/?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	// Create the database. This is a one-time step.
	// Comment out if running multiple times - You may see an error otherwise
	db.Exec("CREATE DATABASE enrolleetrackergorm")
	db.Exec("USE enrolleetrackergorm")

	// Migration to create tables for Enrollee and Dependent schema
	db.AutoMigrate(&Enrollee{}, &Dependent{})
}

func createEnrollee(w http.ResponseWriter, r *http.Request) {
	var enrollee Enrollee
	json.NewDecoder(r.Body).Decode(&enrollee)
	// Creates new enrollee by inserting records in the `enrollees` and `items` table
	db.Create(&enrollee)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(enrollee)
}

func getEnrollees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var enrollees []Enrollee
	db.Preload("Dependents").Find(&enrollees)
	json.NewEncoder(w).Encode(enrollees)
}

func getEnrollee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputEnrolleeID := params["enrolleeId"]

	var enrollee Enrollee
	db.Preload("Dependents").First(&enrollee, inputEnrolleeID)
	json.NewEncoder(w).Encode(enrollee)
}

func updateEnrollee(w http.ResponseWriter, r *http.Request) {
	var updatedEnrollee Enrollee
	json.NewDecoder(r.Body).Decode(&updatedEnrollee)
	db.Save(&updatedEnrollee)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedEnrollee)
}

func updateDependent(w http.ResponseWriter, r *http.Request) {
	var updatedDependent Dependent
	json.NewDecoder(r.Body).Decode(&updatedDependent)
	db.Save(&updatedDependent)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedDependent)
}

func deleteEnrollee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	inputEnrolleeID := params["enrolleeId"]
	// Convert `enrolleeId` string param to uint64
	id64, _ := strconv.ParseUint(inputEnrolleeID, 10, 64)
	// Convert uint64 to uint
	idToDelete := uint(id64)

	db.Where("enrollee_id = ?", idToDelete).Delete(&Dependent{})
	db.Where("enrollee_id = ?", idToDelete).Delete(&Enrollee{})
	w.WriteHeader(http.StatusNoContent)
}

func deleteDependent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	inputDependentID := params["dependentId"]
	// Convert `dependentId` string param to uint64
	id64, _ := strconv.ParseUint(inputDependentID, 10, 64)
	// Convert uint64 to uint
	idToDelete := uint(id64)

	db.Where("dependent_id = ?", idToDelete).Delete(&Dependent{})
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	router := mux.NewRouter()
	// Create
	router.HandleFunc("/enrollees", createEnrollee).Methods("POST")
	// Read
	router.HandleFunc("/enrollees/{enrolleeId}", getEnrollee).Methods("GET")
	// Read-all
	router.HandleFunc("/enrollees", getEnrollees).Methods("GET")
	// Update
	router.HandleFunc("/enrollees/{enrolleeId}", updateEnrollee).Methods("PUT")
	// Delete
	router.HandleFunc("/enrollees/{enrolleeId}", deleteEnrollee).Methods("DELETE")
	// Update
	router.HandleFunc("/dependents/{dependentId}", updateDependent).Methods("PUT")
	// Delete
	router.HandleFunc("/dependents/{dependentId}", deleteDependent).Methods("DELETE")
	// Initialize db connection
	initDB()

	log.Fatal(http.ListenAndServe(":8000", router))
}
