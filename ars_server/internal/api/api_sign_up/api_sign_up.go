package api_sign_up

import (
	"ars_server/internal/helper"
	"ars_server/internal/root"
	"context"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type RequestSignUpWithEmail struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=128"`
}

type ResponseSignUpWithEmail struct {
	Object string `json:"object"`
	Pid    string `json:"pid"`
	Email  string `json:"email"`
}

func SignUpWithEmail(ctx context.Context, authnService root.AuthnService, input RequestSignUpWithEmail) (*root.Customer, error) {
	customer, err := authnService.SignUpWithEmail(ctx, input.Email, input.Password)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func NewRequestSignUpWithEmail(r *http.Request) (RequestSignUpWithEmail, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	req := RequestSignUpWithEmail{}
	err := helper.ReadJson(r, &req)
	if err != nil {
		return req, err
	}

	err = validate.Struct(req)
	if err != nil {
		return req, err
	}

	return req, nil
}

func NewResponseSignUpWithEmail(customer *root.Customer) *ResponseSignUpWithEmail {
	return &ResponseSignUpWithEmail{
		Object: "customer",
		Pid:    customer.Pid,
		Email:  customer.Email,
	}
}
