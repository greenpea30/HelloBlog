# Overview

如果你已经阅读了博客框架的部分，你应该对博客框架开头的**intro**有了更深的理解

为了能够较为深入地实践更多的知识（尤其是传统后端、agent开发、运维方面），我们还是推荐大家尝试搭建从0到1的博客网站，可以调整一下目标：我们不再开发**专属个人博客**，而是开发一款面向**所有人的博客小产品（**类似知乎、小红书、小黑盒这种**）**，可以尝试引入以下功能：

1. AI写作助手
2. MCP Server（让agent能够与你的产品交互）
3. 语义搜索
4. 用户画像和文章推荐（手搓的话有一点难）
5. ...

你可以挑选感兴趣的部分尝试，也欢迎大家**发挥想象力**，尝试更多没有提到的有趣点子，从而摆脱教学项目的**同质化**问题（这也正是我们所苦恼的！）

相信有**vibe-coding** 和文档 的帮助，大家能够顺利完成它！

**教学示例**选择采用React+Golang作为主要开发语言，它们在我们的知识图谱中都有涉及。

同时会使用Gin框架、GORM等

但是大家可以自由选择自己偏好的开发语言和工具

**请不要被教学示例所束缚，你不需要把所有示例都看一遍，它们只是引导，独立自主的学习可能会更高效**





1. 由于是教学示例，需求分析、模块划分这些**流程不会做的非常严谨，仅供参考**
2. 请按验收方案实现基础模块下的**基本功能**（黑色）；验收方案要求的**进阶功能**会用红色标记

# 需求分析

## 用户功能

1. 用户注册和登录
2. 用户信息管理
3. 第三方登录（Github）
4. 数据埋点与用户仪表盘
5. 头像管理与云对象存储

## 博客功能

1. 博客的创建、展示、搜索
2. 支持markdown等格式的文本编辑器
3. 博客热度排行榜
4. 博客的全文检索
5. 博客的语义检索（可以与4合并成混合检索）
6. 在线阅读人数统计
7. 博客推荐或者同风格用户推荐

## 点赞功能

1. 用户对博客和评论能进行点赞或取消点赞‘

## 评论功能

1. 博客评论
2. 划线评论
3. 引用评论（楼中楼）

## 订阅功能

1. 订阅与消息通知
2. 每日新闻/大会关注推送

## 日志功能

1. 日志和日志分析
2. 可观测性

## 运维

1. Docker compose编排
2. CI/CD

## AI自动化方向

1. 文章内容总结
2. AI写作助手
3. Github Trending 自动解读
4. AI 技术论文 Digest

## 其它

1. 移动端适配
2. MCP Server
3. 任务栏与知识树
4. 话题讨论/留言墙
5. 漂流瓶
6. 小桌宠
7. 仿https://neal.fun/（感谢@怀玺 分享的网站，也可以学习这个网站展示大家的小实验和小项目）

# 数据库模型

你要**根据自己想要实现的业务逻辑**自行设计数据模型

题外话：由于不同数据库的sql写法存在差异，大家可以尝试https://dbdiagram.io/通过DBML定义模型，并按照实际需求导出特定数据库的sql语句

DBML示例：

```Plain
Enum like_target_type {
  post
  comment
}

Table likes {
  id integer [primary key, increment]

  user_id integer [not null, note: '点赞用户ID']

  target_type like_target_type [not null, note: '点赞对象类型：post 或 comment']
  target_id integer [not null, note: '点赞对象ID，对应 posts.id 或 comments.id']

  created_at timestamp [not null, default: `now()`]

  indexes {
    (user_id, target_type, target_id) [unique]
    (target_type, target_id)
  }
}
```

一些tips:

1. 由于点赞对象可以是博客也可以是评论，因此可以创建一个枚举型，后面也方便扩展新的点赞对象
2. 根据需求，我们要**创建合适的索引以加速查询**。例如我们要判断一个用户是否已经给该对象点过赞了，那么就创建(user_id, target_type, target_id)联合索引查询

推荐回顾知识图谱中的**数据库索引**部分

# 设计API

推荐回顾知识图谱中的**API设计**部分

接下来要根据具体业务设计前后端通信所需的API

```Plain
// 统一response，之后的response都是对data的描述
{
    "code": 0, // 错误码，0表示成功，其它可以参照网上的一些俗成约定
    "msg": 'success', // 返回信息
    "data": any // 返回数据
}
```

## 用户注册和登录为例：

URL: {host}/api/v1/auth/register

Method: POST

```Plain
//我的注册逻辑涉及这三个
{
  "username": "alice",
  "email": "alice@example.com",
  "password": "password123"
}
//再次声明这是"data"部分
{
  "user": {
    "id": 1,
    "username": "alice",
    "email": "alice@example.com",
    "created_at": "2026-05-25T10:00:00Z"
  },
  "access_token": "jwt-token"
}
```

# 项目结构

根据你的语言设计合理的项目结构和目录树

如果你已经跟随知识图谱对代码规范进行了一定的探索，那么你应该对一些架构已经有所了解，例如MVC、DDD等

教学示例主要依照了MVC的思想，划分了三层：

