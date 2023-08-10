// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	relation "doushen_by_liujun/service/user/api/internal/handler/relation"
	userinfo "doushen_by_liujun/service/user/api/internal/handler/userinfo"
	"doushen_by_liujun/service/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: userinfo.RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: userinfo.LoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/douyin/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtAuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/",
					Handler: userinfo.UserinfoHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/douyin/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtAuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/follow/list",
					Handler: relation.FollowListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/follower/list",
					Handler: relation.FollowerListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/action",
					Handler: relation.FollowHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/friend/list",
					Handler: relation.FriendListHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/douyin/relation"),
	)
}
