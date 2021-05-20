package transport

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jasonbronson/crawlmonster-api/transport/models"
)

func Projects(g *gin.Context) {

	projects, err := models.GetProjects("1243")
	if err != nil {
		log.Println(err)
		g.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	g.JSON(http.StatusOK, projects)

}
