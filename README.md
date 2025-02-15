# LinkMe - 开源论坛项目

![LinkMe](https://socialify.git.ci/wangzijian2002/LinkMe/image?description=1&font=Source%20Code%20Pro&forks=1&issues=1&language=1&logo=https%3A%2F%2Fgithub.com%2Fwangzijian2002%2FLinkMe%2Fassets%2F71474660%2F22ef2063-ab82-481f-898f-29d95fa70236&name=1&pattern=Solid&pulls=1&stargazers=1&theme=Dark)

## 项目简介
LinkMe 是一个使用 Go 语言开发的论坛项目。它旨在为用户提供一个简洁、高效、可扩展的在线交流平台。本项目使用DDD领域设计理念，采用模块化的设计，使得添加新功能或进行定制化修改变得非常容易。LinkMe 支持多种数据库后端，并且可以通过 Kubernetes 进行部署。

## 在线地址
http://www.simplecloud.top/（前端重构中...暂时停用）
超级管理员：admin/admin

## 项目部署文档
[启动文档](./doc/LinkMe项目启动文档.md)

## 微服务项目地址
https://github.com/GoSimplicity/LinkMe-microservices

## 前端项目地址
https://github.com/GoSimplicity/LinkMe-web

## 功能特性
- 用户注册、登录、注销
- 用户角色RBAC
- 用户关系
- 用户评论
- 发布帖子、评论、点赞
- 历史记录
- 帖子审核
- 用户个人资料编辑
- 论坛版块管理
- 自动获取热门榜单
- 用户、帖子搜索
- Kubernetes 一键部署
- 前后端分离架构

## 技术栈
- Go 语言
- Gin Web 框架
- Wire 依赖注入
- Kubernetes 集群管理
- MySQL 数据库
- Redis 缓存数据库
- MongoDB 文档数据库
- Kafka 消息队列
- Prometheus 监控
- ELK 日志收集
- Canal 数据同步
- ElasticSearch 搜索引擎
- Docker 容器化
- 随项目进度技术栈实时更新..

## 目录结构
```
.
├── config           # 项目配置文件目录
├── deploy           # docker及k8s部署文件目录
├── doc              # 项目文档目录
├── internal         # 项目内部包，含核心业务逻辑
├── ioc              # IoC容器配置，负责依赖注入设置
├── pkg              # 自定义工具包与库
├── job              # 定时任务目录
├── logs             # 项目日志目录
├── utils            # 项目工具包目录
├── main.go          # 项目入口文件
├── middleware       # 中间件目录
├── tmp              # 临时文件目录
├── wire_gen.go      # Wire工具生成的代码
├── wire.go          # Wire配置，声明依赖注入关系
```

## 如何贡献
我们欢迎任何形式的贡献，包括但不限于：
- 提交代码（Pull Requests）
- 报告问题（Issues）
- 文档改进
- 功能建议
  请确保在贡献代码之前阅读了我们的[贡献指南](#贡献指南)。

## 贡献指南
- Fork 本仓库
- 创建您的特性分支 (`git checkout -b my-new-feature`)
- 提交您的改动 (`git commit -am 'Add some feature'`)
- 将您的分支推送到 GitHub (`git push origin my-new-feature`)
- 创建一个 Pull Request
## 开始使用

### 克隆项目
```bash
git@github.com:GoSimplicity/LinkMe.git
```

### 环境要求
- Docker 20.10.0+
- Docker Compose 2.0.0+

### 部署步骤(一键部署)
```bash
make deploy
```

#### 一键重新部署
```bash
make redeploy
```

### 部署步骤(手动部署)

#### 创建数据目录并提权
```bash
mkdir -p ./data/kafka/data && chmod -R 777 ./data/kafka
```

#### 启动项目依赖
```bash
docker-compose -f docker-compose-env.yaml up -d
```

#### 构建镜像
```bash
docker build -t linkme/gomodd:v1.22.3 .
```

#### 启动项目
```bash
docker-compose up -d
```

### 项目超级管理员账号
```bash
admin/admin
```


## 许可证
本项目使用 MIT 许可证，详情请见 [LICENSE](./LICENSE) 文件。

## 联系方式
- Email: [wzijian62@gmail.com](mailto:wzijian62@gmail.com)
- 为了方便交流，可以加我vx：GoSimplicity 我拉你进微信群

## 致谢
感谢所有为本项目做出贡献的人！

---
欢迎来到 LinkMe，让我们一起构建更好的论坛社区！
