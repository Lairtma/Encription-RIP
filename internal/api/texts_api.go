package api

import (
	"RIP/internal/app/schemas"
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func (a *Application) GetAllTexts(c *gin.Context) {
	var request schemas.GetAllTextsRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	texts, err := a.repo.GetAllTexts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	texts_cnt := len(texts)
	wrk_text_req, err := a.repo.GetWorkingOrderByUserId(1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var id int
	if len(wrk_text_req) == 0 {
		id = 0
	} else {
		id = wrk_text_req[0].Id
	}
	response := schemas.GetAllTextsResponse{Id: id, Count: texts_cnt, Text: texts}
	c.JSON(http.StatusOK, response)
}

func (a *Application) GetText(c *gin.Context) {
	var request schemas.GetTextRequest
	request.Id = c.Param("Id")
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	text, err := a.repo.GetTextByID(request.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := schemas.GetTextResponse{Text: text}
	c.JSON(http.StatusOK, response)
}

func (a *Application) CreateText(c *gin.Context) {
	var request schemas.CreateTextRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := a.repo.CreateText(request.Text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, "Text was created")
}

func (a *Application) DeleteText(c *gin.Context) {
	var request schemas.GetTextRequest
	request.Id = c.Param("Id")
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := a.repo.DeleteTextByID(request.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Text was deleted")
}

func (a *Application) UpdateText(c *gin.Context) {
	var request schemas.UpdateTextRequest
	request.Id = c.Param("Id")
	if err := c.ShouldBindQuery(&request.Text); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&request.Text); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := a.repo.UpdateTextByID(request.Id, request.Text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Text was updated")
}

func (a *Application) AddTextToOrder(c *gin.Context) {
	var request schemas.AddTextToOrderRequest
	request.Id = c.Param("Id")
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	text, err := a.repo.GetTextByID(request.Id)
	var encType string
	if text.Enc {
		encType = "Тип:/nШифрование с битом чётности"
	} else {
		encType = "Тип:/nДешифрование с битом чётности"
	}

	if err != nil { // если не получилось
		log.Printf("cant transform ind", err)
		c.Error(err)
		c.String(http.StatusBadRequest, "Invalid ID")
		return
	}

	order_work, err := a.repo.GetWorkingOrderByUserId(1)
	var order_ID int
	if len(order_work) == 0 {
		new_order, err := a.repo.CreateOrder(1, 2)
		if err != nil {
			log.Println("unable to create order")
		}
		order_ID = new_order.Id
	} else {
		order_ID = order_work[0].Id
	}
	position, err := a.repo.GetTextIdsByOrderId(order_ID)
	a.repo.AddToOrder(order_ID, text.Id, len(position)+1, encType)

	c.JSON(http.StatusOK, "Text was added")
}

func (a *Application) ChangePic(c *gin.Context) {
	var request schemas.ChangePicRequest
	var err error
	request.Id = c.Param("Id")
	file, err := c.FormFile("img")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	openFile, err := file.Open()
	defer openFile.Close()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	byt, err := ioutil.ReadAll(openFile)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reader := bytes.NewReader(byt)
	_, err = a.minioClient.PutObject(context.Background(), "lab1", file.Filename, reader, int64(file.Size), minio.PutObjectOptions{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	url, err := a.minioClient.PresignedGetObject(context.Background(), "lab1", file.Filename, time.Second*24*60*60, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = a.repo.ChangePicByID(request.Id, strings.Split(url.String(), "?")[0])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Text Pic was updated")
}

func (a *Application) DeletePic(c *gin.Context) {
	var request schemas.DeletePicRequest
	var err error
	request.Id = c.Param("Id")

	text, err := a.repo.GetTextByID(request.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	splitedText := strings.Split(text.Img, "/")
	err = a.minioClient.RemoveObject(context.Background(), "lab1", splitedText[len(splitedText)-1], minio.RemoveObjectOptions{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = a.repo.ChangePicByID(request.Id, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Text Pic was deleted")
}
