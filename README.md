# GO BLOG

## uml图
<img src="./sql/uml.png" alt="uml"/>

## 项目结构
```
├─api # handler 层
│  ├─adverts_api
│  ├─images_api
│  ├─menu_api
│  ├─settings_api
│  └─users_api
├─config # 配置文件 
├─core # 加载配置
├─doc # swagger文档
├─docs # 其余文档
├─flags # 命令行参数
├─global # 全局参数 
├─middleware #中间件
│  └─jwt
├─models # 数据定义
│  ├─common
│  ├─diverseType
│  ├─esmodels
│  ├─req
│  ├─res
│  └─sqlmodels
├─plugins # 插件
│  ├─email_plugin
│  └─freeimg
├─repository # 数据库操作 
│  ├─advert_repo
│  ├─img_repo
│  ├─menu_banner_repo
│  ├─menu_repo
│  └─user_repo
├─router # 路由层
├─service # 服务层
│  ├─advertsService
│  ├─fileService
│  ├─menuService
│  ├─redisService
│  ├─settingsService
│  └─usersService
├─sql # 建表语句(若不使用gorm自动迁移,可直接建表)			
├─uploads # 上传文件存储
└─utils	# 工具

```

