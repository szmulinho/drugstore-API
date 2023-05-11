package get

import (
	"encoding/json"
	"github.com/szmulinho/drugstore/database"
	"github.com/szmulinho/drugstore/internal/model"
	"net/http"
)

func GetAllDrugs(w http.ResponseWriter, r *http.Request) {

	result := database.DB.Find(&model.Drugs)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(model.Drugs)
}

