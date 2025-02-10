package server

import (
	"ars_server/internal/api/api_app_configuration"
	"ars_server/internal/helper"
	"context"
	"net/http"
)

func (s *Server) HandleGetAppConfigurations() http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		data, hasMore, err := api_app_configuration.ListAppConfigurations(ctx, s.ConfigService)
		if err != nil {
			s.Log.Error("Fail to get app configurations", err)
			helper.WriteJson(writer, http.StatusInternalServerError, "")
		}

		response := api_app_configuration.NewResponseListAppConfigurations(data, hasMore)
		helper.WriteJson(writer, http.StatusOK, response)
	}
}
