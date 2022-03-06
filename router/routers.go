package router

import (
	"github.com/gin-gonic/gin"
	"userapi/controllers"
)

type Route struct {
	Method  string
	Pattern string
	Handler gin.HandlerFunc
}

var routes []Route

func init() {
	route("POST", "/v1/user/create", controllers.CreateUser)
	route("POST", "/v1/user/delete", controllers.DeleteUser)
	route("POST", "/v1/user/pwd/change", controllers.ChangePwdByUser)
	route("GET", "/v1/user/login", controllers.Login)
}

func NewRoute() *gin.Engine {
	r := gin.Default()
	for _, route := range routes {
		switch route.Method {
		case "POST":
			r.POST(route.Pattern, route.Handler)
		case "GET":
			r.GET(route.Pattern, route.Handler)
		}
	}
	return r
}

func route(method, pattern string, handler gin.HandlerFunc) {
	routes = append(routes, Route{method, pattern, handler})
}
