package api

import (
	"RIP/internal/app/ds"
	"RIP/internal/app/schemas"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

func (a *Application) GetAllOrdersWithParams(c *gin.Context) {
	var request schemas.GetAllOrdersWithParamsRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if request.FromDate.IsZero() {
		request.FromDate = time.Date(2000, time.January, 1, 0, 0, 0, 396641000, time.UTC)
	}
	if request.ToDate.IsZero() {
		request.ToDate = time.Now()
	}
	orders, err := a.repo.GetAllOrdersWithFilters(request.Status, request.HavingStatus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := schemas.GetAllOrdersWithParamsResponse{Orders: orders}
	c.JSON(http.StatusOK, response)
}

func (a *Application) GetOrder(c *gin.Context) {
	var request schemas.GetOrderRequest
	request.Id = c.Param("Id")
	id_int, err := strconv.Atoi(request.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println("error was there")
		return
	}
	order, err := a.repo.GetOrderByID(id_int)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	texts_in_request, err := a.repo.GetTextIdsByOrderId(order.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	texts := make([]ds.TextToEncOrDec, 0, len(texts_in_request))
	for _, v := range texts_in_request {
		v_string := strconv.Itoa(v)
		text_to_append, err := a.repo.GetTextByID(v_string)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		texts = append(texts, text_to_append)
	}
	response := schemas.GetOrderResponse{Order: order, Count: len(texts_in_request), Texts: texts}
	c.JSON(http.StatusOK, response)
}

func (a *Application) UpdateFieldsOrder(c *gin.Context) {
	var request schemas.UpdateFieldsOrderRequest
	request.Id = c.Param("Id")
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := a.repo.UpdateFieldsOrder(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Fields was updated")
}

func (a *Application) DeleteOrder(c *gin.Context) {
	var request schemas.DeleteOrderRequest
	id := c.Param("Id")
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = a.repo.DeleteOrder(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Order was deleted")
}

func (a *Application) FormOrder(c *gin.Context) {
	var request schemas.FormOrderRequest
	id := c.Param("Id")
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := a.repo.FormOrder(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Order was Formed")
}

func (a *Application) FinishOrder(c *gin.Context) {
	var request schemas.FinishOrderRequest
	id := c.Param("Id")
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := a.repo.FinishOrder(id, request.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Order was Finished")
}
