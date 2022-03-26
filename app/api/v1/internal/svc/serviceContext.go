package svc

import (
	"workbench-server/app/api/v1/internal/config"
	"workbench-server/app/api/v1/internal/middleware"
	"workbench-server/app/model"

	"github.com/casbin/casbin/v2"
	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config          config.Config
	Jwt             rest.Middleware
	Casbin          rest.Middleware
	SafeWall        rest.Middleware
	OperationRecord rest.Middleware
	JwtObject       *middleware.JwtMiddleware
	CasbinObj       *casbin.Enforcer
	RedisClient     *redis.Redis
	//Model
	WebsiteConfigModel  model.WebsiteConfigModel
	UserModel           model.UserModel
	RoleModel           model.SysRoleModel
	LogModel            model.SysLogModel
	CompanyModel        model.CompanyModel
	BrandModel          model.BrandModel
	ProjectModel        model.ProjectModel
	AdChannelModel      model.AdChannelModel
	AdAccountModel      model.AdAccountModel
	AdCampaignModel     model.AdCampaignModel
	AdCampaignCostModel model.AdCampaignCostModel
	AdGroupModel        model.AdGroupModel
	LandingPageModel    model.LandingPageModel
	ArticleModel        model.ArticleModel
	CostModel           model.CostListModel
	FansModel           model.FansListModel
	ServiceAccountModel model.ServiceAccountModel
	MenuModel           model.MenuModel
	TeacherModel        model.TeacherModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	MysqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	RedisClient := redis.New(c.Redis[0].Host, redis.WithPass(c.Redis[0].Pass))
	return &ServiceContext{
		Config:          c,
		Jwt:             middleware.NewJwtMiddleware(c).Handle,
		Casbin:          middleware.NewCasbinMiddleware().Handle,
		OperationRecord: middleware.NewOperationRecordMiddleware().Handle,
		SafeWall:        middleware.NewSafeWallMiddleware().Handle,
		JwtObject:       middleware.NewJwtMiddleware(c),
		CasbinObj:       middleware.NewCasbinMiddleware().CasbinObj,
		RedisClient:     RedisClient,
		//models
		WebsiteConfigModel:  model.NewWebsiteConfigModel(MysqlConn, c.Redis),
		UserModel:           model.NewUserModel(MysqlConn, c.Redis),
		RoleModel:           model.NewSysRoleModel(MysqlConn, c.Redis),
		LogModel:            model.NewSysLogModel(MysqlConn, c.Redis),
		CompanyModel:        model.NewCompanyModel(MysqlConn, c.Redis),
		BrandModel:          model.NewBrandModel(MysqlConn, c.Redis),
		ProjectModel:        model.NewProjectModel(MysqlConn, c.Redis),
		AdChannelModel:      model.NewAdChannelModel(MysqlConn, c.Redis),
		AdAccountModel:      model.NewAdAccountModel(MysqlConn, c.Redis),
		AdCampaignModel:     model.NewAdCampaignModel(MysqlConn, c.Redis),
		AdCampaignCostModel: model.NewAdCampaignCostModel(MysqlConn, c.Redis),
		AdGroupModel:        model.NewAdGroupModel(MysqlConn, c.Redis),
		LandingPageModel:    model.NewLandingPageModel(MysqlConn, c.Redis),
		ArticleModel:        model.NewArticleModel(MysqlConn, c.Redis),
		CostModel:           model.NewCostListModel(MysqlConn, c.Redis),
		FansModel:           model.NewFansListModel(MysqlConn, c.Redis),
		ServiceAccountModel: model.NewServiceAccountModel(MysqlConn, c.Redis),
		MenuModel:           model.NewMenuModel(MysqlConn, c.Redis),
		TeacherModel:        model.NewTeacherModel(MysqlConn, c.Redis),
	}
}

func (s *ServiceContext) ProcessCaptcha(CaptchaKey, Solution string) bool {
	return base64Captcha.DefaultMemStore.Verify(CaptchaKey, Solution, true)
}
