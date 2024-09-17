package app

import (
	"RIP/internal/app/ds"
	"errors"
	"gorm.io/gorm"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"RIP/internal/app/config"
	"RIP/internal/app/dsn"
	"RIP/internal/app/repository"
)

type Application struct {
	repo   *repository.Repository
	config *config.Config
	// dsn string
}

func (a *Application) Run() {
	log.Println("Server start up")

	r := gin.Default()

	r.SetFuncMap(template.FuncMap{
		"replaceNewline": func(text string) template.HTML {
			return template.HTML(strings.ReplaceAll(text, "/n", "<br>"))
		},
		"contains": func(s, substr string) bool {
			return strings.Contains(s, substr)
		},
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.LoadHTMLGlob("templates/*")
	r.Static("/css", "./resources/css")

	r.GET("/cards", func(c *gin.Context) {
		query := c.Query("query") // Получаем поисковый запрос из URL
		log.Printf("query recived %s\n", query)
		var cards []ds.Card
		var err error
		if query != "" {
			var encryption bool
			if query == "en" {
				encryption = true
			} else {
				encryption = false
			}
			cards, err = a.repo.GetCardByType(encryption)
			if err != nil { // если не получилось
				log.Printf("cant get cards by type %v", err)
				c.Error(err)
				return
			}
		} else {
			cards, err = a.repo.GetAllCards()
			if err != nil { // если не получилось
				log.Printf("cant get cards  %v", err)
				c.Error(err)
				return
			}
		}
		var cart_len, cart_id int
		cart, err := a.repo.GetLastCartByCreatorId(1)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			cart_len = 0
			cart_id = 0
		} else if err != nil { // если не получилось
			log.Printf("cant get cart  %v", err)
			c.Error(err)
			return
		} else {
			cart_id = int(cart.Id)
			cartCards, err := a.repo.GetCartCardsByCartId(int(cart.Id))
			if err != nil { // если не получилось
				log.Printf("cant get cartCards  %v", err)
				c.Error(err)
				return
			}
			cart_len = min(len(cartCards), 10)
		}
		c.HTML(http.StatusOK, "home.html", gin.H{
			"title":     "Main website",
			"first_row": cards,
			"query":     query,
			"cart_len":  cart_len,
			"cart_id":   cart_id,
		})
	})

	r.GET("/cart/:id", func(c *gin.Context) {
		id := c.Param("id") // Получаем ID из URL
		index, err := strconv.Atoi(id)
		if err != nil { // если не получилось
			log.Printf("cant get card by id %v", err)
			c.Error(err)
			c.String(http.StatusBadRequest, "Invalid ID")
			return
		}

		if index == 0 {
			c.HTML(http.StatusOK, "cart.html", gin.H{
				"title":     "Main website",
				"card_data": [0]Cart{},
			})
		} else {

			cart, err := a.repo.GetCardByID(index)
			if err != nil { // если не получилось
				log.Printf("cant get card by id %v", err)
				c.Error(err)
				c.String(http.StatusBadRequest, "Invalid ID")
				return
			}

			cartCards, err := a.repo.GetCartCardsByCartId(int(cart.Id))
			if err != nil { // если не получилось
				log.Printf("cant get cartCards  %v", err)
				c.Error(err)
				return
			}
			var cart_html []Cart
			for _, cart_card := range cartCards {
				card, err := a.repo.GetCardByID(int(cart_card.CardId))
				if err != nil { // если не получилось
					log.Printf("cant get cartCards  %v", err)
					c.Error(err)
					return
				}
				cart_html = append(cart_html, Cart{card.Name, card.First_img, card.Second_img, cart_card.Text})
			}

			c.HTML(http.StatusOK, "cart.html", gin.H{
				"title":     "Main website",
				"card_data": cart_html,
			})
		}
	})

	r.GET("/card/:id", func(c *gin.Context) {
		id := c.Param("id") // Получаем ID из URL
		index, err := strconv.Atoi(id)
		if err != nil { // если не получилось
			log.Printf("cant get card by id %v", err)
			c.Error(err)
			c.String(http.StatusBadRequest, "Invalid ID")
			return
		}

		card, err := a.repo.GetCardByID(index)
		if err != nil { // если не получилось
			log.Printf("cant get card by id %v", err)
			c.Error(err)
			c.String(http.StatusBadRequest, "Invalid ID")
			return
		}

		c.HTML(http.StatusOK, "card.html", gin.H{
			"title":     "Main website",
			"card_data": card,
		})
	})

	r.Static("/image", "./resources")

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

	return &app, nil
}

type Cart struct {
	Name       string
	First_img  string
	Second_img string
	Text       string
}
