package handler

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation"
	"net/http"
	"server/internal/core/helper"
	"server/internal/user/model"
)

type signUpReq struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type signUpRes struct {
	Object string `json:"object"`
	PID    string `json:"pid"`
	Phone  string `json:"phone"`
}

func (h *handler) SignUp(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req := signUpReq{}
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

	user, err := h.service.SignUp(ctx, req.Phone, req.Password)
	if err != nil {
		helper.WriteJson(w, http.StatusInternalServerError, err.Error())
		return
	}

	res := &signUpRes{
		Object: model.ObjUser,
		PID:    user.PID,
		Phone:  user.Phone,
	}

	helper.WriteJson(w, http.StatusOK, res)
}
