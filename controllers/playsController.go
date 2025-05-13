package controllers

import (
	"github.com/tomatoCoderq/KachProxyAPI/services"

	"github.com/gin-gonic/gin"
	"net/http"
)

type PlaysController struct {
	playsService *services.PlaysService
}

func NewPlaysController(playsService *services.PlaysService) *PlaysController {
	return &PlaysController {
		playsService: playsService,
	}
}

func (pc *PlaysController) GetAllPlays (ctx *gin.Context) {
	// params := ctx.Request.URL.Query()
	scene := ctx.Query("scene")
	author := ctx.Query("author")
	name := ctx.Query("name")
	age_rating := ctx.Query("age_rating")
	month := ctx.Query("month")

	plays, err := pc.playsService.GetAllPlays(scene, author, name, age_rating, month)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch plays"})
		return
	}

	if len(plays) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No plays found"})
		return
	}

	ctx.JSON(http.StatusOK, plays)		
}

func (pc *PlaysController) GetPlayById (ctx *gin.Context) {
	id := ctx.Param("id")
	play, err := pc.playsService.GetPlayById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch plays"})
		return
	}

	if play == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Play not found"})
		return
	}
	ctx.JSON(http.StatusOK, play)		
}

// func (pc *PlaysController) GetPlayById (ctx *gin.Context) {
// 	id := ctx.Param("id")
// 	play, err := pc.playsService.GetPlayById(id)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch plays"})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, play)		
// }


