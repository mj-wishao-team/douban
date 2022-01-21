# API

## User
### `/api/user/login/pw` `POST`
* `application/x-www-form-urlencoded`
* 密码登录

| 请求参数  | 类型 | 说明                |
| --------- | ---- | ------------------- |
| loginAccount | 必选 | 登录账号 |
| password  |   必选| 密码 |

| 返回参数     | 说明         |
| ------------ | ------------ |
| status       | 状态码     |
| data         | 返回消息     |
| token        | 用户token   |

| status | data | 说明   |
| -------- | ---- | ------ |
| `false` | `"请输入注册时用的邮箱或者手机号"` | `loginAccount` 为空 |
| `false` | `"请输入密码"` | `password` 为空 |
| `false` | `"用户名或密码错误"` | `loginAccount` 不存在 |
| `false` | `"用户名或密码错误"` | `loginAccount` 与 `password` 不匹配 |
| `true` | `"密码登录成功"` | `loginAccount` 与 `password` 匹配 |

### `/api/user/login/sms` `POST`

* `application/x-www-form-urlencoded`
* 短信登录

| 请求参数  | 类型 | 说明     |
| ------- | ---- | ------- |
| phone  |  必选  |  手机号  |
| verify_code | 必选 | 验证码  |

| 返回参数     | 说明         |
| ------------ | ------------ |
| status       | 状态码     |
| data         | 返回消息     |
| token        | 用户token   |

| status | data | 说明   |
| -------- | ---- | ------ |
| `false` | `"手机号不能为空"` | `phone` 为空 |
| `false` | `"短信验证码为空"` | `verify_code` 为空 |
| `false` | `"验证码错误"` | `phone` 与 `verify_code` 不匹配 |
| `false` | `"未发送验证码"` | `verify_code` 无对应验证码 |
| `true` | `false` | 新用户需要跳转到注册界面 |
| `true` | `ture` |  登陆成功|


### `/api/user/register` `POST`
* `application/x-www-form-urlencoded`
* 注册

| 请求参数  | 类型 | ~~说~~明                |
| --------- | ---- | ------------------- |
| password | 必选 |账号密码  |
| username | 必选 | 账号昵称    |


| 返回参数     | 说明         |
| ------------ | ------------ |
| status       | 状态码       |
| data         | 返回消息     |
| token        | 用户token    |

| status | data | 说明   |
| -------- | ---- | ------ |
| `false` | `"密码不能为空"` | `password` 为空 |
| `false` | `"密码不能小于6个字符"` | `password` 长度少于 7个字节 |
| `false` | `"密码不能大于16个字符"` | `password` 长度超过 16 个字节 |
| `false` | `"用户名不能为空"` | `username` 为空 |
| `false` | `"用户名太长了"` | `username` 长度超过 15 个字节 |
| `true` | `"注册成功！"` | 参数合法 |

### `/api/user/setting` `POST`
* `application/x-www-form-urlencoded`
* 个人设置

| 请求参数  | 类型 | 说明                |
| --------- | ---- | ------------------- |
| avatar|可选| 头像 |
| username| 可选 | 昵称|
| domainName|可选| 域名 |
| habitat|可选| 常居地 |
| hometown|可选|家乡 |
| birthday|可选|生日 |

| 返回参数     | 说明         |
| ------------ | ------------ |
| status       | 状态码       |
| data         | 返回消息     |
| token        | 用户token    |

| status | data | 说明   |
| -------- | ---- | ------ |
| `false` | `"昵称30天只能修改一次哦"` | 30天内重复修改 |
| `false` | `"域名不能超过15个字符哦"` | 域名太长 |
| `true` | `ture` |  跟新成功|

### `/api/user/setting/bind_email` `POST`
* `application/x-www-form-urlencoded`
*  绑定邮箱

| 请求参数  | 类型 | 说明                |
| --------- | ---- | ------------------- |
| email|必选| 邮箱 |
| verify_code|必选| 验证码 |

| 返回参数     | 说明         |
| ------------ | ------------ |
| status       | 状态码       |
| data         | 返回消息     |

| status | data | 说明   |
| -------- | ---- | ------ |
| `false` | `"请填写邮箱"` | 邮箱为空 |
| `false` | `"请正确填写邮箱"` | 邮箱格式错误 |
| `false` | `"请填写验证码"` | 验证码为空 |
| `false` | `"验证码错误"` | 验证码错误 |
| `false` | `"验证码过期"` | 验证码过期 |
| `true` | `ture` |  跟新成功|


