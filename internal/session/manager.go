package session

import (
	"filmlibr/internal/entity"
	"net/http"
)

type SessionManager interface {
	Create(user entity.User) (string, error)
	Check(r *http.Request) (Session, error)
}
