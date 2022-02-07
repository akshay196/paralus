package main

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	kclient "github.com/ory/kratos-client-go"
	"github.com/spf13/viper"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"

	"github.com/RafaySystems/rcloud-base/components/common/pkg/auth/interceptors"
	authv3 "github.com/RafaySystems/rcloud-base/components/common/pkg/auth/v3"
	"github.com/RafaySystems/rcloud-base/components/common/pkg/gateway"
	logv2 "github.com/RafaySystems/rcloud-base/components/common/pkg/log"
	configrpc "github.com/RafaySystems/rcloud-base/components/common/proto/rpc/config"
	"github.com/RafaySystems/rcloud-base/components/usermgmt/pkg/providers"
	"github.com/RafaySystems/rcloud-base/components/usermgmt/pkg/server"
	"github.com/RafaySystems/rcloud-base/components/usermgmt/pkg/service"
	pbrpcv3 "github.com/RafaySystems/rcloud-base/components/usermgmt/proto/rpc/v3"
	rpcv3 "github.com/RafaySystems/rcloud-base/components/usermgmt/proto/rpc/v3"
	"google.golang.org/grpc"
	_grpc "google.golang.org/grpc"
)

const (
	rpcPortEnv      = "RPC_PORT"
	apiPortEnv      = "API_PORT"
	debugPortEnv    = "DEBUG_PORT"
	kratosSchemeEnv = "KRATOS_SCHEME"
	kratosAddrEnv   = "KRATOS_ADDR"
	dbAddrEnv       = "DB_ADDR"
	dbNameEnv       = "DB_NAME"
	dbUserEnv       = "DB_USER"
	dbPasswordEnv   = "DB_PASSWORD"
	devEnv          = "DEV"
	configAddrENV   = "CONFIG_ADDR"
)

var (
	rpcPort             int
	apiPort             int
	debugPort           int
	rpcRelayPeeringPort int
	kratosScheme        string
	kratosAddr          string
	kc                  *kclient.APIClient
	dbAddr              string
	dbName              string
	dbUser              string
	dbPassword          string
	db                  *bun.DB
	us                  service.UserService
	gs                  service.GroupService
	rs                  service.RoleService
	rrs                 service.RolepermissionService
	is                  service.IdpService
	ps                  service.OIDCProviderService
	dev                 bool
	_log                = logv2.GetLogger()
	authPool            authv3.AuthPool
	configPool          configrpc.ConfigPool
	configAddr          string
)

func setup() {
	viper.SetDefault(rpcPortEnv, 10000)
	viper.SetDefault(apiPortEnv, 11000)
	viper.SetDefault(debugPortEnv, 12000)
	viper.SetDefault(kratosSchemeEnv, "http")
	viper.SetDefault(kratosAddrEnv, "localhost:4433")
	viper.SetDefault(dbAddr, ":5432")
	viper.SetDefault(dbNameEnv, "admindb")
	viper.SetDefault(dbUserEnv, "admindbuser")
	viper.SetDefault(dbPasswordEnv, "admindbpassword")
	viper.SetDefault(devEnv, true)
	viper.SetDefault(configAddrENV, "localhost:7000")

	viper.BindEnv(rpcPortEnv)
	viper.BindEnv(apiPortEnv)
	viper.BindEnv(debugPortEnv)
	viper.BindEnv(kratosSchemeEnv)
	viper.BindEnv(kratosAddrEnv)
	viper.BindEnv(dbAddrEnv)
	viper.BindEnv(dbNameEnv)
	viper.BindEnv(dbPasswordEnv)
	viper.BindEnv(dbPasswordEnv)
	viper.BindEnv(devEnv)
	viper.BindEnv(configAddrENV)

	rpcPort = viper.GetInt(rpcPortEnv)
	apiPort = viper.GetInt(apiPortEnv)
	debugPort = viper.GetInt(debugPortEnv)
	kratosScheme = viper.GetString(kratosSchemeEnv)
	kratosAddr = viper.GetString(kratosAddrEnv)
	dbAddr = viper.GetString(dbAddrEnv)
	dbName = viper.GetString(dbNameEnv)
	dbUser = viper.GetString(dbUserEnv)
	dbPassword = viper.GetString(dbPasswordEnv)
	dev = viper.GetBool(devEnv)
	configAddr = viper.GetString(configAddrENV)

	rpcRelayPeeringPort = rpcPort + 1
	kratosConfig := kclient.NewConfiguration()
	kratosConfig.Servers[0].URL = kratosScheme + "://" + kratosAddr
	kc = kclient.NewAPIClient(kratosConfig)

	dsn := "postgres://admindbuser:admindbpassword@localhost:5432/admindb?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())

	if dev {
		db.AddQueryHook(bundebug.NewQueryHook(
			bundebug.WithVerbose(true),
			bundebug.FromEnv("BUNDEBUG"),
		))
	}

	us = service.NewUserService(providers.NewKratosAuthProvider(kc), db)
	gs = service.NewGroupService(db)
	rs = service.NewRoleService(db)
	rrs = service.NewRolepermissionService(db)
	is = service.NewIdpService(db)
	ps = service.NewOIDCProviderService(db)

	_log.Infow("usermgmt setup complete")
}

