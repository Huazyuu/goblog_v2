
# goblog_v2

Base URLs:

* <a href="http://dev-cn.your-api-server.com">开发环境: http://dev-cn.your-api-server.com</a>

* <a href="http://test-cn.your-api-server.com">测试环境: http://test-cn.your-api-server.com</a>

# Authentication

- HTTP Authentication, scheme: bearer

# settings

## GET 查看网站信息

GET /apiv1/settings/site

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {
    "created_at": "2025",
    "bei_an": "",
    "title": "yu_blog",
    "qq_image": "",
    "version": "v2",
    "email": "",
    "wechat_image": "",
    "name": "yu",
    "job": "",
    "addr": "bull",
    "slogan": "hello world",
    "slogan_en": "hello world",
    "web": "",
    "bilibili_url": "",
    "gitee_url": "",
    "github_url": ""
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» created_at|string|true|none||none|
|»» bei_an|string|true|none||none|
|»» title|string|true|none||none|
|»» qq_image|string|true|none||none|
|»» version|string|true|none||none|
|»» email|string|true|none||none|
|»» wechat_image|string|true|none||none|
|»» name|string|true|none||none|
|»» job|string|true|none||none|
|»» addr|string|true|none||none|
|»» slogan|string|true|none||none|
|»» slogan_en|string|true|none||none|
|»» web|string|true|none||none|
|»» bilibili_url|string|true|none||none|
|»» gitee_url|string|true|none||none|
|»» github_url|string|true|none||none|
|» msg|string|true|none||none|

## PUT 更新网站信息

PUT /apiv1/settings/site

> Body 请求参数

```json
{
  "created_at": "",
  "bei_an": "",
  "title": "",
  "qq_image": "",
  "version": "",
  "email": "",
  "wechat_image": "",
  "name": "",
  "job": "",
  "addr": "",
  "slogan": "",
  "slogan_en": "",
  "web": "",
  "bilibili_url": "",
  "gitee_url": "",
  "github_url": ""
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» created_at|body|string| 否 |none|
|» bei_an|body|string| 否 |none|
|» title|body|string| 否 |none|
|» qq_image|body|string| 否 |none|
|» version|body|string| 否 |none|
|» email|body|string| 否 |none|
|» wechat_image|body|string| 否 |none|
|» name|body|string| 否 |none|
|» job|body|string| 否 |none|
|» addr|body|string| 否 |none|
|» slogan|body|string| 否 |none|
|» slogan_en|body|string| 否 |none|
|» web|body|string| 否 |none|
|» bilibili_url|body|string| 否 |none|
|» gitee_url|body|string| 否 |none|
|» github_url|body|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {},
  "msg": "网站信息更新成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|» msg|string|true|none||none|

## GET 查看某一项设置信息

GET /apiv1/settings/{name}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|name|path|string| 是 |email qq qiniu jwt chat_group gaode|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {},
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|» msg|string|true|none||none|

## PUT 更新某一项设置

PUT /apiv1/settings/{name}

> Body 请求参数

```json
{
  "host": "",
  "port": 0,
  "user": "",
  "password": "",
  "default_from_email": "",
  "use_ssl": null,
  "user_tls": null
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|name|path|string| 是 |email qq qiniu jwt chat_group gaode|
|body|body|object| 否 |none|
|» host|body|string| 是 |none|
|» port|body|integer| 是 |none|
|» user|body|string| 是 |none|
|» password|body|string| 是 |none|
|» default_from_email|body|string| 是 |none|
|» use_ssl|body|null| 是 |none|
|» user_tls|body|null| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {},
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|» msg|string|true|none||none|

# user

## POST 注册

POST /apiv1/users/register

> Body 请求参数

```json
{
  "nick_name": "",
  "user_name": "",
  "password": "",
  "role": null
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» nick_name|body|string| 是 |none|
|» user_name|body|string| 是 |none|
|» password|body|string| 是 |none|
|» role|body|integer| 是 |none|

> 返回示例

```json
{
  "code": 0,
  "data": {},
  "msg": "用户testuser1创建成功!"
}
```

```json
{
  "code": 7,
  "data": {},
  "msg": "用户名已存在"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|» msg|string|true|none||none|

## POST 登录

POST /apiv1/users/login

> Body 请求参数

```json
{
  "user_name": "",
  "password": ""
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» user_name|body|string| 是 |none|
|» password|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwibmlja19uYW1lIjoiYWRtaW4iLCJyb2xlIjoxLCJ1c2VyX2lkIjo1LCJleHAiOjg0NzY4OTc5NC4zNzQyNTQsImlzcyI6Inp5dXUifQ.diQaCeJ2zUCYYLNE9fFvwg0Erclb2fzsAm96igvC9F4",
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|string|true|none||none|
|» msg|string|true|none||none|

## POST 注销

POST /apiv1/users/logout

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {},
  "msg": "注销成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|» msg|string|true|none||none|

## POST 绑定邮箱

POST /apiv1/users/bind_email

> Body 请求参数

```json
{
  "email": "",
  "code": null,
  "password": ""
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» email|body|string| 是 |none|
|» code|body|integer| 否 |none|
|» password|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {},
  "msg": "验证码已发送"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|» msg|string|true|none||none|

## GET 用户列表

GET /apiv1/users

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page|query|string| 否 |页数|
|key|query|string| 否 |关键词 nick_name user_name查找|
|limit|query|string| 否 |每页条数|
|sort|query|string| 否 |查询排序|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {
    "count": 1,
    "list": [
      {
        "id": 5,
        "created_at": "2025-03-05T20:32:13.637+08:00",
        "nick_name": "admin",
        "user_name": "admin",
        "avatar": "/uploads/avatar/default.png",
        "email": "z****@gmail.com",
        "tel": "",
        "addr": "内网地址",
        "token": "",
        "ip": "127.0.0.1",
        "role": "管理员",
        "sign_status": "邮箱",
        "integral": 0,
        "sign": "",
        "link": "",
        "role_id": 1
      }
    ]
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» count|integer|true|none||none|
|»» list|[object]|true|none||none|
|»»» id|integer|false|none||none|
|»»» created_at|string|false|none||none|
|»»» nick_name|string|false|none||none|
|»»» user_name|string|false|none||none|
|»»» avatar|string|false|none||none|
|»»» email|string|false|none||none|
|»»» tel|string|false|none||none|
|»»» addr|string|false|none||none|
|»»» token|string|false|none||none|
|»»» ip|string|false|none||none|
|»»» role|string|false|none||none|
|»»» sign_status|string|false|none||none|
|»»» integral|integer|false|none||none|
|»»» sign|string|false|none||none|
|»»» link|string|false|none||none|
|»»» role_id|integer|false|none||none|
|» msg|string|true|none||none|

## DELETE 批量删除用户

DELETE /apiv1/users

> Body 请求参数

```json
{
  "id_list": [
    "string"
  ]
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» id_list|body|[string]| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {},
  "msg": "删除成功 共删除 1 个用户"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|» msg|string|true|none||none|

## GET 用户信息

GET /apiv1/user_info

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {
    "id": 5,
    "created_at": "2025-03-05T20:32:13.637+08:00",
    "nick_name": "admin",
    "user_name": "admin",
    "avatar": "/uploads/avatar/default.png",
    "email": "zyuuforyu@gmail.com",
    "tel": "",
    "addr": "内网地址",
    "token": "",
    "ip": "127.0.0.1",
    "role": "管理员",
    "sign_status": "邮箱",
    "integral": 0,
    "sign": "",
    "link": ""
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» id|integer|true|none||none|
|»» created_at|string|true|none||none|
|»» nick_name|string|true|none||none|
|»» user_name|string|true|none||none|
|»» avatar|string|true|none||none|
|»» email|string|true|none||none|
|»» tel|string|true|none||none|
|»» addr|string|true|none||none|
|»» token|string|true|none||none|
|»» ip|string|true|none||none|
|»» role|string|true|none||none|
|»» sign_status|string|true|none||none|
|»» integral|integer|true|none||none|
|»» sign|string|true|none||none|
|»» link|string|true|none||none|
|» msg|string|true|none||none|

## PUT 修改用户信息

PUT /apiv1/user_info

> Body 请求参数

```json
{
  "role": null,
  "nick_name": "",
  "user_id": 0
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» nick_name|body|string| 是 |none|
|» sign|body|string| 是 |none|
|» link|body|string| 是 |none|
|» avatar|body|string| 是 |none|

> 返回示例

```json
{
  "code": 7,
  "data": {},
  "msg": "修改密码成功"
}
```

```json
{
  "code": 7,
  "data": {},
  "msg": "密码错误"
}
```

```json
{
  "code": 7,
  "data": {},
  "msg": "用户不存在 id错误"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|» msg|string|true|none||none|

## PUT 管理员修改用户 昵称 权限

PUT /apiv1/user_role

> Body 请求参数

```json
{
  "role": null,
  "nick_name": "",
  "user_id": 0
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» role|body|null| 是 |none|
|» nick_name|body|string| 是 |none|
|» user_id|body|integer| 是 |none|

> 返回示例

```json
{
  "code": 7,
  "data": {},
  "msg": "修改成功"
}
```

```json
{
  "code": 7,
  "data": {},
  "msg": "权限参数错误"
}
```

```json
{
  "code": 7,
  "data": {},
  "msg": "用户不存在 id错误"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|» msg|string|true|none||none|

## PUT 修改密码

PUT /apiv1/user_password

> Body 请求参数

```json
{
  "role": null,
  "nick_name": "",
  "user_id": 0
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» old_pwd|body|string| 是 |none|
|» pwd|body|string| 是 |none|

> 返回示例

```json
{
  "code": 7,
  "data": {},
  "msg": "修改密码成功"
}
```

```json
{
  "code": 7,
  "data": {},
  "msg": "密码错误"
}
```

```json
{
  "code": 7,
  "data": {},
  "msg": "用户不存在 id错误"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|» msg|string|true|none||none|

# img

## POST 上传图片

POST /apiv1/image

> Body 请求参数

```yaml
file: file://E:\OneDrive -
  exsn4848\图片\Feedback\{5C7107D5-586E-499F-8907-338A8DFC167B}\Capture001.png

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» file|body|string(binary)| 否 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {
    "file_name": "uploads/file/admin/屏幕截图_20250218_094046_duka.png",
    "is_success": true,
    "msg": "上传成功存储在:uploads/file/admin/屏幕截图_20250218_094046_duka.png"
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» file_name|string|true|none||none|
|»» is_success|boolean|true|none||none|
|»» msg|string|true|none||none|
|» msg|string|true|none||none|

## POST 上传图片到freeimg

POST /apiv1/freeimage

> Body 请求参数

```yaml
file: file://E:\OneDrive -
  exsn4848\图片\Feedback\{5C7107D5-586E-499F-8907-338A8DFC167B}\Capture001.png

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» file|body|string(binary)| 否 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {
    "file_name": "https://iili.io/3KYV4fa.md.png",
    "is_success": true,
    "msg": "图床上传成功存储在:https://iili.io/3KYV4fa.md.png"
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» file_name|string|true|none||none|
|»» is_success|boolean|true|none||none|
|»» msg|string|true|none||none|
|» msg|string|true|none||none|

## POST 上传多个图片

POST /apiv1/images

> Body 请求参数

```yaml
images:
  - file://E:\OneDrive - exsn4848\图片\屏幕截图\屏幕截图_20241213_095838.png
  - file://E:\OneDrive - exsn4848\图片\屏幕截图\屏幕截图_20241213_095807.png
  - file://E:\OneDrive - exsn4848\图片\屏幕截图\屏幕截图_20241213_094939.png

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» images|body|string(binary)| 否 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": [
    {
      "file_name": "uploads/file/屏幕截图_20241220_202459.png",
      "is_success": true,
      "msg": "图片上传成功"
    },
    {
      "file_name": "uploads/file/屏幕截图_20241220_202332.png",
      "is_success": true,
      "msg": "图片上传成功"
    },
    {
      "file_name": "uploads/file/屏幕截图_20241220_174643.png",
      "is_success": true,
      "msg": "图片上传成功"
    },
    {
      "file_name": "uploads/file/屏幕截图_20241220_124340.png",
      "is_success": true,
      "msg": "图片上传成功"
    }
  ],
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|[object]|true|none||none|
|»» file_name|string|true|none||none|
|»» is_success|boolean|true|none||none|
|»» msg|string|true|none||none|
|» msg|string|true|none||none|

## GET 图片列表

GET /apiv1/images

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page|query|string| 否 |页数|
|limit|query|string| 否 |每页条数|
|sort|query|string| 否 |查询排序|
|key|query|string| 是 |name path|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {
    "count": 5,
    "list": [
      {
        "id": 81,
        "created_at": "2025-03-10T16:47:42.112+08:00",
        "path": "https://iili.io/3f4XLZJ.md.png",
        "hash": "b803649397468bcc33ed1ece57b3c716",
        "name": "屏幕截图_20241220_124340_usol.png",
        "image_type": "远程图床",
        "bannerCount": 0,
        "articleCount": 2
      },
      {
        "id": 80,
        "created_at": "2025-03-10T16:47:35.435+08:00",
        "path": "uploads/file/admin/屏幕截图_20241225_134442_odip.png",
        "hash": "68ee29431e31d07224ec6109a6781f8a",
        "name": "屏幕截图_20241225_134442_odip.png",
        "image_type": "本地",
        "bannerCount": 0,
        "articleCount": 4
      },
      {
        "id": 79,
        "created_at": "2025-03-10T16:47:35.423+08:00",
        "path": "uploads/file/admin/屏幕截图_20241225_134903_nefi.png",
        "hash": "69ee63ebb39f56f2dcc6475263daf468",
        "name": "屏幕截图_20241225_134903_nefi.png",
        "image_type": "本地",
        "bannerCount": 1,
        "articleCount": 1
      },
      {
        "id": 78,
        "created_at": "2025-03-10T16:47:30.652+08:00",
        "path": "uploads/file/admin/屏幕截图_20241225_140629_duka.png",
        "hash": "463f523eaa20d2674ca404aa819e3838",
        "name": "屏幕截图_20241225_140629_duka.png",
        "image_type": "本地",
        "bannerCount": 1,
        "articleCount": 0
      },
      {
        "id": 64,
        "created_at": "2025-03-07T14:49:55.491+08:00",
        "path": "uploads/file/admin/屏幕截图_20241226_191134_nefi.png",
        "hash": "005d5313acad5ef5eb0a70dfcf301c26",
        "name": "屏幕截图_20241226_191134_nefi.png",
        "image_type": "本地",
        "bannerCount": 1,
        "articleCount": 1
      }
    ]
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» count|integer|true|none||none|
|»» list|[object]|true|none||none|
|»»» id|integer|true|none||none|
|»»» created_at|string|true|none||none|
|»»» path|string|true|none||none|
|»»» hash|string|true|none||none|
|»»» name|string|true|none||none|
|»»» image_type|string|true|none||none|
|»»» bannerCount|integer|true|none||none|
|»»» articleCount|integer|true|none||none|
|» msg|string|true|none||none|

## PUT 根据id修改图片名称

PUT /apiv1/images

> Body 请求参数

```json
{
  "role": null,
  "nick_name": "",
  "user_id": 0
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» id|body|string| 是 |none|
|» name|body|string| 是 |none|

> 返回示例

```json
{
  "code": 7,
  "data": {},
  "msg": "修改密码成功"
}
```

```json
{
  "code": 7,
  "data": {},
  "msg": "密码错误"
}
```

```json
{
  "code": 7,
  "data": {},
  "msg": "用户不存在 id错误"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|» msg|string|true|none||none|

## DELETE 批量删除图片

DELETE /apiv1/images

> Body 请求参数

```json
{
  "id_list": [
    "string"
  ]
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» id_list|body|[string]| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {},
  "msg": "删除成功 共删除 2 张图片"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|» msg|string|true|none||none|

## GET 图片名称列表

GET /apiv1/image_names

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {
    "count": 7,
    "list": [
      {
        "id": 63,
        "path": "uploads/file/admin/屏幕截图_20241226_191223_eyou.png",
        "name": "屏幕截图_20241226_191223_eyou.png"
      },
      {
        "id": 64,
        "path": "uploads/file/admin/屏幕截图_20241226_191134_nefi.png",
        "name": "屏幕截图_20241226_191134_nefi.png"
      },
      {
        "id": 65,
        "path": "https://iili.io/3KYV4fa.md.png",
        "name": "测试图片.png"
      },
      {
        "id": 78,
        "path": "uploads/file/admin/屏幕截图_20241225_140629_duka.png",
        "name": "屏幕截图_20241225_140629_duka.png"
      },
      {
        "id": 79,
        "path": "uploads/file/admin/屏幕截图_20241225_134903_nefi.png",
        "name": "屏幕截图_20241225_134903_nefi.png"
      },
      {
        "id": 80,
        "path": "uploads/file/admin/屏幕截图_20241225_134442_odip.png",
        "name": "屏幕截图_20241225_134442_odip.png"
      },
      {
        "id": 81,
        "path": "https://iili.io/3f4XLZJ.md.png",
        "name": "屏幕截图_20241220_124340_usol.png"
      }
    ]
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» count|integer|true|none||none|
|»» list|[object]|true|none||none|
|»»» id|integer|true|none||none|
|»»» path|string|true|none||none|
|»»» name|string|true|none||none|
|» msg|string|true|none||none|

# advert

## POST 添加广告

POST /apiv1/adverts

> Body 请求参数

```json
{
  "title": "",
  "href": "",
  "images": "",
  "is_show": null
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» title|body|string| 是 |none|
|» href|body|string| 是 |none|
|» images|body|string| 是 |none|
|» is_show|body|boolean| 否 |none|

> 返回示例

```json
{
  "code": 7,
  "data": {},
  "msg": "广告已存在"
}
```

```json
{
  "code": 0,
  "data": {},
  "msg": "增加广告成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|» msg|string|true|none||none|

## GET 广告列表

GET /apiv1/adverts

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page|query|string| 否 |页数|
|limit|query|string| 否 |每页条数|
|sort|query|string| 否 |查询排序|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {
    "count": 1,
    "list": [
      {
        "id": 1,
        "created_at": "2025-03-07T16:16:54.732+08:00",
        "title": "广告一号",
        "href": "https://baidu.com",
        "images": "https://iili.io/3KYV4fa.md.png",
        "is_show": true
      }
    ]
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» count|integer|true|none||none|
|»» list|[object]|true|none||none|
|»»» id|integer|false|none||none|
|»»» created_at|string|false|none||none|
|»»» title|string|false|none||none|
|»»» href|string|false|none||none|
|»»» images|string|false|none||none|
|»»» is_show|boolean|false|none||none|
|»»» nick_name|string|false|none||none|
|»»» user_name|string|false|none||none|
|»»» avatar|string|false|none||none|
|»»» email|string|false|none||none|
|»»» tel|string|false|none||none|
|»»» addr|string|false|none||none|
|»»» token|string|false|none||none|
|»»» ip|string|false|none||none|
|»»» role|string|false|none||none|
|»»» sign_status|string|false|none||none|
|»»» integral|integer|false|none||none|
|»»» sign|string|false|none||none|
|»»» link|string|false|none||none|
|»»» role_id|integer|false|none||none|
|» msg|string|true|none||none|

## DELETE 批量删除广告

DELETE /apiv1/adverts

> Body 请求参数

```json
{
  "id_list": [
    "string"
  ]
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» id_list|body|[string]| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {},
  "msg": "删除成功 共删除 2 条广告"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|» msg|string|true|none||none|

## PUT 修改广告

PUT /apiv1/adverts/{id}

> Body 请求参数

```json
{
  "title": "广告三号",
  "href": "https://baidu.com",
  "images": "https://iili.io/3KYV4fa.md.png",
  "is_show": true
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string| 是 |none|
|body|body|object| 否 |none|
|» title|body|string| 是 |none|
|» href|body|string| 是 |none|
|» images|body|string| 是 |none|
|» is_show|body|boolean| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {},
  "msg": "修改成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|» msg|string|true|none||none|

# menu

## POST 创建菜单

POST /apiv1/menus

> Body 请求参数

```json
{
  "title": "",
  "path": "",
  "slogan": "",
  "abstract": null,
  "abstract_time": 0,
  "banner_time": 0,
  "sort": 0,
  "image_sort_list": [
    null
  ],
  "image_id": 0
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» title|body|string| 是 |none|
|» path|body|string| 是 |none|
|» slogan|body|string| 否 |none|
|» abstract|body|[string]| 否 |none|
|» abstract_time|body|integer| 否 |none|
|» banner_time|body|integer| 否 |none|
|» sort|body|integer| 否 |none|
|» image_sort_list|body|[object]| 否 |none|

> 返回示例

> 200 Response

```json
{
  "code": 7,
  "data": {},
  "msg": "添加菜单成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|» msg|string|true|none||none|

## GET 菜单列表

GET /apiv1/menus

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {
    "count": 2,
    "list": [
      {
        "id": 7,
        "created_at": "2025-03-10T13:51:16.071+08:00",
        "title": "test_title_add1",
        "path": "test_path_add1",
        "slogan": "test_slogan",
        "abstract": [
          "test_abstract1",
          "test_abstract2",
          "test_abstract3"
        ],
        "abstract_time": 7,
        "banner_time": 7,
        "sort": 1,
        "banners": [
          {
            "id": 64,
            "path": "uploads/file/admin/屏幕截图_20241226_191134_nefi.png"
          }
        ]
      },
      {
        "id": 11,
        "created_at": "2025-03-10T20:20:26.671+08:00",
        "title": "test_title_add2",
        "path": "test_path_add2",
        "slogan": "test_slogan",
        "abstract": [
          "test_abstract1",
          "test_abstract2",
          "test_abstract3"
        ],
        "abstract_time": 7,
        "banner_time": 7,
        "sort": 1,
        "banners": [
          {
            "id": 79,
            "path": "uploads/file/admin/屏幕截图_20241225_134903_nefi.png"
          },
          {
            "id": 78,
            "path": "uploads/file/admin/屏幕截图_20241225_140629_duka.png"
          },
          {
            "id": 64,
            "path": "uploads/file/admin/屏幕截图_20241226_191134_nefi.png"
          }
        ]
      }
    ]
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» count|integer|true|none||none|
|»» list|[object]|true|none||none|
|»»» id|integer|true|none||none|
|»»» created_at|string|true|none||none|
|»»» title|string|true|none||none|
|»»» path|string|true|none||none|
|»»» slogan|string|true|none||none|
|»»» abstract|[string]|true|none||none|
|»»» abstract_time|integer|true|none||none|
|»»» banner_time|integer|true|none||none|
|»»» sort|integer|true|none||none|
|»»» banners|[object]|true|none||none|
|»»»» id|integer|true|none||none|
|»»»» path|string|true|none||none|
|» msg|string|true|none||none|

## DELETE 批量删除菜单

DELETE /apiv1/menus

> Body 请求参数

```json
{
  "id_list": [
    "string"
  ]
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» id_list|body|[string]| 是 |none|

> 返回示例

```json
{
  "code": 0,
  "data": {},
  "msg": "删除成功 共删除 2 条菜单"
}
```

```json
{
  "code": 7,
  "data": {},
  "msg": "菜单不存在"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|» msg|string|true|none||none|

## PUT 更新菜单

PUT /apiv1/menus/{id}

> Body 请求参数

```json
{
  "title": "",
  "path": "",
  "slogan": "",
  "abstract": null,
  "abstract_time": 0,
  "banner_time": 0,
  "sort": 0,
  "image_sort_list": [
    null
  ],
  "image_id": 0
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string| 是 |none|
|body|body|object| 否 |none|
|» title|body|string| 是 |none|
|» path|body|string| 是 |none|
|» slogan|body|string| 否 |none|
|» abstract|body|[string]| 否 |none|
|» abstract_time|body|integer| 否 |none|
|» banner_time|body|integer| 否 |none|
|» sort|body|integer| 否 |none|
|» image_sort_list|body|[object]| 否 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {},
  "msg": "菜单更新成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|» msg|string|true|none||none|

## GET 菜单详细

GET /apiv1/menus/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {
    "id": 7,
    "created_at": "2025-03-10T13:51:16.071+08:00",
    "title": "test_title_add1",
    "path": "test_path_add1",
    "slogan": "test_slogan",
    "abstract": [
      "test_abstract1",
      "test_abstract2",
      "test_abstract3"
    ],
    "abstract_time": 7,
    "banner_time": 7,
    "sort": 1,
    "banners": [
      {
        "id": 64,
        "path": "uploads/file/admin/屏幕截图_20241226_191134_nefi.png"
      }
    ]
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» id|integer|true|none||none|
|»» created_at|string|true|none||none|
|»» title|string|true|none||none|
|»» path|string|true|none||none|
|»» slogan|string|true|none||none|
|»» abstract|[string]|true|none||none|
|»» abstract_time|integer|true|none||none|
|»» banner_time|integer|true|none||none|
|»» sort|integer|true|none||none|
|»» banners|[object]|true|none||none|
|»»» id|integer|false|none||none|
|»»» path|string|false|none||none|
|» msg|string|true|none||none|

## GET 菜单名称列表

GET /apiv1/menu_names

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": [
    {
      "id": 7,
      "title": "test_title_add1",
      "path": "test_path_add1"
    },
    {
      "id": 11,
      "title": "test_title_add2",
      "path": "test_path_add2"
    }
  ],
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|[object]|true|none||none|
|»» id|integer|true|none||none|
|»» title|string|true|none||none|
|»» path|string|true|none||none|
|» msg|string|true|none||none|

## GET 菜单详细(通过path查询)

GET /apiv1/menus/detail

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|path|query|string| 否 |none|

> 返回示例

```json
{
  "code": 7,
  "data": {},
  "msg": "record not found"
}
```

```json
{
  "code": 0,
  "data": {
    "id": 7,
    "created_at": "2025-03-10T13:51:16.071+08:00",
    "title": "test_title_add1",
    "path": "test_path_add1",
    "slogan": "test_slogan",
    "abstract": [
      "test_abstract1",
      "test_abstract2",
      "test_abstract3"
    ],
    "abstract_time": 7,
    "banner_time": 7,
    "sort": 1,
    "banners": [
      {
        "id": 64,
        "path": "uploads/file/admin/屏幕截图_20241226_191134_nefi.png"
      }
    ]
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» id|integer|true|none||none|
|»» created_at|string|true|none||none|
|»» title|string|true|none||none|
|»» path|string|true|none||none|
|»» slogan|string|true|none||none|
|»» abstract|[string]|true|none||none|
|»» abstract_time|integer|true|none||none|
|»» banner_time|integer|true|none||none|
|»» sort|integer|true|none||none|
|»» banners|[object]|true|none||none|
|»»» id|integer|false|none||none|
|»»» path|string|false|none||none|
|» msg|string|true|none||none|

# tag

## POST 创建标签

POST /apiv1/tags

> Body 请求参数

```json
{
  "title": "string"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» title|body|string| 是 |none|

> 返回示例

```json
{
  "code": 7,
  "data": {},
  "msg": "标签已存在"
}
```

```json
{
  "code": 0,
  "data": {},
  "msg": "创建标签 tag02 成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|» msg|string|true|none||none|

## GET 标签列表

GET /apiv1/tags

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page|query|string| 否 |1|
|key|query|string| 否 |none|
|limit|query|string| 否 |none|
|sort|query|string| 否 |created_at desc|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {
    "count": 2,
    "list": [
      {
        "id": 2,
        "created_at": "2025-03-11T14:23:16.989+08:00",
        "title": "tag02"
      },
      {
        "id": 1,
        "created_at": "2025-03-11T14:22:49.11+08:00",
        "title": "tag01"
      }
    ]
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» count|integer|true|none||none|
|»» list|[object]|true|none||none|
|»»» id|integer|true|none||none|
|»»» created_at|string|true|none||none|
|»»» title|string|true|none||none|
|» msg|string|true|none||none|

## DELETE 批量删除标签

DELETE /apiv1/tags

> Body 请求参数

```json
{
  "id_list": [
    "string"
  ]
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» id_list|body|[string]| 是 |none|

> 返回示例

```json
{
  "code": 0,
  "data": {},
  "msg": "删除成功 共删除 2 个标签"
}
```

```json
{
  "code": 7,
  "data": {},
  "msg": "标签不存在"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|» msg|string|true|none||none|

## GET 标签名称数量列表

GET /apiv1/tag_names

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": [
    {
      "tag_name": "test",
      "count": "6"
    },
    {
      "tag_name": "go",
      "count": "2"
    },
    {
      "tag_name": "backend",
      "count": "1"
    }
  ],
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|[object]|true|none||none|
|»» tag_name|string|true|none||none|
|»» count|string|true|none||none|
|» msg|string|true|none||none|

## PUT 更新标签

PUT /apiv1/tags/{id}

> Body 请求参数

```json
{
  "title": "string"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string| 是 |none|
|body|body|object| 否 |none|
|» title|body|string| 是 |none|

> 返回示例

```json
{
  "code": 0,
  "data": {},
  "msg": "更新标签成功"
}
```

```json
{
  "code": 7,
  "data": {},
  "msg": "标签不存在"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|» msg|string|true|none||none|

# message

## POST 创建消息

POST /apiv1/messages

> Body 请求参数

```json
{
  "rev_user_id": 0,
  "content": ""
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» rev_user_id|body|integer| 是 |none|
|» content|body|string| 是 |none|

> 返回示例

```json
{
  "code": 7,
  "data": {},
  "msg": "接收者不存在"
}
```

```json
{
  "code": 0,
  "data": {},
  "msg": "发送消息成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|» msg|string|true|none||none|

## GET 与自己相关消息列表

GET /apiv1/messages

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": [
    {
      "send_user_id": 5,
      "send_user_nick_name": "管理员用户",
      "send_user_avatar": "https://ts1.tc.mm.bing.net/th/id/R-C.26fa5434823e0afae3f9b576b61b3df0?rik=1ki5rrqJXLS00w&riu=http%3a%2f%2fpic.52112.com%2f180420%2f180420_32%2fJ9xjxe1jIg_small.jpg&ehk=a8hQQlllEncpFeXgnFZ1a7fIII7lcz2ph6WLdtzS51k%3d&risl=&pid=ImgRaw&r=0",
      "rev_user_id": 6,
      "rev_user_nick_name": "测试用户",
      "rev_user_avatar": "https://ts1.tc.mm.bing.net/th/id/R-C.26fa5434823e0afae3f9b576b61b3df0?rik=1ki5rrqJXLS00w&riu=http%3a%2f%2fpic.52112.com%2f180420%2f180420_32%2fJ9xjxe1jIg_small.jpg&ehk=a8hQQlllEncpFeXgnFZ1a7fIII7lcz2ph6WLdtzS51k%3d&risl=&pid=ImgRaw&r=0",
      "content": "湿润的的西北菜炖菜",
      "created_at": "2025-03-11T16:38:25.23+08:00",
      "message_count": 7
    },
    {
      "send_user_id": 5,
      "send_user_nick_name": "管理员用户",
      "send_user_avatar": "https://ts1.tc.mm.bing.net/th/id/R-C.26fa5434823e0afae3f9b576b61b3df0?rik=1ki5rrqJXLS00w&riu=http%3a%2f%2fpic.52112.com%2f180420%2f180420_32%2fJ9xjxe1jIg_small.jpg&ehk=a8hQQlllEncpFeXgnFZ1a7fIII7lcz2ph6WLdtzS51k%3d&risl=&pid=ImgRaw&r=0",
      "rev_user_id": 7,
      "rev_user_nick_name": "testuser3",
      "rev_user_avatar": "/uploads/avatar/default.png",
      "content": "香菇青菜",
      "created_at": "2025-03-11T16:42:46.884+08:00",
      "message_count": 4
    },
    {
      "send_user_id": 5,
      "send_user_nick_name": "管理员用户",
      "send_user_avatar": "https://ts1.tc.mm.bing.net/th/id/R-C.26fa5434823e0afae3f9b576b61b3df0?rik=1ki5rrqJXLS00w&riu=http%3a%2f%2fpic.52112.com%2f180420%2f180420_32%2fJ9xjxe1jIg_small.jpg&ehk=a8hQQlllEncpFeXgnFZ1a7fIII7lcz2ph6WLdtzS51k%3d&risl=&pid=ImgRaw&r=0",
      "rev_user_id": 8,
      "rev_user_nick_name": "testuser4",
      "rev_user_avatar": "/uploads/avatar/default.png",
      "content": "叉烧包",
      "created_at": "2025-03-11T16:42:52.883+08:00",
      "message_count": 4
    }
  ],
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|[object]|true|none||none|
|»» send_user_id|integer|true|none||none|
|»» send_user_nick_name|string|true|none||none|
|»» send_user_avatar|string|true|none||none|
|»» rev_user_id|integer|true|none||none|
|»» rev_user_nick_name|string|true|none||none|
|»» rev_user_avatar|string|true|none||none|
|»» content|string|true|none||none|
|»» created_at|string|true|none||none|
|»» message_count|integer|true|none||none|
|» msg|string|true|none||none|

## DELETE 批量删除消息

DELETE /apiv1/messages

> Body 请求参数

```json
{
  "id_list": [
    "string"
  ]
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» id_list|body|[string]| 是 |none|

> 返回示例

```json
{
  "code": 0,
  "data": {},
  "msg": "删除成功 共删除 1 个消息"
}
```

```json
{
  "code": 7,
  "data": {},
  "msg": "消息不存在"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|» msg|string|true|none||none|

## GET 消息列表

GET /apiv1/messages_all

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page|query|string| 否 |1|
|key|query|string| 否 |none|
|limit|query|string| 否 |none|
|sort|query|string| 否 |created_at desc|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {
    "count": 1,
    "list": [
      {
        "id": 2,
        "created_at": "2025-03-11T15:50:28.693+08:00",
        "send_user_id": 5,
        "send_user_nick_name": "管理员用户",
        "send_user_avatar": "https://ts1.tc.mm.bing.net/th/id/R-C.26fa5434823e0afae3f9b576b61b3df0?rik=1ki5rrqJXLS00w&riu=http%3a%2f%2fpic.52112.com%2f180420%2f180420_32%2fJ9xjxe1jIg_small.jpg&ehk=a8hQQlllEncpFeXgnFZ1a7fIII7lcz2ph6WLdtzS51k%3d&risl=&pid=ImgRaw&r=0",
        "rev_user_id": 6,
        "rev_user_nick_name": "测试用户",
        "rev_user_avatar": "https://ts1.tc.mm.bing.net/th/id/R-C.26fa5434823e0afae3f9b576b61b3df0?rik=1ki5rrqJXLS00w&riu=http%3a%2f%2fpic.52112.com%2f180420%2f180420_32%2fJ9xjxe1jIg_small.jpg&ehk=a8hQQlllEncpFeXgnFZ1a7fIII7lcz2ph6WLdtzS51k%3d&risl=&pid=ImgRaw&r=0",
        "is_read": false,
        "content": "测试消息"
      }
    ]
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» count|integer|true|none||none|
|»» list|[object]|true|none||none|
|»»» id|integer|true|none||none|
|»»» created_at|string|true|none||none|
|»»» title|string|true|none||none|
|» msg|string|true|none||none|

## GET 两个用户间的聊天记录

GET /apiv1/message_users/record

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page|query|string| 否 |1|
|limit|query|string| 否 |none|
|sort|query|string| 否 |created_at desc|
|sendUserID|query|string| 否 |none|
|revUserID|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {
    "count": 1,
    "list": [
      {
        "id": 18,
        "created_at": "0001-01-01T00:00:00Z",
        "send_user_id": 8,
        "send_user_nick_name": "testuser4",
        "send_user_avatar": "",
        "rev_user_id": 7,
        "rev_user_nick_name": "testuser3",
        "rev_user_avatar": "",
        "is_read": false,
        "content": ""
      }
    ]
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» count|integer|true|none||none|
|»» list|[object]|true|none||none|
|»»» id|integer|true|none||none|
|»»» created_at|string|true|none||none|
|»»» title|string|true|none||none|
|» msg|string|true|none||none|

## GET 有消息的用户列表

GET /apiv1/message_users

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page|query|string| 否 |1|
|limit|query|string| 否 |none|
|sort|query|string| 否 |created_at desc|
|nickename|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {
    "count": 3,
    "list": [
      {
        "userName": "admin",
        "nickName": "管理员用户",
        "userID": 5,
        "avatar": "https://ts1.tc.mm.bing.net/th/id/R-C.26fa5434823e0afae3f9b576b61b3df0?rik=1ki5rrqJXLS00w&riu=http%3a%2f%2fpic.52112.com%2f180420%2f180420_32%2fJ9xjxe1jIg_small.jpg&ehk=a8hQQlllEncpFeXgnFZ1a7fIII7lcz2ph6WLdtzS51k%3d&risl=&pid=ImgRaw&r=0",
        "count": 3
      },
      {
        "userName": "testuser",
        "nickName": "测试用户",
        "userID": 6,
        "avatar": "https://ts1.tc.mm.bing.net/th/id/R-C.26fa5434823e0afae3f9b576b61b3df0?rik=1ki5rrqJXLS00w&riu=http%3a%2f%2fpic.52112.com%2f180420%2f180420_32%2fJ9xjxe1jIg_small.jpg&ehk=a8hQQlllEncpFeXgnFZ1a7fIII7lcz2ph6WLdtzS51k%3d&risl=&pid=ImgRaw&r=0",
        "count": 1
      },
      {
        "userName": "zyuutestuser4",
        "nickName": "testuser4",
        "userID": 8,
        "avatar": "/uploads/avatar/default.png",
        "count": 1
      }
    ]
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» count|integer|true|none||none|
|»» list|[object]|true|none||none|
|»»» userName|string|true|none||none|
|»»» nickName|string|true|none||none|
|»»» userID|integer|true|none||none|
|»»» avatar|string|true|none||none|
|»»» count|integer|true|none||none|
|» msg|string|true|none||none|

## GET 我与某个用户的聊天记录

GET /apiv1/message_users/record/me

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page|query|string| 否 |1|
|limit|query|string| 否 |none|
|sort|query|string| 否 |created_at desc|
|userID|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {
    "count": 4,
    "list": [
      {
        "id": 13,
        "created_at": "2025-03-11T16:42:50.301+08:00",
        "send_user_id": 5,
        "send_user_nick_name": "管理员用户",
        "send_user_avatar": "https://ts1.tc.mm.bing.net/th/id/R-C.26fa5434823e0afae3f9b576b61b3df0?rik=1ki5rrqJXLS00w&riu=http%3a%2f%2fpic.52112.com%2f180420%2f180420_32%2fJ9xjxe1jIg_small.jpg&ehk=a8hQQlllEncpFeXgnFZ1a7fIII7lcz2ph6WLdtzS51k%3d&risl=&pid=ImgRaw&r=0",
        "rev_user_id": 8,
        "rev_user_nick_name": "testuser4",
        "rev_user_avatar": "/uploads/avatar/default.png",
        "is_read": false,
        "content": "茄子炒肉"
      }
    ]
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» count|integer|true|none||none|
|»» list|[object]|true|none||none|
|»»» userName|string|true|none||none|
|»»» nickName|string|true|none||none|
|»»» userID|integer|true|none||none|
|»»» avatar|string|true|none||none|
|»»» count|integer|true|none||none|
|» msg|string|true|none||none|

## GET 某个用户的消息列表

GET /apiv1/message_users/user

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|userID|query|string| 否 |none|

> 返回示例

```json
{
  "code": 0,
  "data": {
    "count": 3,
    "list": [
      {
        "userName": "testuser",
        "nickName": "测试用户",
        "userID": 6,
        "avatar": "https://ts1.tc.mm.bing.net/th/id/R-C.26fa5434823e0afae3f9b576b61b3df0?rik=1ki5rrqJXLS00w&riu=http%3a%2f%2fpic.52112.com%2f180420%2f180420_32%2fJ9xjxe1jIg_small.jpg&ehk=a8hQQlllEncpFeXgnFZ1a7fIII7lcz2ph6WLdtzS51k%3d&risl=&pid=ImgRaw&r=0",
        "count": 8
      },
      {
        "userName": "zyuutestuser3",
        "nickName": "testuser3",
        "userID": 7,
        "avatar": "/uploads/avatar/default.png",
        "count": 4
      },
      {
        "userName": "zyuutestuser4",
        "nickName": "testuser4",
        "userID": 8,
        "avatar": "/uploads/avatar/default.png",
        "count": 4
      }
    ]
  },
  "msg": "成功"
}
```

```json
{
  "code": 7,
  "data": {},
  "msg": "查询失败,没有对应用户"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» count|integer|true|none||none|
|»» list|[object]|true|none||none|
|»»» userName|string|true|none||none|
|»»» nickName|string|true|none||none|
|»»» userID|integer|true|none||none|
|»»» avatar|string|true|none||none|
|»»» count|integer|true|none||none|
|» msg|string|true|none||none|

## GET 我的消息列表

GET /apiv1/message_users/me

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {
    "count": 3,
    "list": [
      {
        "userName": "zyuutestuser4",
        "nickName": "testuser4",
        "userID": 8,
        "avatar": "/uploads/avatar/default.png",
        "count": 4
      },
      {
        "userName": "testuser",
        "nickName": "测试用户",
        "userID": 6,
        "avatar": "https://ts1.tc.mm.bing.net/th/id/R-C.26fa5434823e0afae3f9b576b61b3df0?rik=1ki5rrqJXLS00w&riu=http%3a%2f%2fpic.52112.com%2f180420%2f180420_32%2fJ9xjxe1jIg_small.jpg&ehk=a8hQQlllEncpFeXgnFZ1a7fIII7lcz2ph6WLdtzS51k%3d&risl=&pid=ImgRaw&r=0",
        "count": 8
      },
      {
        "userName": "zyuutestuser3",
        "nickName": "testuser3",
        "userID": 7,
        "avatar": "/uploads/avatar/default.png",
        "count": 4
      }
    ]
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» count|integer|true|none||none|
|»» list|[object]|true|none||none|
|»»» userName|string|true|none||none|
|»»» nickName|string|true|none||none|
|»»» userID|integer|true|none||none|
|»»» avatar|string|true|none||none|
|»»» count|integer|true|none||none|
|» msg|string|true|none||none|

# article

## POST 创建文章

POST /apiv1/articles

{
    "title": "石蒜深度测试（含完整MD语法）333",
    "abstract": "包含代码块/列表/引用/图片的花卉百科测试",
    "content": "# 石蒜（Lycoris radiata）\n\n## 一、植物学特征\n### 1.1 形态描述\n- 多年生草本，鳞茎近球形（直径3-5cm）\n- 叶基生，条形，长约30cm，宽0.5-1cm\n> **注意**：花期无叶（先花后叶现象）\n\n## 二、栽培技术\n### 2.1 种植代码示例（Go/Python双语言）\n```go\n// 模拟鳞茎种植逻辑\nfunc PlantLycoris(bulbSize float64) bool {\n    return bulbSize > 2.5 && soilPH > 6.0\n}\n```\n```python\n# 温湿度监控脚本\ndef check_env(temp, humidity):\n    return 15 < temp < 25 and 60 < humidity < 80\n```\n\n### 2.2 种植步骤（有序列表）\n1. 准备疏松土壤（腐殖土:沙土=3:1）\n2. 鳞茎种植深度：顶部覆土3-5cm\n3. 首次浇水至**完全浸透**（注意排水！）\n\n## 三、文化象征\n![曼殊沙华](https://example.com/lycoris.jpg \"红色彼岸花\")\n> **佛经记载**：\"尔时世尊，四众围绕...雨曼殊沙华\"（《法华经》）\n- 花语：生死相隔（因花叶永不相见）\n- 别称：**彼岸花**、舍子花（日本称呼）\n\n## 四、注意事项\n⚠️ 全株有毒！误食会引起**呕吐/腹泻**（[毒性详情>>](https://botany.gov/poison/lycoris)\n",
    "category": "花卉/植物学",
    "source": "《中国植物志》第16卷",
    "link": "https://flora.cn/lycoris",
    "banner_id": 64,
    "tags": [
        "球根花卉",
        "有毒植物",
        "佛教植物",
        "先花后叶"
    ]
}

> Body 请求参数

```json
{
  "title": "",
  "abstract": "",
  "content": "",
  "category": "",
  "source": "",
  "link": "",
  "banner_id": 0,
  "tags": [
    ""
  ]
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» title|body|string| 是 |none|
|» abstract|body|string| 是 |none|
|» content|body|string| 是 |none|
|» category|body|string| 是 |none|
|» source|body|string| 是 |none|
|» link|body|string| 是 |none|
|» banner_id|body|integer| 是 |none|
|» tags|body|[string]| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {},
  "msg": "文章[gorm scan的使用]创建成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|» msg|string|true|none||none|

## PUT 更新文章

PUT /apiv1/articles

> Body 请求参数

```json
{
  "title": "",
  "abstract": "",
  "content": "",
  "category": "",
  "source": "",
  "link": "",
  "banner_id": 0,
  "tags": [
    ""
  ]
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» id|body|string| 是 |none|
|» title|body|string| 是 |none|
|» abstract|body|string| 是 |none|
|» content|body|string| 是 |none|
|» category|body|string| 是 |none|
|» source|body|string| 是 |none|
|» link|body|string| 是 |none|
|» banner_id|body|integer| 是 |none|
|» tags|body|[string]| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {},
  "msg": "文章[gorm scan的使用]创建成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|» msg|string|true|none||none|

## GET 文章列表

GET /apiv1/articles

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page|query|integer| 否 |页|
|key|query|string| 否 |关键词查询(精准)|
|limit|query|integer| 否 |数量|
|sort|query|string| 否 |created_at desc/asc|
|tags|query|string| 否 |标签|
|category|query|string| 否 |分组|
|is_user|query|boolean| 否 |是否本人|
|date|query|string| 否 |一天时间内发的文章|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {
    "count": 1,
    "list": [
      {
        "id": "UpI5jpUBS5cfAjmlRN3m",
        "created_at": "2025-03-13 14:37:52",
        "updated_at": "2025-03-13 14:37:52",
        "title": "<em>gin</em>的使用",
        "keyword": "gin的使用",
        "abstract": "ShouldBindUrI\n\n\n绑定位置：\n\n\n用于绑定请求URL中的路径参数（URI parameters）。例如，对于一个形如/user/:id的路由，其中:id就是路径参数。当客户端请求这个路由",
        "content": "#  ShouldBindUrI\n\n- **绑定位置**：\n\n  - 用于绑定请求URL中的路径参数（URI parameters）。例如，对于一个形如`/user/:id`的路由",
        "look_count": 0,
        "comment_count": 0,
        "digg_count": 0,
        "collects_count": 0,
        "user_id": 5,
        "user_nick_name": "管理员用户",
        "user_avatar": "https://ts1.tc.mm.bing.net/th/id/R-C.26fa5434823e0afae3f9b576b61b3df0?rik=1ki5rrqJXLS00w&riu=http%3a%2f%2fpic.52112.com%2f180420%2f180420_32%2fJ9xjxe1jIg_small.jpg&ehk=a8hQQlllEncpFeXgnFZ1a7fIII7lcz2ph6WLdtzS51k%3d&risl=&pid=ImgRaw&r=0",
        "category": "go",
        "source": "zyuu",
        "link": "",
        "banner_id": 64,
        "banner_url": "uploads/file/admin/屏幕截图_20241226_191134_nefi.png",
        "tags": [
          "go"
        ]
      }
    ]
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» count|integer|true|none||none|
|»» list|[object]|true|none||none|
|»»» id|string|false|none||none|
|»»» created_at|string|false|none||none|
|»»» updated_at|string|false|none||none|
|»»» title|string|false|none||none|
|»»» keyword|string|false|none||none|
|»»» abstract|string|false|none||none|
|»»» content|string|false|none||none|
|»»» look_count|integer|false|none||none|
|»»» comment_count|integer|false|none||none|
|»»» digg_count|integer|false|none||none|
|»»» collects_count|integer|false|none||none|
|»»» user_id|integer|false|none||none|
|»»» user_nick_name|string|false|none||none|
|»»» user_avatar|string|false|none||none|
|»»» category|string|false|none||none|
|»»» source|string|false|none||none|
|»»» link|string|false|none||none|
|»»» banner_id|integer|false|none||none|
|»»» banner_url|string|false|none||none|
|»»» tags|[string]|false|none||none|
|» msg|string|true|none||none|

## DELETE 批量删除文章

DELETE /apiv1/articles

> Body 请求参数

```json
{
  "id_list": [
    "string"
  ]
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» id_list|body|[string]| 是 |none|

> 返回示例

```json
{
  "code": 0,
  "data": {},
  "msg": "删除成功,共删除3篇文章"
}
```

```json
{
  "code": 7,
  "data": {},
  "msg": "消息不存在"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|» msg|string|true|none||none|

## GET 收藏列表

GET /apiv1/articles/collects

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page|query|integer| 否 |页|
|limit|query|integer| 否 |数量|
|sort|query|string| 否 |created_at desc/asc|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {
    "count": 1,
    "list": [
      {
        "id": "UpI5jpUBS5cfAjmlRN3m",
        "updated_at": "2025-03-13 14:37:52",
        "title": "gin的使用",
        "keyword": "gin的使用",
        "abstract": "ShouldBindUrI\n\n\n绑定位置：\n\n\n用于绑定请求URL中的路径参数（URI parameters）。例如，对于一个形如/user/:id的路由，其中:id就是路径参数。当客户端请求这个路由",
        "content": "",
        "look_count": 0,
        "comment_count": 0,
        "digg_count": 0,
        "collects_count": 0,
        "user_id": 5,
        "user_nick_name": "管理员用户",
        "user_avatar": "https://ts1.tc.mm.bing.net/th/id/R-C.26fa5434823e0afae3f9b576b61b3df0?rik=1ki5rrqJXLS00w&riu=http%3a%2f%2fpic.52112.com%2f180420%2f180420_32%2fJ9xjxe1jIg_small.jpg&ehk=a8hQQlllEncpFeXgnFZ1a7fIII7lcz2ph6WLdtzS51k%3d&risl=&pid=ImgRaw&r=0",
        "category": "go",
        "source": "zyuu",
        "link": "",
        "banner_id": 64,
        "banner_url": "uploads/file/admin/屏幕截图_20241226_191134_nefi.png",
        "tags": [
          "go"
        ],
        "created_at": "2025-03-15 15:30:03"
      }
    ]
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» count|integer|true|none||none|
|»» list|[object]|true|none||none|
|»»» id|string|false|none||none|
|»»» updated_at|string|false|none||none|
|»»» title|string|false|none||none|
|»»» keyword|string|false|none||none|
|»»» abstract|string|false|none||none|
|»»» content|string|false|none||none|
|»»» look_count|integer|false|none||none|
|»»» comment_count|integer|false|none||none|
|»»» digg_count|integer|false|none||none|
|»»» collects_count|integer|false|none||none|
|»»» user_id|integer|false|none||none|
|»»» user_nick_name|string|false|none||none|
|»»» user_avatar|string|false|none||none|
|»»» category|string|false|none||none|
|»»» source|string|false|none||none|
|»»» link|string|false|none||none|
|»»» banner_id|integer|false|none||none|
|»»» banner_url|string|false|none||none|
|»»» tags|[string]|false|none||none|
|»»» created_at|string|false|none||none|
|» msg|string|true|none||none|

## DELETE 批量取消收藏

DELETE /apiv1/articles/collects

> Body 请求参数

```json
{
  "id_list": [
    "string"
  ]
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» id_list|body|[string]| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {},
  "msg": "成功取消收藏 1 篇文章"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|» msg|string|true|none||none|

## GET 文章标签

GET /apiv1/articles/tags

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page|query|integer| 否 |页|
|limit|query|integer| 否 |数量|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {
    "count": 2,
    "list": [
      {
        "tag": "go",
        "count": 2,
        "articleIDList": [
          "gin的使用",
          "gorm scan的使用"
        ],
        "createdAt": "2025-03-14 15:05:43"
      },
      {
        "tag": "test",
        "count": 1,
        "articleIDList": [
          "测试发送"
        ],
        "createdAt": ""
      }
    ]
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» count|integer|true|none||none|
|»» list|[object]|true|none||none|
|»»» tag|string|true|none||none|
|»»» count|integer|true|none||none|
|»»» articleIDList|[string]|true|none||none|
|»»» createdAt|string|true|none||none|
|» msg|string|true|none||none|

## GET 通过文章id查找文章

GET /apiv1/articles/detail/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {
    "id": "wVdDk5UBFZpd6Pp5URrp",
    "created_at": "2025-03-14 14:06:57",
    "updated_at": "2025-03-14 14:06:57",
    "title": "测试发送",
    "keyword": "测试发送",
    "abstract": "测试一下\n",
    "content": "测试一下",
    "look_count": 0,
    "comment_count": 0,
    "digg_count": 0,
    "collects_count": 0,
    "user_id": 5,
    "user_nick_name": "管理员用户",
    "user_avatar": "https://ts1.tc.mm.bing.net/th/id/R-C.26fa5434823e0afae3f9b576b61b3df0?rik=1ki5rrqJXLS00w&riu=http%3a%2f%2fpic.52112.com%2f180420%2f180420_32%2fJ9xjxe1jIg_small.jpg&ehk=a8hQQlllEncpFeXgnFZ1a7fIII7lcz2ph6WLdtzS51k%3d&risl=&pid=ImgRaw&r=0",
    "category": "测试",
    "source": "",
    "link": "",
    "banner_id": 78,
    "banner_url": "uploads/file/admin/屏幕截图_20241225_140629_duka.png",
    "tags": [
      "test"
    ]
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» id|string|true|none||none|
|»» created_at|string|true|none||none|
|»» updated_at|string|true|none||none|
|»» title|string|true|none||none|
|»» keyword|string|true|none||none|
|»» abstract|string|true|none||none|
|»» content|string|true|none||none|
|»» look_count|integer|true|none||none|
|»» comment_count|integer|true|none||none|
|»» digg_count|integer|true|none||none|
|»» collects_count|integer|true|none||none|
|»» user_id|integer|true|none||none|
|»» user_nick_name|string|true|none||none|
|»» user_avatar|string|true|none||none|
|»» category|string|true|none||none|
|»» source|string|true|none||none|
|»» link|string|true|none||none|
|»» banner_id|integer|true|none||none|
|»» banner_url|string|true|none||none|
|»» tags|[string]|true|none||none|
|» msg|string|true|none||none|

## GET 文章id标题列表

GET /apiv1/article_id_title

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": [
    {
      "id": "UpI5jpUBS5cfAjmlRN3m",
      "title": "gin的使用"
    },
    {
      "id": "U5JBjpUBS5cfAjmlx914",
      "title": "gorm scan的使用"
    }
  ],
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|[object]|true|none||none|
|»» id|string|true|none||none|
|»» title|string|true|none||none|
|» msg|string|true|none||none|

## GET 文章category列表

GET /apiv1/article/categories

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": [
    {
      "category": "go",
      "count": 2
    },
    {
      "category": "测试",
      "count": 1
    }
  ],
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|[object]|true|none||none|
|»» category|string|true|none||none|
|»» count|integer|true|none||none|
|»» id|string|true|none||none|
|»» title|string|true|none||none|
|» msg|string|true|none||none|

## GET 通过title查找文章

GET /apiv1/article/detail

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|title|query|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {
    "id": "wVdDk5UBFZpd6Pp5URrp",
    "created_at": "2025-03-14 14:06:57",
    "updated_at": "2025-03-14 14:06:57",
    "title": "测试发送",
    "keyword": "测试发送",
    "abstract": "测试一下\n",
    "content": "测试一下",
    "look_count": 0,
    "comment_count": 0,
    "digg_count": 0,
    "collects_count": 0,
    "user_id": 5,
    "user_nick_name": "管理员用户",
    "user_avatar": "https://ts1.tc.mm.bing.net/th/id/R-C.26fa5434823e0afae3f9b576b61b3df0?rik=1ki5rrqJXLS00w&riu=http%3a%2f%2fpic.52112.com%2f180420%2f180420_32%2fJ9xjxe1jIg_small.jpg&ehk=a8hQQlllEncpFeXgnFZ1a7fIII7lcz2ph6WLdtzS51k%3d&risl=&pid=ImgRaw&r=0",
    "category": "测试",
    "source": "",
    "link": "",
    "banner_id": 78,
    "banner_url": "uploads/file/admin/屏幕截图_20241225_140629_duka.png",
    "tags": [
      "test"
    ]
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» id|string|true|none||none|
|»» created_at|string|true|none||none|
|»» updated_at|string|true|none||none|
|»» title|string|true|none||none|
|»» keyword|string|true|none||none|
|»» abstract|string|true|none||none|
|»» content|string|true|none||none|
|»» look_count|integer|true|none||none|
|»» comment_count|integer|true|none||none|
|»» digg_count|integer|true|none||none|
|»» collects_count|integer|true|none||none|
|»» user_id|integer|true|none||none|
|»» user_nick_name|string|true|none||none|
|»» user_avatar|string|true|none||none|
|»» category|string|true|none||none|
|»» source|string|true|none||none|
|»» link|string|true|none||none|
|»» banner_id|integer|true|none||none|
|»» banner_url|string|true|none||none|
|»» tags|[string]|true|none||none|
|» msg|string|true|none||none|

## GET 文章发布时间列表

GET /apiv1/articles/calendar

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": [
    {
      "date": "2025-03-13",
      "count": 2
    },
    {
      "date": "2025-03-14",
      "count": 1
    }
  ],
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|[object]|true|none||none|
|»» date|string|true|none||none|
|»» count|integer|true|none||none|
|» msg|string|true|none||none|

## POST 收藏文章

POST /apiv1/articles/collects/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {},
  "msg": "收藏成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|» msg|string|true|none||none|

## POST 点赞文章

POST /apiv1/articles/digg

> Body 请求参数

```json
{
  "id": "string"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» id|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {},
  "msg": "文章点赞成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|» msg|string|true|none||none|

## GET 查看文章详细

GET /apiv1/articles/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {
    "id": "BWvRmZUBdKKKQt44rTsQ",
    "created_at": "2025-03-15 20:40:10",
    "updated_at": "2025-03-15 20:40:10",
    "title": "测试文章详细3",
    "keyword": "测试文章详细3",
    "abstract": "测试一下\n",
    "content": "测试一下",
    "look_count": 1,
    "comment_count": 0,
    "digg_count": 0,
    "collects_count": 0,
    "user_id": 5,
    "user_nick_name": "管理员用户",
    "user_avatar": "https://ts1.tc.mm.bing.net/th/id/R-C.26fa5434823e0afae3f9b576b61b3df0?rik=1ki5rrqJXLS00w&riu=http%3a%2f%2fpic.52112.com%2f180420%2f180420_32%2fJ9xjxe1jIg_small.jpg&ehk=a8hQQlllEncpFeXgnFZ1a7fIII7lcz2ph6WLdtzS51k%3d&risl=&pid=ImgRaw&r=0",
    "category": "测试",
    "source": "",
    "link": "",
    "banner_id": 80,
    "banner_url": "uploads/file/admin/屏幕截图_20241225_134442_odip.png",
    "tags": [
      "test"
    ],
    "is_collect": false,
    "next": {
      "id": "BmvRmZUBdKKKQt44uztR",
      "title": "测试文章详细4"
    },
    "prev": {
      "id": "BGvRmZUBdKKKQt44nzuY",
      "title": "测试文章详细2"
    }
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» id|string|true|none||none|
|»» created_at|string|true|none||none|
|»» updated_at|string|true|none||none|
|»» title|string|true|none||none|
|»» keyword|string|true|none||none|
|»» abstract|string|true|none||none|
|»» content|string|true|none||none|
|»» look_count|integer|true|none||none|
|»» comment_count|integer|true|none||none|
|»» digg_count|integer|true|none||none|
|»» collects_count|integer|true|none||none|
|»» user_id|integer|true|none||none|
|»» user_nick_name|string|true|none||none|
|»» user_avatar|string|true|none||none|
|»» category|string|true|none||none|
|»» source|string|true|none||none|
|»» link|string|true|none||none|
|»» banner_id|integer|true|none||none|
|»» banner_url|string|true|none||none|
|»» tags|[string]|true|none||none|
|»» is_collect|boolean|true|none||none|
|»» next|object|true|none||none|
|»»» id|string|true|none||none|
|»»» title|string|true|none||none|
|»» prev|object|true|none||none|
|»»» id|string|true|none||none|
|»»» title|string|true|none||none|
|» msg|string|true|none||none|

## GET 全文搜索

GET /apiv1/articles/text

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page|query|integer| 否 |none|
|key|query|string| 否 |none|
|limit|query|integer| 否 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {
    "count": 5,
    "list": [
      {
        "id": "",
        "key": "a5f6npUBuNwKnQGb1cnn",
        "title": "全文搜索测试",
        "slug": "a5f6npUBuNwKnQGb1cnn#全文搜索测试",
        "body": "全文搜索<em>测</em><em>试</em>"
      },
      {
        "id": "",
        "key": "bZf9npUBuNwKnQGbHsm9",
        "title": "全文搜索测试",
        "slug": "bZf9npUBuNwKnQGbHsm9#全文搜索测试",
        "body": "## 标题1 go hello\n"
      },
      {
        "id": "",
        "key": "cJf9npUBuNwKnQGb1ck8",
        "title": "全文搜索测试",
        "slug": "cJf9npUBuNwKnQGb1ck8#全文搜索测试",
        "body": "## 标题1 go hello  ### 1   ####222\n"
      },
      {
        "id": "",
        "key": "bZf9npUBuNwKnQGbHsm9",
        "title": "全文搜索测试2",
        "slug": "bZf9npUBuNwKnQGbHsm9#全文搜索测试2",
        "body": ""
      },
      {
        "id": "",
        "key": "cJf9npUBuNwKnQGb1ck8",
        "title": "全文搜索测试3",
        "slug": "cJf9npUBuNwKnQGb1ck8#全文搜索测试3",
        "body": ""
      }
    ]
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» count|integer|true|none||none|
|»» list|[object]|true|none||none|
|»»» id|string|true|none||none|
|»»» key|string|true|none||none|
|»»» title|string|true|none||none|
|»»» slug|string|true|none||none|
|»»» body|string|true|none||none|
|» msg|string|true|none||none|

# comment

## POST 创建评论

POST /apiv1/comments

> Body 请求参数

```json
{
  "article_id": "",
  "content": "",
  "parent_comment_id": null
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» article_id|body|integer| 是 |none|
|» content|body|string| 是 |none|
|» parent_comment_id|body|integer| 否 |none|

> 返回示例

```json
{
  "code": 0,
  "data": {},
  "msg": "创建成功"
}
```

```json
{
  "code": 7,
  "data": {},
  "msg": "没有这篇文章"
}
```

```json
{
  "code": 0,
  "data": {},
  "msg": "创建成功"
}
```

```json
{
  "code": 7,
  "data": {},
  "msg": "父评论不存在"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|» msg|string|true|none||none|

## GET 评论列表

GET /apiv1/comments

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page|query|integer| 否 |none|
|limit|query|integer| 否 |none|
|sort|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {
    "count": 3,
    "list": [
      {
        "id": 3,
        "created_at": "2025-03-17T13:07:17.728+08:00",
        "article_title": "石蒜深度测试（含完整MD语法）3",
        "article_banner": "uploads/file/admin/屏幕截图_20241226_191134_nefi.png",
        "parent_comment_id": 2,
        "content": "哇哇哇",
        "digg_count": 0,
        "comment_count": 0,
        "user_nick_name": "管理员用户"
      },
      {
        "id": 2,
        "created_at": "2025-03-17T13:06:36.903+08:00",
        "article_title": "石蒜深度测试（含完整MD语法）3",
        "article_banner": "uploads/file/admin/屏幕截图_20241226_191134_nefi.png",
        "parent_comment_id": null,
        "content": "哇哇哇",
        "digg_count": 0,
        "comment_count": 1,
        "user_nick_name": "管理员用户"
      },
      {
        "id": 1,
        "created_at": "2025-03-17T13:05:31.335+08:00",
        "article_title": "石蒜深度测试（含完整MD语法）3",
        "article_banner": "uploads/file/admin/屏幕截图_20241226_191134_nefi.png",
        "parent_comment_id": null,
        "content": "哇哇哇",
        "digg_count": 0,
        "comment_count": 0,
        "user_nick_name": "管理员用户"
      }
    ]
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» count|integer|true|none||none|
|»» list|[object]|true|none||none|
|»»» id|integer|true|none||none|
|»»» created_at|string|true|none||none|
|»»» article_title|string|true|none||none|
|»»» article_banner|string|true|none||none|
|»»» parent_comment_id|integer¦null|true|none||none|
|»»» content|string|true|none||none|
|»»» digg_count|integer|true|none||none|
|»»» comment_count|integer|true|none||none|
|»»» user_nick_name|string|true|none||none|
|» msg|string|true|none||none|

## GET 有评论的文章列表

GET /apiv1/comments/articles

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page|query|integer| 否 |none|
|limit|query|integer| 否 |none|
|sort|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {
    "count": 1,
    "list": [
      {
        "title": "石蒜深度测试（含完整MD语法）3",
        "id": "H4I7opUBcr4HzqY20e55",
        "count": 3
      }
    ]
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» count|integer|true|none||none|
|»» list|[object]|true|none||none|
|»»» title|string|false|none||none|
|»»» id|string|false|none||none|
|»»» count|integer|false|none||none|
|» msg|string|true|none||none|

## GET 获取某篇文章下的评论

GET /apiv1/comments/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {
    "count": 2,
    "list": [
      {
        "article_id": "H4I7opUBcr4HzqY20e55",
        "comment_count": 0,
        "content": "哇哇哇",
        "created_at": "2025-03-17T13:05:31.335+08:00",
        "digg_count": 0,
        "id": 1,
        "parent_comment_id": null,
        "sub_comments": [],
        "user": {
          "addr": "内网地址",
          "avatar": "https://ts1.tc.mm.bing.net/th/id/R-C.26fa5434823e0afae3f9b576b61b3df0?rik=1ki5rrqJXLS00w&riu=http%3a%2f%2fpic.52112.com%2f180420%2f180420_32%2fJ9xjxe1jIg_small.jpg&ehk=a8hQQlllEncpFeXgnFZ1a7fIII7lcz2ph6WLdtzS51k%3d&risl=&pid=ImgRaw&r=0",
          "ip": "127.0.0.1",
          "nick_name": "管理员用户"
        },
        "user_id": 5
      },
      {
        "article_id": "H4I7opUBcr4HzqY20e55",
        "comment_count": 1,
        "content": "哇哇哇",
        "created_at": "2025-03-17T13:06:36.903+08:00",
        "digg_count": 0,
        "id": 2,
        "parent_comment_id": null,
        "sub_comments": [
          {
            "article_id": "H4I7opUBcr4HzqY20e55",
            "comment_count": 0,
            "content": "哇哇哇",
            "created_at": "2025-03-17T13:07:17.728+08:00",
            "digg_count": 0,
            "id": 3,
            "parent_comment_id": 2,
            "sub_comments": [],
            "user": {
              "addr": "内网地址",
              "avatar": "https://ts1.tc.mm.bing.net/th/id/R-C.26fa5434823e0afae3f9b576b61b3df0?rik=1ki5rrqJXLS00w&riu=http%3a%2f%2fpic.52112.com%2f180420%2f180420_32%2fJ9xjxe1jIg_small.jpg&ehk=a8hQQlllEncpFeXgnFZ1a7fIII7lcz2ph6WLdtzS51k%3d&risl=&pid=ImgRaw&r=0",
              "ip": "127.0.0.1",
              "nick_name": "管理员用户"
            },
            "user_id": 5
          }
        ],
        "user": {
          "addr": "内网地址",
          "avatar": "https://ts1.tc.mm.bing.net/th/id/R-C.26fa5434823e0afae3f9b576b61b3df0?rik=1ki5rrqJXLS00w&riu=http%3a%2f%2fpic.52112.com%2f180420%2f180420_32%2fJ9xjxe1jIg_small.jpg&ehk=a8hQQlllEncpFeXgnFZ1a7fIII7lcz2ph6WLdtzS51k%3d&risl=&pid=ImgRaw&r=0",
          "ip": "127.0.0.1",
          "nick_name": "管理员用户"
        },
        "user_id": 5
      }
    ]
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» count|integer|true|none||none|
|»» list|[object]|true|none||none|
|»»» article_id|string|true|none||none|
|»»» comment_count|integer|true|none||none|
|»»» content|string|true|none||none|
|»»» created_at|string|true|none||none|
|»»» digg_count|integer|true|none||none|
|»»» id|integer|true|none||none|
|»»» parent_comment_id|null|true|none||none|
|»»» sub_comments|[object]|true|none||none|
|»»»» article_id|string|false|none||none|
|»»»» comment_count|integer|false|none||none|
|»»»» content|string|false|none||none|
|»»»» created_at|string|false|none||none|
|»»»» digg_count|integer|false|none||none|
|»»»» id|integer|false|none||none|
|»»»» parent_comment_id|integer|false|none||none|
|»»»» sub_comments|[string]|false|none||none|
|»»»» user|object|false|none||none|
|»»»»» addr|string|true|none||none|
|»»»»» avatar|string|true|none||none|
|»»»»» ip|string|true|none||none|
|»»»»» nick_name|string|true|none||none|
|»»»» user_id|integer|false|none||none|
|»»» user|object|true|none||none|
|»»»» addr|string|true|none||none|
|»»»» avatar|string|true|none||none|
|»»»» ip|string|true|none||none|
|»»»» nick_name|string|true|none||none|
|»»» user_id|integer|true|none||none|
|» msg|string|true|none||none|

## DELETE 删除评论

DELETE /apiv1/comments/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|integer| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {},
  "msg": "共删除 2 条评论"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|» msg|string|true|none||none|

## GET 给某评论点赞

GET /apiv1/comments/digg/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|integer| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {},
  "msg": "评论点赞成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|» msg|string|true|none||none|

# 新闻itab

## GET news接口

GET /apiv1/news

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|query|array[string]| 否 |none|
|size|query|string| 是 |none|
|signaturekey|header|string| 否 |none|
|Version|header|string| 否 |none|
|user-agent|header|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": [
    {
      "hotValue": "",
      "index": "1",
      "link": "https://www.36kr.com/p/3210158561133448",
      "title": "小红书“如接”DeepSeek"
    }
  ],
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|[object]|true|none||none|
|»» hotValue|string|false|none||none|
|»» index|string|false|none||none|
|»» link|string|false|none||none|
|»» title|string|false|none||none|
|» msg|string|true|none||none|

## GET itab_news

GET /api/top/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|query|array[string]| 否 |none|
|size|query|string| 是 |none|
|signaturekey|header|string| 否 |none|
|Version|header|string| 否 |none|
|user-agent|header|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 200,
  "data": [
    {
      "index": "1",
      "title": "小红书“如接”DeepSeek",
      "hotValue": "",
      "link": "https://www.36kr.com/p/3210158561133448"
    }
  ],
  "msg": "请求成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|[object]|true|none||none|
|»» index|string|false|none||none|
|»» title|string|false|none||none|
|»» hotValue|string|false|none||none|
|»» link|string|false|none||none|
|» msg|string|true|none||none|

# chatroom

## GET 聊天记录

GET /apiv1/chat/records

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page|query|integer| 否 |none|
|limit|query|integer| 否 |none|
|key|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {
    "count": 6,
    "list": [
      {
        "avatar": "uploads/chat_avatar/千.png",
        "content": "哈哈哈好呀好呀,你好",
        "created_at": "2025-03-18T17:14:35.212+08:00",
        "id": 47,
        "is_group": true,
        "msg_type": "文本消息",
        "nick_name": "千反田爱瑠进行英灵召唤"
      },
      {
        "avatar": "uploads/chat_avatar/废.png",
        "content": "我是用户一,你好你好",
        "created_at": "2025-03-18T17:14:24.328+08:00",
        "id": 46,
        "is_group": true,
        "msg_type": "文本消息",
        "nick_name": "废柴的炭治郎"
      },
      {
        "avatar": "uploads/chat_avatar/加.png",
        "content": "你好你好",
        "created_at": "2025-03-18T17:02:29.203+08:00",
        "id": 41,
        "is_group": true,
        "msg_type": "文本消息",
        "nick_name": "加藤惠使用王之财宝"
      },
      {
        "avatar": "uploads/chat_avatar/加.png",
        "content": "你好你好",
        "created_at": "2025-03-18T17:02:29.047+08:00",
        "id": 40,
        "is_group": true,
        "msg_type": "文本消息",
        "nick_name": "加藤惠使用王之财宝"
      },
      {
        "avatar": "uploads/chat_avatar/加.png",
        "content": "你好你好",
        "created_at": "2025-03-18T17:02:28.898+08:00",
        "id": 39,
        "is_group": true,
        "msg_type": "文本消息",
        "nick_name": "加藤惠使用王之财宝"
      },
      {
        "avatar": "uploads/chat_avatar/加.png",
        "content": "你好你好",
        "created_at": "2025-03-18T17:02:28.673+08:00",
        "id": 38,
        "is_group": true,
        "msg_type": "文本消息",
        "nick_name": "加藤惠使用王之财宝"
      }
    ]
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» count|integer|true|none||none|
|»» list|[object]|true|none||none|
|»»» avatar|string|true|none||none|
|»»» content|string|true|none||none|
|»»» created_at|string|true|none||none|
|»»» id|integer|true|none||none|
|»»» is_group|boolean|true|none||none|
|»»» msg_type|string|true|none||none|
|»»» nick_name|string|true|none||none|
|» msg|string|true|none||none|

## DELETE 删除记录

DELETE /apiv1/chat/records

> Body 请求参数

```json
{
  "id_list": [
    0
  ]
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» id_list|body|[integer]| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {},
  "msg": "共删除记录9条"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|» msg|string|true|none||none|

# log

## GET 查看日志列表

GET /apiv1/logs

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page|query|integer| 否 |none|
|limit|query|integer| 否 |none|
|key|query|string| 否 |username log_title支持模糊查询|
|sort|query|string| 否 |created_at asc/desc|
|level|query|integer| 否 |1-info 2-warn 3-error|
|type|query|integer| 否 |1-login 2-action|
|ip|query|string| 否 |127.0.0.1|
|addr|query|string| 否 |内网地址|
|user_id|query|integer| 否 |none|
|username|query|string| 否 |none|
|status|query|boolean| 否 |登录成功 1 登录失败0 |
|date|query|string| 否 |2025-03-19|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {
    "count": 1,
    "list": [
      {
        "id": 47,
        "created_at": "2025-03-19T20:59:58.836+08:00",
        "updated_at": "2025-03-19T20:59:58.836+08:00",
        "ip": "127.0.0.1",
        "addr": "内网地址",
        "level": 1,
        "title": "创建广告",
        "content": "",
        "userID": 5,
        "userName": "admin",
        "serviceName": "",
        "status": false,
        "type": 2,
        "readStatus": false
      }
    ]
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» count|integer|true|none||none|
|»» list|[object]|true|none||none|
|»»» id|integer|true|none||none|
|»»» created_at|string|true|none||none|
|»»» updated_at|string|true|none||none|
|»»» ip|string|true|none||none|
|»»» addr|string|true|none||none|
|»»» level|integer|true|none||none|
|»»» title|string|true|none||none|
|»»» content|string|true|none||none|
|»»» userID|integer|true|none||none|
|»»» userName|string|true|none||none|
|»»» serviceName|string|true|none||none|
|»»» status|boolean|true|none||none|
|»»» type|integer|true|none||none|
|»»» readStatus|boolean|true|none||none|
|» msg|string|true|none||none|

## DELETE 删除日志

DELETE /apiv1/logs

> Body 请求参数

```json
{
  "id_list": [
    0
  ]
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» id_list|body|[integer]| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {},
  "msg": "共删除 6 个日志"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|» msg|string|true|none||none|

## GET 手动同步redis日志到mysql

GET /apiv1/logs_refresh

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

# data

## GET 登录注册数据

GET /apiv1/data_login

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|date|query|integer| 否 |	OneWeek = 1 + iota|

#### 详细说明

**date**: 	OneWeek = 1 + iota
	OneMonth
	TwoMonth
	ThreeMonth
	HalfYear
	OneYear

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {
    "date_list": [
      "2025-03-13",
      "2025-03-14",
      "2025-03-15",
      "2025-03-16",
      "2025-03-17",
      "2025-03-18",
      "2025-03-19"
    ],
    "login_data": [
      0,
      0,
      0,
      0,
      0,
      2,
      17
    ],
    "sign_data": [
      0,
      0,
      0,
      0,
      0,
      0,
      0
    ]
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» date_list|[string]|true|none||none|
|»» login_data|[integer]|true|none||none|
|»» sign_data|[integer]|true|none||none|
|» msg|string|true|none||none|

## GET 基本数据列表

GET /apiv1/data_sum

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {
    "user_count": 4,
    "article_count": 13,
    "message_count": 61,
    "chat_group_count": 1,
    "now_login_count": 1,
    "now_sign_count": 0
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» user_count|integer|true|none||用户数量|
|»» article_count|integer|true|none||文章数量|
|»» message_count|integer|true|none||消息数量|
|»» chat_group_count|integer|true|none||群聊记录|
|»» now_login_count|integer|true|none||今日登录|
|»» now_sign_count|integer|true|none||今日注册|
|» msg|string|true|none||none|

## GET 权限id列表

GET /apiv1/role_ids

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": [
    {
      "label": "管理员",
      "value": 1
    },
    {
      "label": "普通用户",
      "value": 2
    },
    {
      "label": "游客",
      "value": 3
    }
  ],
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|[object]|true|none||none|
|»» label|string|true|none||none|
|»» value|integer|true|none||none|
|» msg|string|true|none||none|

# 高德

## GET go blog weather ip

GET /apiv1/gaode/ip_weather

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|ip|query|string| 否 |120.208.127.167 114.247.50.2|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {
    "province": "山西",
    "city": "长治市",
    "adcode": "140400",
    "weather": "晴",
    "temperature": "18",
    "winddirection": "西北",
    "windpower": "≤3",
    "humidity": "6",
    "reporttime": "2025-03-21 14:02:34",
    "temperature_float": "18.0",
    "humidity_float": "6.0"
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» province|string|true|none||none|
|»» city|string|true|none||none|
|»» adcode|string|true|none||none|
|»» weather|string|true|none||none|
|»» temperature|string|true|none||none|
|»» winddirection|string|true|none||none|
|»» windpower|string|true|none||none|
|»» humidity|string|true|none||none|
|»» reporttime|string|true|none||none|
|»» temperature_float|string|true|none||none|
|»» humidity_float|string|true|none||none|
|» msg|string|true|none||none|

## GET go blog weather adcode

GET /apiv1/gaode/adcode_weather

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|adcode|query|string| 否 |110101 140400|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {
    "province": "山西",
    "city": "长治市",
    "adcode": "140400",
    "weather": "晴",
    "temperature": "17",
    "winddirection": "西",
    "windpower": "5",
    "humidity": "6",
    "reporttime": "2025-03-21 13:37:35",
    "temperature_float": "17.0",
    "humidity_float": "6.0"
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» province|string|true|none||none|
|»» city|string|true|none||none|
|»» adcode|string|true|none||none|
|»» weather|string|true|none||none|
|»» temperature|string|true|none||none|
|»» winddirection|string|true|none||none|
|»» windpower|string|true|none||none|
|»» humidity|string|true|none||none|
|»» reporttime|string|true|none||none|
|»» temperature_float|string|true|none||none|
|»» humidity_float|string|true|none||none|
|» msg|string|true|none||none|

# 用户反馈

## POST 用户反馈

POST /apiv1/feedback

> Body 请求参数

```json
{
  "email": "string",
  "content": "string"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» email|body|string| 是 |none|
|» content|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 7,
  "data": {},
  "msg": "存在相同留言"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|» msg|string|true|none||none|

## GET 用户反馈列表

GET /apiv1/feedback

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|limit|query|string| 否 |none|
|page|query|string| 否 |none|
|key|query|string| 否 |content email|
|sort|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {
    "count": 10,
    "list": [
      {
        "id": 11,
        "created_at": "2025-03-21T15:08:14.192+08:00",
        "email": "zyuuforyuhot@mail.com",
        "content": "51"
      }
    ]
  },
  "msg": "成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» count|integer|true|none||none|
|»» list|[object]|true|none||none|
|»»» id|integer|false|none||none|
|»»» created_at|string|false|none||none|
|»»» email|string|false|none||none|
|»»» content|string|false|none||none|
|» msg|string|true|none||none|

## DELETE 批量删除反馈

DELETE /apiv1/feedback

> Body 请求参数

```json
{
  "id_list": [
    "string"
  ]
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» id_list|body|[string]| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {},
  "msg": "共删除 2 条反馈内容"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|» msg|string|true|none||none|



