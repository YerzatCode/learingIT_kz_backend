package controller

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type handler struct {
	db *sql.DB
}

func InitContoller(db *sql.DB, r *gin.Engine) {
	h := handler{db: db}

	api := r.Group("/api")
	{
		api.POST("/users/signup", h.Signup)
		api.POST("/users/login", h.Login)
	}
}
