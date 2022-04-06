package config

type Email struct {
	Contacts []string `mapstructure:"Contacts" json:"Contacts" yaml:"Contacts"` // 联系人
	Port     int      `mapstructure:"Port" json:"Port" yaml:"Port"`             // 端口
	From     string   `mapstructure:"From" json:"From" yaml:"From"`             // 发件人
	Host     string   `mapstructure:"Host" json:"Host" yaml:"Host"`             // 服务器地址
	IsSSL    bool     `mapstructure:"IsSSL" json:"IsSSL" yaml:"IsSSL"`          // 是否SSL
	Secret   string   `mapstructure:"Secret" json:"Secret" yaml:"Secret"`       // 密钥
	Nickname string   `mapstructure:"Nickname" json:"Nickname" yaml:"Nickname"` // 昵称
}
