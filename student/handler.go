package student

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fiscaluno/pandorabox"
	"github.com/fiscaluno/pandorabox/db"

	"github.com/gorilla/mux"
)

// FindAll entitys
func FindAll(w http.ResponseWriter, r *http.Request) {
	entitys := GetAll()
	pandorabox.RespondWithJSON(w, http.StatusOK, entitys)
}

// FindByID find a entity by ID
func FindByID(w http.ResponseWriter, r *http.Request) {

	var msg pandorabox.Message

	msg = pandorabox.Message{
		Content: "Invalid ID, not a int",
		Status:  "ERROR",
		Body:    nil,
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		pandorabox.RespondWithJSON(w, http.StatusOK, msg)
		return
	}

	entity := GetByID(id)

	if entity.ID != 0 {
		pandorabox.RespondWithJSON(w, http.StatusOK, entity)
		return
	}

	msg = pandorabox.Message{
		Content: "Not exist this Student",
		Status:  "ERROR",
		Body:    nil,
	}
	pandorabox.RespondWithJSON(w, http.StatusOK, msg)

}

// Add a entity
func Add(w http.ResponseWriter, r *http.Request) {
	db := db.Conn()
	defer db.Close()

	var entity Entity
	var msg pandorabox.Message

	msg = pandorabox.Message{
		Content: "Invalid request payload",
		Status:  "ERROR",
		Body:    nil,
	}

	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		pandorabox.RespondWithJSON(w, http.StatusBadRequest, msg)
		return
	}

	db.Create(&entity)

	msg = pandorabox.Message{
		Content: "New Student successfully added",
		Status:  "OK",
		Body:    entity,
	}
	pandorabox.RespondWithJSON(w, http.StatusCreated, msg)

}

// DeleteByID find a entity by ID
func DeleteByID(w http.ResponseWriter, r *http.Request) {
	db := db.Conn()
	defer db.Close()

	var entity Entity
	var msg pandorabox.Message

	msg = pandorabox.Message{
		Content: "Invalid ID, not a int",
		Status:  "ERROR",
		Body:    nil,
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		pandorabox.RespondWithJSON(w, http.StatusOK, msg)
		return
	}

	db.Find(&entity, id)
	if entity.ID != 0 {
		db.Delete(&entity)
		msg = pandorabox.Message{
			Content: "Deleted Student successfully",
			Status:  "OK",
			Body:    entity,
		}

		pandorabox.RespondWithJSON(w, http.StatusAccepted, msg)
		return
	}

	msg = pandorabox.Message{
		Content: "Not exist this Student",
		Status:  "ERROR",
		Body:    nil,
	}
	pandorabox.RespondWithJSON(w, http.StatusOK, msg)

}

// UpdateByID find a entity by ID
func UpdateByID(w http.ResponseWriter, r *http.Request) {
	db := db.Conn()
	defer db.Close()

	var entity Entity
	var msg pandorabox.Message

	msg = pandorabox.Message{
		Content: "Invalid ID, not a int",
		Status:  "ERROR",
		Body:    nil,
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		pandorabox.RespondWithJSON(w, http.StatusOK, msg)
		return
	}

	msg = pandorabox.Message{
		Content: "Invalid request payload",
		Status:  "ERROR",
		Body:    nil,
	}

	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		pandorabox.RespondWithJSON(w, http.StatusBadRequest, msg)
		return
	}

	if id >= 0 {
		entity.ID = uint(id)
		db.Save(&entity)

		msg = pandorabox.Message{
			Content: "Update Student successfully ",
			Status:  "OK",
			Body:    entity,
		}
		pandorabox.RespondWithJSON(w, http.StatusAccepted, msg)
		return
	}

	msg = pandorabox.Message{
		Content: "Not exist this Student",
		Status:  "ERROR",
		Body:    nil,
	}
	pandorabox.RespondWithJSON(w, http.StatusOK, msg)

}

// FindByFacebookID find a entity by FacebookID
func FindByFacebookID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	FacebookID := vars["id"]
	entitys := GetByQuery("facebook_id = ?", FacebookID)

	if len(entitys) >= 0 {
		pandorabox.RespondWithJSON(w, http.StatusOK, entitys)
		return
	}

	msg := pandorabox.Message{
		Content: "Not exist this Student",
		Status:  "ERROR",
		Body:    nil,
	}
	pandorabox.RespondWithJSON(w, http.StatusOK, msg)

}
