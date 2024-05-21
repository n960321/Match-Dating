package server

import (
	"Match-Dating/internal/handler"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Server struct {
	app *http.Server
	cfg *Config
}

type Config struct {
	Port    string `mapstructure:"port"`
	Prefork bool   `mapstructure:"prefork"`
}

func NewServer(cfg Config, h *handler.Handler) *Server {
	g := gin.New()
	g.Use(gin.Recovery())
	g.Use(timeMiddleware())

	api := g.Group("/api")
	v1 := api.Group("/v1")

	v1.POST("/Proile", h.CreateProfile)
	v1.DELETE("/Proile/:profile_id", h.DeleteProfile)
	v1.POST("/QuerySinglePeople", h.QuerySinglePeople)
	return &Server{
		app: &http.Server{
			Addr:    cfg.Port,
			Handler: g,
		},
	}
}
func (svr *Server) Run() {
	go func() {
		log.Info().Msgf("Server Start Listening on %s", svr.app.Addr)
		if err := svr.app.ListenAndServe(); err != nil {
			log.Fatal().Err(err)
		}
	}()
}

func (svr *Server) Shutdown(ctx context.Context) {
	if err := svr.app.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server show failed")
	}
}
