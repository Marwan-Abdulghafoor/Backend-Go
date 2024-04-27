package main

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"murwan.net/fiephrs-backend/dao"
	"murwan.net/fiephrs-backend/utils"
)

func getPatientById(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
	var (
		request  utils.StandardIdRequest
		response utils.ResponseDTO
	)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	res, err := dao.FindProfileById(client, request.Id)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Success = true
	response.Message = "Operation Succeeded"
	response.ProfileInfo = *res
	json.NewEncoder(w).Encode(response)
}

func editProfile(w http.ResponseWriter, r *http.Request, client *mongo.Client) {

	// Decode the request body

	var (
		request  utils.EditRequest
		response utils.ResponseDTO
	)

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	println(request.Id)
	println(request.ProfileInfo.IceInstructions)
	request.ProfileInfo.Id = request.Id
	// Update the profile in the database
	err := dao.UpdateProfile(client, request.Id, &request.ProfileInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	println("Here12")
	response.Success = true
	response.Message = "Operation Succeeded"
	response.ProfileInfo = request.ProfileInfo
	json.NewEncoder(w).Encode(response)
}
