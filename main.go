package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/jguerra6/API/controller"
	"github.com/jguerra6/API/infrastructure/datastore"
	"github.com/jguerra6/API/infrastructure/router"
)

var (
	httpRouter       = router.NewMuxRouter()
	db               = datastore.NewFirestoreDB()
	leagueController = controller.NewLeagueController(db, httpRouter)
	teamController   = controller.NewTeamController(db, httpRouter)
)

func homePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Homepage Endpoint Hit")
}

//TODO: Add Patch league functionality

func main() {
	port := flag.String("port", ":8081", "HTTP network address")
	flag.Parse()

	httpRouter.GET("/", homePage)
	httpRouter.GET("/leagues", leagueController.GetAllLeagues)
	//httpRouter.GET("/teams", teamController.GetAllTeams)
	httpRouter.GET("/leagues/{id}", leagueController.GetLeague)
	//httpRouter.GET("/leagues/{id}/teams", teamController.GetAllTeams)
	httpRouter.DELETE("/leagues/{id}", leagueController.DeleteLeague)
	httpRouter.POST("/leagues", leagueController.Addleague)
	httpRouter.PATCH("/leagues/{id}", leagueController.Updateleague)
	//httpRouter.POST("/teams", teamController.AddTeam)
	httpRouter.SERVE(*port)

}
