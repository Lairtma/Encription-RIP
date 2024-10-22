package schemas

import (
	"RIP/internal/app/ds"
	"time"
)

type GetAllTextsRequest struct{}

type GetTextRequest struct {
	Id string
}

type CreateTextRequest struct {
	Text ds.TextToEncOrDec
}

type DeleteTextRequest struct {
	ID string
}

type UpdateTextRequest struct {
	Id   string
	Text ds.TextToEncOrDec
}

type AddTextToOrderRequest struct {
	Id string
}

type ChangePicRequest struct {
	Id  string `json:"id"`
	Img string `json:"image_link"`
}

type DeletePicRequest struct {
	Id string `json:"id"`
}

///MILK REQUESTS MILK REQUESTS MILK REQUESTS MILK REQUESTS MILK REQUESTS MILK REQUESTS MILK REQUESTS MILK REQUESTS MILK REQUESTS///

type GetAllOrdersWithParamsRequest struct {
	HavingStatus bool      `json:"is_status"`
	Status       int       `json:"status"`
	FromDate     time.Time `json:"from_date"`
	ToDate       time.Time `json:"to_date"`
}

type GetOrderRequest struct {
	Id string
}

type UpdateOrderRequest struct {
	ID     int `json:"milk_req_id"`
	MealID int `json:"meal_id"`
	OrderO int `json:"order_o"`
}

type UpdateFieldsOrderRequest struct {
	Id       string `uri:"milk_request" json:"id"`
	Priority int    `json:"priority"`
}

type DeleteOrderRequest struct {
	Id string
}

type FormOrderRequest struct {
	Id string
}

type FinishOrderRequest struct {
	Id     string
	Status int `json:"status"`
}

type DeleteTextFromOrderRequest struct {
	Id     string
	TextId int `json:"text_id"`
}
type UpdatePositionTextInOrderRequest struct {
	Id       string
	TextId   int `json:"text_id"`
	Position int `json:"position"`
}

type CreateUserRequest struct {
	ds.Users
}
