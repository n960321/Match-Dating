package handler

import (
	"Match-Dating/internal/engine"
	"Match-Dating/internal/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Handler struct {
	matchEngine *engine.MatchEngine
}

func NewHandler() *Handler {
	return &Handler{
		matchEngine: engine.NewMatchEngine(),
	}
}

type ReqCreateProile struct {
	Name        string `json:"Name"`
	Gender      int    `json:"Gender"`
	Height      uint   `json:"Height"`
	WantedDates uint   `json:"WantedDates"`
}

func (h *Handler) CreateProfile(c *gin.Context) {
	reqBody := new(ReqCreateProile)
	err := c.Bind(&reqBody)

	if err != nil {
		c.Error(err)
		return
	}

	// Vaild the input

	profile := &model.Profile{
		Name:        reqBody.Name,
		Height:      reqBody.Height,
		Gender:      model.Gender(reqBody.Gender),
		WantedDates: reqBody.WantedDates,
	}

	h.matchEngine.AddProfile(profile)
	result := h.matchEngine.GetAllMatch(profile)
	result.Profile = profile
	c.JSON(http.StatusOK, result)
}

func (h *Handler) DeleteProfile(c *gin.Context) {
	idStr := c.Param("profile_id")
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil || id <= 0 {
		log.Warn().Err(err).Send()
		c.Status(http.StatusBadRequest)
		return
	}

	h.matchEngine.RemoveProfile(&model.Profile{ID: uint(id)})

	c.Status(http.StatusNoContent)
}

func (h *Handler) QuerySinglePeople(c *gin.Context) {
	panic("not implement")
}
