package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf          // rest api配置
	Auth          struct { // jwt鉴权配置
		AccessSecret string // jwt密钥
		AccessExpire int64  // 有效期，单位：秒
	}
	Mysql struct { // 数据库配置，除mysql外，可能还有mongo等其他数据库
		DataSource string // mysql链接地址，满足 $user:$password@tcp($ip:$port)/$db?$queries 格式即可
	}
	Redis   cache.CacheConf // redis缓存
	Storage struct {
		Path   string `json:"Path" yaml:"Path"` //文件仓库目录
		Upload struct {
			Path   string `json:"Path" yaml:"Path"` //上传文件目录
			Images struct {
				Path   string `json:"Path" yaml:"Path"` //上传图片目录
				Avatar struct {
					Path string `json:"Path" yaml:"Path"` //上传头像目录
				} `json:"Avatar" yaml:"Avatar"`
			} `json:"Images" yaml:"Images"`
			Collect struct {
				Path string `json:"Path" yaml:"Path"` //上传收藏夹目录
			} `json:"Collect" yaml:"Collect"`
		} `json:"Upload" yaml:"Upload"`
		Static struct {
			Path string `json:"Path" yaml:"Path"` //静态文件
			Js   struct {
				Path string `json:"Path" yaml:"Path"` // js文件
			} `json:"Js" yaml:"Js"`
			Css struct {
				Path string `json:"Path" yaml:"Path"` // js文件
			} `json:"Css" yaml:"Css"`
			Images struct {
				Path string `json:"Path" yaml:"Path"` // img文件
			} `json:"Images" yaml:"Images"`
			Svg struct {
				Path string `json:"Path" yaml:"Path"` // svg文件
			} `json:"Svg" yaml:"Svg"`
			Json struct {
				Path string `json:"Path" yaml:"Path"` // svg文件
			} `json:"Json" yaml:"Json"`
		} `json:"Static" yaml:"Static"`
	} `json:"Storage" yaml:"Storage"`
}
