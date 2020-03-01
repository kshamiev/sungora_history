package handlers

import (
	"database/sql"

	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"golang.org/x/net/websocket"

	"github.com/kshamiev/sungora/internal/config"
	"github.com/kshamiev/sungora/internal/graphql"
	"github.com/kshamiev/sungora/pkg/app"
	"github.com/kshamiev/sungora/pkg/gql"
	"github.com/kshamiev/sungora/pkg/logger"
)

type Component struct {
	Db  *sql.DB
	Lg  logger.Logger
	Cfg *config.Config
	Wp  *app.Scheduler
}

type component struct {
	db  *sql.DB
	lg  logger.Logger
	cfg *config.Config
	wp  *app.Scheduler
}

func newComponent(comp *Component) *component {
	return &component{
		db:  comp.Db,
		lg:  comp.Lg,
		cfg: comp.Cfg,
		wp:  comp.Wp,
	}
}

type Handler struct {
	*component
	Middleware *Middleware
	General    *General
}

// New инициализация обработчиков запросов (хендлеров)
func New(comp *Component) (router *chi.Mux) {
	hand := &Handler{
		component: newComponent(comp),
	}
	router = chi.NewRouter()

	hand.Middleware = NewMiddleware(hand)

	router.Use(hand.Middleware.Cors().Handler)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(hand.cfg.ServHTTP.RequestTimeout))
	router.Use(hand.Middleware.Logger)
	router.NotFound(hand.Middleware.Static)
	router.Get("/api/docs/*", httpSwagger.Handler()) // swagger

	// GRAPHQL
	router.Handle("/api/playground", handler.Playground("GraphQL playground", "/api/graphql/gql"))
	router.Route("/api/graphql", func(r chi.Router) {
		r.Handle("/gql", handler.GraphQL(gql.NewExecutableSchema(gql.Config{Resolvers: &graphql.Resolver{}})))
	})

	// REST
	hand.General = NewGeneral(hand)
	router.Get("/api/v1/general/ping", hand.General.Ping)
	router.Get("/api/v1/general/version", hand.General.GetVersion)

	// WEBSOCKET
	router.Handle("/api/v1/general/websocket", websocket.Handler(hand.General.WebSocketSample))

	return router
}
