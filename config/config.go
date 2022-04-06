package config

type Server struct {
	//服务配置
	ServerConfig ServerConfig `mapstructure:"ServerConfig" json:"ServerConfig" yaml:"ServerConfig"`
	StaticConfig StaticConfig `mapstructure:"StaticConfig" json:"StaticConfig" yaml:"StaticConfig"`
	//鉴权与权限管理
	Captcha Captcha `mapstructure:"Captcha" json:"Captcha" yaml:"Captcha"`
	JWT     JWT     `mapstructure:"JWT" json:"JWT" yaml:"JWT"`
	Casbin  Casbin  `mapstructure:"Casbin" json:"Casbin" yaml:"Casbin"`
	//storage
	Mysql   Mysql   `mapstructure:"Mysql" json:"Mysql" yaml:"Mysql"`
	Redis   Redis   `mapstructure:"redis" json:"Redis" yaml:"Redis"`
	Cookie  Cookie  `mapstructure:"Cookie" json:"Cookie" yaml:"Cookie"`
	Session Session `mapstructure:"Session" json:"Session" yaml:"Session"`
	//缓存
	Cache Cache `mapstructure:"Cache" json:"Cache" yaml:"Cache"`
	//日志
	Zap Zap `mapstructure:"Zap" json:"Zap" yaml:"Zap"`
	//邮箱
	Email Email `mapstructure:"Email" json:"Email" yaml:"Email"`
	//跨域
	CORS CORS `mapstructure:"CORS" json:"CORS" yaml:"CORS"`
}
