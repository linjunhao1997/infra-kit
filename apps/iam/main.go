package main

import (
	"context"
	"infra-kit/apps/iam/ent"
	"infra-kit/apps/iam/pb"
	"infra-kit/apps/iam/server"
	"infra-kit/apps/iam/service"
	"infra-kit/lib/consullib"
	"infra-kit/lib/redislib"
	"log"
	"net"
	"os"
	"time"

	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

var (
	consulAddr  = os.Getenv("consul_addr")
	consulTag   = os.Getenv("consul_tag")
	serviceName = os.Getenv("service_name")
	serviceIp   = os.Getenv("service_ip")
	port        = os.Getenv("port")
	dsn         = os.Getenv("dsn")
	redisAddr   = os.Getenv("redis_addr")
	redisPasswd = os.Getenv("redis_passwd")
)

func main() {
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	// 获取数据库驱动中的sql.DB对象。
	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	rdb := ent.NewClient(ent.Driver(drv), ent.Log(func(a ...any) {}))

	defer rdb.Close()
	if err := rdb.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	cache, err := redislib.NewRedis(&redislib.RedisConf{
		Addr:   redisAddr,
		Passwd: redisPasswd,
	})
	if err != nil {
		log.Fatal(err)
	}

	var opts = []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(server.RecoveryUnaryServerInterceptor),
	}
	gsv := grpc.NewServer(opts...)
	pb.RegisterIAMServiceServer(gsv, server.NewIAMServiceServer(service.NewIAMService(rdb, cache)))

	if serviceIp == "" {
		serviceIp = "127.0.0.1"
	}
	lis, err := net.Listen("tcp", serviceIp+":"+port)
	if err != nil {
		log.Fatalf("failed listening: %s", err)
	}

	if err := consullib.RegisterService(consulAddr, serviceName, serviceIp, port, []string{consulTag}); err != nil {
		log.Fatal(err)
	}

	// Listen for traffic indefinitely.
	if err := gsv.Serve(lis); err != nil {
		log.Fatalf("server ended: %s", err)
	}
}
