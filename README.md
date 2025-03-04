# ByteScience-WAM-Admin

## 项目背景
`ByteScience-WAM-Admin` 是一个后台管理系统，旨在为业务系统提供管理功能。该项目与其他服务协作，支持多种功能，包括用户管理、权限控制等。该项目使用 Go 编写，并依赖 MySQL 和 Redis 数据库。

## 安装依赖
本项目依赖以下服务：
- **MySQL**: 用于存储系统的持久化数据。
- **Redis**: 用于缓存数据和管理会话等。

### 环境要求
- Go 1.18 及以上
- MySQL 8.0 或更高版本
- Redis 7.2.4 或更高版本

### 安装 MySQL 和 Redis
确保你已经安装并配置了 MySQL 和 Redis。如果没有安装，可以参考以下链接进行安装：
- [MySQL 安装教程](https://dev.mysql.com/doc/refman/8.0/en/installing.html)
- [Redis 安装教程](https://redis.io/docs/getting-started/)

### 使用 Docker 部署
你也可以通过 Docker 来快速部署 **MySQL** 和 **Redis**：

1. **启动 Docker 服务**:
- 执行以下命令来启动容器：
```
docker-compose up -d
```

2. 设置 MySQL 远程连接权限
- 进入 MySQL 容器
```
docker exec -it <容器 ID> bash
```
- 连接到mysql
```
mysql -u root -p
```
- 设置远程连接权限
```
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY '123456' WITH GRANT OPTION; FLUSH PRIVILEGES;
```

## 服务启动
* 安装依赖
```azure
    go get -u
```
* 启动服务
```azure
    go run main.go
```
