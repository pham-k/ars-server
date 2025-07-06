package authn

//
//import (
//	"context"
//	"fmt"
//	"github.com/go-playground/validator/v10"
//	"net/http"
//	"server/internal/core/helper"
//	"server/internal/core/logger"
//	"server/internal/token"
//	"time"
//)
//
//type RequestSignUpWithEmail struct {
//	Email    string `json:"email" validate:"required,email"`
//	Password string `json:"password" validate:"required,min=8,max=128"`
//}
//
//type ResponseSignUpWithEmail struct {
//	Object string `json:"object"`
//	PID    string `json:"pid"`
//	Email  string `json:"email"`
//}
//
//type RequestLogInWithEmail struct {
//	Email    string `json:"email" validate:"required,email"`
//	Password string `json:"password" validate:"required,min=8,max=128"`
//}
//
//type ResponseLogInWithEmail struct {
//	Object              string `json:"object"`
//	PID                 string `json:"pid"`
//	AuthenticationToken string `json:"authentication_token"`
//}
//
//type RequestLogOut struct {
//	PID string `json:"pid"`
//}
//
//type ResponseLogOut struct{}
//
//func NewRequestSignUpWithEmail(r *http.Request) (RequestSignUpWithEmail, error) {
//	validate := validator.New(validator.WithRequiredStructEnabled())
//
//	req := RequestSignUpWithEmail{}
//	err := helper.ReadJson(r, &req)
//	if err != nil {
//		return req, err
//	}
//
//	err = validate.Struct(req)
//	if err != nil {
//		return req, err
//	}
//
//	return req, nil
//}
//
//func NewResponseSignUpWithEmail(user *User) *ResponseSignUpWithEmail {
//	return &ResponseSignUpWithEmail{
//		Object: ObjUser,
//		PID:    user.PID,
//		Email:  user.Email,
//	}
//}
//
//func RegisterWithEmail(ctx context.Context, input RequestSignUpWithEmail, log logger.Logger,
//	service Service, tokenService token.Service) (*User, error) {
//	user, err := service.RegisterWithEmail(ctx, input.Email, input.Password)
//	if err != nil {
//		return nil, err
//	}
//
//	_, err = tokenService.GenerateToken(ctx, token.EmailValidation, 1*time.Hour)
//	if err != nil {
//		return nil, err
//	}
//
//	return user, nil
//}
//
//func LogInWithEmail(ctx context.Context, input RequestLogInWithEmail, log logger.Logger,
//	service Service, tokenService token.Service) (*User, *token.Token, error) {
//	user, err := service.LogInWithEmail(ctx, input.Email, input.Password)
//	if err != nil {
//		return nil, nil, err
//	}
//
//	authnToken, err := tokenService.GenerateToken(ctx, token.Authentication, 24*time.Hour)
//	if err != nil {
//		return nil, nil, err
//	}
//	log.Info(fmt.Sprintf("authnToken: %+v", authnToken))
//
//	authnToken.Data = fmt.Sprintf("%v::%v", "userID", user.ID)
//	log.Info(fmt.Sprintf("authnToken: %+v", authnToken))
//
//	err = tokenService.StoreToken(ctx, authnToken)
//	if err != nil {
//		return nil, nil, err
//	}
//
//	tokenData, err := tokenService.GetToken(ctx, token.Authentication, authnToken.Value)
//	if err != nil {
//		return nil, nil, err
//	}
//
//	log.Info(tokenData)
//
//	return user, authnToken, nil
//}
//
//func NewRequestLogInWithEmail(r *http.Request) (RequestLogInWithEmail, error) {
//	validate := validator.New(validator.WithRequiredStructEnabled())
//
//	req := RequestLogInWithEmail{}
//	err := helper.ReadJson(r, &req)
//	if err != nil {
//		return req, err
//	}
//
//	err = validate.Struct(req)
//	if err != nil {
//		return req, err
//	}
//
//	return req, nil
//}
//
//func NewResponseLogInWithEmail(user *User, token *token.Token) *ResponseLogInWithEmail {
//	return &ResponseLogInWithEmail{
//		Object:              ObjUser,
//		PID:                 user.PID,
//		AuthenticationToken: token.Value,
//	}
//}
//
//func NewRequestLogOut(r *http.Request) (RequestLogOut, error) {
//	validate := validator.New(validator.WithRequiredStructEnabled())
//
//	req := RequestLogOut{}
//	err := helper.ReadJson(r, &req)
//	if err != nil {
//		return req, err
//	}
//
//	err = validate.Struct(req)
//	if err != nil {
//		return req, err
//	}
//
//	return req, nil
//}
//
//func NewResponseLogOut(user *User) *ResponseLogOut {
//	return &ResponseLogOut{}
//}
