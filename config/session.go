package config

type Session struct {
	Name   string `mapstructure:"Name" json:"Name" yaml:"Name"`       //session 名字
	Type   string `mapstructure:"Type" json:"Type" yaml:"Type"`       //驱动方式，支持file、cache
	Store  string `mapstructure:"Store" json:"Store" yaml:"Store"`    //存储连接标识 当type使用cache的时候有效
	Expire int    `mapstructure:"Expire" json:"Expire" yaml:"Expire"` //过期时间
	Prefix string `mapstructure:"Prefix" json:"Prefix" yaml:"Prefix"` //前缀
}
