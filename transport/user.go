package transport

import (
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/jasonbronson/crawlmonster-api/library"
	"github.com/jasonbronson/crawlmonster-api/library/log"
	"github.com/jasonbronson/crawlmonster-api/transport/models"

	"github.com/gin-gonic/gin"
)

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type UserRegister struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	UserLogin
}

func Login(g *gin.Context) {

	userParams := UserLogin{}
	if err := g.BindJSON(&userParams); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := models.GetUserByEmail(userParams.Email)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if library.CheckPasswordHash(userParams.Password, user.Password) {
		g.JSON(http.StatusOK, gin.H{"status": "loginok"})
	} else {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Wrong email/password"})
		return
	}

}

func Register(g *gin.Context) {
	register := UserRegister{}
	if err := g.BindJSON(&register); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uid, _ := uuid.NewV4()
	password, err := library.HashPassword(register.Password)
	if err != nil {
		log.Println("cannot hash user password")
	}
	user := models.User{
		ID:       uid.String(),
		Email:    register.Email,
		Password: password,
	}
	ok, err := models.CreateUser(&user)
	if err != nil {
		log.Println("cannot create user account")
		errf := fmt.Sprintf("Cannot create account error %v", err)
		g.JSON(http.StatusBadRequest, gin.H{"error": errf})
		return
	}
	if ok {
		log.Println(user.Email, " account created")
		g.JSON(http.StatusOK, gin.H{"status": "registerok"})
		return
	}
	g.JSON(http.StatusBadRequest, gin.H{"error": "Cannot create account"})

}
