## GoWeb实战论坛项目
### 技术栈
1. 使用 `MySQL`, `Redis`用作数据库, 管理数据, 同时使用第三方库 `sqlx`, `go-redis`来操作数据库
2. 使用 `Zap` 管理日志, 同时搭建自己的日志中间件
3. 使用 `Viper` 管理配置文件
4. 使用 `gin` 框架进行Web开发, 通过`RESTful`风格的API进行数据通信
5. 使用雪花算法生成不重复ID
6. 使用`validator`进行参数校验, 同时自定义参数校验方法
7. 利用`md5` 包中的加密算法对用户密码进行加密
8. 采用并优化`JWT`中间件实现用户认证, 用户携带`Token`发送请求
9. 编写了`Makefile`实现项目的部署
10. 使用`Air` 实现文件的热加载
11. 利用缓存key减少Zinterstore的执行次数
12. 使用`swagger`生成接口文档, 使用第三方`gin-swagger`库自动生成`RESTful`风格的API文档
    (*代办)
13. 使用`go test`对项目进行测试
14. 采用`go-wrk`对其进行压力测试
15. 使用令牌桶策略, 设计中间件实现限流
16. 使用gin框架的`pprof`工具进行性能调优
17. 利用`Docker` 来部署该项目, 使用分阶段构建的技术, 减少镜像体积

### 设计
![img_2.png](img%2Fimg_2.png)
#### 1. Redis
![img.png](img%2Fimg.png)
![img_1.png](img%2Fimg_1.png)


### 心得体会
1. 可以发现一般dao层不会处理错误, 而是返回到logic层处理, 而logic层也会将错误返回到controller层处理
所以一般处理错误, 会在controller层, 而dao层和logic层顶多对错误进行包装
2. 测试的时候可以运用另外一个数据库, 避免数据污染本来的数据库


### 代做
1. 将其部署到自己的服务器上, 使用`nohup` or `supervisor`来进行部署
2. 将本地的 `mysql` 以及 `redis` 服务改为 `docker` 部署, 增加适用性
