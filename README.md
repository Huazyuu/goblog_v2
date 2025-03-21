# GO BLOG

## uml图

<img src="./doc/sql/uml.png" alt="uml"/>

## 项目结构

```
├─cmd # 运行文件
├─config
├─controller
│  ├─api
│  │  ├─adverts_api
│  │  ├─article_api
│  │  ├─images_api
│  │  ├─menu_api
│  │  ├─message_api
│  │  ├─settings_api
│  │  ├─tag_api
│  │  └─users_api
│  ├─req
│  └─res
├─core
├─doc
├─dump # 导出数据存放路径
│  ├─es
│  └─sql
├─flags # 命令行参数
├─global # 全局变量
├─middleware
│  └─jwt
├─models
│  ├─common
│  ├─diverseType
│  ├─esmodels
│  └─sqlmodels
├─plugins
│  ├─email_plugin
│  ├─freeimg
│  ├─script
│  │  └─cron
│  └─sync
├─repository
│  ├─advert_repo
│  ├─article_repo
│  ├─collect_repo
│  ├─img_repo
│  ├─menu_banner_repo
│  ├─menu_repo
│  ├─msg_repo
│  ├─tag_repo
│  └─user_repo
├─router
├─service
│  ├─advertsService
│  ├─articleService
│  ├─esService
│  │  └─indexService
│  ├─fileService
│  ├─menuService
│  ├─msgService
│  ├─redisService
│  ├─settingsService
│  ├─tagService
│  └─usersService
├─sql
├─uploads
│  └─avatar
└─utils
```

