package controller

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"reflect"

	"github.com/jguerra6/API/errorsHandler"
	"github.com/jguerra6/API/infrastructure/datastore"
	"github.com/jguerra6/API/infrastructure/router"
)

type controller struct{}

var (
	httpRouter = router.NewMuxRouter()
	db         = datastore.NewFirestoreDB()
)

//LeagueController will create an interface to control all the League Operations
type LeagueController interface {
	GetAllLeagues(writer http.ResponseWriter, request *http.Request)
	Addleague(writer http.ResponseWriter, request *http.Request)
	GetLeague(writer http.ResponseWriter, request *http.Request)
	DeleteLeague(writer http.ResponseWriter, request *http.Request)
}

//NewLeagueController returns a League Controller to handle the League Operations
func NewLeagueController() LeagueController {
	return &controller{}
}

func (*controller) GetAllLeagues(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	leagues, err := db.GetAll("leagues")

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(errorsHandler.ServiceError{Message: "Error getting the leagues"})
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(leagues)
}

func validateLeague(league map[string]interface{}) error {

	if league == nil {
		err := errors.New("The league is empty")
		return err
	}

	if league["name"] == "" || league["name"] == nil {
		err := errors.New("The league name can't be empty")
		return err
	}

	//Validate that it's a float parse it to int
	if reflect.TypeOf(league["current_season_id"]) != reflect.TypeOf(1.0) {
		err := errors.New("The current_season_id must be an integer")
		return err
	}
	league["current_season_id"] = int(league["current_season_id"].(float64))
	return nil

}

func (*controller) Addleague(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	tmpLeague := make(map[string]interface{})

	err := json.NewDecoder(request.Body).Decode(&tmpLeague)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(errorsHandler.ServiceError{Message: "Error adding the league"})
		log.Println("Failed decoding item: ", err)
		return
	}
	league := map[string]interface{}{
		"name":              tmpLeague["name"],
		"country":           tmpLeague["country"],
		"current_season_id": tmpLeague["current_season_id"],
	}

	err1 := validateLeague(league)
	if err1 != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(errorsHandler.ServiceError{Message: err1.Error()})
		return
	}

	result, err2 := db.AddItem("leagues", league)

	if err2 != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(errorsHandler.ServiceError{Message: "Error saving the league"})
		return
	}

	writer.WriteHeader(http.StatusOK)

	json.NewEncoder(writer).Encode(result)
}

func (*controller) GetLeague(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	id := httpRouter.GetIDFromRequest(request)
	league, err := db.GetItemByID("leagues", id)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(errorsHandler.ServiceError{Message: "League not found"})
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(league)
}

func (*controller) DeleteLeague(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	id := httpRouter.GetIDFromRequest(request)
	err := db.DeleteItem("leagues", id)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(errorsHandler.ServiceError{Message: "Error getting the leagues"})
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(errorsHandler.ServiceError{Message: "Succesfully deleted league"})
}