1. Controller层，主要解析请求，返回响应
2. Service层，具体的业务逻辑
3. DAO层(Data Access Object），在这里主要是用ORM与数据库交互

(**Repository**要比DAO要更上层一点，可能隐藏了选择内存、本地文件、缓存或者数据库这种细节，即可能使用了多种DAO。在这里出于实际需求和代码量的考虑，只用了一种DAO，但其实也简单地多封装了一层为Repository)

## 后端目录树

```Plain
backend/
├── cmd/                             # 程序入口，main.go
│   ├── api/
│   │   └── main.go                  # HTTP API 启动入口
|   └── mcp/
|       └── main.go                  # MCP server启动入口
│
├── api/                             # 第三方API
│   └── openapi/
│       └── helloblog.v1.yaml
│
├── etc/
│   └── config.example.yaml          # 配置文件
│
├── internal/
│   ├── config/
│   │   └── config.go                # 配置加载
│   │
│   ├── server/
│   │   ├── http.go                  # Gin server 装配
│   │   ├── router.go                # 总路由注册
│   │   ├── user_routes.go           # 用户模块路由注册
│   │   └── middleware/              # Gin 中间件
│   │       ├── auth.go              # JWT 鉴权
│   │       └── cors.go              # 跨域处理
│   │
│   ├── svc/
│   │   └── service_context.go       # 参考 go-zero，将 infra、controller、service 统一装配
│   │
│   ├── controller/                  # Controller 层
│   │   ├── user/
│   │   │   ├── controller.go        # 用户 Controller 结构体与依赖
│   │   │   ├── register.go          # 用户注册接口
│   │   │   ├── login.go             # 用户登录接口
│   │   │   └── me.go                # 当前用户信息接口
│   │   ├── post/
│   │   │   └── ...                  # 文章模块接口
│   │   └── ...                      # comment、like、search 等模块接口
│   │
│   ├── dto/                         # DTO：前后端通信的请求/响应对象
│   │   ├── user.go                  # 用户 DTO
│   │   └── ...                      # post、comment、like、search 等 DTO
│   │
│   ├── service/                     # Service 层
│   │   ├── user/
│   │   │   ├── service.go           # 用户Service
│   │   │   ├── register.go          # 用户注册业务逻辑
│   │   │   ├── login.go             # 用户登录业务逻辑
│   │   │   ├── me.go                # 当前用户查询业务逻辑
│   │   │   └── mapper.go            # 用户模型转换
│   │   ├── post/
│   │   │   └── ...                  # 文章业务逻辑
│   │   ├── comment/
│   │   │   └── ...                  # 评论业务逻辑
│   │   └── ...
│   │
│   ├── dao/                         # DAO 层
│   │   ├── model/                   # GORM Model，要与 migrations 中的表结构对齐
│   │   │   ├── user.go              # 用户表模型
│   │   │   └── ... 
│   │   ├── user.go                  # 用户表数据访问
│   │   └── ...
│   │
│   ├── infra/                       # 基础设施：数据库、Redis、本地模型客户端
│   │   ├── db/
│   │   │   ├── postgres.go          # PostgreSQL 和 GORM 初始化
│   │   │   └── transaction.go       # 数据库事务工具
│   │   ├── redis/
│   │   │   └── redis.go             # Redis 初始化
│   │   └── embedding/
│   │       └── client.go            # 本地词嵌入模型客户端
│   │
│   └── pkg/                         # 这里是一些与业务逻辑关系不大、比较通用容易复用的代码
│       ├── response/                # 统一响应结构、错误码
│       ├── jwt/                     # JWT 签发与校验
│       ├── password/                # 密码哈希
│       └── ...                      # pagination、validator、logger 等通用工具
│
├── migrations/
│   └── 000001_init_schema.sql       # 初始化 SQL 语句
│
└── go.mod
```



注意，该文档面向不擅长使用agent或者没有思路的同学，仅供参考

1. 如果你不选择Go作为你的开发语言，那么了解每一层大概在做什么就可以了，不要深挖代码细节
2. 你可以借助vibe coding先帮助你实现大部分基本模块，之后在理解大致逻辑的基础上尝试完成一部分工作。大部分基础模块都是在做CURD，这非常枯燥，学习效率极低

以实现**注册登录功能**为例，我们来看下它涉及了哪些后端目录文件，梳理后端运行的大致逻辑

```Plain
backend/
├── cmd/                             # 程序入口，main.go
│   └── api/
│       └── main.go                  # HTTP API 启动入口
│
├── etc/
│   └── config.example.yaml          # 配置文件
│
├── internal/
│   ├── config/
│   │   └── config.go                # 配置加载
│   │
│   ├── server/
│   │   ├── http.go                  # Gin server 装配
│   │   ├── router.go                # 总路由注册
│   │   ├── user_routes.go           # 用户模块路由注册
│   │   └── middleware/              # Gin 中间件
│   │       ├── auth.go              # JWT 鉴权
│   │       └── cors.go              # 跨域处理
│   │
│   ├── svc/
│   │   └── service_context.go       # 参考 go-zero，将 infra、controller、service 统一装配
│   │
│   ├── controller/                  # Controller 层
│   │   └── user/
│   │       ├── controller.go        # 用户 Controller 结构体与依赖
│   │       ├── register.go          # 用户注册接口
│   │       ├── login.go             # 用户登录接口
│   │       └── me.go                # 当前用户信息接口
│   │
│   ├── dto/                         # DTO：前后端通信的请求/响应对象
│   │   └── user.go                  # 用户 DTO
│   │
│   ├── service/                     # Service 层
│   │   └── user/
│   │       ├── service.go           # 用户 Service
│   │       ├── register.go          # 用户注册业务逻辑
│   │       ├── login.go             # 用户登录业务逻辑
│   │       ├── me.go                # 当前用户查询业务逻辑
│   │       └── mapper.go            # 用户模型转换
│   │
│   ├── dao/                         # DAO 层
│   │   ├── model/
│   │   │   └── user.go              # users 表 GORM Model
│   │   └── user.go                  # 用户表数据访问
│   │
│   ├── infra/                       # 基础设施
│   │   ├── db/
│   │   │   ├── postgres.go          # PostgreSQL 和 GORM 初始化
│   │   │   └── transaction.go       # 数据库事务工具
│   │   └── redis/
│   │       └── redis.go             # Redis 初始化，用户模块可预留使用
│   │
│   └── pkg/                         # 通用工具代码
│       ├── response/                # 统一响应结构、错误码
│       ├── jwt/                     # JWT 签发与校验
│       └── password/                # 密码哈希
│
├── migrations/
│   └── 000001_init_schema.sql       # 初始化 SQL 语句，包含 users 表结构
│
└── go.mod
```

## 路由

推荐回顾知识图谱中的以下内容：

1. 路由（可能在node.js中介绍的比较多）

```Go
func RegisterRoutes(engine *gin.Engine, svcCtx *svc.ServiceContext) {
    // 所有接口统一加上 /api/v1 前缀，方便后续做版本管理
    api := engine.Group("/api/v1")

    registerUserRoutes(api, svcCtx)
}
func registerUserRoutes(api *gin.RouterGroup, svcCtx *svc.ServiceContext) {
    /*
    认证相关接口：注册、登录
    Group相当于一个分组前缀，你可以在分组上应用统一中间件
    中间件能在处理请求前和返回响应前添加一些处理逻辑，例如校验权限
    详情请参考https://gin-gonic.com/en/docs/middleware
    其它语言也有类似的使用
    */
    auth := api.Group("/auth")
    {
        //当访问/api/v1/auth/register时，请求会由该Controller解析处理
        auth.POST("/register", svcCtx.Controllers.User.Register)
        auth.POST("/login", svcCtx.Controllers.User.Login)
    }

    userRoutes := api.Group("/users")
    /*
    下面这些用户接口需要登录后才能访问，比如获取个人信息
    怎么拦截？使用前面提到的中间件
    */
    userRoutes.Use(middleware.Auth(svcCtx.JWT))
    {
        userRoutes.GET("/me", svcCtx.Controllers.User.Me)
    }
}
```

## Controller

```Go
//这里将Controller层与Service层联系起来
type Controller struct {
    service userservice.UseCase
}

func NewController(service userservice.UseCase) *Controller {
    return &Controller{service: service}
}
```

请求根据前面的路由被交由特定的Controller处理，例如下方是Login接口的处理逻辑，大概做了这些事

1. 对参数进行校验，如果参数不够或者错误就会直接返回错误响应
2. 进入Service层，执行真正的业务逻辑
3. 得到结果后返回响应

```Go
func (ctl *Controller) Login(c *gin.Context) {
    var req dto.LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        response.Fail(c, response.NewError(response.CodeInvalid, "invalid request"))
        return
    }

    resp, err := ctl.service.Login(req)
    if err != nil {
        response.Fail(c, err)
        return
    }

    response.OK(c, resp)
}
```

## Service

推荐回顾知识图谱中的以下内容：

1. 后端语言的interface（Go、Java等都有接口的概念）
2. 登录鉴权（jwt）

接下来是最为核心的逻辑部分

1. 通过DAO层与数据库进行交互，查找对应邮件的用户，校验用户密码等

**userRepository**这里用到了**DDD**的**依赖倒置**（有点大杂烩，可能不是特别好... 主要是为了方便不依赖数据库的测试）

即service不关心接口方法的具体实现逻辑，比如它不在乎GetByID是从哪查询的数据

DAO层或者测试代码可以实现所有的接口方法，之后就能够实现**注入**

1. 生成对应的jwt，用于用户后续请求的鉴权
2. 返回结果

```Go
type UseCase interface {
        Register(req dto.RegisterRequest) (*dto.AuthResponse, error)
        Login(req dto.LoginRequest) (*dto.AuthResponse, error)
        GetMe(userID int64) (*dto.UserResponse, error)
}

type Service struct {
        users userRepository
        jwt   *jwtpkg.Manager
}

type userRepository interface {
        Create(user *model.User) error
        GetByID(id int64) (*model.User, error)
        GetByEmail(email string) (*model.User, error)
        UsernameExists(username string) (bool, error)
        EmailExists(email string) (bool, error)
}

func NewService(users userRepository, jwtManager *jwtpkg.Manager) *Service {
        return &Service{users: users, jwt: jwtManager}
}
func (s *Service) Login(req dto.LoginRequest) (*dto.AuthResponse, error) {
    email := strings.ToLower(strings.TrimSpace(req.Email))
    user, err := s.users.GetByEmail(email)
    if err != nil {
        if dao.IsNotFound(err) {
            return nil, response.NewError(response.CodeUnauthorized, "invalid email or password")
        }
        return nil, response.Wrap(response.CodeInternalError, "internal server error", err)
    }

    if !passwordpkg.Verify(user.PasswordHash, req.Password) {
        return nil, response.NewError(response.CodeUnauthorized, "invalid email or password")
    }

    token, err := s.jwt.Generate(user.ID)
    if err != nil {
        return nil, response.Wrap(response.CodeInternalError, "internal server error", err)
    }

    return &dto.AuthResponse{
        User:        toUserResponse(user),
        AccessToken: token,
    }, nil
}
```

## DTO

同时你也应该注意到了`dto.AuthResponse`这类结构体，它们不是必须的(dto我们之前在**前期准备后端目录树**中提到过)

```Go
type RegisterRequest struct {
        Username string `json:"username" binding:"required,min=3,max=50"`
        Email    string `json:"email" binding:"required,email,max=100"`
        Password string `json:"password" binding:"required,min=8,max=72"`
}

type LoginRequest struct {
        Email    string `json:"email" binding:"required,email,max=100"`
        Password string `json:"password" binding:"required,min=8,max=72"`
}

type UserResponse struct {
        ID        int64     `json:"id"`
        Username  string    `json:"username"`
        Email     string    `json:"email"`
        CreatedAt time.Time `json:"created_at,omitempty"`
}

type AuthResponse struct {
        User        UserResponse `json:"user"`
        AccessToken string       `json:"access_token"`
}
```

dto不是完整的一个实体，它没有业务方法，在项目中仅用于前后端通信

通过定义的`toUserResponse`之类的方法，我们将model转化为dto

以快速实现请求数据的解析、响应时敏感数据和无用数据的过滤等

```Go
func toUserResponse(user *model.User) dto.UserResponse {
        return dto.UserResponse{
                ID:        user.ID,
                Username:  user.Username,
                Email:     user.Email,
                CreatedAt: user.CreatedAt,
        }
}
```

## DAO（Model）

推荐回顾知识图谱中的以下内容：

1. 数据库和ORM

首先我们要根据之前搭建的数据库模型，在代码中定义对应的model

由于我们要通过ORM进行与数据库的交互，我们需要添加对应的tag让ORM如何与数据库交互。

例如通过column来形成结构体字段与数据库列名的映射

```Go
type User struct {
    ID           int64          `gorm:"primaryKey;column:id"`
    Username     string         `gorm:"column:username;size:50;unique;not null"`
    Email        string         `gorm:"column:email;size:100;unique;not null"`
    PasswordHash string         `gorm:"column:password_hash;size:255;not null"`
    CreatedAt    time.Time      `gorm:"column:created_at;autoCreateTime"`
    UpdatedAt    time.Time      `gorm:"column:updated_at;autoUpdateTime"`
    //gorm.DeleteAt告诉ORM这个列仅进行软删除，在删除时仅标注删除时间，之后查询时会忽略这些数据
    DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at"`
}
```

接着使用ORM接口与数据库进行交互，实现根据ID获取用户等函数

```Go
func (d *UserDAO) GetByID(id int64) (*model.User, error) {
        var user model.User
        if err := d.db.First(&user, "id = ?", id).Error; err != nil {
                return nil, err
        }
        return &user, nil
}
```

## Server

推荐回顾知识图谱中的以下内容：

1.  登录鉴权中提到的CORS，我们在这里使用了CORS中间件快速解决它引发的问题

在之后我们使用Gin构建我们的Server：

1. 应用全局中间件
2. 根据基础设施生成service、contrller并注册路由
3. 启动服务器

```Go
type HTTPServer struct {
        cfg    config.Config
        engine *gin.Engine
}

