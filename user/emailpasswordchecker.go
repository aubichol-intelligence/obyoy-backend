package user

import (
	"obyoy-backend/errors"
	"obyoy-backend/model"
	storeuser "obyoy-backend/store/user"
	"obyoy-backend/user/dto"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// EmailAndPasswordChecker is an interface that checks the validity for email and password
type EmailAndPasswordChecker interface {
	EmailAndPasswordCheck(*dto.Login) (*model.User, error)
}

// EmailAndPasswordCheckerFunc implements EmailAndPasswordChecker
type EmailAndPasswordCheckerFunc func(*dto.Login) (*model.User, error)

// EmailAndPasswordCheck implements EmailAndPasswordChecker interface
func (e EmailAndPasswordCheckerFunc) EmailAndPasswordCheck(dto *dto.Login) (*model.User, error) {
	return e(dto)
}

// NewEmailAndPasswordChecker provides EmailAndPasswordChecker
func NewEmailAndPasswordChecker(storeUsers storeuser.Users) EmailAndPasswordChecker {
	f := func(login *dto.Login) (*model.User, error) {
		user, err := storeUsers.FindByEmail(login.Email)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"err":   err,
				"email": login.Email,
			}).Error("to get user by email")

			return nil, &errors.Unknown{
				Base: errors.Base{"Invalid email or password", false},
			}
		}

		if user == nil {
			logrus.WithField("email", login.Email).Debug("Got empty user")
			return nil, &errors.Unknown{
				Base: errors.Base{"Invalid email or password", false},
			}
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
		if err != nil {
			logrus.Debug("Passwords did not match")
			return nil, &errors.Unknown{
				Base: errors.Base{"Invalid email or password", false},
			}
		}

		if user.AccountType != login.AccountType {
			logrus.Debug("Types did not match")
			return nil, &errors.Unknown{
				Base: errors.Base{"Invalid account type", false},
			}
		}
		return user, nil
	}

	return EmailAndPasswordCheckerFunc(f)
}
