package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/jguerra6/API/errorsHandler"

	"github.com/jguerra6/API/infrastructure/datastore"
	"github.com/jguerra6/API/infrastructure/router"
)

type teamController struct {
	router.Router
	datastore.Database
}

//TeamController will create an interface to control all the League Operations
type TeamController interface {
	GetAllTeams(writer http.ResponseWriter, request *http.Request)
	AddTeam(writer http.ResponseWriter, request *http.Request)
	GetTeam(writer http.ResponseWriter, request *http.Request)
	DeleteTeam(writer http.ResponseWriter, request *http.Request)
}

//NewTeamController returns a League Controller to handle the League Operations
func NewTeamController(db datastore.Database, router router.Router) TeamController {

	return &teamController{router, db}
}

func (lt *teamController) GetAllTeams(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	//Get the id from the request url
	vars := lt.GetVarsFromRequest(request)
	id := vars["id"]

	fmt.Println(id)
	teams, err := lt.GetAll("teams")

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(errorsHandler.ServiceError{Message: "Error getting the leagues"})
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(teams)
}

func validateTeam(team map[string]interface{}) error {

	if team == nil {
		err := errors.New("The team is empty")
		return err
	}

	if team["name"] == "" || team["name"] == nil {
		err := errors.New("The team name can't be empty")
		return err
	}

	if team["country"] == "" || team["country"] == nil {
		err := errors.New("The team country can't be empty")
		return err
	}

	return nil

}

func (lt *teamController) AddTeam(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	tmpTeam := make(map[string]interface{})

	err := json.NewDecoder(request.Body).Decode(&tmpTeam)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(errorsHandler.ServiceError{Message: "Error adding the league"})
		log.Println("Failed decoding item: ", err)
		return
	}
	team := map[string]interface{}{
		"name":    tmpTeam["name"],
		"country": tmpTeam["country"],
	}

	err1 := validateTeam(team)
	if err1 != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(errorsHandler.ServiceError{Message: err1.Error()})
		return
	}

	result, err2 := lt.AddItem("teams", team)

	if err2 != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(errorsHandler.ServiceError{Message: "Error saving the team"})
		return
	}

	writer.WriteHeader(http.StatusOK)

	json.NewEncoder(writer).Encode(result)
}

func (lt *teamController) GetTeam(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	//Get the id from the request url
	vars := lt.GetVarsFromRequest(request)
	id := vars["id"]

	league, err := lt.GetItemByID("leagues", id)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(errorsHandler.ServiceError{Message: "League not found"})
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(league)
}

func (lt *teamController) DeleteTeam(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	//Get the id from the request url
	vars := lt.GetVarsFromRequest(request)
	id := vars["id"]

	err := lt.DeleteItem("leagues", id)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(errorsHandler.ServiceError{Message: "Error getting the leagues"})
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(errorsHandler.ServiceError{Message: "Succesfully deleted league"})
}