func run() {

	ctx := signals.SetupSignalHandler()

	var wg sync.WaitGroup

	wg.Add(1)

	go runAPI(&wg, ctx)
	go runRPC(&wg, ctx)
	go runDebug(&wg, ctx)

	<-ctx.Done()
	wg.Wait()
}

func runAPI(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := http.NewServeMux()

	gwHandler, err := gateway.NewGateway(
		ctx,
		fmt.Sprintf(":%d", rpcPort),
		make([]runtime.ServeMuxOption, 0),
		pbrpcv3.RegisterUserHandlerFromEndpoint,
		pbrpcv3.RegisterGroupHandlerFromEndpoint,
		pbrpcv3.RegisterRoleHandlerFromEndpoint,
		pbrpcv3.RegisterRolepermissionHandlerFromEndpoint,
		pbrpcv3.RegisterIdpHandlerFromEndpoint,
		pbrpcv3.RegisterOIDCProviderHandlerFromEndpoint,
	)
	if err != nil {
		_log.Fatalw("unable to create gateway", "error", err)
	}

	mux.Handle("/", gwHandler)

	s := http.Server{
		Addr:    fmt.Sprintf(":%d", apiPort),
		Handler: mux,
	}
	go func() {
		defer s.Shutdown(context.TODO())
		<-ctx.Done()
	}()

	_log.Infow("starting gateway server", "port", apiPort)

	err = s.ListenAndServe()
	if err != nil {
		_log.Fatalw("unable to start gateway", "error", err)
	}
}

func runRPC(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	// defer configPool.Close()
	defer gs.Close()
	defer rs.Close()
	defer rrs.Close()

	userServer := server.NewUserServer(us)
	groupServer := server.NewGroupServer(gs)
	roleServer := server.NewRoleServer(rs)
	rolepermissionServer := server.NewRolePermissionServer(rrs)
	idpServer := server.NewIdpServer(is)
	oidcProviderServer := server.NewOIDCServer(ps)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", rpcPort))
	if err != nil {
		_log.Fatalw("unable to start rpc listener", "error", err)
	}

	var opts []_grpc.ServerOption
	if !dev {
		opts = append(opts, _grpc.UnaryInterceptor(
			interceptors.NewAuthInterceptorWithOptions(
				interceptors.WithLogRequest(),
				interceptors.WithAuthPool(authPool),
				interceptors.WithExclude("POST", "/v2/sentry/bootstrap/:templateToken/register"),
			),
		))
		defer authPool.Close()
	} else {
		opts = append(opts, _grpc.UnaryInterceptor(
			interceptors.NewAuthInterceptorWithOptions(interceptors.WithDummy())),
		)
	}
	s := grpc.NewServer(opts...)
	if err != nil {
		_log.Fatalw("unable to create grpc server", "error", err)
	}

	go func() {
		defer s.GracefulStop()

		<-ctx.Done()
		_log.Infow("context done")
	}()

	rpcv3.RegisterUserServer(s, userServer)
	rpcv3.RegisterGroupServer(s, groupServer)
	rpcv3.RegisterRoleServer(s, roleServer)
	rpcv3.RegisterRolepermissionServer(s, rolepermissionServer)
	rpcv3.RegisterIdpServer(s, idpServer)
	rpcv3.RegisterOIDCProviderServer(s, oidcProviderServer)

	_log.Infow("starting rpc server", "port", rpcPort)
	err = s.Serve(l)
	if err != nil {
		_log.Fatalw("unable to start rpc server", "error", err)
	}
}

func runDebug(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	s := http.Server{
		Addr: fmt.Sprintf(":%d", debugPort),
	}
	_log.Infow("starting debug server", "port", debugPort)
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			_log.Fatalw("unable to start debug server", "error", err)
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	s.Shutdown(ctx)
}

func main() {
	setup()
	run()
}