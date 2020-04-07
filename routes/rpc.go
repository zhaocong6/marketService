package routes

import (
	"google.golang.org/grpc"
	"marketApi/app/rpc/controller"
	marketPd "marketApi/pd/market"
)

func InitRpc(s *grpc.Server) {
	marketC := &controller.Market{}

	marketPd.RegisterMarketServer(s, marketC)
}
