package handlers

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	middlewareChi "github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	websocketGo "golang.org/x/net/websocket"

	"github.com/kshamiev/sungora/internal/config"
	"github.com/kshamiev/sungora/internal/graphql"
	"github.com/kshamiev/sungora/internal/middleware"
	"github.com/kshamiev/sungora/internal/websocket"
	"github.com/kshamiev/sungora/pkg/gql"
)

type handlers struct {
	Middleware *middleware.Middleware
	General    *General
	Websocket  *websocket.Websocket
}

// New инициализация обработчиков запросов (хендлеров)
func New(comp *config.Component) (router *chi.Mux) {
	hand := &handlers{}
	router = chi.NewRouter()

	hand.Middleware = middleware.NewMiddleware(comp)

	router.Use(hand.Middleware.Cors().Handler)
	router.Use(middlewareChi.Recoverer)
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
	router.Handle("/api/v1/general/websocket", websocketGo.Handler(hand.General.GetWebSocketSample))

	hand.Websocket = websocket.NewWebsocket(comp)
	router.HandleFunc("/api/v1/websocket/gorilla/{id}", hand.Websocket.GetWebSocketSample)

	return router
}