# 豆瓣电影

> v1.0.0

# User/login

## POST 短信登录接口

POST /api/user/login/sms

> Body 请求参数

```yaml
phone: string
verify_code: string

```

### 请求参数

| 名称          | 位置 | 类型   | 必选  | 说明                          |
| ------------- | ---- | ------ | ----- | ----------------------------- |
| body          | body | object | false | none                          |
| » phone       | body | string | true  | 电话号码 如果成功会返回电话号 |
| » verify_code | body | string | true  | 验证码                        |

> 返回示例

> 成功

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
  "acess_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7IklkIjo0LCJVc2VybmFtZSI6IiIsIlBhc3N3b3JkIjoiMzZiNjI5MjE2MTZkMjQ3MGQwMzI1YjFkNTk0ZjI0YWEiLCJFbWFpbCI6IiIsIlBob25lIjoiMTU2NTY2NzEwNzAiLCJTYWx0IjoiMTY0MjU5MDA5NCIsIkF2YXRhciI6IiIsIkRvbWFpbk5hbWUiOiIiLCJIYWJpdGF0IjoiIiwiSG9tZXRvd24iOiIiLCJCaXJ0aGRheSI6Ijk5OTktMTItMTJUMDA6MDA6MDArMDg6MDAiLCJSZWdEYXRlIjoiMDAwMS0wMS0wMVQwMDowMDowMFoiLCJTdGF0ZW1lbnQiOiIiLCJGb2xsb3dlcnMiOjAsIkZvbGxvd2luZ3MiOjB9LCJUeXBlIjoiQUNDRVNTX1RPS0VOIiwiVGltZSI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiZXhwIjoxNjQyNTkwOTAwLCJpc3MiOiJkb3ViYW4ifQ.wyzDPzcQDoFXsuQtW-yfNVeS9cU1U-CZVQY5U9PNmAY",
  "data": 4,
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7IklkIjo0LCJVc2VybmFtZSI6IiIsIlBhc3N3b3JkIjoiMzZiNjI5MjE2MTZkMjQ3MGQwMzI1YjFkNTk0ZjI0YWEiLCJFbWFpbCI6IiIsIlBob25lIjoiMTU2NTY2NzEwNzAiLCJTYWx0IjoiMTY0MjU5MDA5NCIsIkF2YXRhciI6IiIsIkRvbWFpbk5hbWUiOiIiLCJIYWJpdGF0IjoiIiwiSG9tZXRvd24iOiIiLCJCaXJ0aGRheSI6Ijk5OTktMTItMTJUMDA6MDA6MDArMDg6MDAiLCJSZWdEYXRlIjoiMDAwMS0wMS0wMVQwMDowMDowMFoiLCJTdGF0ZW1lbnQiOiIiLCJGb2xsb3dlcnMiOjAsIkZvbGxvd2luZ3MiOjB9LCJUeXBlIjoiUkVGUkVTSF9UT0tFTiIsIlRpbWUiOiIwMDAxLTAxLTAxVDAwOjAwOjAwWiIsImV4cCI6MTY0MzE5NTQwMCwiaXNzIjoiZG91YmFuIn0.KGGIybfolq_8Z0xAPUFeMVtmOf8ogLo8y2jthFKDIDM",
  "status": true
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

| 状态码 | 状态码含义                                              | 说明 | 数据模型  |
| ------ | ------------------------------------------------------- | ---- | --------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功 | undefined |

## POST 密码登录接口

POST /api/user/login/pw

> Body 请求参数

```yaml
loginAccount: "17726633740"
password: "123456789"

```

### 请求参数

| 名称           | 位置 | 类型   | 必选  | 说明     |
| -------------- | ---- | ------ | ----- | -------- |
| body           | body | object | false | none     |
| » loginAccount | body | string | true  | 用户账户 |
| » password     | body | string | true  | 密码     |

> 返回示例

> 成功

