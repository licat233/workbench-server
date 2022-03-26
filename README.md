# workbench 服务端

## 下载代码：两个仓库可选
* 腾讯CODING仓库
  >```
  >git clone git@e.coding.net:licat/workbench/workbench-server.git
  >```

* Github仓库
  >```
  >git clone git@github.com:licat233/workbench-server.git
  >```

## 代码提交
1. 先配置 SSH 公钥获取仓库拉取提交权限
   ><a href="https://help.coding.net/docs/repo/ssh/config.html">配置 SSH 公钥教程</a>
2. 先设置多个远程仓库
   >```
   >git remote add origin git@e.coding.net:licat/workbench/workbench-server.git
   >```
   >```
   >git remote set-url --add origin git@github.com:licat233/workbench-server.git
   >```
3. 提交代码
   >```
   >git add .
   >```
   >```
   >git commit -m "update" --no-verify
   >```
   >```
   >git push origin --all
   >```
### ☆重置仓库方法（慎用）
   >1.选择最新的分支
   >```
   >git checkout --orphan latest_branch
   >```
   >2.添加所有变化的文件
   >```
   >git add -A
   >```
   >3.给暂存区的文件添加说明
   >```
   >git commit -am newest
   >```
   >4.删除main分支
   >```
   >git branch -D main
   >```
   >5.新建main分支（相当于重置main分支）
   >```
   >git branch -m main
   >```
   >6.强制将本机的代码库推送到远端仓库，并覆盖！
   >```
   >git push -f origin main
   >```
----------------------------------------  

## 开发进度：
>> 1. Captcha服务已经开发完成；
>> 2. Log服务已经开发完成；
>> 3. user服务开发中...
>>> * 基本的注册登录搜索接口；
>>> * 防爆力破解登录，给login接口添加log；
>>> * 已经添加了jwt黑名单注销功能；
>>> * 实现设备单点登录；
>> 4. 配置logx;
>> 5. OperationRecord中间件暂时不用，但任保留；
>> 6. casbin权限管理系统开发：
>>> * 由于要区分不同的公司人员，所以采用【域/租户模型】；
>>> * rbac api服务开发完成;

---------------------------------------  

## 服务组成
>> 1. captcha;
>> 2. log;
>> 3. user;
>> 4. ...

## 命名要求
###  * 服务包名必须以Service结尾，如captchaService;  
###  * goctl命名风格必须为 -style goZero;  
-----------------------------------------  

## 一. MySQL  
> ### 1. 多表关联查询:  (注意LEFT、RIGHT、INNER JOIN的区别)
>> ```
>> SELECT %s FROM %s AS u LEFT JOIN company AS c ON u.user_company_id = c.cid LEFT JOIN sys_role AS s ON u.user_role_id = s.rid WHERE u.user_mobile = ?;
>> ```
> ### 2. 查询指定日期范围内的数据（需要指定到日，关键词 DATA_FORMAT, BETWEEN)
>> ```
>> SELECT %s FROM %s WHERE date_format(create_time,'%Y-%m-%d') >= date_format(start,'%Y-%m-%d') AND date_format(create_time,'%Y-%m-%d') <= date_format(end,'%Y-%m-%d');
>> ```
>> 或者这样写(推荐)
>> ```
>> SELECT %s FROM %s WHERE `create_time` BETWEEN start AND end;
>> ```
> ### 3. 查询符合条件的条数: (关键函数 COUNT())
>> ```
>> SELECT COUNT(*) FROM `sys_log` WHERE date_format(`create_time`,'%Y-%m-%d')= date;
>> ```
>> ```
>> SELECT COUNT(*) FROM `sys_log` WHERE `ip`= ip AND to_days(`create_time`) = to_days(CURDATE());
>> ```

## 二. 使用基于grpc的go-zero微服务框架  
## 常用gotcl插件命令：
## 1. api 服务
> ### ①. 生成api文件命令
>>
>> ```
>> goctl api -o app.api
>> ```
>>
> ### ②. 根据 api文件 生成 api服务，-dir='api路径'
>>
>> ```
>> goctl api go -api app.api -dir='../v1' -style goZero
>> ```
>>
> ### ③. 启动 api 服务
>>
>> ```
>> go run api/user.go -f api/etc/user-api.yaml
>> ```
>>
> ### ④. 端口测试
>>
>> ```
>> curl -i -X POST \
>>    http://127.0.0.1:8888/user/login \
>>    -H 'content-type: application/json' \
>>    -d '{
>>      "passport":"13570923403",
>>      "password":"123456"
>>      "ticket": "213d22"
>>  }'
>> ```

## 2. rpc 服务
> ### ①. 生成 rpc 模板文件 user.proto
>> ```
>> goctl rpc template -o user.proto
>> ```
>>
> ### ②. 根据 proto文件 生成 rpc代码
>> ```
>> goctl rpc proto -src app.proto -dir='./v1' -style goZero
>> ```
>>
> ### ③. 启动 rpc 服务
>> ```
>> go run user.go -f rpc/etc/user.yaml
>> ```

