package app

import (
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

	r.GET("/cards", func(c *gin.Context) {
		query := c.Query("query") // Получаем поисковый запрос из URL
		log.Printf("query recived %s\n", query)
		cards, err := a.repo.GetAllCards()
		if err != nil { // если не получилось
			log.Printf("cant get product by id %v", err)
			c.Error(err)
			return
		}
		if query == "en" {
			c.HTML(http.StatusOK, "home.html", gin.H{
				"title":      "Main website",
				"first_row":  cards,
				"second_row": "",
				"query":      query,
			})
		} else if query == "de" {
			c.HTML(http.StatusOK, "home.html", gin.H{
				"title":      "Main website",
				"first_row":  cards,
				"second_row": "",
				"query":      query,
			})
		} else {
			c.HTML(http.StatusOK, "home.html", gin.H{
				"title":      "Main website",
				"first_row":  cards,
				"second_row": "",
				"query":      query,
			})
		}
	})

	r.GET("/cart", func(c *gin.Context) {
		c.HTML(http.StatusOK, "cart.html", gin.H{
			"title":     "Main website",
			"card_data": CartInfoFunc(),
		})
	})

	r.GET("/card/:id", func(c *gin.Context) {
		id := c.Param("id") // Получаем ID из URL
		index, err := strconv.Atoi(id)

		if err != nil || index < 0 || index > len(CardsInfoFunc("")) {
			c.String(http.StatusBadRequest, "Invalid ID")
			return
		}

		c.HTML(http.StatusOK, "card.html", gin.H{
			"title":     "Main website",
			"card_data": CardsInfoFunc("")[index-1],
		})
	})

	r.Static("/image", "./resources")

	r.GET("/product", func(c *gin.Context) {
		id := c.Query("id") // получаем из запроса query string

		if id != "" {
			log.Printf("id recived %s\n", id)
			product, err := a.repo.GetCardByType(true)
			if err != nil { // если не получилось
				log.Printf("cant get product by id %v", err)
				c.Error(err)
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"product_price": product[0].Name,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "try with id",
		})
	})
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

type Element struct {
	Id          int
	Name        string
	First_img   string
	Second_img  string
	Encrypting  bool
	Description string
}

type Cart struct {
	Id         int
	Name       string
	First_img  string
	Second_img string
	Text       string
}

func CardsInfoFunc(tpe string) []Element {
	if tpe == "en" {
		return []Element{
			{1, "Шифрование с битом чётности", "http://localhost:9000/lab1/%D1%87%D1%91%D1%82%D0%BD%D0%BE%D1%81%D1%82%D1%8C.png", "http://localhost:9000/lab1/%D1%88%D0%B8%D1%84%D1%80%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5.png", true, ""},
			{2, "Шифрование с контрольной суммой", "http://localhost:9000/lab1/%D0%A1%D1%83%D0%BC%D0%BC%D0%B0.png", "http://localhost:9000/lab1/%D1%88%D0%B8%D1%84%D1%80%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5.png", true, ""},
			{3, "Шифрование с повторением битов", "http://localhost:9000/lab1/%D0%BF%D0%BE%D0%B2%D1%82%D0%BE%D1%80%D0%B5%D0%BD%D0%B8%D0%B5.png", "http://localhost:9000/lab1/%D1%88%D0%B8%D1%84%D1%80%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5.png", true, ""},
			{4, "Шифрование кодом Хэмминга", "http://localhost:9000/lab1/%D1%85%D1%8D%D0%BC%D0%B8%D0%BD%D0%B3.jpg", "http://localhost:9000/lab1/%D1%88%D0%B8%D1%84%D1%80%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5.png", true, ""},
		}
	} else if tpe == "de" {
		return []Element{
			{5, "Дешифрование с битом чётности", "http://localhost:9000/lab1/%D1%87%D1%91%D1%82%D0%BD%D0%BE%D1%81%D1%82%D1%8C.png", "http://localhost:9000/lab1/%D0%B4%D0%B5%D1%88%D0%B8%D1%84%D1%80%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5.png", false, ""},
			{6, "Дешифрование с контрольной суммой", "http://localhost:9000/lab1/%D0%A1%D1%83%D0%BC%D0%BC%D0%B0.png", "http://localhost:9000/lab1/%D0%B4%D0%B5%D1%88%D0%B8%D1%84%D1%80%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5.png", false, ""},
			{7, "Дешифрование с повторением битов", "http://localhost:9000/lab1/%D0%BF%D0%BE%D0%B2%D1%82%D0%BE%D1%80%D0%B5%D0%BD%D0%B8%D0%B5.png", "http://localhost:9000/lab1/%D0%B4%D0%B5%D1%88%D0%B8%D1%84%D1%80%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5.png", false, ""},
			{8, "Дешифрование кодом Хэмминга", "http://localhost:9000/lab1/%D1%85%D1%8D%D0%BC%D0%B8%D0%BD%D0%B3.jpg", "http://localhost:9000/lab1/%D0%B4%D0%B5%D1%88%D0%B8%D1%84%D1%80%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5.png", false, ""},
		}
	}
	return []Element{
		{1, "Шифрование с битом чётности", "http://localhost:9000/lab1/%D1%87%D1%91%D1%82%D0%BD%D0%BE%D1%81%D1%82%D1%8C.png", "http://localhost:9000/lab1/%D1%88%D0%B8%D1%84%D1%80%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5.png", true, "Описание:\nВ каждом пакет данных есть один бит четности, или, так называемый, паритетный бит. Этот бит устанавливается во время записи (или отправки) данных, и затем рассчитывается и сравнивается во время чтения (получения) данных. Он равен сумме по модулю 2 всех бит данных в пакете. То есть число единиц в пакете всегда будет четно . Изменение этого бита (например с 0 на 1) сообщает о возникшей ошибке."},
		{2, "Шифрование с контрольной суммой", "http://localhost:9000/lab1/%D0%A1%D1%83%D0%BC%D0%BC%D0%B0.png", "http://localhost:9000/lab1/%D1%88%D0%B8%D1%84%D1%80%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5.png", true, "Описание: \nВ общем виде контрольная сумма представляет собой некоторое значение, вычисленное по определённой схеме на основе кодируемого сообщения. Проверочная информация при систематическом кодировании приписывается к передаваемым данным. На принимающей стороне абонент знает алгоритм вычисления контрольной суммы: соответственно, программа имеет возможность проверить корректность принятых данных.\nПри передаче пакетов по сетевому каналу могут возникнуть искажения исходной информации вследствие разных внешних воздействий: электрических наводок, плохих погодных условий и многих других. Сущность методики в том, что при хороших характеристиках контрольной суммы в подавляющем числе случаев ошибка в сообщении приведёт к изменению его контрольной суммы. Если исходная и вычисленная суммы не равны между собой, принимается решение о недостоверности принятых данных, и можно запросить повторную передачу пакета."},
		{3, "Шифрование с повторением битов", "http://localhost:9000/lab1/%D0%BF%D0%BE%D0%B2%D1%82%D0%BE%D1%80%D0%B5%D0%BD%D0%B8%D0%B5.png", "http://localhost:9000/lab1/%D1%88%D0%B8%D1%84%D1%80%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5.png", true, "Описание: \n Если с ошибкой произойдет передача одного бита из трех, то ошибка будет исправлена, но если случится двойная или тройная ошибка, то будут получены неправильные данные. Часто коды для исправления ошибок используют совместно с кодами для обнаружения ошибок. При тройном повторении для повышения надежности три бита располагают не подряд, а на фиксированном расстоянии друг от друга. Использование тройного повторения естественно значительно снижает скорость передачи данных."},
		{4, "Шифрование кодом Хэмминга", "http://localhost:9000/lab1/%D1%85%D1%8D%D0%BC%D0%B8%D0%BD%D0%B3.jpg", "http://localhost:9000/lab1/%D1%88%D0%B8%D1%84%D1%80%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5.png", true, "Описание: \n Для каждого числа проверочных символов используется специальная маркировка вида (k, i), где k — количество символов в сообщении, i — количество информационных символов в сообщении. Например, существуют коды (7, 4), (15, 11), (31, 26). Каждый проверочный символ в коде Хэмминга представляет сумму по модулю 2 некоторой подпоследовательности данных."},
		{5, "Дешифрование с битом чётности", "http://localhost:9000/lab1/%D1%87%D1%91%D1%82%D0%BD%D0%BE%D1%81%D1%82%D1%8C.png", "http://localhost:9000/lab1/%D0%B4%D0%B5%D1%88%D0%B8%D1%84%D1%80%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5.png", false, "Описание:\nВ каждом пакет данных есть один бит четности, или, так называемый, паритетный бит. Этот бит устанавливается во время записи (или отправки) данных, и затем рассчитывается и сравнивается во время чтения (получения) данных. Он равен сумме по модулю 2 всех бит данных в пакете. То есть число единиц в пакете всегда будет четно . Изменение этого бита (например с 0 на 1) сообщает о возникшей ошибке."},
		{6, "Дешифрование с контрольной суммой", "http://localhost:9000/lab1/%D0%A1%D1%83%D0%BC%D0%BC%D0%B0.png", "http://localhost:9000/lab1/%D0%B4%D0%B5%D1%88%D0%B8%D1%84%D1%80%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5.png", false, "Описание: \nВ общем виде контрольная сумма представляет собой некоторое значение, вычисленное по определённой схеме на основе кодируемого сообщения. Проверочная информация при систематическом кодировании приписывается к передаваемым данным. На принимающей стороне абонент знает алгоритм вычисления контрольной суммы: соответственно, программа имеет возможность проверить корректность принятых данных.\nПри передаче пакетов по сетевому каналу могут возникнуть искажения исходной информации вследствие разных внешних воздействий: электрических наводок, плохих погодных условий и многих других. Сущность методики в том, что при хороших характеристиках контрольной суммы в подавляющем числе случаев ошибка в сообщении приведёт к изменению его контрольной суммы. Если исходная и вычисленная суммы не равны между собой, принимается решение о недостоверности принятых данных, и можно запросить повторную передачу пакета."},
		{7, "Дешифрование с повторением битов", "http://localhost:9000/lab1/%D0%BF%D0%BE%D0%B2%D1%82%D0%BE%D1%80%D0%B5%D0%BD%D0%B8%D0%B5.png", "http://localhost:9000/lab1/%D0%B4%D0%B5%D1%88%D0%B8%D1%84%D1%80%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5.png", false, "Описание: \n Если с ошибкой произойдет передача одного бита из трех, то ошибка будет исправлена, но если случится двойная или тройная ошибка, то будут получены неправильные данные. Часто коды для исправления ошибок используют совместно с кодами для обнаружения ошибок. При тройном повторении для повышения надежности три бита располагают не подряд, а на фиксированном расстоянии друг от друга. Использование тройного повторения естественно значительно снижает скорость передачи данных."},
		{8, "Дешифрование кодом Хэмминга", "http://localhost:9000/lab1/%D1%85%D1%8D%D0%BC%D0%B8%D0%BD%D0%B3.jpg", "http://localhost:9000/lab1/%D0%B4%D0%B5%D1%88%D0%B8%D1%84%D1%80%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5.png", false, "Описание: \n Для каждого числа проверочных символов используется специальная маркировка вида (k, i), где k — количество символов в сообщении, i — количество информационных символов в сообщении. Например, существуют коды (7, 4), (15, 11), (31, 26). Каждый проверочный символ в коде Хэмминга представляет сумму по модулю 2 некоторой подпоследовательности данных."},
	}
}

func CartInfoFunc() []Cart {
	return []Cart{
		{1, "Шифрование с битом чётности", "http://localhost:9000/lab1/%D1%87%D1%91%D1%82%D0%BD%D0%BE%D1%81%D1%82%D1%8C.png", "http://localhost:9000/lab1/%D1%88%D0%B8%D1%84%D1%80%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5.png", "Текст: \nПривет, мир"},
		{2, "Дешифрование с битом чётности", "http://localhost:9000/lab1/%D1%87%D1%91%D1%82%D0%BD%D0%BE%D1%81%D1%82%D1%8C.png", "http://localhost:9000/lab1/%D0%B4%D0%B5%D1%88%D0%B8%D1%84%D1%80%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5.png", "Текст: \n1010 1000 0000 1111 1001 0011 0101 0110 0000 1100 1010"},
		{2, "Дешифрование с битом чётности", "http://localhost:9000/lab1/%D1%87%D1%91%D1%82%D0%BD%D0%BE%D1%81%D1%82%D1%8C.png", "http://localhost:9000/lab1/%D0%B4%D0%B5%D1%88%D0%B8%D1%84%D1%80%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5.png", "Текст: \n1010 1100 0000 1111 1001 0011 0101 0110 0000 1100 1010"},
	}
}
