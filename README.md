# networker-be

DN42 config generate Sever (Backend) 

一个用于自动收集 DN42 Peer 友友们配置的 Web 服务

带有自动导出朋友们配置数据的管理功能 

> 提供了一个 /admin/export.go 文件 
>
> 编译后运行就可以导出所有的配置数据
>
> 使用类似 [dn42ConfigGenerator](https://github.com/Esonhugh/dn42ConfigGenerator) 的工具
>
> 进行解析生成 就可以一键拿到所有的配置

**前端编写完毕 在位置 [Networker-fe](https://github.com/Esonhugh/Networker-fe)**

## api 文档 docs/README.md

## 使用方法

在 mysql 数据库中 创建对应用户和数据库

创建一个无用的邮箱账号 用于发送验证码

根据说明修改对应的配置条目 application.yaml

```shell
sh buildup.sh
```