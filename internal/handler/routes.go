// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"greet/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/from/:name",
				Handler: GreetHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/hostname",
				Handler: HostnameHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/validate",
				Handler: ValidateHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user-id",
				Handler: UserIdHandler(serverCtx),
			},
		},
	)
}