func New(cfg config.Config, database *gorm.DB, redisClient *redis.Client) *HTTPServer {
        gin.SetMode(cfg.Server.Mode)

        engine := gin.New()
        engine.Use(gin.Logger(), middleware.Recovery(), middleware.RequestID(), middleware.CORS())

        serviceContext := svc.NewServiceContext(cfg, database, redisClient)
        RegisterRoutes(engine, serviceContext)

        return &HTTPServer{
                cfg:    cfg,
                engine: engine,
        }
}

func (s *HTTPServer) Run() error {
        return s.engine.Run(s.cfg.Server.Addr)
}
func main() {
    /*
    加载配置
    什么是配置？例如我们连接数据库、redis都需要password、host、port等等
    1. 如果把这些配置分散在项目各处，就很不利于我们管理
    2. 如果在代码中使用明文，那么在上传代码后别人能够直接看到这些敏感数据，这很不安全。
    因此我们会将敏感数据写在.env/.yaml文件中，通过docker或者其它库加载.env，再在internal/config/config.go中读取或采用默认值
    */
        cfg := config.Load
    
    /*
    利用config对要使用的基础设施进行统一初始化
    */
        database, err := db.NewPostgres(cfg.Database.DSN)
        if err != nil {
                log.Fatalf("connect postgres: %v", err)
        }

        redisClient := redisinfra.New(cfg.Redis)
        if redisClient != nil {
                defer redisClient.Close()
        }

    /*
    在config中我们定义了ServerConfig
    type ServerConfig struct {
            Addr string `yaml:"addr"`
            Mode string `yaml:"mode"`
    }
    在.yaml或者.env中配置的Addr上开始监听
    */
        httpServer := server.New(cfg, database, redisClient)
        if err := httpServer.Run(); err != nil {
                log.Fatalf("run http server: %v", err)
        }
}
```

## Tips

1. 正如开头所说的你可以先让agent生成大致框架，再尝试着手部分工作
2. 你可以使用Postman等工具尝试测试启动好的Server(Agent能够生成Postman导入文件)



推荐回顾知识图谱中的以下内容：

1. 登录鉴权之Oauth

PS:关于登录鉴权有很多第三方服务如clerk，或者开源项目如keycloak，很多时候自己实现一个登录鉴权功能是很不安全的，因此你在项目中完全可以使用它们。但是如果你对其中的原理感兴趣，或许可以尝试自己探索探索

1. 数据库之migrate

# 一种可行的业务逻辑

1. 如果通过第三方登录，获取第三方账号信息后，在数据库中校验博客账户绑定情况
2. 如果已有博客账户，直接登录成功；如果没有博客账户，需选择创建新博客账户或者关联已有博客账户（未与第三方账户绑定）

# 一种可行的设计逻辑

这里使用了数据库迁移（migration）进行数据库结构的调整，相关工具很多例如

1. Gorm的Auto Migrate
2. Go Migrate（示例用的是这个）
3. ...其它语言应该也有对应的工具

1. 创建第三方账号与Hello Blog账号的关联表

```SQL
CREATE TABLE "oauth_identities" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "user_id" integer NOT NULL,
  "provider" varchar(30) NOT NULL,
  "provider_user_id" varchar(100) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE UNIQUE INDEX "oauth_identities_provider_user_id_idx"
