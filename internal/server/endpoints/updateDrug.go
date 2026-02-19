package endpoints

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/szmulinho/drugstore/internal/model"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (h *handlers) UpdateDrug(w http.ResponseWriter, r *http.Request) {
	drugIDStr := mux.Vars(r)["id"]
	DrugID, err := strconv.ParseInt(drugIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid drug ID", http.StatusBadRequest)
		return
	}

	var updatedDrug model.Drug
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Kindly enter data with the drug name and price only in order to update", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody, &updatedDrug)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Keep a map of which fields were present in the JSON so we don't overwrite
	var payloadMap map[string]interface{}
	if err := json.Unmarshal(reqBody, &payloadMap); err != nil {
		payloadMap = map[string]interface{}{}
	}

	var existingDrug model.Drug
	result := h.db.First(&existingDrug, DrugID)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Update only fields present in payload (support lowercase and capitalized keys)
	if _, ok := payloadMap["name"]; ok {
		existingDrug.Name = updatedDrug.Name
	}
	if _, ok := payloadMap["Name"]; ok {
		existingDrug.Name = updatedDrug.Name
	}
	if _, ok := payloadMap["price"]; ok {
		existingDrug.Price = updatedDrug.Price
	}
	if _, ok := payloadMap["Price"]; ok {
		existingDrug.Price = updatedDrug.Price
	}
	if _, ok := payloadMap["image"]; ok {
		existingDrug.Image = updatedDrug.Image
	}
	if _, ok := payloadMap["Image"]; ok {
		existingDrug.Image = updatedDrug.Image
	}
	if _, ok := payloadMap["type"]; ok {
		existingDrug.Type = updatedDrug.Type
	}
	if _, ok := payloadMap["Type"]; ok {
		existingDrug.Type = updatedDrug.Type
	}
	if _, ok := payloadMap["description"]; ok {
		existingDrug.Description = updatedDrug.Description
	}
	if _, ok := payloadMap["Description"]; ok {
		existingDrug.Description = updatedDrug.Description
	}

	result = h.db.Save(&existingDrug)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(existingDrug)
}
