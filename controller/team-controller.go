package controller

/*
type controller struct{}

var (
	db = datastore.NewFirestoreDB()
)

//LeagueController will create an interface to control all the League Operations
type TeamController interface {
	GetAllTeams(writer http.ResponseWriter, request *http.Request)
	AddTeam(writer http.ResponseWriter, request *http.Request)
	GetTeam(writer http.ResponseWriter, request *http.Request)
	DeleteTeam(writer http.ResponseWriter, request *http.Request)
}

//NewLeagueController returns a League Controller to handle the League Operations
func NewTeamController() TeamController {
	return &controller{}
}

func (*controller) GetAllTeams(writer http.ResponseWriter, request *http.Request) {
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

func validateTeam(league map[string]interface{}) error {

	if league == nil {
		err := errors.New("The league is empty")
		return err
	}

	if league["name"] == "" || league["name"] == nil {
		err := errors.New("The league name can't be empty")
		return err
	}
	return nil
}

func (*controller) AddTeam(writer http.ResponseWriter, request *http.Request) {
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

	err1 := validateTeam(league)
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

func (*controller) GetTeam(writer http.ResponseWriter, request *http.Request, string id) {
	writer.Header().Set("Content-Type", "application/json")

	league, err := db.GetItemByID("leagues", id)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(errorsHandler.ServiceError{Message: "League not found"})
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(league)
}

func (*controller) DeleteTeam(writer http.ResponseWriter, request *http.Request) {
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
*/
