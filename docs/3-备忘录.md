# Golang 第三轮考核

## 目的

- 掌握http协议和Web工作原理
- 掌握Go语言的Hertz/Gin 的Web框架
- 掌握使用关系型数据库，如：mysql
- 学习**RESTful API**接口规范
- 学习编写文档

## 背景
FanOne准备放寒假了，但是他是一个**摸鱼怪**，请你写一个**备忘录**（写 API 接口即可），让FanOne记录下寒假要完成的事项，能在寒假完成弯道超车！

因为字节和心脏只有一个能跳！我们希望你可以使用字节跳动开源社区 Cloudwego 开源的 HTTP 框架——[Hertz](https://www.cloudwego.io/zh/docs/hertz/)

但是如果你感到吃力，你也可以先从 Gin 框架来写这一轮项目！可以看参考资料哦（os：其实感觉 hertz 写的会更快）

## 任务
> 编写以下API，并编写接口文档 （推荐使用postman）

### 用户模块

- 实现基本的用户注册登录 ( 用token实现 )

**提醒一下：** FanOne创建的待办事务是不能让FanTwo看到的噢~

**注意事项：** 可以使用 jwt token，但是`github.com/dgrijalva/jwt-go` 这个包存在安全问题，请不要使用。可以通过寻找它的升级版本来替换它，例如`github.com/golang-jwt/jwt`（如果你用 hertz 框架，你可以阅读文档，他们的框架自带 jwt-auth）

### 事务模块

增

- 添加一条新的待办事项

改

- 将 一条/所有 代办事项设置为已完成
- 将 一条/所有 已完成事项设置为待办

查

- 查看所有 已完成/未完成/所有 事项。 (需要分页)
- 输入**关键词**查询事项。（需要分页）

删

- 删除 一条/所有已经完成/所有待办/所有 事项



> 一条事务至少需要这些属性：id、标题、内容、完成状态、添加时间、截止时间

### Bonus

1. 自动生成接口文档
2. 使用三层架构设计
3. 考虑数据库交互安全性
4. 思考一个比要求中的结构更优秀的返回结构
5. 对项目使用Redis

## 要求

1. 接口满足**RESTful API**规范
2. 接口文档可以不写**参数描述**
3. 数据返回建议使用JSON格式。如下所示

```
{
  "status": 200,                    // 200 表示正常/成功，500 代表错误。自行了解HTTP状态码。
  "data": {	                    // 业务数据。所有的业务信息都应该放到 data 对象上。
    "items": [
      {
        "id": 1,// 待办事项ID
        "title": "更改好了！",        // 主题
        "content": "好耶！",	    // 内容
        "view": 0,	            // 访问次数
        "status": 1,	            // 状态（正在进行/已完成/其他）
        "created_at": 1638257438,   // 创建时间
        "start_time": 1638257437,   // 开始时间
        "end_time": 0	            // 结束时间
      }
    ],
    "total": 1	                    // 检索出的匹配全部条目数（不是items的len值）
  },
  "msg": "ok",	                    // 返回信息
  "error": ""	                    // 错误信息
}
```

## 提示
- hertz 的官方文档比较晦涩难懂（也就是说需要一定的基础才可以看的比较顺利，不过大家都是这么过来的），你可以结合网上的资料来进行学习，当然，他们有给样例库——[hertz-example](https://github.com/cloudwego/hertz-examples)
- hertz 需要安装命令行工具**hz!** (具体看hertz官方文档)，不安装也可以，但是写起来会和 gin 几乎一致
```
go install github.com/cloudwego/hertz/cmd/hz@latest
```
- hertz的很多使用其实和gin差不多,可以先看看gin的一些简明教程,hertz的使用主要就是看官方文档 (**示例部分**可以多看看)

## 参考

- Hertz中文文档 ：https://www.cloudwego.io/zh/docs/hertz/overview
- Hertz入门 ：https://juejin.cn/post/7124337913352945672
- B站Hertz入门 ：https://www.bilibili.com/video/BV1ta411H7pe?t=1348.7
- 使用 Gin 设计 RESTful API：https://blog.csdn.net/flysnow_org/article/details/103520881
- B站Gin教程：https://www.bilibili.com/video/BV1fA411F7aM?p=1
- Gin知识点总结：https://blog.csdn.net/weixin_45304503/article/details/120381359
- Gorm中文文档：https://learnku.com/docs/gorm/v2
- B站教程：https://www.bilibili.com/video/BV1GT4y1R7tX
