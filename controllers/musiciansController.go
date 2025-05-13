package controllers

import (
	"KachProxyAPI/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MusiciansController struct {
	musiciansService *services.MusiciansService
}

func NewMusiciansController(musiciansService *services.MusiciansService) *MusiciansController {
	return &MusiciansController {
		musiciansService: musiciansService,
	}
}

func (ac *MusiciansController) GetAllMusicians(ctx *gin.Context) {
	actors, err := ac.musiciansService.GetAllMusicians()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch actors"})
		return
	}
	ctx.JSON(http.StatusOK, actors)
}

func (ac *MusiciansController) GetMusicianById(ctx *gin.Context) {
	id := ctx.Param("id")
	actor, err := ac.musiciansService.GetMusicianById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch actors"})
		return
	}
	if actor == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Actor not found"})
		return
	}
	ctx.JSON(http.StatusOK, actor)
}