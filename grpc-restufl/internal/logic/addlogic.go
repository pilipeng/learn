package logic

import (
	"context"

	"grpc-restufl/internal/svc"
	"grpc-restufl/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *pb.SumRequest) (*pb.SumReponse, error) {
	return &pb.SumReponse{
		Result: in.Numbers.A + in.Numbers.B,
	}, nil
}
