package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/jguerra6/API/controller"
	"github.com/jguerra6/API/infrastructure/router"
)

var (
	httpRouter       = router.NewMuxRouter()
	leagueController = controller.NewLeagueController()
)

func homePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Homepage Endpoint Hit")
}

//TODO: Add Patch league functionality
//Create another endpoint for teams

func main() {
	port := flag.String("port", ":8081", "HTTP network address")
	flag.Parse()

	httpRouter.GET("/", homePage)
	httpRouter.GET("/leagues", leagueController.GetAllLeagues)
	httpRouter.GET("/leagues/{id}", leagueController.GetLeague)
	httpRouter.DELETE("/leagues/{id}", leagueController.DeleteLeague)
	httpRouter.POST("/leagues", leagueController.Addleague)
	httpRouter.SERVE(*port)

}
