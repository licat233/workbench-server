package config

type Casbin struct { //Casbin Model文件
	ModelPath string `mapstructure:"ModelPath" json:"ModelPath" yaml:"ModelPath"` //model文件路径
}
