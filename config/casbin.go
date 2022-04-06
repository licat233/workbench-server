package config

type Casbin struct { //Casbin Model文件
	ModelPath string `json:"ModelPath" yaml:"ModelPath"` //model文件路径
}
