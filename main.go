package main

//go:generate swag i -o api/docs

import (
	"context"
	"database/sql"
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
		err    error
		cfg    *config.Config
		db     *sql.DB
		wp     *app.Scheduler
		server *app.Server
	)
	// Logging
	lg := logger.CreateLogger(nil)

	flagConfigPath := flag.String("c", "config.yaml", "used for set path to config file")
	flag.Parse()

	// ConfigApp загрузка конфигурации
	if cfg, err = config.Get(*flagConfigPath); err != nil {
		lg.WithError(err).Fatal("couldn't get config")
	}

	lg = logger.CreateLogger(&cfg.Lg)

	// ConnectDB SqlBoiler
	if db, err = app.NewConnectPostgres(&cfg.Postgresql); err != nil {
		lg.WithError(err).Fatal("couldn't connect to postgres")
	}

	// Workflow
	wp = tasks(db, cfg, lg)
	defer wp.Wait()

	// Server Web & Handlers
	h := handlers.NewMain(db, lg, cfg)
	if server, err = app.NewServer(&cfg.ServHTTP, h, lg); err != nil {
		lg.WithError(err).Fatal("new web server error")
	}

	defer server.Wait(lg)

	// Server GRPC (sample)
	// srv := grpc.NewServer()
	// pb.RegisterNameServer(srv, grpcname.NewServer(db))
	//
	// if grpcServer, err = app.NewGRPC(&cfg.GRPCServer, srv, lg); err != nil {
	// 	lg.WithError(err).Fatal("new grpc server error")
	// }
	//
	// defer grpcServer.Wait(lg)

	// general
	var o models.GooseDBVersion

	if err = queries.Raw(model.SQLAppVersion.String()).Bind(context.Background(), db, &o); err != nil {
		lg.WithError(err).Fatal("couldn't get version DB")
	}

	cfg.App.Version = strconv.FormatInt(o.VersionID, 10) + time.Now().Format("-2006-01-02-15-04-05")

	if cfg.Lg.Level > 4 {
		boil.DebugMode = true
	}

	app.Lock(make(chan os.Signal, 1))
}

func tasks(db *sql.DB, cfg *config.Config, lg logger.Logger) *app.Scheduler {
	wp := app.NewScheduler(lg)
	wp.Add(workers.NewUserOnlineOff(db, cfg))
	wp.Start(workers.UserOnlineOffName)

	return wp
}
