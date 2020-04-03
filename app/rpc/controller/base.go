package controller

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Base struct{}

func (b *Base) InvalidArgumentResponse(msg string) error {
	return status.Error(codes.InvalidArgument, msg)
}
