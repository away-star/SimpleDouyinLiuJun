// Code generated by goctl. DO NOT EDIT.
// Source: content.proto

package server

import (
	"context"

	"doushen_by_liujun/service/content/rpc/internal/logic"
	"doushen_by_liujun/service/content/rpc/internal/svc"
	"doushen_by_liujun/service/content/rpc/pb"
)

type ContentServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedContentServer
}

func NewContentServer(svcCtx *svc.ServiceContext) *ContentServer {
	return &ContentServer{
		svcCtx: svcCtx,
	}
}

// -----------------------璇勮淇℃伅-----------------------
func (s *ContentServer) AddComment(ctx context.Context, in *pb.AddCommentReq) (*pb.AddCommentResp, error) {
	l := logic.NewAddCommentLogic(ctx, s.svcCtx)
	return l.AddComment(in)
}

func (s *ContentServer) UpdateComment(ctx context.Context, in *pb.UpdateCommentReq) (*pb.UpdateCommentResp, error) {
	l := logic.NewUpdateCommentLogic(ctx, s.svcCtx)
	return l.UpdateComment(in)
}

func (s *ContentServer) DelComment(ctx context.Context, in *pb.DelCommentReq) (*pb.DelCommentResp, error) {
	l := logic.NewDelCommentLogic(ctx, s.svcCtx)
	return l.DelComment(in)
}

func (s *ContentServer) GetCommentById(ctx context.Context, in *pb.GetCommentByIdReq) (*pb.GetCommentByIdResp, error) {
	l := logic.NewGetCommentByIdLogic(ctx, s.svcCtx)
	return l.GetCommentById(in)
}

func (s *ContentServer) SearchComment(ctx context.Context, in *pb.SearchCommentReq) (*pb.SearchCommentResp, error) {
	l := logic.NewSearchCommentLogic(ctx, s.svcCtx)
	return l.SearchComment(in)
}

// -----------------------鐐硅禐淇℃伅-----------------------
func (s *ContentServer) AddFavorite(ctx context.Context, in *pb.AddFavoriteReq) (*pb.AddFavoriteResp, error) {
	l := logic.NewAddFavoriteLogic(ctx, s.svcCtx)
	return l.AddFavorite(in)
}

func (s *ContentServer) UpdateFavorite(ctx context.Context, in *pb.UpdateFavoriteReq) (*pb.UpdateFavoriteResp, error) {
	l := logic.NewUpdateFavoriteLogic(ctx, s.svcCtx)
	return l.UpdateFavorite(in)
}

func (s *ContentServer) DelFavorite(ctx context.Context, in *pb.DelFavoriteReq) (*pb.DelFavoriteResp, error) {
	l := logic.NewDelFavoriteLogic(ctx, s.svcCtx)
	return l.DelFavorite(in)
}

func (s *ContentServer) GetFavoriteById(ctx context.Context, in *pb.GetFavoriteByIdReq) (*pb.GetFavoriteByIdResp, error) {
	l := logic.NewGetFavoriteByIdLogic(ctx, s.svcCtx)
	return l.GetFavoriteById(in)
}

func (s *ContentServer) SearchFavorite(ctx context.Context, in *pb.SearchFavoriteReq) (*pb.SearchFavoriteResp, error) {
	l := logic.NewSearchFavoriteLogic(ctx, s.svcCtx)
	return l.SearchFavorite(in)
}

// -----------------------瑙嗛淇℃伅-----------------------
func (s *ContentServer) AddVideo(ctx context.Context, in *pb.AddVideoReq) (*pb.AddVideoResp, error) {
	l := logic.NewAddVideoLogic(ctx, s.svcCtx)
	return l.AddVideo(in)
}

func (s *ContentServer) UpdateVideo(ctx context.Context, in *pb.UpdateVideoReq) (*pb.UpdateVideoResp, error) {
	l := logic.NewUpdateVideoLogic(ctx, s.svcCtx)
	return l.UpdateVideo(in)
}

func (s *ContentServer) DelVideo(ctx context.Context, in *pb.DelVideoReq) (*pb.DelVideoResp, error) {
	l := logic.NewDelVideoLogic(ctx, s.svcCtx)
	return l.DelVideo(in)
}

func (s *ContentServer) GetVideoById(ctx context.Context, in *pb.GetVideoByIdReq) (*pb.GetVideoByIdResp, error) {
	l := logic.NewGetVideoByIdLogic(ctx, s.svcCtx)
	return l.GetVideoById(in)
}

func (s *ContentServer) SearchVideo(ctx context.Context, in *pb.SearchVideoReq) (*pb.SearchVideoResp, error) {
	l := logic.NewSearchVideoLogic(ctx, s.svcCtx)
	return l.SearchVideo(in)
}

func (s *ContentServer) GetFeedList(ctx context.Context, in *pb.FeedListReq) (*pb.FeedListResp, error) {
	l := logic.NewGetFeedListLogic(ctx, s.svcCtx)
	return l.GetFeedList(in)
}

func (s *ContentServer) GetWorkCountByUserId(ctx context.Context, in *pb.GetWorkCountByUserIdReq) (*pb.GetWorkCountByUserIdResp, error) {
	l := logic.NewGetWorkCountByUserIdLogic(ctx, s.svcCtx)
	return l.GetWorkCountByUserId(in)
}
