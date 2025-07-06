package handler

import (
	"net/http"
	"server/internal/core/helper"
)

type signOutReq struct {
	PID string `json:"pid"`
}

type signOutRes struct {
	PID string `json:"pid"`
}

func (h *handler) SignOut(w http.ResponseWriter, r *http.Request) {
	helper.WriteJson(w, http.StatusOK, "")
}
