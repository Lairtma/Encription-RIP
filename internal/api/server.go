package api

import (
	"RIP/internal/app/config"
	"RIP/internal/app/ds"
	"RIP/internal/app/dsn"
	"RIP/internal/app/repository"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
)

type Application struct {
	repo        *repository.Repository
	minioClient *minio.Client
	config      *config.Config
}

// @title DevIntApp
// @version 1.1
// @description This is API for Text en/decryption requests
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func (a *Application) Run() {
	log.Println("Server start up")

	r := gin.Default()

	r.GET("/api/texts", a.RoleMiddleware(ds.Users{IsModerator: false}, ds.Users{IsModerator: true}), a.GetAllTexts)
	r.GET("/api/texts/:Id", a.RoleMiddleware(ds.Users{IsModerator: false}, ds.Users{IsModerator: true}), a.GetText)
	r.POST("/api/texts", a.CreateText)
	r.DELETE("/api/texts/:Id", a.DeleteText)
	r.PUT("/api/texts/:Id", a.UpdateText)
	r.POST("/api/text_to_order/:Id", a.AddTextToOrder)
	r.POST("api/text/pic/:Id", a.ChangePic)

	r.GET("/api/order", a.RoleMiddleware(ds.Users{IsModerator: false}, ds.Users{IsModerator: true}), a.GetAllOrdersWithParams)
	r.GET("/api/order/:Id", a.RoleMiddleware(ds.Users{IsModerator: false}, ds.Users{IsModerator: true}), a.GetOrder)
	r.PUT("/api/order/:Id", a.UpdateFieldsOrder)
	r.DELETE("/api/order/:Id", a.DeleteOrder)
	r.PUT("/api/order/form/:Id", a.FormOrder)
	r.PUT("/api/order/finish/:Id", a.RoleMiddleware(ds.Users{IsModerator: true}), a.FinishOrder)

	r.DELETE("/api/order_text/:Id", a.DeleteTextFromOrder)
	r.PUT("/api/order_text/:Id", a.UpdatePositionTextInOrder)

	r.POST("/api/register_user", a.RegisterUser)
	r.POST("/api/login_user", a.LoginUser)
	r.POST("/api/logout", a.LogoutUser)
	r.GET("/protected", a.RoleMiddleware(ds.Users{IsModerator: true}), func(c *gin.Context) {
		userID := c.MustGet("userID").(float64)
		c.JSON(http.StatusOK, gin.H{"message": "Пользователь авторизован с правами модератора", "userID": userID})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
