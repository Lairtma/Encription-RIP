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

type CreateMealResponse struct{}

type DeleteMealResponse struct{}

///MILK REQUESTS MILK REQUESTS MILK REQUESTS MILK REQUESTS MILK REQUESTS MILK REQUESTS MILK REQUESTS MILK REQUESTS MILK REQUESTS///

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

type DeleteMealFromMilkReqResponse struct{}

type UpdateOrderMilkReqMealsResponse struct{}
