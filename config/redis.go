package config

type Redis struct {
	DB       int    `mapstructure:"DB" json:"DB" yaml:"DB"`                   // redis的哪个数据库
	Addr     string `mapstructure:"Addr" json:"Addr" yaml:"Addr"`             // 服务器地址:端口
	Password string `mapstructure:"Password" json:"Password" yaml:"Password"` // 密码
	Type     string `mapstructure:"Type" json:"Type" yaml:"Type"`             //类型
}
