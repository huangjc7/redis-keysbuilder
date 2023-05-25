## redis-keysbuilder版本日志
* version 1.0.2
  * redis集群版本key生成工具 快速并发生成随机key写入redis集群用来测试
  * 新增参数来认证redis-cluster密码
  * 使用方法参数变化 更redis-cli命令使用尽可能保持一致
  * 重新使用新的并发逻辑 优化了执行效率
  * 新增build.sh异步构建脚本
  * 适用版本
    * redis 7xx
  * 使用帮助
    * `./keysbuilder_linux_amd64 -help`
  * 使用方法
    * `./keysbuilder_linux_amd64 -h 127.0.0.1:6379 -k 50000 -a 123456 //写入50000随机keys`
  * 构建方法(amd64架构的linux版本)
    * `./build.sh amd64 linux`
---
* version 1.0.1
  * redis集群版本key生成工具 快速并发生成随机key写入redis集群用来测试
  * 新增参数来认证redis-cluster密码
  * 适用版本
    * redis 7xx
  * 使用方法
    * `./keysbuilder_linux_amd64 -h 127.0.0.1:6379 -n 50000 -p 123456 //写入50000随机keys`
---
* version v1.0.0
  * redis集群版本key生成工具 快速并发生成随机key写入redis集群用来测试
  * 适用版本
    * redis 7xx
  * 使用方法
      * `./keysbuilder_linux_amd64 -h 127.0.0.1:6379 -n 50000 //写入50000随机keys`
