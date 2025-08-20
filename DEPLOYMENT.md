# 部署指南

本文档提供了使用 Docker 和 Docker Compose 构建并运行整个 V2Ray SaaS 应用的技术栈的说明。

## 系统要求

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

请确保您的系统上已安装 Docker 和 Docker Compose，并且 Docker 守护进程正在运行。

## 运行应用程序

整个应用程序技术栈，包括所有的 Go 微服务、MySQL 数据库和 Redis 缓存，都已在 `docker-compose.yml` 文件中定义。

要构建并以分离模式（在后台）运行所有服务，请导航到项目的根目录并执行以下命令：

```bash
docker-compose up --build -d
```

### 命令解析

- `docker-compose up`: 这是启动 `docker-compose.yml` 文件中定义的服务集的标准命令。
- `--build`: 此标志会强制 Docker Compose 在启动容器之前，从各自的 `Dockerfile` 构建 Go 服务的镜像。当您首次运行或代码发生变更后，都应使用此标志。
- `-d`: 此标志以分离模式运行容器，意味着它们将在后台运行，您不会在终端中直接看到它们的日志。

## 验证部署状态

运行命令后，您可以通过以下方式检查正在运行的容器的状态：

```bash
docker-compose ps
```

您应该能看到所有服务（`db`, `redis`, `portal-api`, `admin-api`, `node-service`, `node-agent`）的状态（`State`）都为 `Up`。

您也可以查看特定服务的日志：

```bash
# 示例：查看 portal-api 的日志
docker-compose logs -f portal-api
```

## 访问服务

容器启动并运行后，您可以通过本地主机的以下默认端口访问各项服务：

- **门户 API**: `http://localhost:8080`
- **后台管理 API**: `http://localhost:8081`
- **节点服务**: `http://localhost:8082`
- **节点代理**: `http://localhost:8083`
- **MySQL 数据库**: `localhost:3306`
- **Redis**: `localhost:6379`

## 停止应用程序

要停止所有正在运行的服务，请使用以下命令：

```bash
docker-compose down
```

如果您还希望移除数据库的持久化数据卷（**这将删除您的所有数据**），可以添加 `-v` 标志：

```bash
docker-compose down -v
```