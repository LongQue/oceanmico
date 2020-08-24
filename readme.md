## Proto 生成 GO 和 Go-micro 文件

protoc  --micro_out=. --go_out=. proto\rand\rand.proto
8.23
B站up主 https://www.bilibili.com/video/av201306940 一周双更照着学微服务
搭建consul 配置proto 调通微服务调用
出事了 github Desktop 创建本地仓库不忽略.idea

8.24
增加token生成验证，用户登录返回token，rand接口增加验证中间件
rand模块增加sentinel限流