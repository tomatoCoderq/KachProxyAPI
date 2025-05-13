package controllers

import (
	"github.com/tomatoCoderq/KachProxyAPI/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ActorsController struct {
	actorsService *services.ActorsService
}

func NewActorsController(actorsService *services.ActorsService) *ActorsController {
	return &ActorsController {
		actorsService: actorsService,
	}
}

func (ac *ActorsController) GetAllActors(ctx *gin.Context) {
	actors, err := ac.actorsService.GetAllActors()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch actors"})
		return
	}
	ctx.JSON(http.StatusOK, actors)
}

func (ac *ActorsController) GetActorById(ctx *gin.Context) {
	id := ctx.Param("id")
	actor, err := ac.actorsService.GetActorById(id)
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