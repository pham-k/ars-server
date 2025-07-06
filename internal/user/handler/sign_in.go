package handler

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation"
	"net/http"
	"server/internal/core/helper"
	"server/internal/user/model"
)

type signInReq struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type signInRes struct {
	Object string `json:"object"`
	PID    string `json:"pid"`
}

func (h *handler) SignIn(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req := signInReq{}
	err := helper.ReadJson(r, &req)
	if err != nil {
		helper.WriteJson(w, http.StatusBadRequest, err.Error())
		return
	}

	err = validation.ValidateStruct(&req,
		validation.Field(&req.Phone, validation.Required),
		validation.Field(&req.Password, validation.Required, validation.Length(8, 128)),
	)

	if err != nil {
		helper.WriteJson(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.service.SignIn(ctx, req.Phone, req.Password)
	if err != nil {
		helper.WriteJson(w, http.StatusInternalServerError, err.Error())
		return
	}

	res := &signInRes{
		Object: model.ObjUser,
		PID:    user.PID,
	}

	helper.WriteJson(w, http.StatusOK, res)
}
