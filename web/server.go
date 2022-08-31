package web

import (
	"net/http"

	"github.com/eurofurence/reg-export-service/internal/repository/config"
	"github.com/eurofurence/reg-export-service/internal/repository/logging"
	"github.com/eurofurence/reg-export-service/web/controller/healthctl"
	"github.com/eurofurence/reg-export-service/web/filter/corsfilter"
	"github.com/eurofurence/reg-export-service/web/filter/logreqid"
	"github.com/eurofurence/reg-export-service/web/filter/reqid"
	"github.com/go-chi/chi/v5"
)

func Create() chi.Router {
	logging.NoCtx().Info("Building routers...")
	server := chi.NewRouter()

	server.Use(reqid.RequestIdMiddleware())
	server.Use(logreqid.LogRequestIdMiddleware())
	server.Use(corsfilter.CorsHeadersMiddleware())

	healthctl.Create(server)
	// add your controllers here
	return server
}

func Serve(server chi.Router) {
	address := config.ServerAddr()
	logging.NoCtx().Info("Listening on " + address)
	err := http.ListenAndServe(address, server)
	if err != nil {
		logging.NoCtx().Error(err)
	}
}
