package user

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"workbench-server/common/respx"
	"workbench-server/common/utils"
	"workbench-server/service/app/api/v1/internal/svc"
	"workbench-server/service/app/api/v1/internal/types"
	"workbench-server/service/app/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) SearchUserLogic {
	return SearchUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchUserLogic) SearchUser(req types.UserSearchReq) (interface{}, error) {
	var total int64 = 100
	keyword := strings.TrimSpace(req.KeyWord)
	if len(keyword) == 0 {
		return nil, errors.New("请输入搜索内容")
	}

	var userinfo = make([]*model.UserInfo, 0)
	var err error
	if utils.VerifyEmailFormat(keyword) == nil {
		userinfo, err = l.svcCtx.UserModel.FindAllByEmail(req.Current, req.PageSize, keyword)
	} else if utils.VerifyMobileFormat(keyword) == nil {
		userinfo, err = l.svcCtx.UserModel.FindAllByMobile(req.Current, req.PageSize, keyword)
	} else if utils.VerifyRealnameFormat(keyword) == nil {
		userinfo, err = l.svcCtx.UserModel.FindAllByRealName(req.Current, req.PageSize, keyword)
	} else if utils.VerifyUidFormat(keyword) == nil {
		uid, _ := strconv.ParseInt(keyword, 10, 64)
		res, e := l.svcCtx.UserModel.FindOneById(uid)
		return respx.NewListResponse(e, res, total, req.PageSize)
	} else {
		userinfo, err = l.svcCtx.UserModel.FindAllByPassport(req.Current, req.PageSize, keyword)
	}
	return respx.NewListResponse(err, userinfo, total, req.PageSize)
}