## 3. model 基础代码生成
### * 例如：根据 user 表生成 model
> ### ①.  常规生成
>> ```
>> goctl model mysql datasource -url="root:zxcvbnm3.14LICAT@tcp(127.0.0.1:3306)/workbench" -table="user"  -dir='./model' 
>> ```
> ### ②.  带有缓存
>> ```
>> goctl model mysql datasource -url="root:zxcvbnm3.14LICAT@tcp(127.0.0.1:3306)/workbench" -table="user"  -dir='./model' -cache=true
>> ```
> ### ③.  使用驼峰法命名
>> ```
>> goctl model mysql datasource -url="root:zxcvbnm3.14LICAT@tcp(127.0.0.1:3306)/workbench" -table="user"  -dir='./model' -style goZero
>> ```
> ### ④.  带有缓存且使用驼峰法命名（推荐）
>> ```
>> goctl model mysql datasource -url="workbench:8kMRrkDB2aKkf5Mk@tcp(127.0.0.1:3306)/workbench" -table="brand"  -dir='./model' -cache=true -style goZero
>> ```

## 4. swagger生成
###
>①生成swagger.json 文件
>>```
>> goctl api plugin -plugin goctl-swagger="swagger -filename user.json" -api user.api -dir .
>>```
>②指定Host，basePath api-host-and-base-path
>>```
>>goctl api plugin -plugin goctl-swagger="swagger -filename user.json -host 0.0.0.0 -basepath /api" -api user.api -dir .
>>```
>③swagger ui 查看生成的文档
>>```
>>docker run --rm -p 8083:8080 -e SWAGGER_JSON=/foo/user.json -v $PWD:/foo swaggerapi/swagger-ui
>>```
>④Swagger Codegen 生成客户端调用代码(go,javascript,php)
>>```
>>for l in javascript; do
>>  docker run --rm -v "$(pwd):/go-work" swaggerapi/swagger-codegen-cli generate \
>>    -i "/go-work/rest.swagger.json" \
>>    -l "$l" \
>>    -o "/go-work/clients/$l"
>>done
>>```
----------------------------------------------------------------

# 三、日志打印关闭
> ```
> logx.Disable()
> ```

# 四、自定义错误处理
> ```
> httpx.SetErrorHandler(errorx.ErrorHandler)
> ```
> ```
> func ErrorHandler(err error) (int, interface{}) {
>	if err == nil {
>		return http.StatusInternalServerError, nil
>	}
>	if v,b := err.(*CodeError); b == true {
>		return http.StatusOK, v.Data()
>	}
>	return http.StatusOK, NewErrorResponse(fmt.Sprint(err),nil)
> }
> ```

# 五、数据库连接
> ```
> mysqlConn := sqlx.NewMysql(c.Mysql.DataSource)
> redisConn := redis.NewRedis(c.CacheRedis[0].Host, c.CacheRedis[0].Type, c.CacheRedis[0].Pass)
> ```

# 六、grpc连接
> ```
> captchaClient.NewCaptcha(zrpc.MustNewClient(c.CaptchaRpc))
> ```

# 七、请求头设置
> ```
> Authorization：jwt-token
> ```

# 八、关于客户端IP地址的获取，需要配置Nginx
> ```
> location /v1/user/ {
      add_header 'Access-Control-Allow-Origin' '*';
      add_header 'Access-Control-Allow-Methods' '*';
      proxy_pass http://127.0.0.1:9002;
      proxy_set_header X-Real_IP $remote_addr;
      proxy_set_header Host $host;
      proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
      proxy_http_version 1.1;
    }
> ```

# 九、通过api文件生成swagger文档
> ①生成swagger.json 文件
>> ```
>> goctl api plugin -plugin goctl-swagger="swagger -filename foo/app.json" -api app.api -dir .
>> ```
>
> ②指定Host，basePath api-host-and-base-path
>> ```
>> goctl api plugin -plugin goctl-swagger="swagger -filename app.json -host api.catli.net -basepath /" -api app.api -dir .
>> ```
>
> ③swagger ui 查看生成的文档
>> ```
>> docker run --rm -p 8083:8080 -e SWAGGER_JSON=/foo/app.json -v $PWD:/foo swaggerapi/swagger-ui
>> ```
>
> ④Swagger Codegen 生成客户端调用代码(go,javascript,php)
>> ```
>> for l in ts; do
>>  docker run --rm -v "$(pwd):/go-work" swaggerapi/swagger-codegen-cli generate \
>>    -i "/go-work/rest.swagger.json" \
>>    -l "$l" \
>>    -o "/go-work/clients/$l"
>> done
>> ```

# 十、Nignx静态资源配置

自行根据资源位置更改root路径
>```
>#静态资源配置-图片
>location ~ /storage/.*\.(jpg|png|jpeg|gif|ico|svg|webp|bmp|swf)$ 
>{
>   root /www/wwwroot/workbench/workbench-server/;
>   expires      30d;
>   error_log /dev/null;
>   access_log /dev/null;
>}
>
>#静态资源配置-文件
>location ~ /storage/.*\.(js|css|json|txt)$
>{
>   root /www/wwwroot/workbench/workbench-server/;
>   expires      12h;
>   error_log /dev/null;
>   access_log /dev/null; 
>}
>    
>#静态文件夹配置（可不要）
>location /storage 
>{
>   root /www/wwwroot/workbench/workbench-server/;
>   #autoindex on;
>}
>```