package session

import (
	"filmlibr/internal/entity"
	"net/http"
)

type SessionManager interface {
	Create(user entity.User) (string, error)
	Check(r *http.Request) (Session, error)
}

type SessionManagerImpl struct {
}

func NewSessionManagerImpl() *SessionManagerImpl {
	return &SessionManagerImpl{}
}

func (sm *SessionManagerImpl) Create(user entity.User) (string, error) {
	return "", nil
}

func (sm *SessionManagerImpl) Check(r *http.Request) (Session, error) {
	return Session{}, nil
}
