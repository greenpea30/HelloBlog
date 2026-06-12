# HelloBlog

一个小小博客网站！

## 功能

- 写文章、编辑、删除（Markdown 支持）
- 点赞（红/灰切换）、评论
- 模糊搜索文章
- 个人主页（修改昵称、头像、简介）
- 评论消息通知
- 友情链接管理
- 浙大学号登录（ZJU PASS 验证）
- 只看校友筛选
- 浏览量统计

## 技术栈

| 层 | 技术 |
|------|------|
| 前端 | React 18 + Vite + React Router |
| 后端 | Go 1.25 + Gin + GORM |
| 数据库 | PostgreSQL 16 + pgvector |
| 缓存 | Redis 7 |
| 认证 | JWT + bcrypt |
| 容器 | Docker Compose |

## 本地运行

### 前置要求

- Go 1.21+
- Node.js 18+
- Docker & Docker Compose

### 启动数据库

```bash
docker-compose up -d
```

启动 PostgreSQL（端口 5432）和 Redis（端口 6379）。

### 初始化数据库表

```bash
docker exec -i helloblog-postgres psql -U helloblog -d helloblog < backend/migrations/000001_init_schema.sql
```

### 添加 ZJU 学号字段（可选）

如果用浙大学号登录，需要执行：

```bash
docker exec -i helloblog-postgres psql -U helloblog -d helloblog -c \
  "ALTER TABLE users ADD COLUMN IF NOT EXISTS zju_id varchar(20) UNIQUE; ALTER TABLE users ALTER COLUMN email DROP NOT NULL;"
```

### 配置后端

```bash
cp backend/etc/config.example.yaml backend/etc/config.yaml
```

默认配置即可本地运行。可按需修改 `jwt.secret`、数据库连接等。

### 启动后端

```bash
cd backend
go run cmd/api/main.go
```

看到 `Listening and serving HTTP on :8080` 表示启动成功。

### 安装前端依赖并启动

新开终端：

```bash
cd frontend
npm install
npm run dev
```

看到 `VITE ready in ... ms` 表示启动成功。

### 访问

打开浏览器访问 **http://localhost:3000**

## 测试账号

| 方式 | 账号 |
|------|------|
| 邮箱登录 | `test@test.com` / `test123456` |
| ZJU 学号登录 | 你的浙大学号 + ZJU PASS |

## 常用命令

| 命令 | 作用 |
|------|------|
| `docker-compose down` | 停止数据库 |
| `docker-compose down -v` | 停止数据库并清空所有数据 |
| `lsof -ti:8080 \| xargs kill -9` | 强制杀掉后端进程 |
| `pkill -f vite` | 强制杀掉前端进程 |
| `cd backend && go build ./...` | 编译检查后端 |

## 项目结构

```
blog2.0/
├── docker-compose.yml          # 数据库容器编排
├── backend/
│   ├── cmd/api/main.go         # 后端入口
│   ├── internal/
│   │   ├── config/             # 配置读取
│   │   ├── controller/         # HTTP 处理器
│   │   ├── dao/                # 数据访问层
│   │   ├── dto/                # 数据传输对象
│   │   ├── service/            # 业务逻辑层
│   │   ├── server/             # HTTP 服务 + 路由
│   │   ├── svc/                # 服务上下文
│   │   ├── infra/              # 基础设施（DB, Redis）
│   │   ├── pkg/                # 工具包（JWT, 密码, 响应）
│   │   └── zjulogin/           # 浙大登录 SDK
│   ├── migrations/             # 数据库迁移 SQL
│   └── uploads/                # 上传的文件
├── frontend/
│   ├── src/
│   │   ├── api/client.js       # API 调用
│   │   ├── pages/              # 页面组件
│   │   └── components/         # 通用组件
│   └── vite.config.js
└── README.md
```

## API 概览

| 方法 | 路径 | 说明 | 认证 |
|------|------|------|------|
| POST | `/api/v1/auth/register` | 注册 | - |
| POST | `/api/v1/auth/login` | 邮箱登录 | - |
| POST | `/api/v1/auth/zju-login` | 学号登录 | - |
| GET | `/api/v1/posts` | 文章列表 | - |
| GET | `/api/v1/posts/:id` | 文章详情 | - |
| POST | `/api/v1/posts` | 创建文章 | ✅ |
| PUT | `/api/v1/posts/:id` | 更新文章 | ✅ |
| DELETE | `/api/v1/posts/:id` | 删除文章 | ✅ |
| POST | `/api/v1/likes/toggle` | 点赞/取消 | ✅ |
| GET | `/api/v1/search` | 搜索文章 | - |
| GET | `/api/v1/notifications` | 通知列表 | ✅ |
| GET | `/api/v1/users/me` | 个人信息 | ✅ |
| PUT | `/api/v1/users/me` | 修改资料 | ✅ |
| GET | `/api/v1/links` | 友情链接 | - |
| POST | `/api/v1/links` | 添加链接 | ✅ |
| GET | `/api/v1/upload/avatar` | 上传头像 | ✅ |

