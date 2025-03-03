package main

import (
	"context"
	"time"

	auth "github.com/ChaoJiCaiNiao3/dymall/app/auth/kitex_gen/auth"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct {
}

// DeliverTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) DeliverTokenByRPC(ctx context.Context, req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	// 创建一个新的 JWT 令牌
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   string(req.UserId),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}

	return &auth.DeliveryResp{Token: tokenString}, nil
}

// VerifyTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) VerifyTokenByRPC(ctx context.Context, req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	// 验证 JWT 令牌
	token, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return &auth.VerifyResp{Res: false}, nil
	}

	if !token.Valid {
		return &auth.VerifyResp{Res: false}, nil
	}

	return &auth.VerifyResp{Res: true}, nil
}