```json
{
  "acess_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7IklkIjo0LCJVc2VybmFtZSI6IiIsIlBhc3N3b3JkIjoiMzZiNjI5MjE2MTZkMjQ3MGQwMzI1YjFkNTk0ZjI0YWEiLCJFbWFpbCI6IiIsIlBob25lIjoiMTU2NTY2NzEwNzAiLCJTYWx0IjoiMTY0MjU5MDA5NCIsIkF2YXRhciI6IiIsIkRvbWFpbk5hbWUiOiIiLCJIYWJpdGF0IjoiIiwiSG9tZXRvd24iOiIiLCJCaXJ0aGRheSI6Ijk5OTktMTItMTJUMDA6MDA6MDArMDg6MDAiLCJSZWdEYXRlIjoiMDAwMS0wMS0wMVQwMDowMDowMFoiLCJTdGF0ZW1lbnQiOiIiLCJGb2xsb3dlcnMiOjAsIkZvbGxvd2luZ3MiOjB9LCJUeXBlIjoiQUNDRVNTX1RPS0VOIiwiVGltZSI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiZXhwIjoxNjQyNTkwOTAwLCJpc3MiOiJkb3ViYW4ifQ.wyzDPzcQDoFXsuQtW-yfNVeS9cU1U-CZVQY5U9PNmAY",
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

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功 | Inline   |

### 返回数据结构

状态码 **200**

| 名称            | 类型    | 必选 | 约束 | 说明 |
| --------------- | ------- | ---- | ---- | ---- |
| » acess_token   | string  | true | none | none |
| » data          | integer | true | none | none |
| » refresh_token | string  | true | none | none |
| » status        | boolean | true | none | none |

# User/注册

## POST 短信注册

POST /api/user/register

> Body 请求参数

```yaml
phone: "1725014728"
username: MJ
password: "123456"

```

### 请求参数

| 名称       | 位置 | 类型   | 必选  | 说明                              |
| ---------- | ---- | ------ | ----- | --------------------------------- |
| body       | body | object | false | none                              |
| » phone    | body | string | true  | 电话号码由短信界面返回的phone填入 |
| » username | body | string | true  | 用户名                            |
| » password | body | string | true  | 密码                              |

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

| 状态码 | 状态码含义                                                   | 说明       | 数据模型 |
| ------ | ------------------------------------------------------------ | ---------- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)      | 成功       | Inline   |
| 500    | [Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1) | 服务器错误 | Inline   |

### 返回数据结构

状态码 **200**

| 名称   | 类型   | 必选 | 约束 | 说明 |
| ------ | ------ | ---- | ---- | ---- |
| » data | string | true | none | none |

状态码 **500**

| 名称   | 类型        | 必选 | 约束 | 说明 |
| ------ | ----------- | ---- | ---- | ---- |
| » info | string¦null | true | none | none |

# User/发送短信

## POST 发送短信

POST /api/verify/sms

> Body 请求参数

```yaml
phone: "17726633740"

```

### 请求参数

| 名称    | 位置 | 类型   | 必选  | 说明 |
| ------- | ---- | ------ | ----- | ---- |
| body    | body | object | false | none |
| » phone | body | string | true  | none |

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

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功 | Inline   |

### 返回数据结构

# User/邮箱

## POST 发送邮箱验证码

POST /api/verify/emial

> Body 请求参数

```yaml
email: 1725014728@qq.com

```

### 请求参数

| 名称    | 位置 | 类型   | 必选  | 说明                                   |
| ------- | ---- | ------ | ----- | -------------------------------------- |
| body    | body | object | false | none                                   |
| » email | body | string | true  | 邮箱账户 前端要校验邮箱格式 正常再请求 |

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

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功 | Inline   |

### 返回数据结构

状态码 **200**

| 名称     | 类型   | 必选 | 约束 | 说明 |
| -------- | ------ | ---- | ---- | ---- |
| » data   | string | true | none | none |
| » status | string | true | none | none |

## POST 绑定邮箱

POST /api/user/bind_email

> Body 请求参数

```yaml
email: 1725014728@qq.com
verify_code: "049514"

```

### 请求参数

| 名称          | 位置 | 类型   | 必选  | 说明 |
| ------------- | ---- | ------ | ----- | ---- |
| body          | body | object | false | none |
| » email       | body | string | true  | none |
| » verify_code | body | string | true  | none |

> 返回示例

> 成功

```json
{
  "data": "绑定成功",
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

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功 | Inline   |

### 返回数据结构

状态码 **200**

| 名称     | 类型   | 必选 | 约束 | 说明 |
| -------- | ------ | ---- | ---- | ---- |
| » data   | string | true | none | none |
| » status | string | true | none | none |

