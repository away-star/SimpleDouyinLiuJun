package relation

import (
	"context"
	"doushen_by_liujun/internal/common"
	"doushen_by_liujun/internal/util"
	"doushen_by_liujun/service/user/api/internal/svc"
	"doushen_by_liujun/service/user/api/internal/types"
	"doushen_by_liujun/service/user/rpc/pb"
	"github.com/juju/ratelimit"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type FollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	bucket *ratelimit.Bucket
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		bucket: ratelimit.NewBucketWithRate(10, 10), //每秒钟生成 10 个令牌——令牌桶限流1s内最多处理10个请求
	}
}

func (l *FollowLogic) Follow(req *types.FollowReq) (resp *types.FollowResp, err error) {
	logger, e := util.ParseToken(req.Token)
	if e != nil {
		if err := l.svcCtx.KqPusherClient.Push("user_api_relation_followLogic_Follow_ParseToken_false"); err != nil {
			log.Fatal(err)
		}
		return &types.FollowResp{
			StatusCode: common.TOKEN_EXPIRE_ERROR,
			StatusMsg:  "token失效",
		}, err
	}
	if l.bucket.TakeAvailable(1) == 0 {
		// 令牌不足，限流处理
		//判断是关注还是取消关注
		if req.ActionType == 1 {
			go func(userId, followId string) { //新开协程执行延迟写的操作
				randomInterval := time.Duration(rand.Int63n(int64(10 * time.Minute))) //10分钟内的随机时间
				ticker := time.NewTicker(randomInterval)                              //延迟执行写入数据库
			OuterLoop:
				for {
					select {
					case _ = <-ticker.C:
						_, err := l.svcCtx.UserRpcClient.AddFollows(l.ctx, &pb.AddFollowsReq{
							UserId:   userId,
							FollowId: followId,
						})
						if err == nil { //确实写进去了再结束
							ticker.Stop()
							break OuterLoop
						}
					}
				}
			}(strconv.Itoa(int(logger.UserID)), strconv.FormatInt(req.ToUserId, 10))
			return &types.FollowResp{
				StatusCode: common.OK,
				StatusMsg:  "关注成功",
			}, nil
		} else {
			go func(userId, followId string) { //新开协程执行延迟写的操作
				randomInterval := time.Duration(rand.Int63n(int64(10 * time.Minute))) //10分钟内的随机时间
				ticker := time.NewTicker(randomInterval)                              //延迟执行写入数据库
			OuterLoop:
				for {
					select {
					case _ = <-ticker.C:
						_, err := l.svcCtx.UserRpcClient.DelFollows(l.ctx, &pb.DelFollowsReq{
							UserId:   userId,
							FollowId: followId,
						})
						if err == nil { //确实写进去了再结束
							ticker.Stop()
							break OuterLoop
						}
					}
				}
			}(strconv.Itoa(int(logger.UserID)), strconv.FormatInt(req.ToUserId, 10))

			return &types.FollowResp{
				StatusCode: common.OK,
				StatusMsg:  "取关成功",
			}, err
		}
	} else {
		//判断是关注还是取消关注
		if req.ActionType == 1 {
			_, err := l.svcCtx.UserRpcClient.AddFollows(l.ctx, &pb.AddFollowsReq{
				UserId:   strconv.Itoa(int(logger.UserID)),
				FollowId: strconv.FormatInt(req.ToUserId, 10),
			})
			if err != nil {
				if err := l.svcCtx.KqPusherClient.Push("user_api_relation_followLogic_Follow_AddFollows_false"); err != nil {
					log.Fatal(err)
				}
				return &types.FollowResp{
					StatusCode: common.DB_ERROR,
					StatusMsg:  "关注失败",
				}, err
			}
			if err := l.svcCtx.KqPusherClient.Push("user_api_relation_followLogic_Follow_AddFollows_success"); err != nil {
				log.Fatal(err)
			}
			return &types.FollowResp{
				StatusCode: common.OK,
				StatusMsg:  "关注成功",
			}, nil
		} else {
			_, err := l.svcCtx.UserRpcClient.DelFollows(l.ctx, &pb.DelFollowsReq{
				UserId:   strconv.Itoa(int(logger.UserID)),
				FollowId: strconv.FormatInt(req.ToUserId, 10),
			})
			if err != nil {
				if err := l.svcCtx.KqPusherClient.Push("user_api_relation_followLogic_Follow_DelFollows_false"); err != nil {
					log.Fatal(err)
				}
				return &types.FollowResp{
					StatusCode: common.DB_ERROR,
					StatusMsg:  "删除关注失败",
				}, err
			}
			if err := l.svcCtx.KqPusherClient.Push("user_api_relation_followLogic_Follow_DelFollows_success"); err != nil {
				log.Fatal(err)
			}
			return &types.FollowResp{
				StatusCode: common.OK,
				StatusMsg:  "取关成功",
			}, err
		}
	}
}
