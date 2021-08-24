package server

import "github.com/gin-gonic/gin"

// NewRouter is the constructor of the router
func NewRouter() (r *gin.Engine) {
	r = gin.New()

	// MIDDLEWARES
	r.Use(logger())
	r.Use(cors())
	r.Use(jsonify())

	return r
}
