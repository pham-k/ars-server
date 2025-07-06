package handler

import (
	"context"
	"net/http"
	"server/internal/core/helper"
)

func (h *handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	result, err := h.service.GetUsers(ctx)
	if err != nil {
		helper.WriteJson(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := result
	helper.WriteJson(w, http.StatusOK, response)
}
