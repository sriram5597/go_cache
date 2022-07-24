package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	goCache "github.com/sriram5597/go_cache/src/cache"
)

func GetKey(w http.ResponseWriter, r *http.Request) {
	parameter := r.URL.Query().Get("key")

	key, err := strconv.Atoi(parameter)

	if err != nil {
		errResponse := errorResponse{Message: err.Error()}
		SendResponse(w, http.StatusBadRequest, "Bad Request", errResponse)
		return
	}
	log.Println("GET")
	result := goCache.GetKey(key)
	response := getKeyResponse{Value: result}
	SendResponse(w, 200, "Success", response)
}


func SetKey(w http.ResponseWriter, r *http.Request) {
	var payload setKeyPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	
	if err != nil {
		errResponse := errorResponse{Message: err.Error()}
		SendResponse(w, http.StatusBadRequest, "Bad Request", errResponse)
		return
	}
	log.Println("SET")

	error := goCache.SetKey(payload.Key, payload.Value, payload.Expiry)
	if error != nil {
		SendResponse(w, 500, "Something went wrong", errorResponse{Message: error.Error()})
		return
	}
	response := successResponse{Message: "Key Cached"}
	SendResponse(w, 200, "success", response)
}
