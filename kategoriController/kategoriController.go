package kategoriController

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"../database"
	"../models"
	"github.com/gorilla/mux"
)

// showAllItems showing all field table of item
func ShowAllKategori(w http.ResponseWriter, r *http.Request) {
	var Kategori models.Kategori
	var arrKategori []models.Kategori
	var response models.ResponseKategori

	db := database.Connect()
	defer db.Close()

	SQL := "select * from kategori"

	queryID := r.URL.Query().Get("id")
	querySearch := r.URL.Query().Get("search")

	if queryID != "" {

		SQL = "select * from kategori where id = '" + queryID + "'"

	} else if querySearch != "" {

		SQL = "select * from kategori where kategori like '%" + querySearch + "%'"

	}

	rows, err := db.Query(SQL)

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&Kategori.Id, &Kategori.Kategori); err != nil {
			log.Fatal(err.Error())

		} else {
			arrKategori = append(arrKategori, Kategori)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arrKategori

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

// InsertItemData insert data into item tables
func InsertKategoriData(w http.ResponseWriter, r *http.Request) {
	// var arrItem []models.Item
	var response models.ResponseKategori

	db := database.Connect()
	defer db.Close()

	var Kategori models.Kategori
	decode := json.NewDecoder(r.Body)
	errBody := decode.Decode(&Kategori)
	if errBody != nil {
		panic(errBody)
	}

	_, errBody = db.Exec("insert into kategori (id, kategori) values (?, ?)", Kategori.Id, Kategori.Kategori)

	// check error or not
	if errBody != nil {
		log.Print(errBody)
		w.WriteHeader(401)
		w.Write([]byte("Something when error"))
	} else {
		response.Status = 1
		response.Message = "Data Successfully added"
		log.Print("Data inserted to Item")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}

}

// getVarsID get an params from URLs
func getVarsID(r *http.Request) (id int, err error) {
	vars := mux.Vars(r)
	if val, ok := vars["id"]; ok {
		convertedVal, err := strconv.Atoi(val)
		if err != nil {
			return id, err
		}
		id = convertedVal
	}
	return
}

// updateItemData update an item data
func UpdateKategoriData(w http.ResponseWriter, r *http.Request) {
	var response models.ResponseKategori

	db := database.Connect()
	defer db.Close()

	var Kategori models.Kategori
	decode := json.NewDecoder(r.Body)
	errBody := decode.Decode(&Kategori)
	if errBody != nil {
		panic(errBody)
	}

	// get the params from URL
	ParamsID, errParams := getVarsID(r)

	if errParams != nil {
		w.WriteHeader(401)
		w.Write([]byte("ID not inserted or something wrong with inputs"))
		log.Panic(errParams)
	} else {
		// execute update item
		_, errBody = db.Exec("UPDATE kategori set kategori = ? where id = ?", Kategori.Kategori, ParamsID)

		// check error or not
		if errBody != nil {
			log.Print(errBody)
			w.WriteHeader(401)
			w.Write([]byte("Something when error"))
			log.Panic(errBody)
		} else {
			response.Status = 1
			response.Message = "Data Successfully updated"
			log.Print("Data updated to Item")

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}
	}

}

func DeleteKategoriData(w http.ResponseWriter, r *http.Request) {
	var response models.ResponseKategori

	db := database.Connect()
	defer db.Close()

	ParamsID, err := getVarsID(r)

	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("ID doesnt inputed or something wrong with inputs"))
		log.Panic(err)
	} else {

		_, err = db.Exec("DELETE from kategori where id = ?", ParamsID)

		if err != nil {

			w.WriteHeader(500)
			w.Write([]byte("failed to delete data"))
			log.Panic(err)

		} else {

			response.Status = 1
			response.Message = "Success Delete Data"
			log.Print("Delete data to database")

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		}
	}

}
