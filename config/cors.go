package config

type CORS struct {
	Mode      string          `mapstructure:"Mode" json:"Mode" yaml:"Mode"`                //cors模式：Dev | Pro
	Whitelist []CORSWhitelist `mapstructure:"Whitelist" json:"Whitelist" yaml:"Whitelist"` //请求源白名单列表
}

type CORSWhitelist struct {
	AllowOrigin      string `mapstructure:"AllowOrigin" json:"AllowOrigin" yaml:"AllowOrigin"`                //允许请求源
	AllowMethods     string `mapstructure:"AllowMethods" json:"AllowMethods" yaml:"AllowMethods"`             //允许请求方法
	AllowHeaders     string `mapstructure:"AllowHeaders" json:"AllowHeaders" yaml:"AllowHeaders"`             //允许请求中携带字段
	ExposeHeaders    string `mapstructure:"ExposeHeaders" json:"ExposeHeaders" yaml:"ExposeHeaders"`          //允许暴露的response头部信息
	AllowCredentials bool   `mapstructure:"AllowCredentials" json:"AllowCredentials" yaml:"AllowCredentials"` //允许客户端携带验证信息
}
