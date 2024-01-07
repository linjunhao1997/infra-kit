package main

import (
	"context"
	"infra-kit/apps/iam/ent"
	"infra-kit/apps/iam/pb"
	"infra-kit/apps/iam/server"
	"infra-kit/apps/iam/service"
	"infra-kit/lib/consullib"
	"infra-kit/lib/grpclib"
	"infra-kit/lib/loglib"
	"infra-kit/lib/metriclib"
	"infra-kit/lib/redislib"
	"infra-kit/lib/tracelib"
	"log"
	"log/slog"
	"net"
	"os"
	"time"

	"github.com/oklog/run"

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
	loglib.Init()
	err := tracelib.Init("iam-service")
	if err != nil {
		slog.Error("tracelib Init failed", "err", err)
		return
	}
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
		slog.Error("failed create db schema", "err", err)
	}

	cache, err := redislib.NewRedis(&redislib.RedisConf{
		Addr:   redisAddr,
		Passwd: redisPasswd,
	})
	if err != nil {
		slog.Error("failed connect redis", "err", err)
	}

	var opts = []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(grpclib.LoggingUnaryServerInterceptor, grpclib.MetricUnaryServerInterceptor, grpclib.RecoveryUnaryServerInterceptor),
		grpc.ChainStreamInterceptor(grpclib.LoggingStreamServerInterceptor, grpclib.MetricStreamServerInterceptor, grpclib.RecoveryStreamServerInterceptor),
		grpc.StatsHandler(grpclib.StatsServerHandler),
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

	g := &run.Group{}
	g.Add(func() error {
		slog.Info("grpc server listening", "port", port)
		return gsv.Serve(lis)
	}, func(err error) {
		gsv.GracefulStop()
		gsv.Stop()
	})

	httpSrv := metriclib.InitMetricsHttpServer(":8081")
	g.Add(func() error {
		return httpSrv.ListenAndServe()
	}, func(err error) {
		if err := httpSrv.Close(); err != nil {
			slog.Error("failed to stop web server", "err", err)
		}
	})

	if err := g.Run(); err != nil {
		slog.Error("server ended", "err", err)
		os.Exit(1)
	}
}
