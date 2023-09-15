# 微服务链路追踪-Jaeger
## 准备两个服务
### auth
鉴权服务。用于token的生成、销毁、验证。本服务功能请求采用http方式。为了提高程序性能与降低网络负载可以使用rpc的方式，这里为了演示，所以采取该方式。
|功能说明|请求接口|参数|请求示例(json)|
|----|----|----|----|
|token生成|/genToken|"accountID":账号ID<br>"password":密码|<code>{"accountID": "leebai", "password": "123456"}</code>|
|token删除|/deleteToken|"accountID": 账号ID|<code>{"accountID": "leebai"}</code>|
|token验证|/verifyToken|"accountID": 账号ID<br>"token": 服务端返回的token|<code>{"accountID": "leebai", "token": "token-leebai"}</code>|

### user
用户服务。暂时提供用户登录、退出两个功能。本服务功能请求采用http方式。
|功能说明|请求接口|参数|请求示例(json)|
|----|----|----|----|
|登录|/login|"accountID": 账号ID<br> "password": 密码|<code>{"accountID": "leebai", "password": "123456"}</code>|
|退出登录|/logout|"accountID": 账号ID|<code>{"accountID": "leebai"}</code>|

## 调用基本逻辑
### 客户端登录逻辑
* 客户端向`user`服务发送`/login`请求进行登录
* `user`服务处理`/login`请求时向`auth`服务发送`/genToken`请求生成`token`
* `user`服务收到`/genToken`请求的响应后再返回给客户端
### 客户端退出登录逻辑
* 客户端向`user`服务发送`/logout`请求来退出登录
* `user`服务处理`/logout`请求时向`auth`服务发送`/deleteToken`请求删除该账号的信息
* `user`服务收到`/deleteToken`请求的响应后再返回给客户端

### 确定链路追踪目标
就是`user`服务向`auth`服务之间的调用逻辑链路需要进行串联。

### 引入追踪框架jaeger
