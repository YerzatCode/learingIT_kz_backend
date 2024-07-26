package controller

import (
	"net/http"

	"github.com/YerzatCode/learing_kz_backend/pkg/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (h handler) Signup(c *gin.Context) {
	body := model.UserModel{}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect data!",
		})
		return
	}

	if body.Name == "" || body.Email == "" || body.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect data!",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to hash password!",
		})
		return
	}

	result, err := h.db.Exec("insert into User (name, email, password) values ($1, $2, $3)",
		body.Name, body.Email, hash)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create user!",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created!",
		"data":    result,
	})

}

func (h handler) Login(c *gin.Context) {}
