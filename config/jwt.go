package config

type JWT struct {
	AccessSecret string `mapstructure:"AccessSecret" json:"AccessSecret" yaml:"AccessSecret"` // jwt签名
	AccessExpire int64  `mapstructure:"AccessExpire" json:"AccessExpire" yaml:"AccessExpire"` // 过期时间
	BufferTime   int64  `mapstructure:"BufferTime" json:"BufferTime" yaml:"BufferTime"`       // 缓冲时间
	Issuer       string `mapstructure:"Issuer" json:"Issuer" yaml:"Issuer"`                   // 签发者
}
