package app

import (
	"github.com/alokgit007/FIRSTRESTAPIGO/controllers"
)


func mapUrls() {
	  router.GET(relativePath: "/ping", controllers.Ping)
	  
	  router.GET(relativePath: "/users/:user_id", controllers.GetUser)
	  router.GET(relativePath: "/users/:search", controllers.SearchUser)
	  router.POST(relativePath: "/users", controllers.CreateUser)
}
,