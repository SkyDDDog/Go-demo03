## 全局公共参数
#### 全局Header参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 全局Query参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 全局Body参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 全局认证方式
```text
noauth
```
#### 全局预执行脚本
```javascript
暂无预执行脚本
```
#### 全局后执行脚本
```javascript
暂无后执行脚本
```
## /Golang二轮
```text
暂无描述
```
#### Header参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### Query参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### Body参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /Golang二轮/用户接口
```text
暂无描述
```
#### Header参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### Query参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### Body参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /Golang二轮/用户接口/注册
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:1234/api/user/register

#### 请求方式
> POST

#### Content-Type
> json

#### 请求Body参数
```javascript
{
    "username": "test",
    "password": "123456"
}
```
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /Golang二轮/用户接口/登录
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:1234/api/user/login

#### 请求方式
> POST

#### Content-Type
> json

#### 请求Body参数
```javascript
{
    "username": "test",
    "password": "123456"
}
```
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 10000,
	"msg": "success",
	"data": {
		"token": "661fe75115e45a3520ec74121898e2af"
	}
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 10000 | Number | 
msg | success | String | 返回文字描述
data | - | Object | 返回数据
data.token | 661fe75115e45a3520ec74121898e2af | Number | 
## /Golang二轮/事务接口
```text
暂无描述
```
#### Header参数
参数名 | 示例值 | 参数描述
--- | --- | ---
token | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QiLCJwYXNzd29yZCI6IjEyMzQ1NiIsImV4cCI6MTcwMzUxODc2NiwiaXNzIjoiTXJIdWFuZyJ9.HHSyw9oDCkdWdoluiD69hU8IDukiBliahkzblOhEDkE | 用户访问token
#### Query参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### Body参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /Golang二轮/事务接口/添加一条代办事项
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:1234/api/note

#### 请求方式
> POST

#### Content-Type
> json

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QiLCJwYXNzd29yZCI6IjEyMzQ1NiIsImV4cCI6MTcwMzUxODc2NiwiaXNzIjoiTXJIdWFuZyJ9.HHSyw9oDCkdWdoluiD69hU8IDukiBliahkzblOhEDkE | String | 是 | -
#### 请求Body参数
```javascript
{
    "title": "标题",
    "content": "内容",
    "ddl": "@now('yyyy-MM-dd HH:mm:ss')"
}
```
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /Golang二轮/事务接口/将 一条 代办事项设置为已完成
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:1234/api/note/done/1607252161287163904

#### 请求方式
> PUT

#### Content-Type
> none

#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /Golang二轮/事务接口/将 一条 完成事项设置为待办
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:1234/api/note/undo/1607252161287163904

#### 请求方式
> PUT

#### Content-Type
> none

#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /Golang二轮/事务接口/将 多条 代办事项设置为已完成
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:1234/api/note/done

#### 请求方式
> PUT

#### Content-Type
> none

#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /Golang二轮/事务接口/将 多条 完成事项设置为待办
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:1234/api/note/undo

#### 请求方式
> PUT

#### Content-Type
> none

#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /Golang二轮/事务接口/删除 一条 事项
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:1234/api/note/1607252161287163904

#### 请求方式
> DELETE

#### Content-Type
> none

#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /Golang二轮/事务接口/删除 所有已经完成 事项
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:1234/api/note/done

#### 请求方式
> DELETE

#### Content-Type
> none

#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /Golang二轮/事务接口/删除 所有待办 事项
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:1234/api/note/undo

#### 请求方式
> DELETE

#### Content-Type
> none

#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /Golang二轮/事务接口/删除 所有 事项
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:1234/api/note/all

#### 请求方式
> DELETE

#### Content-Type
> none

#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /Golang二轮/事务接口/查看所有事项
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:1234/api/note?pageNum=1&pageSize=10

#### 请求方式
> GET

#### Content-Type
> none

#### 请求Query参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
pageNum | 1 | String | 否 | 第几页
pageSize | 10 | String | 否 | 一页几条数据
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /Golang二轮/事务接口/查看所有已完成事项
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:1234/api/note/done?pageNum=1&pageSize=10

#### 请求方式
> GET

#### Content-Type
> none

#### 请求Query参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
pageNum | 1 | String | 否 | 第几页
pageSize | 10 | String | 否 | 一页几条数据
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /Golang二轮/事务接口/查看所有待办事项
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:1234/api/note/undo?pageNum=1&pageSize=10

#### 请求方式
> GET

#### Content-Type
> none

#### 请求Query参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
pageNum | 1 | String | 否 | 第几页
pageSize | 10 | String | 否 | 一页几条数据
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /Golang二轮/事务接口/输入关键词查询事项
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:1234/api/note/2?pageNum=1&pageSize=10

#### 请求方式
> GET

#### Content-Type
> none

#### 请求Query参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
pageNum | 1 | String | 否 | 第几页
pageSize | 10 | String | 否 | 一页几条数据
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```