package logic

import (
	"context"
	"os"

	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HostnameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHostnameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HostnameLogic {
	return &HostnameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HostnameLogic) Hostname() (resp *types.HostNameResponse, err error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	resp = &types.HostNameResponse{Hostname: hostname}
	return
}
