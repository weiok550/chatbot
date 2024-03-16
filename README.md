# chatbot
聊天机器人，一个面试作业

## 功能概述
- 前端【实现】
  - 功能页面
    - 登陆页
    - 聊天页
  - 使用框架
    - Vue
    - Bootstrap
- 服务端
  - 功能及服务
    - 聊天【实现】
      - 聊天策略接口
        - 聊天策略服务实现
          - 选择机器人
      - 机器人接口
        - mock机器人服务
        - gpt机器人服务
          - 聊天响应
      - 聊天记录服务接口
        - 聊天记录服务实现
          - 聊天记录异步入库(规划)
          - 聊天记录同步入库
    - 用户【实现】
      - 用户服务接口
        - 登陆检测用户名密码是否正确
        - 根据用户名获取账户信息
        - 检测登录态
        - 记录登陆态及用户信息
      - 用户登陆
    - 活动接口【实现】
      - mock活动
        - 检测活动中是否有需要向用户推送的聊天消息
    - 消息通道【规划】
      - 用户连接服务
        - 获取用户的连接服务器及连接句柄
        - 记录用户的连接服务器及连接句柄 
      - 消息通道-websocket
        - 接受用户消息
        - 给用户主动推送消息
        
  - 使用框架
    - gin
  - 使用库
    - zap
    - lumberjack
    - viper
    - gorm
    - go-redis


## 存储
- mysql 落地聊天记录及用户信息
- 聊天记录可以预见数据量会很大根据用户id取模hash分表,提前评估，单表数据量控制在2000万+定时老数据归档
- 用户登陆态信息redis存储，单机本地可以缓存少量(控制内存消耗),优先从本地缓存校验

## 安全
- 对于用户侧数据持怀疑态度不能直接拿来渲染避免xss攻击，服务端前端均需要根据业务场景进行特殊字符的过滤/拦截/替换
- 密码等敏感数据落地使用密文存储
- 手写sql注意类型校验及绑定避免sql注入
## 部署
- 存储及服务均可分布式部署，可以水平扩展
- 随着业务规模扩大可以针对每个服务单独拆分部署 + rpc调用改造

## 业务流程
- todo

## 线上体验
- https://chatbot.putianxia.fun
- 用户名:neo1 密码:123456

## 附件[作业要求]

```
Messaging Platform: Review Bot
Instructions
Hello! We are excited to share this take-home assignment with you. Please read the problem
statement below and raise questions if any. We are happy to clarify over email or a quick call.
Please try to get back to us in a week. After your solution is ready, please share it with us and
schedule a 1-hour meeting to discuss your solution.
Overview
It is challenging for SMBs to leverage the tooling available for all the large companies, due to a
fundamental constraint in resources. As a few-person SMB you want to be focusing on
operations and providing excellent service to customers. We want to help them by providing
these services through the most ubiquitous tech platform available: Messaging.
As part of this exercise assume that your system is integrated with a 3rd party messaging
platform and build a basic chatbot that allows an SMB to gather personal reviews from
customers. The chatbot should be triggered based on some conditions/context. (E.g. recently
completed transaction outside the thread, Thank you notes inside the thread), select the right
template to send to the customer and manage/store the response.
You should structure your code so that it's testable and also take care to properly handle errors
returned from downstream dependencies (or other errors like timeouts) and fallback graciously.
We are not looking for exact coverage numbers, but your testing philosophy should be
represented in your coding decisions. Your code should also be designed to be deployable in a
PaaS provider (heroku or AWS Lambda or other services), but also able to run in a development
machine. You can use a container to create a deployable image as well.
You will be required to show a working demo of the chatbot during the live interview. You can
showcase the bot functionality either through Postman API calls, a simple web UI or use a real
API.
Expect to spend 4 ~ 5 hours on this exercise. Plan accordingly and use a simple architecture to
maximize your output. Feel free to mock any downstream dependencies that you would
consider utilizing.
Details
Take a look into the Telegram documentation to draw inspiration about how a 3rd party
messaging platform api is structured. You may hook up to the real api to make your bot more
engaging, but may also choose to simulate the 3rd party provider.
Some Principles
Use What You Know: We’re looking to see you performing your best work, so feel free to use
the frameworks that you are most comfortable with. Our backend tech stack is based on Go and
Python so we prefer to see your work in either of those 2 programming languages.
Showcase Your Strengths: As you’re working on your bot, focus your time and energy on your
areas of expertise. If you’re a great architect, focus on that. If you’re good at building UX, focus
on that. If you’re strong in ML, showcase that. We want to see what you’re best at.
Pair Programming Interview
When we meet for the pair programming interview you will be asked to demo the basic
requirements outlined above and we will work to add some more functionality to your
application, so please don’t build more features than we outline here! During the interview we
will also ask you about architecture choices you made while building the bot. Be prepared to
answer these questions as they come up.
```