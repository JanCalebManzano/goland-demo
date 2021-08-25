package server

import (
	"goland-demo/api/controllers"
	"log"

	"github.com/gin-gonic/gin"
)

// NewRouter is the constructor of the router
func NewRouter() (r *gin.Engine) {
	r = gin.New()

	// MIDDLEWARES
	r.Use(logger())
	r.Use(cors())
	r.Use(jsonify())

	userCtrl, err := controllers.NewUserController()
	if err != nil {
		log.Fatalln(err)
	}
	userRtr := r.Group("/users")
	userRtr.GET("", userCtrl.GetUsers())
	userRtr.GET("/0", userCtrl.GetUser())

	todoCtrl, err := controllers.NewTodoController()
	if err != nil {
		log.Fatalln(err)
	}
	todoRtr := r.Group("/todos")
	todoRtr.GET("", todoCtrl.GetTodos())
	todoRtr.GET("/0", todoCtrl.GetTodo())

	return r
}
