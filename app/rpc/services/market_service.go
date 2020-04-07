package services

import (
	"github.com/zhaocong6/market"
	pd "marketApi/pd/market"
)

type MarketService struct{}

func (m *MarketService) GetMarketData(in *pd.MarketRequest) *pd.MarketResponse {
	data := market.Find(in.Organize, in.Symbol)
	if len(data) == 0 {
		return &pd.MarketResponse{}
	}
	return m.marketToResponse(data)[in.Symbol]
}

func (m *MarketService) marketToResponse(data map[string]*market.Marketer) map[string]*pd.MarketResponse {
	rep := make(map[string]*pd.MarketResponse)

	for k, v := range data {
		rep[k] = &pd.MarketResponse{
			Organize:  string(v.Organize),
			Symbol:    v.Symbol,
			BuyFirst:  v.BuyFirst,
			SellFirst: v.SellFirst,
			BuyDepth:  m.depthToArray(v.BuyDepth),
			SellDepth: m.depthToArray(v.SellDepth),
			Timestamp: int64(v.Timestamp),
			Temporize: int64(v.Temporize),
		}
	}

	return rep
}

func (m *MarketService) depthToArray(depth market.Depth) []*pd.Depth {
	d := make([]*pd.Depth, len(depth))

	for k, v := range depth {
		d[k] = &pd.Depth{
			Price:  v[0],
			Amount: v[1],
		}
	}

	return d
}
