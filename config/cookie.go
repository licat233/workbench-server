package config

type Cookie struct {
	Expire int    `json:"expire" yaml:"expire"` //保存时间
	Path   string `json:"Path" yaml:"Path"`     //保存路径
	Domain string `json:"Domain" yaml:"Domain"` //有效域名
}