ON "oauth_identities" ("provider", "provider_user_id");

CREATE UNIQUE INDEX "oauth_identities_user_provider_idx"
ON "oauth_identities" ("user_id", "provider");

CREATE INDEX "oauth_identities_user_id_idx"
ON "oauth_identities" ("user_id");

ALTER TABLE "oauth_identities"
ADD CONSTRAINT "user_oauth_identities"
FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE DEFERRABLE INITIALLY IMMEDIATE;
```

可以尝试解释这三条索引分别是为了加速哪些操作

Tips：例如第一个索引：快速找出某一个第三方账号绑定的用户/是否已绑定用户

1. 邮箱密码允许为空字段

```Go
ALTER TABLE "users" ALTER COLUMN "email" DROP NOT NULL;
ALTER TABLE "users" ALTER COLUMN "password_hash" DROP NOT NULL;
```

# Oauth流程

## 注册应用

首先在[Github Oauth Apps](https://github.com/settings/developers)注册你的应用，确定应用的HomePage URL和Authorization callback URL，并获取clientID和client secret

将它们配置到你的项目中去

## 一种可行的路由和重定向设计

教学示例使用了 golang.org/x/oauth2 包，以简化并规范Oauth流程

大致流程是

1. 前端点击"使用Github账户登陆"，朝后端的/authorize接口发起一个GET请求
2. 后端出于安全性设计，生成一个随机字符串state（使用redis进行保存）， 并将浏览器重定向至一个由 **`clientID`****、****`redirect_uri`****、****`scope`** **和刚刚生成的** **`state`** **拼接而成的 GitHub 官方授权 URL**。当授权成功返回redirect_url时需要携带state

```Go
/*
加载配置文件中的配置，用它们组装一个oauth2包的GithubClient，用于处理一系列授权事宜
RedirectURL:授权成功后的，必须跟注册应用时配置的Authorization callback URL一致
Scope:权限范围，例如"read:user"表示该用户的公开个人资料的读权限
State:用于应对CSRF等安全问题，请自行了解
*/
githubClient := oauthpkg.NewGitHubClient(
    cfg.OAuth.GitHub.ClientID,
    cfg.OAuth.GitHub.ClientSecret,
    cfg.OAuth.GitHub.RedirectURL,
)
...

