package api

import (
	"RIP/internal/app/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Удалить текст из запроса на шифрование
// @Description Удалить текст из запроса на шифрование по ID запроса и TextId
// @Tags text_and_order
// @Accept json
// @Produce json
// @Param Id path string true "ID запроса на шифрование"
// @Param TextId query string true "Id текста"
// @Success 200 {string} string "Text was deleted from order"
// @Failure 400 {object} schemas.ResponseMessage
// @Failure 500 {object} schemas.ResponseMessage
// @Router /api/order_text/{Id} [delete]
func (a *Application) DeleteTextFromOrder(c *gin.Context) {
	var request schemas.DeleteTextFromOrderRequest
	request.Id = c.Param("Id")
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ID, err := strconv.Atoi(request.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	errr := a.repo.DeleteTextFromOrder(ID, request.TextId)
	if errr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Text was deleted from order")
}

func (a *Application) UpdatePositionTextInOrder(c *gin.Context) {
	var request schemas.UpdatePositionTextInOrderRequest
	request.Id = c.Param("Id")
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ID, err := strconv.Atoi(request.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	errr := a.repo.UpdatePositionTextInOrder(ID, request.TextId, request.Position)
	if errr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Text position was changed in order")
}
