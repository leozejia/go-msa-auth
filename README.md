# go-msa-auth

`go-msa-auth` 是一个使用 Go 编写的微服务，用于验证授权码的有效性。客户端通过 HTTP POST 请求与服务交互，服务会根据存储的授权码验证结果返回授权码是否有效。

---

## 功能

- 接收客户端发送的授权码。
- 验证授权码是否有效。
- 返回授权码的验证结果：
  - **有效**：返回 `{"status": "valid"}`
  - **无效**：返回 `{"status": "invalid"}`

---

## 技术栈

- **语言**：Go
- **Web 框架**：Gin
- **容器化**：Docker
- **依赖管理**：Go Modules

---

## 项目结构
go-msa-auth/
├── cmd/                   # 服务入口点
│   └── main.go            # 主入口文件
├── config/                # 配置文件和加载逻辑
│   └── config.go          # 配置加载逻辑
├── internal/              # 核心业务逻辑模块
│   ├── handlers/          # HTTP 路由处理器
│   │   └── auth_handler.go # 授权处理逻辑
│   ├── models/            # 数据库模型和数据访问层
│   │   └── database.go    # 数据库初始化和操作
│   ├── services/          # 业务逻辑层
│   │   └── auth_service.go # 授权服务逻辑
│   └── utils/             # 通用工具
│       └── time_utils.go  # 时间工具函数
├── server/                # 服务器初始化
│   └── server.go          # 路由绑定和服务器启动
├── go.mod                 # Go 模块文件
└── README.md              # 项目说明文档