func NewGitHubClient(clientID, clientSecret, redirectURL string) *GitHubClient {
    return &GitHubClient{
        config: &oauth2.Config{
            ClientID:     clientID,
            ClientSecret: clientSecret,
            RedirectURL:  redirectURL,
            Scopes:       []string{"read:user"},
            Endpoint:     github.Endpoint,
        },
        httpClient: http.DefaultClient,
    }
}

func (c *GitHubClient) AuthCodeURL(state string) string {
    return c.config.AuthCodeURL(state, oauth2.AccessTypeOnline)
}
...

func (s *StateStore) Create(ctx context.Context) (string, error) {
    if s.redis == nil {
        return "", errRedisRequired
    }

    state, err := randomToken(32)
    if err != nil {
        return "", err
    }
    return state, s.redis.Set(ctx, "oauth:state:"+state, "1", s.ttl).Err()
}
```

1. github授权成功后回调redirect_url并携带code和state。通过检验state来判断请求的安全性，调用**ExchangeCode**用code与github交换token
2. 接着调用**GetUser**拿token从Github获取用户信息，并在我们的关联表中确认是否已有博客账户。如果已有账户，那么直接返回博客应用的token；否则返回一个ticket，前端以ticket为凭证选择创建新博客账户或者关联已有博客账户（未与第三方账户绑定）

```Go
func (s *Service) GitHubCallback(ctx context.Context, code, state string) (*dto.OAuthCallbackResponse, error) {
    if s.githubOAuth == nil || s.oauthStates == nil || s.oauthTickets == nil || s.oauthIdentities == nil {
        return nil, response.NewError(response.CodeInternalError, "github oauth is not configured")
    }

    code = strings.TrimSpace(code)
    state = strings.TrimSpace(state)
    if code == "" || state == "" {
        return nil, response.NewError(response.CodeInvalid, "invalid request")
    }

    if err := s.oauthStates.Consume(ctx, state); err != nil {
        return nil, response.NewError(response.CodeUnauthorized, "invalid oauth state")
    }

    accessToken, err := s.githubOAuth.ExchangeCode(ctx, code)
    if err != nil {
        return nil, response.Wrap(response.CodeUnauthorized, "github oauth failed", err)
    }

    githubUser, err := s.githubOAuth.GetUser(ctx, accessToken)
    if err != nil {
        return nil, response.Wrap(response.CodeUnauthorized, "github oauth failed", err)
    }

    providerUserID := oauthpkg.ProviderUserIDFromGitHubID(githubUser.ID)
    identity, err := s.oauthIdentities.GetByProviderUserID(oauthpkg.GitHubProvider, providerUserID)
    if err == nil {
        return s.authenticatedOAuthResponse(identity.UserID)
    }
    if !dao.IsNotFound(err) {
        return nil, response.Wrap(response.CodeInternalError, "internal server error", err)
    }

    ticket, err := s.oauthTickets.Create(ctx, oauthpkg.PendingIdentity{
        Provider:       oauthpkg.GitHubProvider,
        ProviderUserID: providerUserID,
    })
    if err != nil {
        return nil, response.Wrap(response.CodeInternalError, "internal server error", err)
    }

    return &dto.OAuthCallbackResponse{
        Status:      "binding_required",
        OAuthTicket: ticket,
        Provider:    oauthpkg.GitHubProvider,
    }, nil
}
func (c *GitHubClient) ExchangeCode(ctx context.Context, code string) (string, error) {
        token, err := c.config.Exchange(ctx, code)
        if err != nil {
                return "", err
        }
        if !token.Valid() || token.AccessToken == "" {
                return "", fmt.Errorf("github token exchange failed: invalid token")
        }
        return token.AccessToken, nil
}

