package request

type MarketRequest struct {
	Organize   string `json:"organize" binding:"required"`
	MarketType int8   `json:"market_type" binding:"required,min=1,max=4"`
	Symbol     string `json:"symbol" binding:"required"`
}
