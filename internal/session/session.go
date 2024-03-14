package session

import (
	"context"
	"errors"
	"filmlibr/internal/entity"
)

type sessKey string

type Session struct {
	Token string
	User  entity.User
}

var SessionKey sessKey = "sessionKey"

var (
	ErrNoAuth   = errors.New("no session found")
	ErrBadToken = errors.New("bad token")
)

func SessionFromContext(ctx context.Context) (Session, error) {
	sess, ok := ctx.Value(SessionKey).(Session)
	if !ok {
		return Session{}, ErrNoAuth
	}
	return sess, nil
}

func ContextWithSession(ctx context.Context, sess Session) context.Context {
	return context.WithValue(ctx, SessionKey, sess)
}
