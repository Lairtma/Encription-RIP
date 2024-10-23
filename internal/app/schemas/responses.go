package schemas

import (
	"RIP/internal/app/ds"
)

type GetAllTextsResponse struct {
	Id    int                 `json:"text_req_ID"`
	Count int                 `json:"count"`
	Text  []ds.TextToEncOrDec `json:"texts"`
}

type GetTextResponse struct {
	Text ds.TextToEncOrDec `json:"text"`
}

type GetAllOrdersWithParamsResponse struct {
	Orders []ds.EncOrDecOrder
}

type GetAllOrdersResponse struct {
	Orders []ds.EncOrDecOrder `json:"orders"`
}

type GetOrderResponse struct {
	Order ds.EncOrDecOrder    `json:"order"`
	Count int                 `json:"count"`
	Texts []ds.TextToEncOrDec `json:"texts"`
}
type CreateTextResponse struct {
	Id              int
	MessageResponse string
}

type AddTextToOrderResponce struct {
	TextId          int
	OrderId         int
	MessageResponse string
}

type DeleteTextResponse struct {
	Id              int
	MessageResponse string
}

type DeleteMealFromMilkReqResponse struct{}

type UpdateOrderMilkReqMealsResponse struct{}

type ResponseMessage struct {
}
