package svc

import (
	"workbench-server/common/global"
	"workbench-server/service/v1/app/api/internal/config"
	"workbench-server/service/v1/app/api/internal/middleware"
	"workbench-server/service/v1/app/model"

	"github.com/casbin/casbin/v2"
	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config config.Config
	//Middleware
	JwtMiddleware             rest.Middleware
	CasbinMiddleware          rest.Middleware
	SafeWallMiddleware        rest.Middleware
	OperationRecordMiddleware rest.Middleware
	//实例
	Jwt    *middleware.JwtMiddleware
	Casbin *casbin.SyncedEnforcer
	Redis  *redis.Redis
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
	MysqlConn := sqlx.NewMysql(global.CONFIG.Mysql.DataSource())
	RedisCacheConf := global.CONFIG.Cache.Redis
	return &ServiceContext{
		Config: c,
		//middleware
		JwtMiddleware:             middleware.NewJwtMiddleware().Handle,
		CasbinMiddleware:          middleware.NewCasbinMiddleware().Handle,
		SafeWallMiddleware:        middleware.NewSafeWallMiddleware().Handle,
		OperationRecordMiddleware: middleware.NewOperationRecordMiddleware().Handle,
		//实例
		Jwt:    middleware.NewJwtMiddleware(),
		Casbin: middleware.NewCasbinMiddleware().CasbinService,
		Redis:  global.RedisX,
		//models
		WebsiteConfigModel:  model.NewWebsiteConfigModel(MysqlConn, RedisCacheConf),
		UserModel:           model.NewUserModel(MysqlConn, RedisCacheConf),
		RoleModel:           model.NewSysRoleModel(MysqlConn, RedisCacheConf),
		LogModel:            model.NewSysLogModel(MysqlConn, RedisCacheConf),
		CompanyModel:        model.NewCompanyModel(MysqlConn, RedisCacheConf),
		BrandModel:          model.NewBrandModel(MysqlConn, RedisCacheConf),
		ProjectModel:        model.NewProjectModel(MysqlConn, RedisCacheConf),
		AdChannelModel:      model.NewAdChannelModel(MysqlConn, RedisCacheConf),
		AdAccountModel:      model.NewAdAccountModel(MysqlConn, RedisCacheConf),
		AdCampaignModel:     model.NewAdCampaignModel(MysqlConn, RedisCacheConf),
		AdCampaignCostModel: model.NewAdCampaignCostModel(MysqlConn, RedisCacheConf),
		AdGroupModel:        model.NewAdGroupModel(MysqlConn, RedisCacheConf),
		LandingPageModel:    model.NewLandingPageModel(MysqlConn, RedisCacheConf),
		ArticleModel:        model.NewArticleModel(MysqlConn, RedisCacheConf),
		CostModel:           model.NewCostListModel(MysqlConn, RedisCacheConf),
		FansModel:           model.NewFansListModel(MysqlConn, RedisCacheConf),
		ServiceAccountModel: model.NewServiceAccountModel(MysqlConn, RedisCacheConf),
		MenuModel:           model.NewMenuModel(MysqlConn, RedisCacheConf),
		TeacherModel:        model.NewTeacherModel(MysqlConn, RedisCacheConf),
	}
}

func (s *ServiceContext) ProcessCaptcha(CaptchaKey, Solution string) bool {
	return base64Captcha.DefaultMemStore.Verify(CaptchaKey, Solution, true)
}
