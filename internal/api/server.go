package api

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type TextToEncOrDec struct {
	Id      int
	Enc     bool
	Text    string
	Img     string
	ByteLen int
}

type EncOrDecOrder struct {
	Id      int
	Img     string
	Text    string
	EncType string
	Res     string
	Ord     int
}

func CardsInfoFunc() []TextToEncOrDec {
	return []TextToEncOrDec{
		{1, true, "Текст:\nПривет, мир!", "http://localhost:9000/lab1/hello_world_ru.png", 21},
		{2, true, "Текст:\nHello, world!", "http://localhost:9000/lab1/hello_world_en.png", 13},
		{3, true, "Текст:\nКак дела?", "http://localhost:9000/lab1/how_are_you_ru.png", 16},
		{4, true, "Текст:\nHow are you?", "http://localhost:9000/lab1/how_are_you_en.png", 12},
		{5, false, "Текст:\n1111 0001 1111 0000 1111 1010 1100", "http://localhost:9000/lab1/1111_0001_1111_0000_1111_1010_1100.png", 28},
		{6, false, "Текст:\n1111 0001 1111 0000 1111 1010 1100", "http://localhost:9000/lab1/1111_0001_1111_0000_1111_1010_1100.png", 28},
		{7, false, "Текст:\n1111 0001 1111 0000 1111 1010 1100", "http://localhost:9000/lab1/1111_0001_1111_0000_1111_1010_1100.png", 28},
		{8, false, "Текст:\n1111 0001 1111 0000 1111 1010 1100", "http://localhost:9000/lab1/1111_0001_1111_0000_1111_1010_1100.png", 28},
	}
}

func CartInfoFunc() []EncOrDecOrder {
	return []EncOrDecOrder{
		{1, "http://localhost:9000/lab1/hello_world_ru.png", "Текст:\nПривет, мир!", "Тип: Шифрование с битом чётности", "", 1},
		{2, "http://localhost:9000/lab1/how_are_you_ru.png", "Текст:\nКак дела?", "Тип: Шифрование с битом чётности", "", 2},
		{3, "http://localhost:9000/lab1/1111_0001_1111_0000_1111_1010_1100.png", "Текст:\n1111 0001 1111 0000 1111 1010 1100", "Тип: Дешифрование с битом чётности", "", 3},
	}
}

func StartServer() {
	log.Println("Server start up")

	r := gin.Default()

	r.SetFuncMap(template.FuncMap{
		"replaceNewline": func(text string) template.HTML {
			return template.HTML(strings.ReplaceAll(text, "\n", "<br>"))
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

	r.GET("/textsencordec", func(c *gin.Context) {
		query := c.Query("query") // Получаем поисковый запрос из URL
		fmt.Println(query)
		arr := CardsInfoFunc()
		var new_arr []TextToEncOrDec
		for i := 0; i < (len(arr)); {
			if arr[i].Enc && query == "en" || !arr[i].Enc && query == "de" || query == "" {
				new_arr = append(new_arr, arr[i])
			}
			i++
		}
		fmt.Println(new_arr)
		c.HTML(http.StatusOK, "textsencordec.html", gin.H{
			"title":     "Main website",
			"first_row": new_arr,
			"query":     query,
			"len":       len(CartInfoFunc()),
		})
	})

	r.GET("/encordecorder/:id", func(c *gin.Context) {
		c.HTML(http.StatusOK, "encordecorder.html", gin.H{
			"title":     "Main website",
			"card_data": CartInfoFunc(),
		})
	})

	r.GET("/text/:id", func(c *gin.Context) {
		id := c.Param("id") // Получаем ID из URL
		index, err := strconv.Atoi(id)

		if err != nil || index < 0 || index > len(CardsInfoFunc()) {
			c.String(http.StatusBadRequest, "Invalid ID")
			return
		}

		c.HTML(http.StatusOK, "text.html", gin.H{
			"title":     "Main website",
			"card_data": CardsInfoFunc()[index-1],
		})
	})

	r.Static("/image", "./resources")

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	log.Println("Server down")
}
