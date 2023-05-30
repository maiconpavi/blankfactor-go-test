package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/maiconpavi/blankfactor-go-test/repository"
	"github.com/maiconpavi/blankfactor-go-test/schema"
)

const (
	eventNotFound string = "event not found"
)

// @Summary
// @Description
// @Accept  json
// @Produce  json
// @Tags event
// @Success 200 {object} []schema.Event
// @Router /event/list [get]
func EventList(ctx *gin.Context) {
	eventRepository, err := repository.NewEventRepository()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	events, err := eventRepository.List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

// @Summary
// @Description
// @Accept  json
// @Produce  json
// @Tags event
// @Success 200 {object} []schema.EventPair
// @Router /event/list-overlap-pairs [get]
func EventListOverlapPairs(ctx *gin.Context) {
	eventRepository, err := repository.NewEventRepository()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	events, err := eventRepository.ListOverlapPairs()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

// @Summary
// @Description
// @Accept  json
// @Produce  json
// @Param event body schema.Event true "event"
// @Tags event
// @Success 200 {object} int
// @Router /event [post]
func EventPost(ctx *gin.Context) {
	eventRepository, err := repository.NewEventRepository()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var event schema.Event
	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := eventRepository.Insert(event)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, id)
}

// @Summary
// @Description
// @Accept  json
// @Produce  json
// @Tags event
// @Param id path int true "Event ID"
// @Success 200 {object} schema.Event
// @Router /event/{id} [get]
func EventGet(ctx *gin.Context) {
	eventRepository, err := repository.NewEventRepository()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rawID := ctx.Param("id")
	id, err := strconv.Atoi(rawID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event, err := eventRepository.Get(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if event.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": eventNotFound})
		return
	}
	ctx.JSON(http.StatusOK, event)
}

// @Summary
// @Description
// @Accept  json
// @Produce  json
// @Tags event
// @Param id path int true "Event ID"
// @Success 200 {object} int
// @Router /event/{id} [delete]
func EventDelete(ctx *gin.Context) {
	eventRepository, err := repository.NewEventRepository()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rawID := ctx.Param("id")
	id, err := strconv.Atoi(rawID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if event, err := eventRepository.Get(id); err != nil || event.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": eventNotFound})
		return
	}

	if err := eventRepository.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, id)
}

// @Summary
// @Description
// @Accept  json
// @Produce  json
// @Tags event
// @Param event body schema.Event true "event"
// @Success 200 {object} schema.Event
// @Router /event [put]
func EventPut(ctx *gin.Context) {
	eventRepository, err := repository.NewEventRepository()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var event schema.Event
	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if event, err := eventRepository.Get(event.ID); err != nil || event.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": eventNotFound})
		return
	}

	if err := eventRepository.Update(event); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, event)
}
