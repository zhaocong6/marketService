package request

type MarketRequest struct {
	requester

	Organize   string `json:"organize" validate:"required"`
	MarketType int8   `json:"market_type" validate:"required,min=1,max=4"`
	Symbol     string `json:"symbol" validate:"required"`
}
