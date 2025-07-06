package service

import (
	"context"
	"errors"
	"github.com/alexedwards/argon2id"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"server/internal/user/model"
	"server/internal/user/repository"
	"server/internal/user/repository/sgen"
)

type Service interface {
	GetUsers(ctx context.Context) ([]*model.User, error)
	SignUp(ctx context.Context, phone string, password string) (*model.User, error)
	SignIn(ctx context.Context, phone string, password string) (*model.User, error)
	//SignOut(ctx context.Context, pID string) (*model.User, error)
}

type service struct {
	Repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{
		Repo: repo,
	}
}

func (s *service) GetUsers(ctx context.Context) ([]*model.User, error) {
	query := s.Repo.NewQuery()
	data, err := query.ListUsers(ctx)
	if err != nil {
		return nil, err
	}
	var users []*model.User
	for _, user := range data {
		users = append(users, &model.User{
			PID:    user.Pid,
			Phone:  user.Phone,
			Object: "user",
		})
	}
	return users, nil
}

func (s *service) SignUp(ctx context.Context, phone string, password string) (*model.User, error) {
	passwordHash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return nil, err
	}

	pID, err := s.newUserPID()
	if err != nil {
		return nil, err
	}

	q, tx, err := s.Repo.NewQueryTx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	repoUser, err := q.CreateUser(ctx, sgen.CreateUserParams{
		Pid:          pID,
		Phone:        phone,
		CountryCode:  "VN",
		PasswordHash: passwordHash,
	})
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	output := &model.User{
		PID:    repoUser.Pid,
		Object: model.ObjUser,
		Phone:  repoUser.Phone,
	}
	return output, nil
}

func (s *service) SignIn(ctx context.Context, phone string, password string) (*model.User, error) {
	q := s.Repo.NewQuery()
	repoUser, err := q.GetUserByPhone(ctx, phone)
	if err != nil {
		return nil, err
	}

	match, err := argon2id.ComparePasswordAndHash(password, repoUser.PasswordHash)
	if !match {
		return nil, errors.New("invalid credentials")
	}

	user := &model.User{
		Object: model.ObjUser,
		PID:    repoUser.Pid,
		Phone:  repoUser.Phone,
	}

	return user, nil
}

func (s *service) newUserPID() (string, error) {
	nanoID, err := gonanoid.New()
	if err != nil {
		return "", err
	}

	pID := model.PIDPrefixUser + "_" + nanoID
	return pID, nil
}
