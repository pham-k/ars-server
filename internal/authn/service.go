package authn

//
//import (
//	"context"
//	"database/sql"
//	"errors"
//	"github.com/alexedwards/argon2id"
//	gonanoid "github.com/matoous/go-nanoid/v2"
//	"server/internal/core/logger"
//	"server/internal/core/database"
//)
//
//type service struct {
//	logger logger.logger
//	rDB database.Database
//}
//
//func NewService(logger logger.logger, rDB database.Database) Service {
//	return &service{
//		logger: logger,
//		rDB: rDB,
//	}
//}
//
//func (s service) RegisterWithEmail(ctx context.Context, email, password string) (*User, error) {
//	passwordHash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
//	if err != nil {
//		return nil, err
//	}
//
//	pID, err := s.newUserPID()
//	if err != nil {
//		return nil, err
//	}
//
//	repo, tx, err := s.rDB.NewRepoWithTx()
//	if err != nil {
//		return nil, err
//	}
//	defer tx.Rollback()
//
//	user, err := repo.RegisterUserWithEmail(ctx, repository.RegisterUserWithEmailParams{
//		Pid:          pID,
//		AuthnType:    sql.NullString{String: string(AuthnTypeEmail), Valid: true},
//		Email:        sql.NullString{String: email, Valid: true},
//		PasswordHash: passwordHash,
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	if err = tx.Commit(); err != nil {
//		return nil, err
//	}
//
//	output := &User{
//		ID:        user.ID,
//		PID:       user.Pid,
//		Object:    ObjUser,
//		AuthnType: Type(user.AuthnType.String),
//		Email:     user.Email.String,
//	}
//	return output, nil
//}
//
//func (s service) LogInWithEmail(ctx context.Context, email, password string) (*User, error) {
//	repo := s.rDB.NewRepo()
//	result, err := repo.GetUserFromEmail(ctx, repository.GetUserFromEmailParams{
//		Email:     sql.NullString{String: email, Valid: true},
//		AuthnType: sql.NullString{String: string(AuthnTypeEmail), Valid: true},
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	match, err := argon2id.ComparePasswordAndHash(password, result.PasswordHash)
//	if !match {
//		return nil, errors.New("invalid credentials")
//	}
//
//	user := &User{
//		ID:     result.ID,
//		PID:    result.Pid,
//		Object: ObjUser,
//	}
//
//	return user, nil
//}
//
//func (s service) LogOut(ctx context.Context, userPID string) error {
//	return nil
//}
//
//func (s service) newUserPID() (string, error) {
//	nanoID, err := gonanoid.New()
//	if err != nil {
//		return "", err
//	}
//
//	pID := PIDPrefixUser + "_" + nanoID
//	return pID, nil
//}
//
////func (s service) ValidateEmail(ctx context.Context, token string) (bool, error) {
////	repo, tx, err := s.rDB.NewRepoWithTx()
////	if err != nil {
////		return false, err
////	}
////	defer tx.Rollback()
////
////	userID, err := repo.GetUserIDFromEmailValidationToken(ctx, token)
////	if err != nil {
////		return false, err
////	}
////	if !userID.Valid {
////		return false, errors.New("invalid user ID")
////	}
////
////	authnEmail, err := repo.GetAuthnEmail(ctx, userID.Int64)
////	if err != nil {
////		return false, err
////	}
////	if len(authnEmail) != 1 {
////		return false, errors.New("multiple authn email")
////	}
////
////	activated, err := repo.ValidateEmail(ctx, authnEmail[0].Email)
////	if err != nil {
////		return false, err
////	}
////
////	if err = tx.Commit(); err != nil {
////		return false, nil
////	}
////	return activated, nil
////}