func (c *GitHubClient) GetUser(ctx context.Context, accessToken string) (*GitHubUser, error) {
        clientContext := context.WithValue(ctx, oauth2.HTTPClient, c.httpClient)
        client := oauth2.NewClient(clientContext, oauth2.StaticTokenSource(&oauth2.Token{
                AccessToken: accessToken,
                TokenType:   "Bearer",
        }))

        req, err := http.NewRequestWithContext(clientContext, http.MethodGet, "https://api.github.com/user", nil)
        if err != nil {
                return nil, err
        }
        req.Header.Set("Accept", "application/vnd.github+json")

        resp, err := client.Do(req)
        if err != nil {
                return nil, err
        }
        defer resp.Body.Close()

        if resp.StatusCode < 200 || resp.StatusCode >= 300 {
                return nil, fmt.Errorf("github user request failed: status %d", resp.StatusCode)
        }

        var user GitHubUser
        if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
                return nil, err
        }
        if user.ID <= 0 {
                return nil, fmt.Errorf("github user id is empty")
        }

        return &user, nil
}
```





推荐回顾知识图谱相关内容：

1. 搜索引擎
2. 向量数据库

教学选择采用PostgreSQL，因为它支持多种插件，能同时兼顾基本的数据存储、全文检索、语义检索等多种功能。但是更多的功能也意味着更多的概念和耦合...

总之没有标准方案

你也可以使用其他方案，例如使用mysql实现基本存储，再使用milvus来实现语义检索，这或许对于新手更加清晰

# 可以参考的资料

**飞书知识图谱及图谱中的参考资料**

[PostgreSQL中文文档](https://postgresql.ac.cn/docs/current/index.html)PostgreSQL中文文档但不推荐（有广告、搜索功能会让你回到英文官网、不支持中文检索）

[PostgreSQL官方英文文档](https://www.postgresql.org/docs/current/index.html)

[一篇关于PostgreSQL混合检索的博客](https://jkatz05.com/post/postgres/hybrid-search-postgres-pgvector/)（如果力竭了这篇那么没太必要往下看了0.o）

# 全文检索

教学采用`ts_vector`来实现全文检索：[官方文档](https://www.postgresql.org/docs/current/textsearch-controls.html?utm_source=chatgpt.com)

```SQL
/*
zhparser是一个PostgreSQL中文分词扩展
如果当前数据库里还没有启用zhparser扩展，就启用它
*/
CREATE EXTENSION IF NOT EXISTS zhparser;

/*
创建一个全文搜索配置，之后有的函数就可以采用"chinese"这个配置对文本进行分词和判断词性
将 n, v, a, i, e, l这些词性的词保留下来
*/
CREATE TEXT SEARCH CONFIGURATION chinese (PARSER = zhparser);
ALTER TEXT SEARCH CONFIGURATION chinese ADD MAPPING FOR n,v,a,i,e,l WITH simple;

/*
除了常规的一些字段外，我们还有一个新字段search_vector，存储我们的分词处理结果
*/
CREATE TABLE "posts" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "user_id" integer NOT NULL,
  "title" varchar(200) NOT NULL,
  "summary" text,
  "content" text NOT NULL,
  "search_vector" tsvector,
  "like_count" integer NOT NULL DEFAULT 0,
  "comment_count" integer NOT NULL DEFAULT 0,
  "status" record_status NOT NULL DEFAULT 'normal',
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

/*
这里创建了一个触发器（数据库的一个概念）
当插入帖子或帖子的title、summary、content发生变动时，计算search_vector
*/
CREATE FUNCTION "posts_search_sync"()
RETURNS trigger AS $$
BEGIN
  NEW."search_vector" =
    setweight(to_tsvector('chinese', coalesce(NEW."title", '')), 'A') ||
    setweight(to_tsvector('chinese', coalesce(NEW."summary", '')), 'B') ||
    setweight(to_tsvector('chinese', coalesce(NEW."content", '')), 'C');
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER "posts_search_sync_trigger"
BEFORE INSERT OR UPDATE OF "title", "summary", "content"
ON "posts"
FOR EACH ROW
EXECUTE FUNCTION "posts_search_sync"();

/*
跟B+树索引类似，GIN索引也是用来加速搜索的
关于原理，大家可以参考以下文档：
https://www.postgresql.org/docs/18/gin.html
https://www.postgresql.org/docs/18/textsearch-indexes.html
*/
CREATE INDEX "posts_search_vector_idx" ON "posts" USING GIN ("search_vector");

/*
接下来定义全局检索函数，大致逻辑如下：
1. 确定输入和输出
2，对输入进行清晰，例如通过trim过滤掉query的空格
3. 将query转化为ts_query
4. 按照ts_query和search_vector的匹配程度进行rank
*/
CREATE FUNCTION "full_text_search_posts"(
  "query_text" text,
  "match_count" integer DEFAULT 20
)
RETURNS TABLE (
  "post_id" integer,
  "title" varchar(200),
  "summary" text,
  "created_at" timestamp,
  "rank" integer,
  "score" double precision
) AS $$
WITH "input" AS (
  SELECT
    nullif(trim("query_text"), '') AS "q",
    greatest("match_count", 1) AS "limit_n"
),
"text_query" AS (
  SELECT websearch_to_tsquery('chinese', "q") AS "tsq"
  FROM "input"
  WHERE "q" IS NOT NULL
)
SELECT
  p."id",
  p."title",
  p."summary",
  p."created_at",
  row_number() OVER (
    ORDER BY ts_rank_cd(p."search_vector", tq."tsq", 32) DESC, p."created_at" DESC
  )::integer AS "rank",
  ts_rank_cd(p."search_vector", tq."tsq", 32)::double precision AS "score"
FROM "posts" p
CROSS JOIN "text_query" tq
WHERE p."status" = 'normal'
  AND p."search_vector" @@ tq."tsq"
