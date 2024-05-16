package user

import (
	"fmt"
	"time"

	"obyoy-backend/cache"
	cfg "obyoy-backend/config"
	"obyoy-backend/errors"
	"obyoy-backend/model"
	"obyoy-backend/user/dto"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

// Login interface defines the functionalities for logging in
type Login interface {
	CreateToken(*dto.Login) (*dto.Token, error)
}

// longin holds login related
type login struct {
	emailAndPasswordChecker EmailAndPasswordChecker
	cacheSession            cache.Session
	cfg                     cfg.Session
}

// CreateToken implements Login interface
func (l *login) CreateToken(loginDto *dto.Login) (*dto.Token, error) {
	user, err := l.emailAndPasswordChecker.EmailAndPasswordCheck(loginDto)
	if err != nil {
		return nil, fmt.Errorf("Failed email and password check %w", err)
	}

	session := model.Session{
		Key:       fmt.Sprintf("%s:%s", uuid.New().String(), user.ID),
		UserID:    user.ID,
		ExpiredAt: time.Now().Add(l.cfg.Length),
		CreatedAt: time.Now(),
	}

	if err = l.cacheSession.Create(&session); err != nil {
		logrus.WithField("user email", user.Email).Error("Could not create session")
		return nil, &errors.Unknown{
			Base: errors.Base{"Unable to create token", false},
		}
	}

	return &dto.Token{
		Token:       session.Key,
		ID:          user.ID,
		AccountType: user.AccountType,
	}, nil
}

// NewLoginParams lists all the parameters for NewLogin
type NewLoginParams struct {
	dig.In
	EmailAndPasswordChecker EmailAndPasswordChecker
	CacheSession            cache.Session
	Cfg                     cfg.Session
}

// NewLogin provides Login
func NewLogin(params NewLoginParams) Login {
	return &login{
		emailAndPasswordChecker: params.EmailAndPasswordChecker,
		cacheSession:            params.CacheSession,
		cfg:                     params.Cfg,
	}
}
