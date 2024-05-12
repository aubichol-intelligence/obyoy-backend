package user

import (
	"obyoy-backend/cache"
	"obyoy-backend/errors"
	"obyoy-backend/model"

	"github.com/sirupsen/logrus"
)

// SessionVerifier is an interface that verifies a user session
type SessionVerifier interface {
	VerifySession(string) (*model.Session, error)
}

// SessionVerifierFunc is a function type
type SessionVerifierFunc func(string) (*model.Session, error)

// VerifySession implements SessionVerifier
func (s SessionVerifierFunc) VerifySession(sessionID string) (*model.Session, error) {
	return s(sessionID)
}

// NewSessionVerifier is function that returns a SessionVerifier
func NewSessionVerifier(cacheSession cache.Session) SessionVerifier {
	f := func(sessionID string) (*model.Session, error) {
		session, err := cacheSession.GetByKey(sessionID)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"err":        err,
				"session_id": sessionID,
			}).Error("could not get session by id")
			return nil, &errors.Unknown{
				Base: errors.Base{"Invalid token", false},
			}
		}

		if session == nil {
			logrus.WithField("session_id", sessionID).Debug("got empty session")
			return nil, &errors.Unauthorized{
				Base: errors.Base{"Invalid token", false},
			}
		}

		return session, nil
	}

	return SessionVerifierFunc(f)
}
