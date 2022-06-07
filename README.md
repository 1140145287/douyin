# dousheng

## 4399　菜鸟驿站队

### 项目启动方式
```shell
go run main.go
```

底层服务：mysql、Redis<br>
ORM框架：Gorm、Gorm拓展soft-delete<br>
HTTP框架：Gin、Gin拓展Gin-Jwt<br>
配置管理：Viper<br>
日志管理：Zap<br>

### 代码结构

>configs　存放配置文件<br>
>controller　存储各个部分控制器<br>
>dao　数据访问层<br>
>global　全局变量<br>
>logic　存放实际业务逻辑<br>
>models　存放数据库结构体<br>
>pkg　项目相关模块包<br>
>
>>middleware　存放拦截器等中间件<br>
>>jwt　鉴权模块<br>
>>seting　全局信息相关<br>
>
>public　暂时存放视频，随后视频会自动传送到aliyun-oss<br>
>router　存放路由器<br>
>storage　项目生成的临时文件<br>

