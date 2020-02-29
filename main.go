package main

//go:generate swag i -o api/docs

import (
	"context"
	"flag"
	"os"
	"strconv"
	"time"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"

	_ "github.com/kshamiev/sungora/api/docs"
	"github.com/kshamiev/sungora/internal/config"
	"github.com/kshamiev/sungora/internal/handlers"
	"github.com/kshamiev/sungora/internal/model"
	"github.com/kshamiev/sungora/internal/workers"
	"github.com/kshamiev/sungora/pkg/app"
	"github.com/kshamiev/sungora/pkg/logger"
	"github.com/kshamiev/sungora/pkg/models"
)

// @title Sungora API
// @description sungora
// @version 1.0.0 DB migrate: 20190921191020
// @contact.name API Support
// @contact.email konstantin@shamiev.ru
// @license.name Sample License
// @termsOfService http://swagger.io/terms/
//
// @host
// @BasePath /
// @schemes http https
//
// @tag.name General
// @tag.description Общие запросы
//
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	var (
		err       error
		server    *app.Server
		component = &handlers.Component{}
	)
	// Logging
	component.Lg = logger.CreateLogger(nil)

	flagConfigPath := flag.String("c", "config.yaml", "used for set path to config file")
	flag.Parse()

	// ConfigApp загрузка конфигурации
	if component.Cfg, err = config.Get(*flagConfigPath); err != nil {
		component.Lg.WithError(err).Fatal("couldn't get config")
	}

	component.Lg = logger.CreateLogger(&component.Cfg.Lg)

	// ConnectDB SqlBoiler
	if component.Db, err = app.NewPostgresBoiler(&component.Cfg.Postgresql); err != nil {
		component.Lg.WithError(err).Fatal("couldn't connect to postgres")
	}

	// Client GRPC (sample)
	// if grpcClientName, err = app.NewGRPCClient(&component.Cfg.GRPCClientName); err != nil {
	// 	component.Lg.WithError(err).Fatal("new grpc client error")
	// }
	// defer grpcClientName.Wait()
	// component.ClientName = pb.NewNameClient(grpcClientName.Conn)

	// Server GRPC (sample)
	// srv := grpc.NewServer()
	// pb.RegisterNameServer(srv, namepkg.NewServer(component.Db))
	//
	// if grpcServer, err = app.NewGRPCServer(&component.Cfg.GRPCServer, srv, component.Lg); err != nil {
	// 	component.Lg.WithError(err).Fatal("new grpc server error")
	// }
	// defer grpcServer.Wait(component.Lg)

	// Workflow
	if component.Wp, err = workers.New(component); err != nil {
		component.Lg.WithError(err).Fatal("workers init")
	}

	defer component.Wp.Wait()

	// Server Web & Handlers
	h := handlers.New(component)
	if server, err = app.NewServer(&component.Cfg.ServHTTP, h, component.Lg); err != nil {
		component.Lg.WithError(err).Fatal("new web server error")
	}

	defer server.Wait(component.Lg)

	// general
	var o models.GooseDBVersion

	if err = queries.Raw(model.SQLAppVersion.String()).Bind(context.Background(), component.Db, &o); err != nil {
		component.Lg.WithError(err).Fatal("couldn't get version DB")
	}

	component.Cfg.App.Version = strconv.FormatInt(o.VersionID, 10) + time.Now().Format("-2006-01-02-15-04-05")

	if component.Cfg.Lg.Level > 4 {
		boil.DebugMode = true
	}

	app.Lock(make(chan os.Signal, 1))
}