ORDER BY "score" DESC, p."created_at" DESC
LIMIT (SELECT "limit_n" FROM "input");
$$ LANGUAGE sql STABLE;
```

# 语义检索

教学采用pg_vector来实现语义检索：[Github-pgvector](https://github.com/pgvector/pgvector)

```SQL
/*
vector扩展提供了vector类型、`<=>`cosine distance运算符、HNSW/IVFFlat向量索引能力
*/
CREATE EXTENSION IF NOT EXISTS vector;

/*
单独维护一张post的嵌入向量的表
*/
CREATE TABLE "post_embeddings" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "post_id" integer UNIQUE NOT NULL,
  "embedding" vector(1024) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

/*
HNSW能帮助我们快速找出邻近向量
详情请参考知识图谱的向量数据库部分
*/
CREATE INDEX "post_embeddings_embedding_hnsw_idx"
ON "post_embeddings" USING hnsw ("embedding" vector_cosine_ops)
WITH (m = 16, ef_construction = 64);

/*
如果帖子被删除了，它的嵌入向量也可以删除了，因此添加级联删除
*/
ALTER TABLE "post_embeddings" ADD CONSTRAINT "post_embedding_ref" FOREIGN KEY ("post_id") REFERENCES "posts" ("id") ON DELETE CASCADE DEFERRABLE INITIALLY IMMEDIATE;

/*
接下来定义向量检索函数，大致逻辑如下：
1. 确定输入和输出
2，对输入进行清晰，例如通过trim过滤掉query的空格
3. 按照post向量和query向量的相似度进行rank
*/
CREATE FUNCTION "vector_search_posts"(
  "query_embedding" vector(1024),
  "match_count" integer DEFAULT 20
)
RETURNS TABLE (
  "post_id" integer,
  "title" varchar(200),
  "summary" text,
  "created_at" timestamp,
  "rank" integer,
  "distance" double precision
) AS $$
WITH "input" AS (
  SELECT
    "query_embedding" AS "embedding",
    greatest("match_count", 1) AS "limit_n"
)
SELECT
  p."id",
  p."title",
  p."summary",
  p."created_at",
  row_number() OVER (
    ORDER BY pe."embedding" <=> i."embedding", p."created_at" DESC
  )::integer AS "rank",
  (pe."embedding" <=> i."embedding")::double precision AS "distance"
FROM "input" i
JOIN "post_embeddings" pe ON i."embedding" IS NOT NULL
JOIN "posts" p ON p."id" = pe."post_id"
WHERE p."status" = 'normal'
ORDER BY pe."embedding" <=> i."embedding", p."created_at" DESC
LIMIT (SELECT "limit_n" FROM "input");
$$ LANGUAGE sql STABLE;
```

# 混合检索

现在我们既能进行全局检索，也能进行向量检索，是时候将它们的结果合并起来了！

然而全局检索和混合检索的评分标准完全不同，如何合并？

## RRF

这里采取的是RRF(Reciprocal Rank Fusion)

大致逻辑如下：

```
score = 1 / (k + rank)
```

1. 针对每一种检索，我们根据它的排名取倒数，得到一个新的分数
2. k是平滑参数，为了避免排名靠前和靠后的分数差距过大，例如1/1和1/60平滑为1/(60+1)和1/(60+60)
3. 最后将每一种检索的分数加起来重新排序

```SQL

CREATE FUNCTION "hybrid_search_posts"(
  "query_text" text,
  "query_embedding" vector(1024) DEFAULT NULL,
  "match_count" integer DEFAULT 20,
  "candidate_count" integer DEFAULT 100,
  "text_weight" double precision DEFAULT 0.45,
  "vector_weight" double precision DEFAULT 0.55,
  "rrf_k" integer DEFAULT 60
)
RETURNS TABLE (
  "post_id" integer,
  "title" varchar(200),
  "summary" text,
  "created_at" timestamp,
  "score" double precision
) AS $$
WITH "input" AS (
  SELECT
    greatest("match_count", 1) AS "limit_n",
    greatest("candidate_count", "match_count", 1) AS "candidate_n",
    greatest("rrf_k", 1) AS "rank_k",
    greatest("text_weight", 0.0) AS "text_w",
    greatest("vector_weight", 0.0) AS "vector_w"
),
"ranked_candidates" AS (
  SELECT
    ft."post_id",
    i."text_w" / (i."rank_k" + ft."rank") AS "score_part"
  FROM "input" i
  CROSS JOIN "full_text_search_posts"("query_text", i."candidate_n") ft

  UNION ALL

  SELECT
    vs."post_id",
    i."vector_w" / (i."rank_k" + vs."rank") AS "score_part"
  FROM "input" i
  CROSS JOIN "vector_search_posts"("query_embedding", i."candidate_n") vs
),
"combined" AS (
  SELECT
    "post_id",
    sum("score_part")::double precision AS "score"
  FROM "ranked_candidates"
  GROUP BY "post_id"
)
SELECT
  p."id",
  p."title",
  p."summary",
  p."created_at",
  c."score"
