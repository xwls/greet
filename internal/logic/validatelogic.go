package logic

import (
	"context"
	"github.com/go-playground/validator/v10"

	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ValidateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewValidateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidateLogic {
	return &ValidateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ValidateLogic) Validate(req *types.User) (resp *types.User, err error) {
	validate := validator.New()
	//err = validate.StructCtx(l.ctx, req)
	err = validate.Struct(req)
	if err != nil {
		return nil, err
	}
	resp = req
	return
}
