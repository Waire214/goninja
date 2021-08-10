package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Waire214/goninja/models"
	"github.com/go-chi/chi/v5"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	paramID := chi.URLParam(r, "productbyid")
	productdetail := models.GetSingleProduct(paramID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(productdetail)
}
