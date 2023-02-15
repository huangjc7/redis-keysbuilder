## redis-keysbuilder-cluster版本日志
* version 1.0.1
  * 新增参数来认证redis-cluster密码
  * 适用版本
    * redis 7xx
  * 使用方法
    * `./keysbuilder_linux_amd64 -h 127.0.0.1:6379 -h 50000 -p 123456 //写入50000随机keys`
---
* version v1.0.0
  * redis集群版本key生成工具 快速并发生成随机key写入redis集群用来测试
  * 适用版本
    * redis 7xx
  * 使用方法
      * `./keysbuilder_linux_amd64 -h 127.0.0.1:6379 -h 50000 //写入50000随机keys`
