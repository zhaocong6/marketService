package routes

import (
	"google.golang.org/grpc"
	"ws/marketApi/app/rpc/controller"
	marketPd "ws/marketApi/pd/market"
)

func InitRpc(s *grpc.Server) {
	marketC := &controller.Market{}

	marketPd.RegisterMarketServer(s, marketC)
}
