# IM 通信聊天系统

## 拷贝代码
git clone https://github.com/PokemanMaster/IM-System/
cd IM-System

# 本地运行

项目运行后启动在 9000 端口
项目支持移动端
项目的前端地址: https://github.com/PokemanMaster/IM-System-Client

## 项目描述
本项目是基于websocket实时通信的，采用udp进行通信机制，使用 redis 消息队列缓冲，来实现
创建和添加好友、群，与好友聊天，多个好友可进入群来聊天的系统。

## 图片展示

### 登录页面

![home](server/pkg/docs/image/login.png)

### 好友/群列表页面

![login](server/pkg/docs/image/chat.png)
![login](server/pkg/docs/image/chat2.png)

### 聊天页面

![login](server/pkg/docs/image/friend.png)
![login](server/pkg/docs/image/group2.png)

### 个人页面

![product](server/pkg/docs/image/personal.png)


# 后端

- api：用于定义接口函数
- conf：配置参数
- model：应用数据库模型
- pkg / e：封装错误码
- pkg / util：工具函数
- serializer：将数据序列化为 json 的函数
- route 路由逻辑处理
- service：接口函数的实现
- utils：第三方工具


