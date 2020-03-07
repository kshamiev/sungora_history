package handlers

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"golang.org/x/net/websocket"

	"github.com/kshamiev/sungora/internal/config"
	"github.com/kshamiev/sungora/internal/graphql"
	"github.com/kshamiev/sungora/pkg/gql"
)

type handlers struct {
	Middleware *Middleware
	General    *General
	Websocket  *Websocket
}

// New инициализация обработчиков запросов (хендлеров)
func New(comp *config.Component) (router *chi.Mux) {
	hand := &handlers{}
	router = chi.NewRouter()

	hand.Middleware = NewMiddleware(comp)

	router.Use(hand.Middleware.Cors().Handler)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(comp.Cfg.ServHTTP.RequestTimeout))
	router.Use(hand.Middleware.Logger)
	router.NotFound(hand.Middleware.Static)
	router.Get("/api/docs/*", httpSwagger.Handler()) // swagger

	// GRAPHQL
	router.Handle("/api/playground", handler.Playground("GraphQL playground", "/api/graphql/gql"))
	router.Route("/api/graphql", func(r chi.Router) {
		r.Handle("/gql", handler.GraphQL(gql.NewExecutableSchema(gql.Config{Resolvers: &graphql.Resolver{}})))
	})

	// REST
	hand.General = NewGeneral(comp)
	router.Get("/api/v1/general/ping", hand.General.Ping)
	router.Get("/api/v1/general/version", hand.General.GetVersion)

	// WEBSOCKET
	router.Handle("/api/v1/general/websocket", websocket.Handler(hand.General.WebSocketSample))

	hand.Websocket = NewWebsocket(comp)
	router.HandleFunc("/api/v1/websocket/gorilla/{id}", hand.Websocket.WebSocketSample)

	return router
}