FROM "combined" c
JOIN "posts" p ON p."id" = c."post_id"
ORDER BY c."score" DESC, p."created_at" DESC
LIMIT (SELECT "limit_n" FROM "input");
$$ LANGUAGE sql STABLE;
```





推荐回顾知识图谱中的以下内容：

1. Agent
2. MCP
3. 登录鉴权之Session
4. 实时通信技术之SSE

在教学案例中，MCP Server 是一个独立进程，它和 HTTP API 共用同一套业务层

# 可以参考的资料

谷歌官方开发并维护了一批[MCP-SDK](https://modelcontextprotocol.io/docs/sdk)，大家可以根据自己的语言选择对应的SDK搭建MCP Server以及Client

你也可以先熟悉一下：[MCP Server官方教程](https://modelcontextprotocol.io/docs/develop/build-server#python)

[Go MCP Server Auth示例](https://github.com/modelcontextprotocol/go-sdk/blob/main/examples/server/auth-middleware/main.go#L287)

# Server

创建一个MCP server非常方便

```Go
func NewServer(svcCtx *svc.ServiceContext) *sdkmcp.Server {
        /*
    https://pkg.go.dev/github.com/modelcontextprotocol/go-sdk/mcp#NewServer
    func NewServer(impl *Implementation, options *ServerOptions) *Server
    1. Implementation用来描述程序本身的信息
    2. 一些可选配置：https://pkg.go.dev/github.com/modelcontextprotocol/go-sdk/mcp#ServerOptions
    */
    server := sdkmcp.NewServer(&sdkmcp.Implementation{
                Name:    "helloblog-mcp",//用于代码识别和使用
                Version: "0.1.0",//定义你的程序版本
        }, nil)

        registerTools(server, svcCtx)
        return server
}
```

## 添加Tools

除了添加Tools还可以添加Prompts、Resources原语，可以自行了解

```Go
func Register(server *sdkmcp.Server, svcCtx *svc.ServiceContext) {
    /*
    1.需要添加tool的server
    2.一个Tool的描述
    3.最后需要一个handler
    */
    sdkmcp.AddTool(server, &sdkmcp.Tool{
        Name:        ToolHealth,
        Description: "Check Helloblog PostgreSQL and Redis connectivity.",
    }, handler(svcCtx))
}
func handler(svcCtx *svc.ServiceContext) sdkmcp.ToolHandlerFor[Input, Output] {
    /*
    这里使用了泛型
    根据SDK描述 CallToolRequest和CallToolResult会自动从Input和Output中生辰，大多数情况下可以忽略
    */
    return func(ctx context.Context, _ *sdkmcp.CallToolRequest, _ Input) (*sdkmcp.CallToolResult, Output, error) {
        status := Output{
            Postgres: "unknown",
            Redis:    "disabled",
        }

        sqlDB, err := svcCtx.DB.DB()
        if err != nil {
            status.Postgres = fmt.Sprintf("error: %v", err)
        } else if err := sqlDB.PingContext(ctx); err != nil {
            status.Postgres = fmt.Sprintf("error: %v", err)
        } else {
            status.Postgres = "ok"
        }

        if svcCtx.Redis != nil {
            if err := svcCtx.Redis.Ping(ctx).Err(); err != nil {
                status.Redis = fmt.Sprintf("error: %v", err)
            } else {
                status.Redis = "ok"
            }
        }

        return nil, status, nil
    }
}
```

# Session

Session这个概念我们在知识图谱的登录鉴权中提到过

在这里，我们希望同一个client和server之间的交流可以保持一个有状态的会话，就像微信聊天一样，而不是网游中的小世界喇叭（莫名想起来的比喻）

为什么？以SSE为例：有时候Server需要较长时间才能返回数据，这时为了减轻连接压力，它会断开长连接，要求Client在一定时间后再来获取结果。这时候server就要进行会话管理，以针对性的恢复上下文

[MCP Session Management](https://modelcontextprotocol.io/specification/2025-11-25/basic/transports#session-management)

# Transports

[MCP Transports](https://modelcontextprotocol.io/specification/2025-11-25/basic/transports#streamable-http)

MCP Server和MCP Client必须有一种通信的手段，官方提供了两种方式：

1. stdio：Server和Client经过标准输入输出流进行通信。很显然这个不是我们想要的
2. Streamable HTTP： Server 作为一个可以处理多个 client 连接的独立 HTTP 服务运行，并且通过 HTTP POST / GET 传输 JSON-RPC 消息 。

如何得到HTTP服务？

```Go
func NewHTTPHandler(server *sdkmcp.Server, cfg config.MCPConfig) http.Handler {
    /*重点就是NewStreamableHTTPHandler这个函数，func NewStreamableHTTPHandler(
        getServer func(*http.Request) *Server,
        opts *StreamableHTTPOptions,
    ) *StreamableHTTPHandler
    1.会话管理函数，例如始终返回同一个会话或针对不同请求来源给予不同会话
    Client                                                   Server
     ⇅                          (jsonrpc2)                     ⇅
    ClientSession ⇄ Client Transport ⇄ Server Transport ⇄ ServerSession
    2.一些可选配置：详情参考https://pkg.go.dev/github.com/modelcontextprotocol/go-sdk/mcp#NewStreamableHTTPHandler
        a.Stateless:无状态，因此基本上只有client请求,server回复，sever无法请求
        b.JSONResponse：是否将可流式传播的流打包为json返回
    */
    handler := sdkmcp.NewStreamableHTTPHandler(func(*http.Request) *sdkmcp.Server {
        return server
    }, &sdkmcp.StreamableHTTPOptions{
        Stateless:    true,
        JSONResponse: true,
    })

    protection := http.NewCrossOriginProtection()
    for _, origin := range cfg.AllowedOrigins {
        if origin == "" {
            continue
        }
        _ = protection.AddTrustedOrigin(origin)
    }

    return protection.Handler(handler)
}
```

# Auth

关于鉴权保护，官方推荐的是使用oauth，像vscode与mcp server形成连接采取的也是对应的一套端点发现、client注册、鉴权流程，比较复杂

[MCP: The Authorization Flow: Step by Step](https://modelcontextprotocol.io/docs/tutorials/security/authorization#the-authorization-flow-step-by-step)

[Go MCP Server Auth示例](https://github.com/modelcontextprotocol/go-sdk/blob/main/examples/server/auth-middleware/main.go#L287)

大家可以参照官方文档进行自行探索