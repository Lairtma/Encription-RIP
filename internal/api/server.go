package api

import (
	"RIP/internal/app/config"
	"RIP/internal/app/dsn"
	"RIP/internal/app/repository"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

type Application struct {
	repo        *repository.Repository
	minioClient *minio.Client
	config      *config.Config
}

func (a *Application) Run() {
	log.Println("Server start up")

	r := gin.Default()

	r.GET("/api/texts", a.GetAllTexts)
	r.GET("/api/texts/:Id", a.GetText)
	r.POST("/api/texts", a.CreateText)
	r.DELETE("/api/texts/:Id", a.DeleteText)
	r.PUT("/api/texts/:Id", a.UpdateText)
	r.POST("/api/text_to_order/:Id", a.AddTextToOrder)
	r.POST("api/text/change_pic/:Id", a.ChangePic)

	r.GET("/api/order", a.GetAllOrdersWithParams)
	r.GET("/api/order/:Id", a.GetOrder)
	r.PUT("/api/order/:Id", a.UpdateFieldsOrder)
	r.DELETE("/api/order/:Id", a.DeleteOrder)
	r.PUT("/api/order/form/:Id", a.FormOrder)
	r.PUT("/api/order/finish/:Id", a.FinishOrder)

	r.DELETE("/api/order_text/:Id", a.DeleteTextFromOrder)
	r.PUT("/api/order_text/:Id", a.UpdatePositionTextInOrder)

	r.POST("/api/registration", a.CreateUser)

	r.Static("/css", "./resources")

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	log.Println("Server down")
}

func New() (*Application, error) {
	var err error
	app := Application{}
	app.config, err = config.NewConfig()
	if err != nil {
		return nil, err
	}

	app.repo, err = repository.New(dsn.FromEnv())
	if err != nil {
		return nil, err
	}

	endpoint := "localhost:9000"
	accessKeyID := "minio"
	secretAccessKey := "minio124"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Println(err)
	}

	app.minioClient = minioClient

	return &app, nil
}
