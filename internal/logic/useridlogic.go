package logic

import (
	"context"
	"github.com/go-playground/validator/v10"

	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserIdLogic {
	return &UserIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserIdLogic) UserId(req *types.UserIdReq) (resp *types.UserIdReq, err error) {
	validate := validator.New()
	err = validate.StructCtx(l.ctx, req)
	if err != nil {
		return nil, err
	}
	resp = req
	return
}
