# infra-kit
```
// https://github.com/protocolbuffers/protobuf/releases
// 下载protoc包和protobuf包
// 将protoc包解压，将protoc程序保存到GOPATH的bin目录下
// 将protobuf包解压，将src目录覆盖到GOPATH的src目录
// 注意GOPATH全路径不要存在中文，否则影响-I查找路径
// 安装grpc,protobuf的go代码生成程序,grpc-gateway代码生成程序
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

// 安装ent相关程序和依赖
go install entgo.io/ent/cmd/ent@latest
```

## 目录
apps 微服务集
- iam 微服务-访问与控制
    - ent 模型相关
    - pb protobuf模型及代码
    - server grpc实现
    - service 业务实现
gateway 微服务HTTP代理网关
lib 通用包
## 配置
参照.vscode目录配置启动参数
## 启动
apps/<微服务>/main.go启动grpc服务
gateway/main.go启动微服务http代理网关
