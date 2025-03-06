package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/nikagar4epm/go_api/api"
	"github.com/nikagar4epm/go_api/internal/tools"
	log "github.com/sirupsen/logrus"
)

func handleClientProfile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetClientProfile(w, r)
	case http.MethodPatch:
		UpdateClientProfile(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetClientProfile(w http.ResponseWriter, r *http.Request) {
	var params = api.ClientProfileParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()

	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var clientProfile *tools.UserDetails
	clientProfile = (*database).GetUserProfile(params.Username)
	if clientProfile == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.ClientProfileResponse{
		Profile: clientProfile,
		Code:    http.StatusOK,
	}

	// var respuesta = api.ClientProfileResponse{
	// 	Profile: *tools.ClientProfile{
	// 		Email: "Something",
	// 		Id:    "Something",
	// 		Name:  "Something",
	// 		Token: "Something",
	// 		// Email: (*tokenDetails).Email,
	// 		// Id: (*tokenDetails).Id,
	// 		// Name: (*tokenDetails).Name,
	// 		// Token: (*tokenDetails).Token,
	// 	},
	// 	Code: http.StatusOK,
	// }

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}

func UpdateClientProfile(w http.ResponseWriter, r *http.Request) {
	var username = r.URL.Query().Get("username")
	var err error

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()

	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	clientProfile := (*database).GetUserProfile(username)
	if username == "" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	var payloadData *tools.UserDetails
	if err := json.NewDecoder(r.Body).Decode(&payloadData); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	clientProfile.Email = payloadData.Email
	clientProfile.Name = payloadData.Name
	(*database).SetUserProfile(clientProfile.Id, *clientProfile)

	w.WriteHeader(http.StatusOK)

	var response = api.ClientProfileResponse{
		Profile: clientProfile,
		Code:    http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
