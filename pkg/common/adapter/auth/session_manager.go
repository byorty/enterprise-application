package auth

import (
	"context"
	"github.com/byorty/enterprise-application/pkg/common/adapter/application"
	"github.com/byorty/enterprise-application/pkg/common/adapter/log"
	"github.com/byorty/enterprise-application/pkg/common/adapter/server/grpc"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type SessionManager interface {
	CreateTokenBySession(context.Context, pbv1.Session) (string, error)
	CreateToken(ctx context.Context, userUUID string, group pbv1.UserGroup) (string, error)
	GetSessionByToken(ctx context.Context, token string) (pbv1.Session, error)
}

func NewFxSessionManager(
	provider application.Provider,
	logger log.Logger,
	jwtHelper JWTHelper,
) (SessionManager, error) {
	var cfg SessionConfig
	err := provider.PopulateByKey("session", &cfg)
	if err != nil {
		return nil, err
	}

	return &sessionManager{
		logger:    logger.Named("session_manager"),
		cfg:       cfg,
		jwtHelper: jwtHelper,
	}, nil
}

type sessionManager struct {
	logger    log.Logger
	cfg       SessionConfig
	jwtHelper JWTHelper
}

func (s *sessionManager) CreateTokenBySession(ctx context.Context, session pbv1.Session) (string, error) {
	claims := &SessionClaims{
		StandardClaims: jwt.StandardClaims{
			Audience:  s.cfg.Audience,
			Issuer:    s.cfg.Issuer,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(s.cfg.Duration).Unix(),
		},
		Session: session,
	}

	return s.jwtHelper.CreateToken(claims)
}

func (s *sessionManager) CreateToken(ctx context.Context, userUUID string, group pbv1.UserGroup) (string, error) {
	return s.CreateTokenBySession(
		ctx,
		pbv1.Session{
			Uuid:  userUUID,
			Group: group,
		},
	)
}

func (s *sessionManager) GetSessionByToken(ctx context.Context, token string) (pbv1.Session, error) {
	logger := s.logger.WithCtx(ctx)
	if len(token) == 0 {
		logger.Debug("guest session")
		return pbv1.Session{
			Group: pbv1.UserGroupGuest,
		}, nil
	}

	claims := new(SessionClaims)
	err := s.jwtHelper.Parse(token, claims)
	if err != nil {
		logger.Error(err)
		return pbv1.Session{}, grpc.ErrUnauthenticated(grpc.ErrSessionNotFound)
	}

	logger.Debugf("b2c session=%s with group=%s", claims.Session.Uuid, claims.Session.Group)
	return claims.Session, nil
}
