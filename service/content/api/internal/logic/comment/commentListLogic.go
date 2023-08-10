package comment

import (
	"context"
	"doushen_by_liujun/internal/common"
	"doushen_by_liujun/internal/util"
	"doushen_by_liujun/service/content/api/internal/svc"
	"doushen_by_liujun/service/content/api/internal/types"
	"doushen_by_liujun/service/content/rpc/pb"
	userPB "doushen_by_liujun/service/user/rpc/pb"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"time"
)

type CommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentListLogic) CommentList(req *types.CommentListReq) (resp *types.CommentListResp, err error) {
	// todo: add your logic here and delete this line
	logger, e := util.ParseToken(req.Token)
	if e != nil {
		return &types.CommentListResp{
			StatusCode:  common.TOKEN_EXPIRE_ERROR,
			StatusMsg:   "无效token",
			CommentList: []types.Comment{},
		}, e
	}
	follows, e := l.svcCtx.ContentRpcClient.GetCommentById(l.ctx, &pb.GetCommentByIdReq{
		Id: req.VideoId,
	})
	var comments []types.Comment
	if e != nil {
		return &types.CommentListResp{
			StatusCode:  common.DB_ERROR,
			StatusMsg:   "查询评论列表失败",
			CommentList: []types.Comment{},
		}, e
	}
	for _, item := range follows.Comment {
		//查询用户信息
		info, e := l.svcCtx.UserRpcClient.GetUserinfoById(l.ctx, &userPB.GetUserinfoByIdReq{
			Id: item.UserId,
		})
		fmt.Println("用户信息", info, e)
		var user types.User
		if e != nil {
			return &types.CommentListResp{
				StatusCode:  common.DB_ERROR,
				StatusMsg:   "查询评论列表失败",
				CommentList: []types.Comment{},
			}, e
		}
		followCount, e := l.svcCtx.UserRpcClient.GetFollowsCountById(l.ctx, &userPB.GetFollowsCountByIdReq{
			Id: item.UserId,
		})
		if e != nil {
			return &types.CommentListResp{
				StatusCode:  common.DB_ERROR,
				StatusMsg:   "查询评论列表失败",
				CommentList: []types.Comment{},
			}, e
		}
		followerCount, e := l.svcCtx.UserRpcClient.GetFollowersCountById(l.ctx, &userPB.GetFollowersCountByIdReq{
			Id: item.UserId,
		})
		if e != nil {
			return &types.CommentListResp{
				StatusCode:  common.DB_ERROR,
				StatusMsg:   "查询评论列表失败",
				CommentList: []types.Comment{},
			}, e
		}
		isFollow, e := l.svcCtx.UserRpcClient.CheckIsFollow(l.ctx, &userPB.CheckIsFollowReq{
			Userid:   logger.ID,
			Followid: strconv.Itoa(int(info.Userinfo.Id)),
		})
		if e != nil {
			return &types.CommentListResp{
				StatusCode:  common.DB_ERROR,
				StatusMsg:   "查询评论列表失败",
				CommentList: []types.Comment{},
			}, e
		}
		user = types.User{
			Id:              info.Userinfo.Id,
			Name:            info.Userinfo.Name,
			FollowCount:     followCount.Count,
			FollowerCount:   followerCount.Count,
			IsFollow:        isFollow.IsFollowed, //我对这个的理解就是当前用户对这条数据的用户是否关注
			Avatar:          info.Userinfo.Avatar,
			BackgroundImage: info.Userinfo.BackgroundImage,
			Signature:       info.Userinfo.Signature,
			WorkCount:       0, //查表
			FavoriteCount:   0, //查表
		}
		comments = append(comments, types.Comment{
			Id:         item.Id,
			User:       user,
			Content:    item.Content,
			CreateDate: time.Unix(item.CreateTime, 0).Format("01-02"),
		})
	}
	return &types.CommentListResp{
		StatusCode:  common.OK,
		StatusMsg:   "查询评论列表成功",
		CommentList: comments,
	}, nil
}
