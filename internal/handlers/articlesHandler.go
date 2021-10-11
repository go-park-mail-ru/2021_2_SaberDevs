package handlers

import (
	"github.com/go-park-mail-ru/2021_2_SaberDevs/internal/data"
	"github.com/go-park-mail-ru/2021_2_SaberDevs/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ArticlesHandler struct {
}

func NewArticlesHandler() *ArticlesHandler {
	var handler ArticlesHandler
	return &handler
}

const chunkSize = 5

func (api *ArticlesHandler) Getfeed(c echo.Context) error {
	rec := c.QueryParam("idLastLoaded")
	if rec == "" {
		rec = "0"
	}

	from, err := strconv.Atoi(rec)
	if err != nil {
		c.Logger().Printf("Error: %s", err.Error())
		return c.JSON(http.StatusNotFound, models.ErrNotFeedNumber)
	}
	var ChunkData []models.NewsRecord
	// Возвращаем записи
	testData := data.TestData
	if from >= 0 && from+chunkSize < len(testData) {
		ChunkData = testData[from : from+chunkSize]
	} else {
		start := 0
		if len(testData) > chunkSize {
			start = len(testData) - chunkSize
		}
		ChunkData = testData[start : len(testData)-1]

	}
	// формируем ответ
	response := models.ChunkResponse{
		Status:    http.StatusOK,
		ChunkData: ChunkData,
	}
	return c.JSON(http.StatusOK, response)
}
