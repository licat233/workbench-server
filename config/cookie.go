package config

type Cookie struct {
	Expire int    `mapstructure:"Expire" json:"Expire" yaml:"Expire"` //保存时间
	Path   string `mapstructure:"Path" json:"Path" yaml:"Path"`       //保存路径
	Domain string `mapstructure:"Domain" json:"Domain" yaml:"Domain"` //有效域名
}
