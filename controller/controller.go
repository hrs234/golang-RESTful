package controller

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
func ShowAllItems(w http.ResponseWriter, r *http.Request) {

	// c, err := r.Cookie("access-token")
	// if err != nil {
	// 	if err == http.ErrNoCookie {
	// 		w.WriteHeader(401)
	// 		w.Write([]byte("Token not provided"))
	// 	}
	// 	w.WriteHeader(400)
	// 	log.Panic(err)
	// }

	// storeTkn := c.Value
	// claims := &Signin.Claim{}

	// Tkn, err := jwt.ParseWithClaims(storeTkn, claims, func(token *jwt.Token) (interface{}, error) {
	// 	return Signin.MyJWT, nil
	// })
	// if err != nil {
	// 	if err == jwt.ErrSignatureInvalid {
	// 		w.WriteHeader(http.StatusUnauthorized)
	// 		w.Write([]byte("Unauthorized..."))
	// 	}
	// 	w.WriteHeader(400)
	// }

	// if !Tkn.Valid {
	// 	w.WriteHeader(401)
	// 	w.Write([]byte("Token is incorrect"))
	// }

	/////////////////////////
	var Item models.Item
	var arrItem []models.Item
	var response models.ResponseItem

	db := database.Connect()
	defer db.Close()

	queryID := r.URL.Query().Get("id")
	querySearch := r.URL.Query().Get("search")

	rows, err := db.Query("select * from item")

	if queryID != "" {
		rows, err = db.Query("select * from item where id = ?", queryID)

	} else if querySearch != "" {
		rows, err = db.Query("select * from item where item_id like '%" + querySearch + "%'")

	}

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&Item.Id, &Item.Item_id, &Item.Kategori_id); err != nil {
			log.Fatal(err.Error())

		} else {
			arrItem = append(arrItem, Item)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arrItem

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

// InsertItemData insert data into item tables
func InsertItemData(w http.ResponseWriter, r *http.Request) {
	// var arrItem []models.Item
	var response models.ResponseItem

	db := database.Connect()
	defer db.Close()

	var Item models.Item
	decode := json.NewDecoder(r.Body)
	errBody := decode.Decode(&Item)
	if errBody != nil {
		panic(errBody)
	}

	_, errBody = db.Exec("insert into item (id, item_id, kategori_id) values (?, ?, ?)", Item.Id, Item.Item_id, Item.Kategori_id)

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
func UpdateItemData(w http.ResponseWriter, r *http.Request) {
	var response models.ResponseItem

	db := database.Connect()
	defer db.Close()

	var Item models.Item
	decode := json.NewDecoder(r.Body)
	errBody := decode.Decode(&Item)
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
		_, errBody = db.Exec("UPDATE item set item_id = ?, kategori_id = ? where id = ?", Item.Item_id, Item.Kategori_id, ParamsID)

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

func DeleteItemData(w http.ResponseWriter, r *http.Request) {
	var response models.ResponseItem

	db := database.Connect()
	defer db.Close()

	ParamsID, err := getVarsID(r)

	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("ID doesnt inputed or something wrong with inputs"))
		log.Panic(err)
	} else {

		_, err = db.Exec("DELETE from item where id = ?", ParamsID)

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
