---
title: 豆瓣电影 v1.0.0
language_tabs:
- shell: Shell
- http: HTTP
- javascript: JavaScript
- ruby: Ruby
- python: Python
- php: PHP
- java: Java
- go: Go
  toc_footers: []
  includes: []
  search: true
  code_clipboard: true
  highlight_theme: darkula
  headingLevel: 2
  generator: "@tarslib/widdershins v4.0.4"

---

# 豆瓣电影

> v1.0.0

# User

## GET 获取个人信息

GET /api/user/get_user

> Body 请求参数

```yaml
access_token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7IklkIjo0LCJVc2VybmFtZSI6IiIsIlBhc3N3b3JkIjoiMzZiNjI5MjE2MTZkMjQ3MGQwMzI1YjFkNTk0ZjI0YWEiLCJFbWFpbCI6IiIsIlBob25lIjoiMTc3MjY2MzM3NDAiLCJTYWx0IjoiMTY0MjU5MDA5NCIsIkF2YXRhciI6Imh0dHBzOi8vZG91YmFuLWF2YXRhci0xMzA4NzU3Mzg1LmNvcy5hcC1jaG9uZ3FpbmcubXlxY2xvdWQuY29tL2F2YXRhci80ZjA5MjM1YjRlMTc0YzQ1OGY0NzFkN2IxNzM5N2IzNi5wbmciLCJEb21haW5OYW1lIjoiIiwiSGFiaXRhdCI6IiIsIkhvbWV0b3duIjoiIiwiQmlydGhkYXkiOiI5OTk5LTEyLTEyVDAwOjAwOjAwKzA4OjAwIiwiUmVnRGF0ZSI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiU3RhdGVtZW50IjoiIiwiRm9sbG93ZXJzIjowLCJGb2xsb3dpbmdzIjowfSwiVHlwZSI6IkFDQ0VTU19UT0tFTiIsIlRpbWUiOiIwMDAxLTAxLTAxVDAwOjAwOjAwWiIsImV4cCI6MTY0NTQzNDk2MCwiaXNzIjoiZG91YmFuIn0.v9WYGEXyiwbVFEUWdUE9heOaFePX1b8pnsDOYq6G1EE
refresh_token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7IklkIjo0LCJVc2VybmFtZSI6IiIsIlBhc3N3b3JkIjoiMzZiNjI5MjE2MTZkMjQ3MGQwMzI1YjFkNTk0ZjI0YWEiLCJFbWFpbCI6IiIsIlBob25lIjoiMTc3MjY2MzM3NDAiLCJTYWx0IjoiMTY0MjU5MDA5NCIsIkF2YXRhciI6Imh0dHBzOi8vZG91YmFuLWF2YXRhci0xMzA4NzU3Mzg1LmNvcy5hcC1jaG9uZ3FpbmcubXlxY2xvdWQuY29tL2F2YXRhci80ZjA5MjM1YjRlMTc0YzQ1OGY0NzFkN2IxNzM5N2IzNi5wbmciLCJEb21haW5OYW1lIjoiIiwiSGFiaXRhdCI6IiIsIkhvbWV0b3duIjoiIiwiQmlydGhkYXkiOiI5OTk5LTEyLTEyVDAwOjAwOjAwKzA4OjAwIiwiUmVnRGF0ZSI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiU3RhdGVtZW50IjoiIiwiRm9sbG93ZXJzIjowLCJGb2xsb3dpbmdzIjowfSwiVHlwZSI6IlJFRlJFU0hfVE9LRU4iLCJUaW1lIjoiMDAwMS0wMS0wMVQwMDowMDowMFoiLCJleHAiOjE2NDYwMzM3NjAsImlzcyI6ImRvdWJhbiJ9.18g5w-wr4Au4AqekDSDtK70_lnUWg31w2jqFtT6jw5A

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object|false|none|
|» access_token|body|string|true|none|
|» refresh_token|body|string|true|none|

> 返回示例

> 成功

```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7IklkIjo0LCJVc2VybmFtZSI6IiIsIlBhc3N3b3JkIjoiMzZiNjI5MjE2MTZkMjQ3MGQwMzI1YjFkNTk0ZjI0YWEiLCJFbWFpbCI6IiIsIlBob25lIjoiMTc3MjY2MzM3NDAiLCJTYWx0IjoiMTY0MjU5MDA5NCIsIkF2YXRhciI6Imh0dHBzOi8vZG91YmFuLWF2YXRhci0xMzA4NzU3Mzg1LmNvcy5hcC1jaG9uZ3FpbmcubXlxY2xvdWQuY29tL2F2YXRhci80ZjA5MjM1YjRlMTc0YzQ1OGY0NzFkN2IxNzM5N2IzNi5wbmciLCJEb21haW5OYW1lIjoiIiwiSGFiaXRhdCI6IiIsIkhvbWV0b3duIjoiIiwiQmlydGhkYXkiOiI5OTk5LTEyLTEyVDAwOjAwOjAwKzA4OjAwIiwiUmVnRGF0ZSI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiU3RhdGVtZW50IjoiIiwiRm9sbG93ZXJzIjowLCJGb2xsb3dpbmdzIjowfSwiVHlwZSI6IkFDQ0VTU19UT0tFTiIsIlRpbWUiOiIwMDAxLTAxLTAxVDAwOjAwOjAwWiIsImV4cCI6MTY0NTQzNDk2MCwiaXNzIjoiZG91YmFuIn0.v9WYGEXyiwbVFEUWdUE9heOaFePX1b8pnsDOYq6G1EE",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7IklkIjo0LCJVc2VybmFtZSI6IiIsIlBhc3N3b3JkIjoiMzZiNjI5MjE2MTZkMjQ3MGQwMzI1YjFkNTk0ZjI0YWEiLCJFbWFpbCI6IiIsIlBob25lIjoiMTc3MjY2MzM3NDAiLCJTYWx0IjoiMTY0MjU5MDA5NCIsIkF2YXRhciI6Imh0dHBzOi8vZG91YmFuLWF2YXRhci0xMzA4NzU3Mzg1LmNvcy5hcC1jaG9uZ3FpbmcubXlxY2xvdWQuY29tL2F2YXRhci80ZjA5MjM1YjRlMTc0YzQ1OGY0NzFkN2IxNzM5N2IzNi5wbmciLCJEb21haW5OYW1lIjoiIiwiSGFiaXRhdCI6IiIsIkhvbWV0b3duIjoiIiwiQmlydGhkYXkiOiI5OTk5LTEyLTEyVDAwOjAwOjAwKzA4OjAwIiwiUmVnRGF0ZSI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiU3RhdGVtZW50IjoiIiwiRm9sbG93ZXJzIjowLCJGb2xsb3dpbmdzIjowfSwiVHlwZSI6IlJFRlJFU0hfVE9LRU4iLCJUaW1lIjoiMDAwMS0wMS0wMVQwMDowMDowMFoiLCJleHAiOjE2NDYwMzM3NjAsImlzcyI6ImRvdWJhbiJ9.18g5w-wr4Au4AqekDSDtK70_lnUWg31w2jqFtT6jw5A",
  "status": "true",
  "userInfo": {
    "Id": 4,
    "Username": "mj",
    "Email": "",
    "Phone": "17726633740",
    "Avatar": "https://douban-avatar-1308757385.cos.ap-chongqing.myqcloud.com/avatar/4f09235b4e174c458f471d7b17397b36.png",
    "DomainName": "",
    "Habitat": "",
    "Hometown": "",
    "Birthday": "9999-12-12T00:00:00+08:00",
    "RegDate": "2022-01-19T00:00:00+08:00",
    "Statement": "",
    "Followers": 0,
    "Followings": 0
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## PUT 跟换常住地

PUT /api/user/change_habitat

Token 是access_token+" "+refresh_token

> Body 请求参数

```yaml
habitat: 重庆

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object|false|none|
|» habitat|body|string|true|none|

> 返回示例

> 成功

```json
{
  "data": "修改成功",
  "status": "true"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# User/login

## POST 短信登录接口

POST /api/user/login/sms

> Body 请求参数

```yaml
phone: string
verify_code: string

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object|false|none|
|» phone|body|string|true|电话号码 如果成功会返回电话号|
|» verify_code|body|string|true|验证码|

> 返回示例

> 成功登录

```json
{
  "data": "未发送验证码",
  "status": "false"
}
```

```json
{
  "data": "电话号码不能为空",
  "status": "false"
}
```

```json
{
  "data": "手机号格式错误",
  "status": "false"
}
```

```json
{
  "data": "验证码错误或者过期",
  "status": "false"
}
```

```json
{
  "data": "验证码不能为空",
  "status": "false"
}
```

```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7IklkIjo0LCJVc2VybmFtZSI6IiIsIlBhc3N3b3JkIjoiMzZiNjI5MjE2MTZkMjQ3MGQwMzI1YjFkNTk0ZjI0YWEiLCJFbWFpbCI6IiIsIlBob25lIjoiMTU2NTY2NzEwNzAiLCJTYWx0IjoiMTY0MjU5MDA5NCIsIkF2YXRhciI6IiIsIkRvbWFpbk5hbWUiOiIiLCJIYWJpdGF0IjoiIiwiSG9tZXRvd24iOiIiLCJCaXJ0aGRheSI6Ijk5OTktMTItMTJUMDA6MDA6MDArMDg6MDAiLCJSZWdEYXRlIjoiMDAwMS0wMS0wMVQwMDowMDowMFoiLCJTdGF0ZW1lbnQiOiIiLCJGb2xsb3dlcnMiOjAsIkZvbGxvd2luZ3MiOjB9LCJUeXBlIjoiQUNDRVNTX1RPS0VOIiwiVGltZSI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiZXhwIjoxNjQyNTkwOTAwLCJpc3MiOiJkb3ViYW4ifQ.wyzDPzcQDoFXsuQtW-yfNVeS9cU1U-CZVQY5U9PNmAY",
  "data": 4,
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7IklkIjo0LCJVc2VybmFtZSI6IiIsIlBhc3N3b3JkIjoiMzZiNjI5MjE2MTZkMjQ3MGQwMzI1YjFkNTk0ZjI0YWEiLCJFbWFpbCI6IiIsIlBob25lIjoiMTU2NTY2NzEwNzAiLCJTYWx0IjoiMTY0MjU5MDA5NCIsIkF2YXRhciI6IiIsIkRvbWFpbk5hbWUiOiIiLCJIYWJpdGF0IjoiIiwiSG9tZXRvd24iOiIiLCJCaXJ0aGRheSI6Ijk5OTktMTItMTJUMDA6MDA6MDArMDg6MDAiLCJSZWdEYXRlIjoiMDAwMS0wMS0wMVQwMDowMDowMFoiLCJTdGF0ZW1lbnQiOiIiLCJGb2xsb3dlcnMiOjAsIkZvbGxvd2luZ3MiOjB9LCJUeXBlIjoiUkVGUkVTSF9UT0tFTiIsIlRpbWUiOiIwMDAxLTAxLTAxVDAwOjAwOjAwWiIsImV4cCI6MTY0MzE5NTQwMCwiaXNzIjoiZG91YmFuIn0.KGGIybfolq_8Z0xAPUFeMVtmOf8ogLo8y2jthFKDIDM",
  "status": "ture"
}
```

```json
{
  "info": "新用户",
  "status": "true",
  "data": "phone"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功登录|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» access_token|string|true|none|none|
|» data|integer|true|none|none|
|» refresh_token|string|true|none|none|
|» status|boolean|true|none|none|

## POST 密码登录接口

POST /api/user/login/pw

> Body 请求参数

```yaml
loginAccount: "17726633740"
password: "123456"

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object|false|none|
|» loginAccount|body|string|true|用户账户|
|» password|body|string|true|密码|

> 返回示例

> 成功登录

```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7IklkIjo0LCJVc2VybmFtZSI6IiIsIlBhc3N3b3JkIjoiMzZiNjI5MjE2MTZkMjQ3MGQwMzI1YjFkNTk0ZjI0YWEiLCJFbWFpbCI6IiIsIlBob25lIjoiMTU2NTY2NzEwNzAiLCJTYWx0IjoiMTY0MjU5MDA5NCIsIkF2YXRhciI6IiIsIkRvbWFpbk5hbWUiOiIiLCJIYWJpdGF0IjoiIiwiSG9tZXRvd24iOiIiLCJCaXJ0aGRheSI6Ijk5OTktMTItMTJUMDA6MDA6MDArMDg6MDAiLCJSZWdEYXRlIjoiMDAwMS0wMS0wMVQwMDowMDowMFoiLCJTdGF0ZW1lbnQiOiIiLCJGb2xsb3dlcnMiOjAsIkZvbGxvd2luZ3MiOjB9LCJUeXBlIjoiQUNDRVNTX1RPS0VOIiwiVGltZSI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiZXhwIjoxNjQyNTkwOTAwLCJpc3MiOiJkb3ViYW4ifQ.wyzDPzcQDoFXsuQtW-yfNVeS9cU1U-CZVQY5U9PNmAY",
  "data": 4,
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7IklkIjo0LCJVc2VybmFtZSI6IiIsIlBhc3N3b3JkIjoiMzZiNjI5MjE2MTZkMjQ3MGQwMzI1YjFkNTk0ZjI0YWEiLCJFbWFpbCI6IiIsIlBob25lIjoiMTU2NTY2NzEwNzAiLCJTYWx0IjoiMTY0MjU5MDA5NCIsIkF2YXRhciI6IiIsIkRvbWFpbk5hbWUiOiIiLCJIYWJpdGF0IjoiIiwiSG9tZXRvd24iOiIiLCJCaXJ0aGRheSI6Ijk5OTktMTItMTJUMDA6MDA6MDArMDg6MDAiLCJSZWdEYXRlIjoiMDAwMS0wMS0wMVQwMDowMDowMFoiLCJTdGF0ZW1lbnQiOiIiLCJGb2xsb3dlcnMiOjAsIkZvbGxvd2luZ3MiOjB9LCJUeXBlIjoiUkVGUkVTSF9UT0tFTiIsIlRpbWUiOiIwMDAxLTAxLTAxVDAwOjAwOjAwWiIsImV4cCI6MTY0MzE5NTQwMCwiaXNzIjoiZG91YmFuIn0.KGGIybfolq_8Z0xAPUFeMVtmOf8ogLo8y2jthFKDIDM",
  "status": true
}
```

```json
{
  "data": "请输入注册时用的邮箱或者手机号",
  "status": "false"
}
```

```json
{
  "data": "请输入注册时用的邮箱或者手机号",
  "status": "false"
}
```

```json
{
  "data": "请输入密码",
  "status": "false"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功登录|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» acess_token|string|true|none|none|
|» data|integer|true|none|none|
|» refresh_token|string|true|none|none|
|» status|boolean|true|none|none|

# User/register

## POST 短信注册

POST /api/user/register

> Body 请求参数

```yaml
phone: "1725014728"
username: MJ
password: "123456"

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object|false|none|
|» phone|body|string|true|电话号码由短信界面返回的phone填入|
|» username|body|string|true|用户名|
|» password|body|string|true|密码|

> 返回示例

> 成功

```json
{
  "data": "注册成功",
  "status": "ture"
}
```

```json
{
  "data": "参数解析失败",
  "status": "false"
}
```

```json
{
  "data": "该用户名已经注册",
  "status": "false"
}
```

```json
{
  "data": "用户名不能大于14个字符",
  "status": "false"
}
```

```json
{
  "data": " 密码不能小于6个字符",
  "status": "false"
}
```

```json
{
  "data": " 密码不能大于20个字符",
  "status": "false"
}
```

> 服务器错误

```json
{
  "info": "服务器出错"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|服务器错误|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» data|string|true|none|none|

状态码 **500**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» info|string¦null|true|none|none|

# User/sendSms

## POST 发送短信

POST /api/verify/sms

> Body 请求参数

```yaml
phone: "13360647237"

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object|false|none|
|» phone|body|string|true|none|

> 返回示例

> 成功

```json
{
  "data": "电话号码不能为空",
  "status": "false"
}
```

```json
{
  "data": "手机号格式错误",
  "status": "flase"
}
```

```json
{
  "data": "phone(手机号)",
  "info": "短信发送成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# User/email

## PUT 解绑邮箱

PUT /api/user/unbind_email

> Body 请求参数

```yaml
verify_code: string
password: string

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object|false|none|
|» verify_code|body|string|true|none|
|» password|body|string|true|none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## PUT 绑定邮箱

PUT /api/user/bind_email

> Body 请求参数

```yaml
email: mjgopher@163.com
verify_code: "241306"
access_token: string
refresh_token: string

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object|false|none|
|» email|body|string|true|none|
|» verify_code|body|string|true|none|
|» access_token|body|string|true|none|
|» refresh_token|body|string|true|none|

> 返回示例

> 成功

```json
{
  "data": "修改成功",
  "status": "ture"
}
```

```json
{
  "data": "验证码不能为空",
  "status": "ture"
}
```

```json
{
  "data": "验证码错误或者过期",
  "status": "ture"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» data|string|true|none|none|
|» status|string|true|none|none|

## POST 发送邮箱验证码

POST /api/verify/emial

> Body 请求参数

```yaml
email: 1725014728@qq.com

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object|false|none|
|» email|body|string|true|邮箱账户 前端要校验邮箱格式 正常再请求|

> 返回示例

> 成功

```json
{
  "data": "发送成功",
  "status": "ture"
}
```

```json
{
  "data": "邮箱格式错误",
  "status": "false"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|说明|
|---|---|---|---|---|
|» data|string|true|none|none|
|» status|string|true|none|none|

# User/phone

## PUT 绑定电话号码

PUT /api/user/bind_phone

> Body 请求参数

```yaml
verify_code: string
phone: string
access_token: string
refresh_token: string

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object|false|none|
|» verify_code|body|string|true|none|
|» phone|body|string|true|none|
|» access_token|body|string|true|none|
|» refresh_token|body|string|true|none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## PUT 解绑电话号码

PUT /api/user/unbind_phone

> Body 请求参数

```yaml
verify_code: string
id: "5"
phone: "17726633740"

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object|false|none|
|» verify_code|body|string|true|none|
|» id|body|string|true|none|
|» phone|body|string|true|none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# User/suicideAccount

## DELETE 注销账户

DELETE /api/user/suicide

> Body 请求参数

```yaml
access_token: string
refresh_token: string

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object|false|none|
|» access_token|body|string|true|none|
|» refresh_token|body|string|true|none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# User/changeAccount

## PUT 修改账户信息

PUT /api/user/change_account

> Body 请求参数

```yaml
access_token: string
refresh_token: string
new_username: string
birthday: string
habitat: string
hometown: string
hometown_public: string
birthday_public: string

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object|false|none|
|» access_token|body|string|true|none|
|» refresh_token|body|string|true|none|
|» new_username|body|string|false|none|
|» birthday|body|string|false|none|
|» habitat|body|string|false|none|
|» hometown|body|string|false|none|
|» hometown_public|body|string|false|none|
|» birthday_public|body|string|false|none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# User/avatar

## POST 修改头像

POST /api/user/change_avatar

> Body 请求参数

```yaml
id: "4"
avatar: file://C:\Users\17250\Desktop\QQ摸鱼.png

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object|false|none|
|» id|body|string|false|none|
|» avatar|body|string(binary)|true|none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# Comment

## GET 获取单个影评

GET /api/movie/review/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|

> 返回示例

> 成功

```json
{
  "data": "true",
  "reply": [
    {
      "Id": 1,
      "Uid": 4,
      "Pid": 3,
      "Like": 0,
      "Ptable": "review",
      "Date": "2022-02-21T16:48:52+08:00",
      "Username": "mj",
      "Content": "确实很好看",
      "Avatar": "https://douban-avatar-1308757385.cos.ap-chongqing.myqcloud.com/avatar/4f09235b4e174c458f471d7b17397b36.png",
      "RepCnt": 0
    }
  ],
  "review": [
    {
      "Id": 3,
      "Mid": 1291543,
      "Uid": 4,
      "Username": "mj",
      "Avatar": "https://douban-avatar-1308757385.cos.ap-chongqing.myqcloud.com/avatar/4f09235b4e174c458f471d7b17397b36.png",
      "Title": "好看的不得了",
      "Comment": "非常非常非常非常非常非常非常非常非常好看",
      "Time": "2022-02-20T16:16:45+08:00",
      "People": 1,
      "Likes": 0,
      "Unlikes": 0,
      "Report": 0,
      "Star": 4
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## PUT 短评点赞 只能点赞

PUT /api/movie/comment/like/{id}

1 点赞

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取短评

GET /api/movie/comments/{mid}

> Body 请求参数

```yaml
{}

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|mid|path|string|true|none|
|body|body|object|false|none|

> 返回示例

> 成功

```json
{
  "data": [
    {
      "Id": 8,
      "Mid": 1291543,
      "Uid": 4,
      "Username": "mj",
      "Avatar": "https://douban-avatar-1308757385.cos.ap-chongqing.myqcloud.com/avtar.png\r\n",
      "Static": "看过",
      "Comment": "很好看",
      "Time": "2022-02-19T23:34:37+08:00",
      "Help": 0,
      "Report": 0,
      "Star": 4
    }
  ],
  "status": "true"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 发表短评

POST /api/movie/comment/put

> Body 请求参数

```yaml
star: "4"
comment: 很好看
type: 看过
mid: "1291543"

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object|false|none|
|» star|body|string|true|none|
|» comment|body|string|true|none|
|» type|body|string|true|none|
|» mid|body|string|true|none|

> 返回示例

> 成功

```json
{
  "data": "评论成功",
  "status": "true"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## PUT 影评点赞

PUT /api/movie/review/like/{id}

-1 踩 1 点赞

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取影评

GET /api/movie/reviews/{mid}

> Body 请求参数

```yaml
{}

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|mid|path|string|true|none|
|body|body|object|false|none|

> 返回示例

> 成功

```json
{
  "data": [
    {
      "Id": 3,
      "Mid": 1291543,
      "Uid": 4,
      "Username": "mj",
      "Avatar": "https://douban-avatar-1308757385.cos.ap-chongqing.myqcloud.com/avtar.png\r\n",
      "Title": "好看的不得了",
      "Comment": "非常非常非常非常非常非常非常非常非常好看",
      "Time": "2022-02-20T16:16:45+08:00",
      "Likes": 0,
      "Unlikes": 0,
      "Report": 0,
      "Star": 4
    }
  ],
  "status": "true"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 发表影评

POST /api/movie/review/put

> Body 请求参数

```yaml
mid: "1291543"
comment: 非常非常非常非常非常非常非常非常非常好看
title: 好看的不得了
star: "4"

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object|false|none|
|» mid|body|string|true|none|
|» comment|body|string|true|none|
|» title|body|string|true|none|
|» star|body|string|true|none|

> 返回示例

> 成功

```json
{
  "data": "评论成功",
  "status": "true"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# Movie

## GET 分类找电影

GET /api/movie/sort

> Body 请求参数

```yaml
tag: string
sort: string
start: string

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object|false|none|
|» tag|body|string|true|none|
|» sort|body|string|true|none|
|» start|body|string|true|none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取电影信息

GET /api/movie/subject/{id}

> Body 请求参数

```yaml
access_token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7IklkIjo0LCJVc2VybmFtZSI6IiIsIlBhc3N3b3JkIjoiMzZiNjI5MjE2MTZkMjQ3MGQwMzI1YjFkNTk0ZjI0YWEiLCJFbWFpbCI6IiIsIlBob25lIjoiMTc3MjY2MzM3NDAiLCJTYWx0IjoiMTY0MjU5MDA5NCIsIkF2YXRhciI6Imh0dHBzOi8vZG91YmFuLWF2YXRhci0xMzA4NzU3Mzg1LmNvcy5hcC1jaG9uZ3FpbmcubXlxY2xvdWQuY29tL2F2dGFyLnBuZ1xyXG4iLCJEb21haW5OYW1lIjoiIiwiSGFiaXRhdCI6IiIsIkhvbWV0b3duIjoiIiwiQmlydGhkYXkiOiI5OTk5LTEyLTEyVDAwOjAwOjAwKzA4OjAwIiwiUmVnRGF0ZSI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiU3RhdGVtZW50IjoiIiwiRm9sbG93ZXJzIjowLCJGb2xsb3dpbmdzIjowfSwiVHlwZSI6IkFDQ0VTU19UT0tFTiIsIlRpbWUiOiIwMDAxLTAxLTAxVDAwOjAwOjAwWiIsImV4cCI6MTY0NTI5MjcxNiwiaXNzIjoiZG91YmFuIn0.YWOHmkOvk0qmLFe_qDtARbuliUyQUsPT6wr2Pc2UeDw
refresh_token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7IklkIjo0LCJVc2VybmFtZSI6IiIsIlBhc3N3b3JkIjoiMzZiNjI5MjE2MTZkMjQ3MGQwMzI1YjFkNTk0ZjI0YWEiLCJFbWFpbCI6IiIsIlBob25lIjoiMTc3MjY2MzM3NDAiLCJTYWx0IjoiMTY0MjU5MDA5NCIsIkF2YXRhciI6Imh0dHBzOi8vZG91YmFuLWF2YXRhci0xMzA4NzU3Mzg1LmNvcy5hcC1jaG9uZ3FpbmcubXlxY2xvdWQuY29tL2F2dGFyLnBuZ1xyXG4iLCJEb21haW5OYW1lIjoiIiwiSGFiaXRhdCI6IiIsIkhvbWV0b3duIjoiIiwiQmlydGhkYXkiOiI5OTk5LTEyLTEyVDAwOjAwOjAwKzA4OjAwIiwiUmVnRGF0ZSI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiU3RhdGVtZW50IjoiIiwiRm9sbG93ZXJzIjowLCJGb2xsb3dpbmdzIjowfSwiVHlwZSI6IlJFRlJFU0hfVE9LRU4iLCJUaW1lIjoiMDAwMS0wMS0wMVQwMDowMDowMFoiLCJleHAiOjE2NDU4OTcyMTYsImlzcyI6ImRvdWJhbiJ9.TGcRVzK62YuXtS03xm4segKTIQwz23ROqXsTFFA7764

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|
|body|body|object|false|none|
|» access_token|body|string|false|none|
|» refresh_token|body|string|false|none|

> 返回示例

> 成功

```json
{
  "Movies": {
    "mid": 1291543,
    "name": "功夫",
    "stars": 4,
    "date": "2004-12-23T00:00:00+08:00",
    "tags": "喜剧,动作,犯罪,奇幻",
    "avatar": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2219011938.jpg",
    "detail": {
      "nicknames": [
        "功夫3D",
        "Kung Fu Hustle"
      ],
      "director": "周星驰",
      "writers": [
        "曾瑾昌",
        "陈文强",
        "周星驰",
        "霍昕"
      ],
      "characters": [
        "周星驰",
        "元秋",
        "元华",
        "黄圣依",
        "梁小龙",
        "陈国坤",
        "田启文",
        "林子聪",
        "林雪",
        "冯克安",
        "释彦能",
        "冯小刚",
        "袁祥仁",
        "张一白",
        "赵志凌",
        "董志华",
        "何文辉",
        "陈凯师",
        "贾康熙",
        "林子善",
        "任珈锐",
        "王仕颖"
      ],
      "type": [
        "喜剧",
        "动作",
        "犯罪",
        "奇幻"
      ],
      "website": "",
      "region": "中国大陆 / 中国香港",
      "language": "粤语 / 汉语普通话 / 手语",
      "release": "2004-12-23 00:00:00",
      "period": 100,
      "IMDb": "tt0373074"
    },
    "score": {
      "score": "8.70",
      "total_cnt": 942145,
      "five": "50.7%",
      "four": "3630.01%",
      "three": "11.8%",
      "two": "0.9%",
      "one": "0.2%"
    },
    "plot": "1940年代的上海，自小受尽欺辱的街头混混阿星（周星驰）为了能出人头地，可谓窥见机会的缝隙就往里钻，今次他盯上行动日益猖獗的黑道势力“斧头帮”，想借之大名成就大业。阿星假冒“斧头帮”成员试图在一个叫“猪笼城寨”的地方对居民敲诈，不想引来真的“斧头帮”与“猪笼城寨”居民的恩怨。“猪笼城寨”原是藏龙卧虎之处，居民中有许多身怀绝技者（元华、梁小龙等），他们隐藏于此本是为远离江湖恩怨，不想麻烦自动上身，躲都躲不及。而在观战正邪两派的斗争中，阿星逐渐领悟功夫的真谛。",
    "celebrities": [
      1048026,
      1299569,
      1280434,
      1287182,
      1304876,
      1050240,
      1005110,
      1274267,
      1229775,
      1315882,
      1314402,
      1274279,
      1283563,
      1202926,
      1274255,
      1301574,
      1275554,
      1314248,
      1274936,
      1322073,
      1319710,
      1333924,
      1314861,
      1321155,
      1364721
    ]
  },
  "MyShortComment": [
    {
      "Id": 8,
      "Mid": 1291543,
      "Uid": 4,
      "Username": "",
      "Avatar": "",
      "Static": "看过",
      "Comment": "很好看",
      "Time": "2022-02-19T23:34:37+08:00",
      "Help": 0,
      "Report": 0,
      "Star": 4
    }
  ],
  "acess_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7IklkIjo0LCJVc2VybmFtZSI6IiIsIlBhc3N3b3JkIjoiMzZiNjI5MjE2MTZkMjQ3MGQwMzI1YjFkNTk0ZjI0YWEiLCJFbWFpbCI6IiIsIlBob25lIjoiMTc3MjY2MzM3NDAiLCJTYWx0IjoiMTY0MjU5MDA5NCIsIkF2YXRhciI6Imh0dHBzOi8vZG91YmFuLWF2YXRhci0xMzA4NzU3Mzg1LmNvcy5hcC1jaG9uZ3FpbmcubXlxY2xvdWQuY29tL2F2dGFyLnBuZ1xyXG4iLCJEb21haW5OYW1lIjoiIiwiSGFiaXRhdCI6IiIsIkhvbWV0b3duIjoiIiwiQmlydGhkYXkiOiI5OTk5LTEyLTEyVDAwOjAwOjAwKzA4OjAwIiwiUmVnRGF0ZSI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiU3RhdGVtZW50IjoiIiwiRm9sbG93ZXJzIjowLCJGb2xsb3dpbmdzIjowfSwiVHlwZSI6IkFDQ0VTU19UT0tFTiIsIlRpbWUiOiIwMDAxLTAxLTAxVDAwOjAwOjAwWiIsImV4cCI6MTY0NTI5NDIyOCwiaXNzIjoiZG91YmFuIn0.fewaXHzlrSY9yEg7e6nw83AFw8FTOvgJvh0CqD8I20M",
  "comment": [
    {
      "Id": 8,
      "Mid": 1291543,
      "Uid": 4,
      "Username": "mj",
      "Avatar": "https://douban-avatar-1308757385.cos.ap-chongqing.myqcloud.com/avtar.png\r\n",
      "Static": "",
      "Comment": "很好看",
      "Time": "2022-02-19T23:34:37+08:00",
      "Help": 0,
      "Report": 0,
      "Star": 4
    }
  ],
  "discussion": null,
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7IklkIjo0LCJVc2VybmFtZSI6IiIsIlBhc3N3b3JkIjoiMzZiNjI5MjE2MTZkMjQ3MGQwMzI1YjFkNTk0ZjI0YWEiLCJFbWFpbCI6IiIsIlBob25lIjoiMTc3MjY2MzM3NDAiLCJTYWx0IjoiMTY0MjU5MDA5NCIsIkF2YXRhciI6Imh0dHBzOi8vZG91YmFuLWF2YXRhci0xMzA4NzU3Mzg1LmNvcy5hcC1jaG9uZ3FpbmcubXlxY2xvdWQuY29tL2F2dGFyLnBuZ1xyXG4iLCJEb21haW5OYW1lIjoiIiwiSGFiaXRhdCI6IiIsIkhvbWV0b3duIjoiIiwiQmlydGhkYXkiOiI5OTk5LTEyLTEyVDAwOjAwOjAwKzA4OjAwIiwiUmVnRGF0ZSI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiU3RhdGVtZW50IjoiIiwiRm9sbG93ZXJzIjowLCJGb2xsb3dpbmdzIjowfSwiVHlwZSI6IlJFRlJFU0hfVE9LRU4iLCJUaW1lIjoiMDAwMS0wMS0wMVQwMDowMDowMFoiLCJleHAiOjE2NDU4OTg3MjgsImlzcyI6ImRvdWJhbiJ9.M2aPDtnyB-jTPmQBHdLOGBfu9Nr-kXJgAuQ5xW_ziX0",
  "reviews": null,
  "status": "true"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 电影排行榜

GET /api/movie/chart

> Body 请求参数

```yaml
start: "1"

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object|false|none|
|» start|body|string|true|none|

> 返回示例

> 成功

```json
{
  "data": [
    {
      "mid": 34960130,
      "name": "深宅",
      "tags": "剧情,悬疑,惊悚,恐怖",
      "avatar": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2583403599.jpg",
      "detail": {
        "nicknames": [
          "毒魔：血战大屠杀(港)"
        ],
        "director": "亚历山大·布斯蒂罗,朱利安·莫利",
        "writers": [
          "赵婷",
          "帕特里克·伯利"
        ],
        "characters": [
          "嘉玛·陈",
          "理查德·麦登",
          "安吉丽娜·朱莉"
        ],
        "type": [
          "喜剧",
          "动画",
          "冒险",
          "古装"
        ],
        "website": "https://www.unifrance.org/film/48781/the-deep-house",
        "region": "法国 / 比利时",
        "language": "英语 / 法语",
        "release": "2021-06-30 00:00:00",
        "period": 85,
        "IMDb": "tt11686490"
      },
      "score": {
        "score": "5.9",
        "total_cnt": 6840,
        "five": "5.0%",
        "four": "16.2%",
        "three": "53.0%",
        "two": "21.6%",
        "one": "4.2%"
      }
    },
    {
      "mid": 30223888,
      "name": "永恒族",
      "tags": "动作,科幻,奇幻,冒险",
      "avatar": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2677303737.jpg",
      "detail": {
        "nicknames": [
          "毒魔：血战大屠杀(港)",
          "猛毒2：血蜘蛛(台)"
        ],
        "director": "赵婷",
        "writers": [
          "詹姆斯·卡梅隆",
          "查尔斯·H·伊格里",
          "乔什·弗莱德曼",
          "大卫·S·高耶",
          "贾斯汀·罗德斯"
        ],
        "characters": [
          "汤姆·哈迪",
          "伍迪·哈里森",
          "米歇尔·威廉姆斯",
          "娜奥米·哈里斯",
          "瑞德·斯科特",
          "斯蒂芬·格拉汉姆",
          "汤姆·赫兰德",
          "里斯·谢尔史密斯",
          "肖恩·德兰尼",
          "佩吉·陆",
          "劳伦斯·斯佩尔曼",
          "杰克·班戴拉",
          "斯克鲁比斯·皮普",
          "阿姆·阿勒卡迪",
          "斯图尔特·亚历山大",
          "克里斯托弗戈德温",
          "格雷戈·洛基特",
          "桑尼·阿什本·瑟金斯",
          "雷切尔·汉德肖",
          "小田部阿基",
          "埃里克·西格蒙德森",
          "何塞‧帕尔马",
          "Rosie Marcel",
          "艾略特·凯博尔",
          "威廉·W·巴伯",
          "埃尔文·费利西达"
        ],
        "type": [
          "喜剧",
          "动画",
          "冒险",
          "古装"
        ],
        "website": "",
        "region": "美国",
        "language": "英语 / 美国手语 / 西班牙语 / 拉丁语 / 古希腊语",
        "release": "2021-11-05 00:00:00",
        "period": 157,
        "IMDb": "tt9032400"
      },
      "score": {
        "score": "5.9",
        "total_cnt": 88051,
        "five": "3.7%",
        "four": "18.9%",
        "three": "49.1%",
        "two": "22.8%",
        "one": "5.5%"
      }
    },
    {
      "mid": 35073565,
      "name": "门锁",
      "tags": "剧情,惊悚,犯罪",
      "avatar": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2685382025.jpg",
      "detail": {
        "nicknames": [
          "毒魔：血战大屠杀(港)"
        ],
        "director": "别克",
        "writers": [
          "詹姆斯·卡梅隆",
          "查尔斯·H·伊格里",
          "乔什·弗莱德曼",
          "大卫·S·高耶"
        ],
        "characters": [
          "汤姆·哈迪",
          "伍迪·哈里森",
          "米歇尔·威廉姆斯",
          "娜奥米·哈里斯",
          "瑞德·斯科特",
          "斯蒂芬·格拉汉姆",
          "汤姆·赫兰德",
          "里斯·谢尔史密斯",
          "肖恩·德兰尼",
          "佩吉·陆",
          "劳伦斯·斯佩尔曼",
          "杰克·班戴拉",
          "斯克鲁比斯·皮普"
        ],
        "type": [
          "喜剧",
          "动画",
          "冒险"
        ],
        "website": "",
        "region": "中国大陆",
        "language": "汉语普通话",
        "release": "2021-11-19 00:00:00",
        "period": 105,
        "IMDb": "tt13010918"
      },
      "score": {
        "score": "4.4",
        "total_cnt": 111465,
        "five": "1.1%",
        "four": "4.4%",
        "three": "26.9%",
        "two": "46.0%",
        "one": "21.5%"
      }
    },
    {
      "mid": 27605563,
      "name": "鹿角",
      "tags": "恐怖",
      "avatar": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2587966857.jpg",
      "detail": {
        "nicknames": [
          "毒魔：血战大屠杀(港)",
          "猛毒2：血蜘蛛(台)"
        ],
        "director": "斯科特·库珀",
        "writers": [
          "詹姆斯·卡梅隆",
          "查尔斯·H·伊格里"
        ],
        "characters": [
          "汤姆·哈迪",
          "伍迪·哈里森",
          "米歇尔·威廉姆斯",
          "娜奥米·哈里斯",
          "瑞德·斯科特",
          "斯蒂芬·格拉汉姆",
          "汤姆·赫兰德",
          "里斯·谢尔史密斯",
          "肖恩·德兰尼",
          "佩吉·陆",
          "劳伦斯·斯佩尔曼",
          "杰克·班戴拉",
          "斯克鲁比斯·皮普",
          "阿姆·阿勒卡迪",
          "斯图尔特·亚历山大",
          "克里斯托弗戈德温",
          "格雷戈·洛基特",
          "桑尼·阿什本·瑟金斯"
        ],
        "type": [
          "喜剧"
        ],
        "website": "",
        "region": "美国 / 墨西哥 / 加拿大",
        "language": "英语",
        "release": "2021-10-29 00:00:00",
        "period": 99,
        "IMDb": "tt7740510"
      },
      "score": {
        "score": "5.3",
        "total_cnt": 4217,
        "five": "3.2%",
        "four": "8.9%",
        "three": "47.1%",
        "two": "33.4%",
        "one": "7.4%"
      }
    },
    {
      "mid": 30382416,
      "name": "毒液2",
      "tags": "动作,科幻,惊悚",
      "avatar": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2675102928.jpg",
      "detail": {
        "nicknames": [
          "小黄人",
          "迷你兵团(港)",
          "小小兵(台)",
          "小黄人大电影",
          "终结者2019",
          "终结者6",
          "Terminator 6"
        ],
        "director": "安迪·瑟金斯",
        "writers": [
          "詹姆斯·卡梅隆",
          "查尔斯·H·伊格里"
        ],
        "characters": [
          "基努·里维斯",
          "凯瑞-安·莫斯",
          "叶海亚·阿卜杜勒-迈丁",
          "乔纳森·格罗夫",
          "杰西卡·亨维克",
          "尼尔·帕特里克·哈里斯",
          "贾达·萍克·史密斯",
          "佩丽冉卡·曹帕拉",
          "克里斯蒂娜·里奇",
          "朗贝尔·维尔森",
          "安德鲁·卡德威尔",
          "托比·奥伍梅尔",
          "马克思·雷迈特",
          "约书亚·格罗斯",
          "布莱恩·J·史密斯",
          "埃伦迪拉·伊瓦拉",
          "迈克尔·X·萨默斯",
          "马克斯·毛夫",
          "帕鲁布·科里",
          "弗莉玛·阿吉曼",
          "安德鲁·罗斯尼",
          "里奥·盛",
          "特尔玛·霍普金斯",
          "约翰·盖耶塔",
          "Donald Mustard",
          "基姆·利布莱利",
          "查德·斯塔赫斯基",
          "朱利安·格雷",
          "陈虎",
          "史蒂芬·邓利维",
          "艾伦·霍尔曼",
          "伊恩·皮里",
          "Nicolas de Pruyssenaere",
          "梅西·麦克利"
        ],
        "type": [
          "喜剧",
          "动画",
          "冒险"
        ],
        "website": "",
        "region": "美国",
        "language": "英语",
        "release": "2021-10-01 00:00:00",
        "period": 97,
        "IMDb": "tt7097896"
      },
      "score": {
        "score": "5.0",
        "total_cnt": 114043,
        "five": "1.2%",
        "four": "5.6%",
        "three": "44.1%",
        "two": "40.4%",
        "one": "8.7%"
      }
    },
    {
      "mid": 26140265,
      "name": "新生化危机",
      "tags": "动作,科幻,恐怖",
      "avatar": "https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2692391480.jpg",
      "detail": {
        "nicknames": [
          "小黄人",
          "迷你兵团(港)",
          "小小兵(台)",
          "小黄人大电影"
        ],
        "director": "约翰内斯·罗伯茨",
        "writers": [
          "詹姆斯·卡梅隆"
        ],
        "characters": [
          "基努·里维斯",
          "凯瑞-安·莫斯",
          "叶海亚·阿卜杜勒-迈丁",
          "乔纳森·格罗夫",
          "杰西卡·亨维克",
          "尼尔·帕特里克·哈里斯",
          "贾达·萍克·史密斯",
          "佩丽冉卡·曹帕拉",
          "克里斯蒂娜·里奇",
          "朗贝尔·维尔森",
          "安德鲁·卡德威尔",
          "托比·奥伍梅尔",
          "马克思·雷迈特",
          "约书亚·格罗斯",
          "布莱恩·J·史密斯",
          "埃伦迪拉·伊瓦拉",
          "迈克尔·X·萨默斯",
          "马克斯·毛夫",
          "帕鲁布·科里"
        ],
        "type": [
          "喜剧",
          "动画",
          "冒险"
        ],
        "website": "",
        "region": "德国 / 美国 / 加拿大 / 法国",
        "language": "英语",
        "release": "2021-11-24 00:00:00",
        "period": 107,
        "IMDb": "tt6920084"
      },
      "score": {
        "score": "4.6",
        "total_cnt": 15559,
        "five": "3.1%",
        "four": "6.5%",
        "three": "29.6%",
        "two": "37.9%",
        "one": "23.0%"
      }
    },
    {
      "mid": 34801038,
      "name": "黑客帝国：矩阵重启",
      "tags": "动作,科幻",
      "avatar": "https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2844387600.jpg",
      "detail": {
        "nicknames": [
          "小黄人",
          "迷你兵团(港)",
          "小小兵(台)",
          "小黄人大电影",
          "终结者2019",
          "终结者6"
        ],
        "director": "拉娜·沃卓斯基",
        "writers": [
          "詹姆斯·卡梅隆",
          "查尔斯·H·伊格里",
          "乔什·弗莱德曼",
          "大卫·S·高耶"
        ],
        "characters": [
          "桑德拉·布洛克",
          "乔恩·哈姆",
          "迈克尔·基顿",
          "艾莉森·珍妮",
          "史蒂夫·库根",
          "珍妮弗·桑德斯",
          "杰弗里·拉什",
          "史蒂夫·卡瑞尔",
          "皮埃尔·柯芬",
          "凯蒂·米克松",
          "迈克尔·贝亚蒂",
          "真田广之",
          "大卫·罗森鲍姆",
          "亚历克斯·道丁",
          "保罗·索恩利",
          "约翰·盖蒂尔",
          "彼得·奥蒙德",
          "洛娜·布朗",
          "彼得·舒勒",
          "李绍祥",
          "迪恩·查尔斯·查普曼",
          "大卫·巴特勒",
          "安德烈·雅各布斯",
          "泰瑞·诺顿",
          "李恬洁",
          "贝乐乐",
          "肯尼斯·以色列",
          "小佩里·祖鲁",
          "肖恩·雷卡特尔",
          "Rodney L. James",
          "艾伦·霍尔曼",
          "伊恩·皮里",
          "Nicolas de Pruyssenaere",
          "梅西·麦克利",
          "詹姆斯·麦克特格",
          "威廉·W·巴伯",
          "丹尼尔·伯哈特",
          "埃尔文·费利西达",
          "迈克尔·J·格温",
          "Linda Joy Henry",
          "约翰·洛巴托",
          "安妮-玛丽·奥尔森",
          "艾蒂安·维克",
          "克莱顿·华生",
          "詹姆斯·D·韦斯顿二世"
        ],
        "type": [
          "喜剧",
          "动画"
        ],
        "website": "thechoiceisyours.whatisthematrix.com",
        "region": "美国",
        "language": "英语",
        "release": "2022-01-14 00:00:00",
        "period": 148,
        "IMDb": "tt10838180"
      },
      "score": {
        "score": "5.7",
        "total_cnt": 84099,
        "five": "3.9%",
        "four": "15.0%",
        "three": "48.4%",
        "two": "27.0%",
        "one": "5.7%"
      }
    },
    {
      "mid": 6982558,
      "name": "长城",
      "tags": "动作,奇幻,冒险",
      "avatar": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2400082527.jpg",
      "detail": {
        "nicknames": [
          "小黄人"
        ],
        "director": "张艺谋",
        "writers": [
          "詹姆斯·卡梅隆",
          "查尔斯·H·伊格里",
          "乔什·弗莱德曼",
          "大卫·S·高耶",
          "贾斯汀·罗德斯",
          "比利·雷"
        ],
        "characters": [
          "桑德拉·布洛克",
          "乔恩·哈姆",
          "迈克尔·基顿",
          "艾莉森·珍妮",
          "史蒂夫·库根",
          "珍妮弗·桑德斯",
          "杰弗里·拉什",
          "史蒂夫·卡瑞尔",
          "皮埃尔·柯芬",
          "凯蒂·米克松",
          "迈克尔·贝亚蒂",
          "真田广之",
          "大卫·罗森鲍姆",
          "亚历克斯·道丁",
          "保罗·索恩利",
          "约翰·盖蒂尔",
          "彼得·奥蒙德",
          "洛娜·布朗",
          "彼得·舒勒",
          "李绍祥",
          "迪恩·查尔斯·查普曼"
        ],
        "type": [
          "喜剧",
          "动画",
          "冒险"
        ],
        "website": "",
        "region": "中国大陆 / 美国",
        "language": "英语 / 汉语普通话",
        "release": "2016-12-16 00:00:00",
        "period": 104,
        "IMDb": "tt2034800"
      },
      "score": {
        "score": "4.9",
        "total_cnt": 372441,
        "five": "5.2%",
        "four": "11.0%",
        "three": "31.4%",
        "two": "26.9%",
        "one": "25.6%"
      }
    },
    {
      "mid": 30378158,
      "name": "秘密访客",
      "tags": "悬疑,惊悚",
      "avatar": "https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2629408730.jpg",
      "detail": {
        "nicknames": [
          "小黄人"
        ],
        "director": "陈正道",
        "writers": [
          "詹姆斯·卡梅隆"
        ],
        "characters": [
          "桑德拉·布洛克",
          "乔恩·哈姆",
          "迈克尔·基顿",
          "艾莉森·珍妮",
          "史蒂夫·库根",
          "珍妮弗·桑德斯",
          "杰弗里·拉什",
          "史蒂夫·卡瑞尔",
          "皮埃尔·柯芬"
        ],
        "type": [
          "喜剧",
          "动画"
        ],
        "website": "",
        "region": "中国大陆",
        "language": "汉语普通话",
        "release": "2021-05-01 00:00:00",
        "period": 111,
        "IMDb": "tt10097384"
      },
      "score": {
        "score": "5.3",
        "total_cnt": 167421,
        "five": "1.8%",
        "four": "11.4%",
        "three": "45.5%",
        "two": "33.7%",
        "one": "7.6%"
      }
    },
    {
      "mid": 30459571,
      "name": "明日之战",
      "tags": "动作,科幻,冒险",
      "avatar": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2678875868.jpg",
      "detail": {
        "nicknames": [
          "小黄人",
          "迷你兵团(港)",
          "小小兵(台)"
        ],
        "director": "克里斯·麦凯",
        "writers": [
          "詹姆斯·卡梅隆"
        ],
        "characters": [
          "桑德拉·布洛克",
          "乔恩·哈姆",
          "迈克尔·基顿",
          "艾莉森·珍妮",
          "史蒂夫·库根",
          "珍妮弗·桑德斯",
          "杰弗里·拉什",
          "史蒂夫·卡瑞尔",
          "皮埃尔·柯芬",
          "凯蒂·米克松",
          "迈克尔·贝亚蒂",
          "真田广之",
          "大卫·罗森鲍姆",
          "亚历克斯·道丁",
          "保罗·索恩利",
          "约翰·盖蒂尔",
          "彼得·奥蒙德",
          "洛娜·布朗",
          "彼得·舒勒",
          "李绍祥",
          "迪恩·查尔斯·查普曼",
          "大卫·巴特勒",
          "安德烈·雅各布斯",
          "泰瑞·诺顿",
          "李恬洁",
          "贝乐乐",
          "肯尼斯·以色列",
          "小佩里·祖鲁",
          "肖恩·雷卡特尔",
          "Rodney L. James"
        ],
        "type": [
          "喜剧",
          "动画",
          "冒险"
        ],
        "website": "",
        "region": "美国",
        "language": "英语",
        "release": "2021-09-03 00:00:00",
        "period": 138,
        "IMDb": "tt9777666"
      },
      "score": {
        "score": "5.9",
        "total_cnt": 95719,
        "five": "2.4%",
        "four": "17.0%",
        "three": "56.0%",
        "two": "20.8%",
        "one": "3.7%"
      }
    },
    {
      "mid": 4840388,
      "name": "新喜剧之王",
      "tags": "剧情,喜剧",
      "avatar": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2555440969.jpg",
      "detail": {
        "nicknames": [
          "小黄人",
          "迷你兵团(港)",
          "小小兵(台)",
          "小黄人大电影"
        ],
        "director": "周星驰,邱礼涛",
        "writers": [
          "詹姆斯·卡梅隆",
          "查尔斯·H·伊格里",
          "乔什·弗莱德曼",
          "大卫·S·高耶"
        ],
        "characters": [
          "桑德拉·布洛克",
          "乔恩·哈姆",
          "迈克尔·基顿",
          "艾莉森·珍妮",
          "史蒂夫·库根",
          "珍妮弗·桑德斯",
          "杰弗里·拉什",
          "史蒂夫·卡瑞尔",
          "皮埃尔·柯芬",
          "凯蒂·米克松"
        ],
        "type": [
          "喜剧",
          "动画"
        ],
        "website": "",
        "region": "中国大陆 / 中国香港",
        "language": "汉语普通话 / 粤语",
        "release": "2019-02-05 00:00:00",
        "period": 91,
        "IMDb": "tt9368628"
      },
      "score": {
        "score": "5.7",
        "total_cnt": 380346,
        "five": "9.5%",
        "four": "16.8%",
        "three": "35.3%",
        "two": "25.2%",
        "one": "13.1%"
      }
    },
    {
      "mid": 30377703,
      "name": "来电狂响",
      "tags": "剧情,喜剧",
      "avatar": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2542268337.jpg",
      "detail": {
        "nicknames": [
          "小黄人",
          "迷你兵团(港)",
          "小小兵(台)"
        ],
        "director": "于淼",
        "writers": [
          "詹姆斯·卡梅隆",
          "查尔斯·H·伊格里",
          "乔什·弗莱德曼"
        ],
        "characters": [
          "桑德拉·布洛克",
          "乔恩·哈姆",
          "迈克尔·基顿",
          "艾莉森·珍妮",
          "史蒂夫·库根",
          "珍妮弗·桑德斯",
          "杰弗里·拉什",
          "史蒂夫·卡瑞尔",
          "皮埃尔·柯芬",
          "凯蒂·米克松",
          "迈克尔·贝亚蒂",
          "真田广之",
          "大卫·罗森鲍姆",
          "亚历克斯·道丁",
          "保罗·索恩利",
          "约翰·盖蒂尔",
          "彼得·奥蒙德",
          "洛娜·布朗",
          "彼得·舒勒",
          "李绍祥",
          "迪恩·查尔斯·查普曼",
          "大卫·巴特勒",
          "安德烈·雅各布斯",
          "泰瑞·诺顿",
          "李恬洁",
          "贝乐乐"
        ],
        "type": [
          "喜剧",
          "动画"
        ],
        "website": "",
        "region": "中国大陆",
        "language": "汉语普通话",
        "release": "2018-12-28 00:00:00",
        "period": 103,
        "IMDb": "tt9408490"
      },
      "score": {
        "score": "5.7",
        "total_cnt": 268845,
        "five": "3.6%",
        "four": "17.0%",
        "three": "47.9%",
        "two": "25.6%",
        "one": "5.9%"
      }
    },
    {
      "mid": 26935283,
      "name": "侍神令",
      "tags": "奇幻",
      "avatar": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2629260713.jpg",
      "detail": {
        "nicknames": [
          "小黄人",
          "迷你兵团(港)",
          "小小兵(台)"
        ],
        "director": "李蔚然",
        "writers": [
          "詹姆斯·卡梅隆",
          "查尔斯·H·伊格里"
        ],
        "characters": [
          "桑德拉·布洛克",
          "乔恩·哈姆",
          "迈克尔·基顿",
          "艾莉森·珍妮",
          "史蒂夫·库根",
          "珍妮弗·桑德斯",
          "杰弗里·拉什",
          "史蒂夫·卡瑞尔",
          "皮埃尔·柯芬",
          "凯蒂·米克松"
        ],
        "type": [
          "喜剧"
        ],
        "website": "",
        "region": "中国大陆",
        "language": "汉语普通话",
        "release": "2021-02-12 00:00:00",
        "period": 120,
        "IMDb": "tt12151820"
      },
      "score": {
        "score": "5.3",
        "total_cnt": 173455,
        "five": "2.6%",
        "four": "11.3%",
        "three": "44.6%",
        "two": "32.8%",
        "one": "8.6%"
      }
    },
    {
      "mid": 26575103,
      "name": "捉妖记2",
      "tags": "喜剧,动作,奇幻,古装",
      "avatar": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2509643816.jpg",
      "detail": {
        "nicknames": [
          "小黄人"
        ],
        "director": "许诚毅",
        "writers": [
          "詹姆斯·卡梅隆",
          "查尔斯·H·伊格里",
          "乔什·弗莱德曼"
        ],
        "characters": [
          "桑德拉·布洛克",
          "乔恩·哈姆",
          "迈克尔·基顿",
          "艾莉森·珍妮",
          "史蒂夫·库根",
          "珍妮弗·桑德斯",
          "杰弗里·拉什",
          "史蒂夫·卡瑞尔",
          "皮埃尔·柯芬",
          "凯蒂·米克松",
          "迈克尔·贝亚蒂",
          "真田广之",
          "大卫·罗森鲍姆",
          "亚历克斯·道丁",
          "保罗·索恩利",
          "约翰·盖蒂尔",
          "彼得·奥蒙德",
          "洛娜·布朗",
          "彼得·舒勒",
          "李绍祥",
          "迪恩·查尔斯·查普曼",
          "大卫·巴特勒",
          "安德烈·雅各布斯"
        ],
        "type": [
          "喜剧",
          "动画",
          "冒险",
          "古装"
        ],
        "website": "",
        "region": "中国大陆 / 中国香港",
        "language": "汉语普通话",
        "release": "2018-02-16 00:00:00",
        "period": 110,
        "IMDb": "tt6170484"
      },
      "score": {
        "score": "4.9",
        "total_cnt": 334855,
        "five": "2.4%",
        "four": "7.7%",
        "three": "37.7%",
        "two": "37.6%",
        "one": "14.6%"
      }
    },
    {
      "mid": 26806528,
      "name": "一呼一吸",
      "tags": "剧情,爱情,传记",
      "avatar": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2493495877.jpg",
      "detail": {
        "nicknames": [
          "小黄人",
          "迷你兵团(港)",
          "小小兵(台)"
        ],
        "director": "安迪·瑟金斯",
        "writers": [
          "詹姆斯·卡梅隆"
        ],
        "characters": [
          "桑德拉·布洛克",
          "乔恩·哈姆",
          "迈克尔·基顿",
          "艾莉森·珍妮",
          "史蒂夫·库根",
          "珍妮弗·桑德斯",
          "杰弗里·拉什",
          "史蒂夫·卡瑞尔",
          "皮埃尔·柯芬",
          "凯蒂·米克松",
          "迈克尔·贝亚蒂",
          "真田广之",
          "大卫·罗森鲍姆",
          "亚历克斯·道丁",
          "保罗·索恩利",
          "约翰·盖蒂尔",
          "彼得·奥蒙德",
          "洛娜·布朗",
          "彼得·舒勒",
          "李绍祥",
          "迪恩·查尔斯·查普曼",
          "大卫·巴特勒",
          "安德烈·雅各布斯",
          "泰瑞·诺顿"
        ],
        "type": [
          "喜剧",
          "动画",
          "冒险"
        ],
        "website": "",
        "region": "英国",
        "language": "英语",
        "release": "2017-09-11 00:00:00",
        "period": 118,
        "IMDb": "tt5716464"
      },
      "score": {
        "score": "7.9",
        "total_cnt": 31535,
        "five": "25.8%",
        "four": "48.0%",
        "three": "24.4%",
        "two": "1.6%",
        "one": "0.2%"
      }
    },
    {
      "mid": 26679552,
      "name": "推销员",
      "tags": "剧情,悬疑,家庭",
      "avatar": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2386724199.jpg",
      "detail": {
        "nicknames": [
          "小黄人",
          "迷你兵团(港)",
          "小小兵(台)",
          "小黄人大电影",
          "终结者2019"
        ],
        "director": "阿斯哈·法哈蒂",
        "writers": [
          "詹姆斯·卡梅隆"
        ],
        "characters": [
          "桑德拉·布洛克",
          "乔恩·哈姆",
          "迈克尔·基顿",
          "艾莉森·珍妮",
          "史蒂夫·库根"
        ],
        "type": [
          "喜剧",
          "动画",
          "冒险"
        ],
        "website": "",
        "region": "伊朗 / 法国",
        "language": "波斯语",
        "release": "2016-05-21 00:00:00",
        "period": 125,
        "IMDb": "tt5186714"
      },
      "score": {
        "score": "7.6",
        "total_cnt": 30125,
        "five": "15.6%",
        "four": "54.3%",
        "three": "27.5%",
        "two": "2.1%",
        "one": "0.5%"
      }
    },
    {
      "mid": 1978369,
      "name": "我的机器人女友",
      "tags": "喜剧,爱情,科幻",
      "avatar": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p637412789.jpg",
      "detail": {
        "nicknames": [
          "小黄人",
          "迷你兵团(港)",
          "小小兵(台)"
        ],
        "director": "郭在容",
        "writers": [
          "詹姆斯·卡梅隆"
        ],
        "characters": [
          "桑德拉·布洛克",
          "乔恩·哈姆",
          "迈克尔·基顿",
          "艾莉森·珍妮",
          "史蒂夫·库根",
          "珍妮弗·桑德斯",
          "杰弗里·拉什",
          "史蒂夫·卡瑞尔",
          "皮埃尔·柯芬",
          "凯蒂·米克松",
          "迈克尔·贝亚蒂"
        ],
        "type": [
          "喜剧",
          "动画",
          "冒险"
        ],
        "website": "",
        "region": "日本",
        "language": "日语",
        "release": "2009-07-28 00:00:00",
        "period": 115,
        "IMDb": "tt0929860"
      },
      "score": {
        "score": "7.5",
        "total_cnt": 157357,
        "five": "19.3%",
        "four": "43.0%",
        "three": "33.2%",
        "two": "3.9%",
        "one": "0.6%"
      }
    },
    {
      "mid": 26366465,
      "name": "我的少女时代",
      "tags": "喜剧,爱情",
      "avatar": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2285115802.jpg",
      "detail": {
        "nicknames": [
          "小黄人"
        ],
        "director": "陈玉珊",
        "writers": [
          "詹姆斯·卡梅隆"
        ],
        "characters": [
          "桑德拉·布洛克",
          "乔恩·哈姆",
          "迈克尔·基顿",
          "艾莉森·珍妮",
          "史蒂夫·库根",
          "珍妮弗·桑德斯",
          "杰弗里·拉什",
          "史蒂夫·卡瑞尔",
          "皮埃尔·柯芬",
          "凯蒂·米克松",
          "迈克尔·贝亚蒂",
          "真田广之",
          "大卫·罗森鲍姆",
          "亚历克斯·道丁",
          "保罗·索恩利",
          "约翰·盖蒂尔",
          "彼得·奥蒙德",
          "洛娜·布朗",
          "彼得·舒勒",
          "李绍祥"
        ],
        "type": [
          "喜剧",
          "动画"
        ],
        "website": "",
        "region": "中国台湾",
        "language": "汉语普通话",
        "release": "2015-11-19 00:00:00",
        "period": 134,
        "IMDb": "tt4967094"
      },
      "score": {
        "score": "7.8",
        "total_cnt": 556305,
        "five": "25.5%",
        "four": "45.6%",
        "three": "25.0%",
        "two": "3.0%",
        "one": "0.8%"
      }
    },
    {
      "mid": 27109633,
      "name": "终结者：黑暗命运",
      "tags": "动作,科幻,冒险",
      "avatar": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2571762536.jpg",
      "detail": {
        "nicknames": [
          "小黄人",
          "迷你兵团(港)",
          "小小兵(台)",
          "小黄人大电影",
          "终结者2019",
          "终结者6",
          "Terminator 6"
        ],
        "director": "蒂姆·米勒",
        "writers": [
          "布莱恩·林奇",
          "查尔斯·H·伊格里",
          "乔什·弗莱德曼",
          "大卫·S·高耶",
          "贾斯汀·罗德斯",
          "比利·雷",
          "盖尔·安妮·赫德"
        ],
        "characters": [
          "桑德拉·布洛克",
          "乔恩·哈姆",
          "迈克尔·基顿",
          "艾莉森·珍妮",
          "史蒂夫·库根",
          "珍妮弗·桑德斯",
          "杰弗里·拉什",
          "史蒂夫·卡瑞尔",
          "皮埃尔·柯芬",
          "凯蒂·米克松",
          "迈克尔·贝亚蒂",
          "真田广之",
          "大卫·罗森鲍姆",
          "亚历克斯·道丁",
          "保罗·索恩利",
          "约翰·盖蒂尔",
          "彼得·奥蒙德",
          "洛娜·布朗",
          "彼得·舒勒"
        ],
        "type": [
          "喜剧",
          "动画",
          "冒险"
        ],
        "website": "",
        "region": "美国 / 西班牙 / 匈牙利",
        "language": "英语 / 西班牙语",
        "release": "2019-11-01 00:00:00",
        "period": 128,
        "IMDb": "tt6450804"
      },
      "score": {
        "score": "6.8",
        "total_cnt": 157185,
        "five": "7.9%",
        "four": "34.4%",
        "three": "47.9%",
        "two": "8.7%",
        "one": "1.1%"
      }
    },
    {
      "mid": 11624706,
      "name": "小黄人大眼萌",
      "tags": "喜剧,动画",
      "avatar": "https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2265761240.jpg",
      "detail": {
        "nicknames": [
          "小黄人",
          "迷你兵团(港)",
          "小小兵(台)",
          "小黄人大电影"
        ],
        "director": "凯尔·巴尔达,皮埃尔·柯芬",
        "writers": [
          "布莱恩·林奇"
        ],
        "characters": [
          "桑德拉·布洛克",
          "乔恩·哈姆",
          "迈克尔·基顿",
          "艾莉森·珍妮",
          "史蒂夫·库根",
          "珍妮弗·桑德斯",
          "杰弗里·拉什",
          "史蒂夫·卡瑞尔",
          "皮埃尔·柯芬",
          "凯蒂·米克松",
          "迈克尔·贝亚蒂",
          "真田广之",
          "大卫·罗森鲍姆",
          "亚历克斯·道丁",
          "保罗·索恩利"
        ],
        "type": [
          "喜剧",
          "动画"
        ],
        "website": "www.minionsmovie.com",
        "region": "美国",
        "language": "英语 / 西班牙语",
        "release": "2015-09-13 00:00:00",
        "period": 91,
        "IMDb": "tt2293640"
      },
      "score": {
        "score": "7.7",
        "total_cnt": 255875,
        "five": "22.7%",
        "four": "43.6%",
        "three": "29.5%",
        "two": "3.6%",
        "one": "0.6%"
      }
    }
  ],
  "status": "true"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# Search

## GET 搜索

GET /api/movie/search

可以搜索电影 和影人

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|key|query|string|true|none|

> 返回示例

> 成功

```json
{
  "celebrity": null,
  "movie": [
    {
      "mid": 1295644,
      "name": "这个杀手不太冷",
      "tags": "剧情,动作,犯罪",
      "avatar": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p511118051.jpg",
      "detail": {
        "nicknames": [
          "地球不死人(港)",
          "这个人来自洞穴",
          "来自地穴的男人",
          "穴居人",
          "地底奇人"
        ],
        "director": "吕克·贝松",
        "writers": [
          "杰罗姆·比克斯比"
        ],
        "characters": [
          "大卫·李·史密斯",
          "托尼·托德",
          "约翰·比灵斯列",
          "安妮卡·彼得森",
          "阿丽西丝·索普",
          "威廉姆·卡特",
          "理查德·雷西尔",
          "艾伦·克劳福德",
          "史蒂文·利特尔",
          "罗比·布赖恩",
          "麦温",
          "乔治·马丁",
          "罗伯特·拉萨多",
          "亚当·布斯奇",
          "马里奥·托迪斯科",
          "萨米·纳塞利"
        ],
        "type": [
          "剧情",
          "科幻",
          "犯罪"
        ],
        "website": "",
        "region": "法国 / 美国",
        "language": "英语 / 意大利语 / 法语",
        "release": "1994-09-14 00:00:00",
        "period": 110,
        "IMDb": "tt0110413"
      },
      "score": {
        "score": "9.4",
        "total_cnt": 2077617,
        "five": "73.9%",
        "four": "22.7%",
        "three": "3.2%",
        "two": "0.2%",
        "one": "0.1%"
      }
    },
    {
      "mid": 2300586,
      "name": "这个男人来自地球",
      "tags": "剧情,科幻",
      "avatar": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p513303986.jpg",
      "detail": {
        "nicknames": [
          "地球不死人(港)",
          "这个人来自洞穴",
          "来自地穴的男人",
          "穴居人",
          "地底奇人",
          "长生不老",
          "来自地球的男人"
        ],
        "director": "理查德·沙因克曼",
        "writers": [
          "杰罗姆·比克斯比"
        ],
        "characters": [
          "大卫·李·史密斯",
          "托尼·托德",
          "约翰·比灵斯列",
          "安妮卡·彼得森",
          "阿丽西丝·索普",
          "威廉姆·卡特",
          "理查德·雷西尔",
          "艾伦·克劳福德",
          "史蒂文·利特尔",
          "罗比·布赖恩"
        ],
        "type": [
          "剧情",
          "科幻"
        ],
        "website": "",
        "region": "美国",
        "language": "英语",
        "release": "2007-11-13 00:00:00",
        "period": 87,
        "IMDb": "tt0756683"
      },
      "score": {
        "score": "8.5",
        "total_cnt": 336867,
        "five": "45.7%",
        "four": "37.9%",
        "three": "13.6%",
        "two": "2.0%",
        "one": "0.8%"
      }
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# Celebrity

## GET 获取影人信息

GET /api/movie/celebrity/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|

> 返回示例

> 成功

```json
{
  "data": [
    {
      "Id": 1000002,
      "Name": "乔纳森·格罗夫",
      "Avatar": "Jonathan",
      "EnglishName": "男",
      "Gender": "白羊座",
      "Sign": "1985年03月26日",
      "Birth": "美国,宾夕法尼亚州,兰开斯特",
      "Hometown": "演员 / 配音",
      "Job": "nm2676147",
      "IMDb": "乔纳森·格罗夫，美国演员，歌手。1985年3月26日出生于美国宾夕法尼亚州兰开斯特，他在阿米什社区长大，父亲是门诺教徒。乔纳森·格罗夫的职业生涯开始于舞台剧，他在2007年横扫8项托尼奖的音乐剧《春之觉醒》中饰演男一号 Melchior，22岁即得到了托尼奖音乐剧男主角提名。2015年他出演了大热百老汇音乐剧《汉密尔顿》，收获了格莱美奖和托尼奖音乐剧男配角提名。 电视剧方面，他在《欢乐合唱团》中饰演Jesse St. James；HBO台的同志题材电视剧《寻》中饰演男主角Patrick Murray；N...(展开全部)\n乔纳森·格罗夫，美国演员，歌手。1985年3月26日出生于美国宾夕法尼亚州兰开斯特，他在阿米什社区长大，父亲是门诺教徒。乔纳森·格罗夫的职业生涯开始于舞台剧，他在2007年横扫8项托尼奖的音乐剧《春之觉醒》中饰演男一号 Melchior，22岁即得到了托尼奖音乐剧男主角提名。2015年他出演了大热百老汇音乐剧《汉密尔顿》，收获了格莱美奖和托尼奖音乐剧男配角提名。 电视剧方面，他在《欢乐合唱团》中饰演Jesse St. James；HBO台的同志题材电视剧《寻》中饰演男主角Patrick Murray；Netflix出品的《心灵猎人》中饰演男主角Holden Ford。电影方面出演了李安执导的《制造伍德斯托克音乐节》，为《冰雪奇缘》中男主角Kristoff配音。乔纳森在2009年时出柜。",
      "Brief": "https://img1.doubanio.com/view/celebrity/raw/public/p1391831466.59.jpg"
    }
  ],
  "status": "true"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# MyLook

## GET 获取电影主页信息

GET /api/movie/mine

> Body 请求参数

```yaml
access_token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7IklkIjo0LCJVc2VybmFtZSI6IiIsIlBhc3N3b3JkIjoiMzZiNjI5MjE2MTZkMjQ3MGQwMzI1YjFkNTk0ZjI0YWEiLCJFbWFpbCI6IiIsIlBob25lIjoiMTc3MjY2MzM3NDAiLCJTYWx0IjoiMTY0MjU5MDA5NCIsIkF2YXRhciI6Imh0dHBzOi8vZG91YmFuLWF2YXRhci0xMzA4NzU3Mzg1LmNvcy5hcC1jaG9uZ3FpbmcubXlxY2xvdWQuY29tL2F2YXRhci80ZjA5MjM1YjRlMTc0YzQ1OGY0NzFkN2IxNzM5N2IzNi5wbmciLCJEb21haW5OYW1lIjoiIiwiSGFiaXRhdCI6IiIsIkhvbWV0b3duIjoiIiwiQmlydGhkYXkiOiI5OTk5LTEyLTEyVDAwOjAwOjAwKzA4OjAwIiwiUmVnRGF0ZSI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiU3RhdGVtZW50IjoiIiwiRm9sbG93ZXJzIjowLCJGb2xsb3dpbmdzIjowfSwiVHlwZSI6IkFDQ0VTU19UT0tFTiIsIlRpbWUiOiIwMDAxLTAxLTAxVDAwOjAwOjAwWiIsImV4cCI6MTY0NTQzNTQ4NCwiaXNzIjoiZG91YmFuIn0.r7dJipjD2sZDytAMrwNSy8MrUsqqXRUu9pOSvIWhsng
refresh_token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7IklkIjo0LCJVc2VybmFtZSI6IiIsIlBhc3N3b3JkIjoiMzZiNjI5MjE2MTZkMjQ3MGQwMzI1YjFkNTk0ZjI0YWEiLCJFbWFpbCI6IiIsIlBob25lIjoiMTc3MjY2MzM3NDAiLCJTYWx0IjoiMTY0MjU5MDA5NCIsIkF2YXRhciI6Imh0dHBzOi8vZG91YmFuLWF2YXRhci0xMzA4NzU3Mzg1LmNvcy5hcC1jaG9uZ3FpbmcubXlxY2xvdWQuY29tL2F2YXRhci80ZjA5MjM1YjRlMTc0YzQ1OGY0NzFkN2IxNzM5N2IzNi5wbmciLCJEb21haW5OYW1lIjoiIiwiSGFiaXRhdCI6IiIsIkhvbWV0b3duIjoiIiwiQmlydGhkYXkiOiI5OTk5LTEyLTEyVDAwOjAwOjAwKzA4OjAwIiwiUmVnRGF0ZSI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiU3RhdGVtZW50IjoiIiwiRm9sbG93ZXJzIjowLCJGb2xsb3dpbmdzIjowfSwiVHlwZSI6IlJFRlJFU0hfVE9LRU4iLCJUaW1lIjoiMDAwMS0wMS0wMVQwMDowMDowMFoiLCJleHAiOjE2NDYwMzQyODQsImlzcyI6ImRvdWJhbiJ9.lM8GFcje4NLMi5oV7fQRwnPKci6hmVIeo3o127IfC8o

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object|false|none|
|» access_token|body|string|true|none|
|» refresh_token|body|string|true|none|

> 返回示例

> 成功

```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7IklkIjo0LCJVc2VybmFtZSI6IiIsIlBhc3N3b3JkIjoiMzZiNjI5MjE2MTZkMjQ3MGQwMzI1YjFkNTk0ZjI0YWEiLCJFbWFpbCI6IiIsIlBob25lIjoiMTc3MjY2MzM3NDAiLCJTYWx0IjoiMTY0MjU5MDA5NCIsIkF2YXRhciI6Imh0dHBzOi8vZG91YmFuLWF2YXRhci0xMzA4NzU3Mzg1LmNvcy5hcC1jaG9uZ3FpbmcubXlxY2xvdWQuY29tL2F2dGFyLnBuZ1xyXG4iLCJEb21haW5OYW1lIjoiIiwiSGFiaXRhdCI6IiIsIkhvbWV0b3duIjoiIiwiQmlydGhkYXkiOiI5OTk5LTEyLTEyVDAwOjAwOjAwKzA4OjAwIiwiUmVnRGF0ZSI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiU3RhdGVtZW50IjoiIiwiRm9sbG93ZXJzIjowLCJGb2xsb3dpbmdzIjowfSwiVHlwZSI6IkFDQ0VTU19UT0tFTiIsIlRpbWUiOiIwMDAxLTAxLTAxVDAwOjAwOjAwWiIsImV4cCI6MTY0NTM4NDgwOCwiaXNzIjoiZG91YmFuIn0.d_q5NNxDZ96CDodTl_XsvpsUZRp4CqD8OE4kHsyizz8",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7IklkIjo0LCJVc2VybmFtZSI6IiIsIlBhc3N3b3JkIjoiMzZiNjI5MjE2MTZkMjQ3MGQwMzI1YjFkNTk0ZjI0YWEiLCJFbWFpbCI6IiIsIlBob25lIjoiMTc3MjY2MzM3NDAiLCJTYWx0IjoiMTY0MjU5MDA5NCIsIkF2YXRhciI6Imh0dHBzOi8vZG91YmFuLWF2YXRhci0xMzA4NzU3Mzg1LmNvcy5hcC1jaG9uZ3FpbmcubXlxY2xvdWQuY29tL2F2dGFyLnBuZ1xyXG4iLCJEb21haW5OYW1lIjoiIiwiSGFiaXRhdCI6IiIsIkhvbWV0b3duIjoiIiwiQmlydGhkYXkiOiI5OTk5LTEyLTEyVDAwOjAwOjAwKzA4OjAwIiwiUmVnRGF0ZSI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiU3RhdGVtZW50IjoiIiwiRm9sbG93ZXJzIjowLCJGb2xsb3dpbmdzIjowfSwiVHlwZSI6IlJFRlJFU0hfVE9LRU4iLCJUaW1lIjoiMDAwMS0wMS0wMVQwMDowMDowMFoiLCJleHAiOjE2NDU5MDMyMDgsImlzcyI6ImRvdWJhbiJ9.HzvxsSs_d4EtSNZ8cfgkhej-5re2WWNI4pMzQcPliOc",
  "status": "true",
  "影评": null,
  "想看": null,
  "看过": [
    {
      "Mid": 1291543,
      "Uid": 4,
      "Type": "看过",
      "MovieName": "功夫",
      "MovieAvatar": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2219011938.jpg"
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取自己的影评

GET /api/people/reviews

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# Dicussion

## DELETE 删除讨论

DELETE /api/movie/dicussion/delele_discuss/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 讨论点赞

POST /api/movie/discussion/like/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 发布讨论

POST /api/movie/discussion/put_discuss

> Body 请求参数

```yaml
mid: "1291543"
title: 怎么说呢
value: 很好看

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object|false|none|
|» mid|body|string|true|none|
|» title|body|string|true|none|
|» value|body|string|true|none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取讨论列表

GET /api/movie/discussions/{mid}

> Body 请求参数

```yaml
sort: host

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|mid|path|string|true|none|
|body|body|object|false|none|
|» sort|body|string|true|host  time|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## PUT 跟新讨论

PUT /api/movie/discussion/updata

> Body 请求参数

```yaml
id: string
title: string
value: string

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object|false|none|
|» id|body|string|true|none|
|» title|body|string|true|none|
|» value|body|string|true|none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取单个讨论

GET /api/movie/discussion/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# Reply

## POST 发布回复

POST /api/movie/reply/post

> Body 请求参数

```yaml
pid: "3"
type: review
value: 确实好看

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object|false|none|
|» pid|body|string|true|none|
|» type|body|string|true|none|
|» value|body|string|true|none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取回复

GET /api/movie/reply/{id}

> Body 请求参数

```yaml
start: "0"

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string|true|none|
|type|query|string|true|review discussion reply|
|body|body|object|false|none|
|» start|body|string|true|none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 数据模型

